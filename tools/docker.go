package tools

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/cli/opts"
	"github.com/docker/go-connections/nat"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var cli, err = client.NewClientWithOpts(client.WithHost("tcp://127.0.0.1:2375"))

type ContainerUsage struct {
	containerId string
	cpuUsage    string
	memUsage    string
}

type ContainerPS struct {
	Id      string                   `json:"id"`
	Names   []string                 `json:"names"`
	ImageID string                   `json:"imageID"`
	Ports   []map[string]interface{} `json:"ports"`
	State   string                   `json:"state"`
	Status  string                   `json:"status"`
}

func PullImage(imageName string) error {
	defer func() {
		err := recover()
		if err != nil {
			errorf("Exec faild")
		}
	}()
	ctx := context.Background()

	events, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	d := json.NewDecoder(events)

	type Event struct {
		Status         string `json:"status"`
		Error          string `json:"error"`
		Progress       string `json:"progress"`
		ProgressDetail struct {
			Current int `json:"current"`
			Total   int `json:"total"`
		} `json:"progressDetail"`
	}

	var event *Event
	for {
		if err := d.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}

		}
	}
	if event != nil {
		if strings.Contains(event.Status, fmt.Sprintf("Downloaded newer image for %s", imageName)) {
			// new
			fmt.Println("new")
		}
		if strings.Contains(event.Status, fmt.Sprintf("Image is up to date for %s", imageName)) {
			// up-to-date
			fmt.Println("up-to-date")
		}
	}
	return nil

}

func ImageList() {
	ctx := context.Background()
	images, err := cli.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		println(err)
	}

	list, _ := json.Marshal(images)
	fmt.Println(string(list))
	fmt.Println("image size:", len(images))
}

func CreateContainer(imageName string, containerName string, cmd string, env []string, memoryGB int, cpuCore int, rportmap map[string]string, diskSizeGB int, gpuCount int) (container.ContainerCreateCreatedBody, error) {
	defer func() {
		err := recover()
		if err != nil {
			errorf("exec faild")
		}
	}()
	ctx := context.Background()
	var portMap = nat.PortMap{}
	var storage map[string]string = make(map[string]string)
	storage["size"] = strconv.Itoa(diskSizeGB) + "G"
	for k, v := range rportmap {
		portMap[nat.Port(v)] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: k,
			},
		}
	}
	var config = container.Config{
		Image:      imageName,
		Tty:        true,
		WorkingDir: "/",
	}
	cmd = strings.Trim(cmd, " ")
	if cmd != "" {
		cmds := strings.Split(cmd, " ")
		config.Cmd = cmds
	}

	if env != nil && len(env) > 0 {
		if len(env) == 1 && env[0] != "" {
			config.Env = env
		}
		if len(env) > 1 {
			config.Env = env
		}
	}
	var resource container.Resources
	if gpuCount > 0 {
		var gpuOpts = opts.GpuOpts{}
		gpuOpts.Set("all")
		resource = container.Resources{
			Memory:         int64(memoryGB) * 1024 * 1024 * 1024,
			NanoCPUs:       int64(cpuCore) * 1000000000,
			DeviceRequests: gpuOpts.Value(),
		}
	} else {
		resource = container.Resources{
			Memory:   int64(memoryGB) * 1024 * 1024 * 1024,
			NanoCPUs: int64(cpuCore) * 1000000000,
		}
	}

	resp, err := cli.ContainerCreate(ctx, &config, &container.HostConfig{
		Resources:    resource,
		PortBindings: portMap,
		StorageOpt:   storage,
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
	}, nil, nil, containerName)

	if err != nil {
		log.Printf("%v", err.Error())
		return resp, err
	}

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Printf("%v", err)
		return resp, err
	}

	//wait running
	//statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	//select {
	//case err := <-errCh:
	//	log.Printf()("container start error ", err.Error())
	//case status := <-statusCh:
	//	fmt.Println("container start success ", status.StatusCode)
	//}

	return resp, err

	//out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	//if err != nil {
	//	log.Printf()("get container log error ", err.Error())
	//}
	//io.Copy(os.Stdout, out)
}

func GetContainers() []ContainerPS {

	containerPs := new([]ContainerPS)
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return *containerPs
	}
	cbytes, _ := json.Marshal(containers)
	err2 := json.Unmarshal(cbytes, &containerPs)
	if err2 != nil {
		return *containerPs
	}
	return *containerPs
}

func StopContainer(containerID string) {
	ctx := context.Background()

	timeout := time.Minute
	err := cli.ContainerStop(ctx, containerID, &timeout)
	if err != nil {
		log.Printf("%v", err.Error())
	}
	fmt.Println("success")
}

func RemoveContainer(containerID string) error {
	defer func() {
		err := recover()
		if err != nil {
			errorf("RemoveContainer faild")
		}
	}()
	ctx := context.Background()

	err := cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true})
	if err != nil {
		log.Printf("%v", err.Error())
		return err
	}
	fmt.Println("success")
	return nil
}

func RemoveContainerByName(containerName string) {
	containerPS := GetContainers()
	for _, i := range containerPS {
		if i.Names[0] == "/"+containerName {
			_ = Retry(3, 1*time.Second, func() error {
				err2 := RemoveContainer(i.Id)
				if err2 != nil {
					return err2
				}
				return nil
			})

		}
	}
}

func ContainerLogs(containerID string) {
	ctx := context.Background()
	out, err := cli.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		log.Println("get container log error")
	}
	io.Copy(os.Stdout, out)
}

func ContainerCommit(containerID string) {
	ctx := context.Background()

	commitResp, err := cli.ContainerCommit(ctx, containerID, types.ContainerCommitOptions{Reference: "helloworld"})
	if err != nil {
		println(err)
	}

	b, _ := json.Marshal(commitResp)
	fmt.Println(string(b))

}
func ExecCheck(ctx context.Context, container string) bool {
	ir, err := cli.ContainerExecCreate(ctx, container, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          []string{"/bin/sh", "-c", "cat /etc/shells|grep /bin/bash"},
		Tty:          true,
		WorkingDir:   "/",
	})
	if err != nil {
		return false
	}

	// 附加到上面创建的/bin/bash进程中
	hr, err := cli.ContainerExecAttach(ctx, ir.ID, types.ExecStartCheck{Detach: false, Tty: true})
	if err != nil {
		return false
	}
	defer hr.Close()
	data, _ := ioutil.ReadAll(hr.Reader)
	res := string(data)
	if res == "" {
		return false
	} else if strings.Contains(res, "/bin/bash") {
		return true
	}
	return false
}
func Exec(container string) (hr types.HijackedResponse, err error) {
	defer func() {
		err := recover()
		if err != nil {
			errorf("Exec faild")
		}
	}()
	ctx := context.Background()
	command := "/bin/sh"
	if ExecCheck(ctx, container) {
		command = "/bin/bash"
	}
	ir, err := cli.ContainerExecCreate(ctx, container, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          []string{command},
		Tty:          true,
	})
	if err != nil {
		return
	}

	// 附加到上面创建的/bin/bash进程中
	hr, err = cli.ContainerExecAttach(ctx, ir.ID, types.ExecStartCheck{Detach: false, Tty: true})
	if err != nil {
		return
	}
	return
}

func statsContainer() ([]ContainerUsage, error) {
	if err != nil {
		println(err.Error())
	}
	defer func() {
		err := recover()
		if err != nil {
			errorf("Exec faild")
		}
	}()
	ctx := context.Background()
	println("")
	res, err := Command(ctx, "docker stats --all --no-stream|awk 'NR>2{print line}{line=$0} END{print line}'|awk '{print$1\" \"$3\" \"$7}'")
	if err != nil {
		return nil, err
	}
	var infos []string
	infos = strings.Split(strings.ReplaceAll(res, "\r\n", "\n"), "\n")
	if len(infos) < 3 {
		return nil, errors.New("docker check error")
	}
	var containerUsages = make([]ContainerUsage, len(infos))
	for i := 0; i < len(infos); i++ {
		if infos[i] == "" {
			continue
		}
		info := strings.Split(infos[i], " ")
		containerUsage := new(ContainerUsage)
		containerUsage.containerId = info[0]
		containerUsage.cpuUsage = info[1]
		containerUsage.memUsage = info[2]
		containerUsages[i] = *containerUsage
	}
	return containerUsages, nil

}

func read(ctx context.Context, wg *sync.WaitGroup, std io.ReadCloser, res *string) {

	reader := bufio.NewReader(std)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			readString, err := reader.ReadString(' ')
			*res = *res + readString
			if err != nil || err == io.EOF {
				return
			}
		}
	}
}

func Command(ctx context.Context, cmd string) (string, error) {
	var res = ""
	//c := exec.CommandContext(ctx, "cmd", "/C", cmd) // windows
	c := exec.CommandContext(ctx, "bash", "-c", cmd) // mac linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return res, err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return res, err
	}
	var wg sync.WaitGroup
	wg.Add(2)

	go read(ctx, &wg, stderr, &res)
	go read(ctx, &wg, stdout, &res)
	err = c.Start()
	wg.Wait()
	println(res)
	return res, err
}

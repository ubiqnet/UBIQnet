package tools

import (
	"bytes"
	"encoding/json"
	"github.com/ethereum/go-ethereum/log"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Netinfo struct {
	IP          string `json:"ip"`
	IsPublic    bool
	CountryCode string `json:"country_code2"`
	Country     string `json:"country_name"`
	RegionName  string `json:"state_prov"`
	City        string `json:"city"`
	Isp         string `json:"isp"`
	Lat         string `json:"latitude"`
	Lon         string `json:"longitude"`
	ERR         string
}

var rep = RandString(20)

type HTTP struct {
	listener net.Listener
}

func (h *HTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(rep))
	if err != nil {
		return
	}
}
func RandString(codeLen int) string {

	rawStr := "jkwangagDGFHGSERKILMJHSNOPQR546413890_"

	buf := make([]byte, 0, codeLen)
	b := bytes.NewBuffer(buf)

	rand.Seed(time.Now().UnixNano())
	for rawStrLen := len(rawStr); codeLen > 0; codeLen-- {
		randNum := rand.Intn(rawStrLen)
		b.WriteByte(rawStr[randNum])
	}
	return b.String()
}

func startServer(p *int) {

	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:0")
	if err != nil {
		return
	}
	l, _ := net.ListenTCP("tcp", addr)
	*p = l.Addr().(*net.TCPAddr).Port
	h := new(HTTP)
	h.listener = l
	go func() {
		err := http.Serve(h.listener, h)
		if err != nil {
			return
		}
	}()
	time.Sleep(time.Second * 60)
	err = h.listener.Close()
	if err != nil {
		return
	}
}

func checkPorts(ipPorts []string) int {
	for _, ipPort := range ipPorts {

		conn, err := net.DialTimeout("tcp", ipPort, 3*time.Second)
		if err != nil {
			log.Info("prot:" + ipPort + " check fail!~")
			return 0
		} else {
			if conn != nil {
				log.Info("prot:" + ipPort + " check success!~")
				_ = conn.Close()
			} else {
				log.Info("prot:" + ipPort + " check fail!~")
				return 0
			}
		}
	}
	return 1
}

func checkAccess() Netinfo {
	timeout := 10 * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	netInfo := Netinfo{
		IP:          "Unknown",
		IsPublic:    false,
		CountryCode: "Unknown",
		Country:     "Unknown",
		RegionName:  "Unknown",
		City:        "Unknown",
		Isp:         "Unknown",
		Lat:         "0",
		Lon:         "0",
		ERR:         "",
	}

	url := "https://api.ipgeolocation.io/ipgeo"
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		netInfo.ERR = err.Error()
		return netInfo
	}
	req.Header.Add("authority", "api.ipgeolocation.io")
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,ru;q=0.8")
	req.Header.Add("dnt", "1")
	req.Header.Add("origin", "https://ipgeolocation.io")
	req.Header.Add("referer", "https://ipgeolocation.io/")
	req.Header.Add("sec-ch-ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	var ipResponse *http.Response
	err = Retry(3, time.Second*3, func() error {

		ipResponse, err = client.Do(req)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		netInfo.ERR = err.Error()
		println("error:no public ip!")
		netInfo.ERR = err.Error()
		return netInfo
	}
	res, _ := ioutil.ReadAll(ipResponse.Body)
	err = json.Unmarshal(res, &netInfo)
	if err != nil {
		netInfo.ERR = err.Error()
	}
	println("public ip:" + string(netInfo.IP))
	var p1 int = 0
	var p2 int = 0
	var p3 int = 0
	go startServer(&p1)
	go startServer(&p2)
	go startServer(&p3)
	time.Sleep(time.Second * 1)
	var ps = [3]int{p1, p2, p3}
	for p := range ps {
		address := "http://" + string(netInfo.IP) + ":" + strconv.Itoa(p)
		ipResponse, err = client.Get(address)
		if err == nil {
			res, _ := ioutil.ReadAll(ipResponse.Body)
			if string(res) == rep {
				println("Test Success")
				netInfo.IsPublic = true
				continue
			}
		}
		println("")
		println("Public IP tested Failed,Maybe you should turn off the firewall")
		netInfo.IsPublic = false
	}
	return netInfo
}

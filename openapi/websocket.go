package openapi

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"ubiqnet/constants"
	"ubiqnet/service"
	"ubiqnet/store/localstore"
	"ubiqnet/tools"
)

func StartWebsocket() {
	http.HandleFunc("/terminal", terminal)
	srv := &http.Server{
		Addr: "0.0.0.0:8066",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("pong"))
}

func terminal(w http.ResponseWriter, r *http.Request) {
	// websocket握手
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}
	defer conn.Close()

	err = r.ParseForm()
	if err != nil {
		println("form error")
		return
	}
	// 获取容器ID或name
	container := r.Form.Get("container")
	msg := r.Form.Get("msg")
	timestamp := r.Form.Get("timestamp")
	sign := r.Form.Get("sign")
	address, err := service.VerifyMessage(msg+timestamp, sign)
	if err != nil {
		println("VerifyMessage error")
		return
	}
	infoStr := localstore.Get(constants.STOP_CONTAONER_KEY_PRIFIX + container)
	imageInfo := service.ToImageInfo(infoStr)
	maxTimestaamp, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		println("timestaamp error")
		return
	}
	ownerAddress := imageInfo.OwnerAddress
	if ownerAddress == "" {
		return
	}
	if !strings.EqualFold(ownerAddress, address) || (maxTimestaamp+30*60*1000) < time.Now().UTC().Unix() {
		println(ownerAddress)
		println(address)
		println(time.Now().UTC().Unix())
		println((maxTimestaamp + 30*60))
		text := "error!"
		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			return
		}
		return
	}
	// 执行exec，获取到容器终端的连接
	hr, err := tools.Exec(container)
	if err != nil {
		log.Error(err)
		return
	}
	// 关闭I/O流
	defer hr.Close()
	// 退出进程
	defer func() {
		hr.Conn.Write([]byte("exit\r"))
	}()

	go func() {
		wsWriterCopy(hr.Conn, conn)
	}()
	wsReaderCopy(conn, hr.Conn)
}

func wsWriterCopy(reader io.Reader, writer *websocket.Conn) {
	buf := make([]byte, 8192)
	for {
		nr, err := reader.Read(buf)
		if nr > 0 {
			err := writer.WriteMessage(websocket.BinaryMessage, buf[0:nr])
			if err != nil {
				return
			}
		}
		if err != nil {
			return
		}
	}
}

func wsReaderCopy(reader *websocket.Conn, writer io.Writer) {
	for {
		messageType, p, err := reader.ReadMessage()
		if err != nil {
			return
		}
		if messageType == websocket.TextMessage {
			writer.Write(p)
		}
	}
}

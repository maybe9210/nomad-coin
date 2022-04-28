package p2p

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/maybe9210/nomad-coin/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	openPort := r.URL.Query().Get("openPort")
	ip := utils.Splitter(r.RemoteAddr, ":", 0)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return openPort != "" && ip != ""
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)

	initPeer(conn, ip, openPort)
}

func AddPeer(address, port, openPort string) {
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws/openPort=%s", address, port, openPort[1:]), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}

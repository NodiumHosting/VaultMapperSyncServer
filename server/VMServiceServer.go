package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"regexp"

	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var data []byte

var uuidRegex, _ = regexp.Compile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
var vaultIDRegex, _ = regexp.Compile("^vault_[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")

func handshakeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("upgrade") != "websocket" { // guard against non websocket requests
		return
	}

	uuid := r.URL.Query().Get("uuid")
	vaultID := r.URL.Query().Get("vaultID") // if checks pass, upgrade
	log.Printf(vaultID + ": " + uuid)

	if !uuidRegex.MatchString(uuid) || !vaultIDRegex.MatchString(vaultID) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		log.Println("Check not passed: " + uuid)
		return // close the ws basically..
	}
	log.Printf("Connection successful: " + uuid)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	// after this point is the websocket loop
	// conn.ReadMessage() reads the message, works like onMessage
	// use onClose to do stuff after closing socket

	defer onClose(uuid, vaultID)

	// this should basically be the onMessage thingy
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			return
		}
	}

}

func onClose(uuid string, vaultID string) {
	log.Println(uuid + " closed connection to vault: " + vaultID)
}

func Run(ip string, port int) {
	fmt.Println("HELLO FROM SERVER")

	http.HandleFunc("/", handshakeHandler)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", ip, port), nil); err != nil {
		log.Fatal(err)
	}
}

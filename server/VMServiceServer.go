package server

import (
	"fmt"
	pb "github.com/NodiumHosting/VaultMapperSyncServer/proto"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
	"regexp"
	"time"

	"net/http"
)

// HUB stores all Vault structures
var HUB = Hub{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

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

	SendVault(vaultID, conn) // send vault to client

	ok := HUB.AddConnectionToVault(vaultID, uuid, conn)
	_ = AddPlayerToVault(uuid, vaultID) // add player to vault db

	if !ok { // if not ok -> connection exists -> return/close connection
		_ = conn.WriteMessage(websocket.CloseMessage, nil)
		err := conn.Close()
		if err != nil {
			return
		}
		return
	}

	defer onClose(uuid, vaultID)

	// this should basically be the onMessage thingy
	for {
		err := conn.SetReadDeadline(time.Now().Add(15 * time.Second))
		if err != nil {
			return
		}
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read: " + err.Error())
			return
		}

		if string(data) == "keep_me_alive" {
			//log.Println("Keep alive received")
			continue
		}

		var msg pb.Message
		err2 := proto.Unmarshal(data, &msg)
		if err2 != nil {
			log.Println("Marshal problem")
			return
		}
		onMessage(vaultID, uuid, &msg)
	}
}

func onMessage(vaultID string, uuid string, msg *pb.Message) {
	//log.Printf("\nOn message from %s\ntype: %v\n", uuid, msg.GetType())
	msgType := msg.GetType()
	switch msgType {
	case pb.MessageType_VAULT_PLAYER:
		// this case handles accepted player packet
		HandlePlayerMovement(vaultID, uuid, msg)
		break
	case pb.MessageType_VAULT_CELL:
		// This case should handle accepted VaultCell
		HandleVaultCell(vaultID, uuid, msg)
		break
	case pb.MessageType_PLAYER_DISCONNECT:
		// this shouldn't happen as PlayerDisconnect is only S2C
		log.Println(uuid + " tried to send PlayerDisconnect which shouldn't happen")
		break
	case pb.MessageType_VAULT:
		// this shouldn't happen as the Vault is only S2C
		log.Println(uuid + " tried to send Vault which shouldn't happen")
		break
	case pb.MessageType_TOAST:
		// this shouldn't happen as the Toast is only S2C
		log.Println(uuid + " tried to send Toast which shouldn't happen")
		break
	default:
		log.Println(uuid + " sent unknown packet")
		break
	}
}

func onClose(uuid string, vaultID string) { // need to send down PlayerDisconnect to others in vault here
	log.Println(uuid + " closed connection to vault: " + vaultID)
	msg := pb.Message{
		Type:             pb.MessageType_PLAYER_DISCONNECT,
		PlayerDisconnect: &pb.PlayerDisconnect{Uuid: uuid},
	}
	BroadcastMessage(vaultID, uuid, &msg)
	HUB.RemoveConnectionFromVault(vaultID, uuid)
}

// HandlePlayerMovement handles incoming PlayerMovement packets from clients and broadcasts them to the other players
func HandlePlayerMovement(vaultID string, uuid string, msg *pb.Message) {
	//log.Println("Handling PlayerMovement")
	BroadcastMessage(vaultID, uuid, msg)
}

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// HandleVaultCell handles incoming VaultCell packets from clients, broadcasts them to the other players and adds them to internal structures
func HandleVaultCell(vaultID string, uuid string, msg *pb.Message) {
	log.Println("Handling VaultCell")
	cell := msg.GetVaultCell()
	if cell.GetCellType() == pb.CellType_CELLTYPE_ROOM {
		if cell.GetX()%2 != 0 || cell.GetZ()%2 != 0 {
			log.Println("Sent invalid room cell", uuid)
			return
		}
	} else if cell.GetCellType() == pb.CellType_CELLTYPE_TUNNEL_X {
		if Abs(cell.GetX())%2 != 1 || cell.GetZ()%2 != 0 {
			log.Println("Sent invalid tunnel cellz", uuid)
			return
		}
	} else if cell.GetCellType() == pb.CellType_CELLTYPE_TUNNEL_Z {
		if cell.GetX()%2 != 0 || Abs(cell.GetZ())%2 != 1 {
			log.Println("Sent invalid tunnel cellx", uuid)
			return
		}
	}

	BroadcastMessage(vaultID, uuid, msg)
	vault := HUB.GetVault(vaultID)
	if vault == nil {
		return
	}

	vault.AddOrReplaceCell(cell)
}

// BroadcastMessage is used to broadcast Message to a vault, with excludeUUID being excluded
func BroadcastMessage(vaultID string, excludeUUID string, msg *pb.Message) {
	vault := HUB.GetVault(vaultID) // get vault
	if vault == nil {
		return
	}
	messageBuffer, err := proto.Marshal(msg) // serialize message into buffer
	if err != nil {
		return
	}
	//log.Println("Buffer ready, broadcasting")
	vault.Connections.Range(func(key, val interface{}) bool { // go through connections and add to their Send channels
		if key != excludeUUID {
			conn := val.(*Connection)
			conn.Send <- messageBuffer
			//log.Println("sent to buffer")
		}
		return true
	})
}

// SendVault sends all the Vault.Cells using the Vault message type initially to sync vault to client if joined after start
func SendVault(vaultID string, conn *websocket.Conn) {
	vault := HUB.GetVault(vaultID)
	if vault == nil {
		log.Println("Tried to send vault that doesn't exist, it will be created now")
		return // if vault doesn't exist, do nothing - this can happen when this is the first player joining a fresh vault
	}
	log.Println("Sending vault to client")
	var cells []*pb.VaultCell
	vault.Cells.Range(func(key, val interface{}) bool {
		cells = append(cells, val.(*pb.VaultCell))
		//log.Println("appended cell")
		return true
	})

	msg := pb.Message{
		Type:  pb.MessageType_VAULT,
		Vault: &pb.Vault{Cells: cells},
	}

	messageBuffer, err := proto.Marshal(&msg)
	if err != nil {
		return
	}

	errr := conn.WriteMessage(websocket.BinaryMessage, messageBuffer)
	if errr != nil {
		return
	}
}

func Run(ip string, port int) {
	fmt.Println("HELLO FROM SERVER")

	http.HandleFunc("/", handshakeHandler)
	http.Handle("/stats", rateLimit(http.HandlerFunc(statsHandler)))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", ip, port), nil); err != nil {
		log.Fatal(err)
	}
}

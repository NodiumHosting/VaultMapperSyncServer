package server

import (
	"fmt"
	pb "github.com/NodiumHosting/VaultMapperSyncServer/proto"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

// Vault is a helper struct that is used to store connections to vault and other runtime data
//
// If other data is added, there should be a Mutex created with it to make sure writing is thread-safe
type Vault struct {
	UUID        string
	Connections sync.Map // stores a map of current connections inside the vault, key is uuid, value is Connection
	Cells       sync.Map // stores a map of cells inside the vault, key is x,z, value is pb.VaultCell
}

// AddConnection adds the connection to Vault structure and starts up the WritePump
//
// ok is false if the connection already exists, else false
func (v *Vault) AddConnection(playerUUID string, conn *websocket.Conn) bool {
	_, ok := v.Connections.Load(playerUUID)
	if ok {
		log.Println("Tried to add connection but it already exists")
		return false // connection already exists
	}
	c := &Connection{ // create connection
		uuid: playerUUID,
		conn: conn,
		Send: make(chan []byte, 256), // buffered channel of 256 bytes
	}

	v.Connections.Store(playerUUID, c) // store the connection inside vault
	go c.WritePump()                   // Start the write pump
	return true
}

// RemoveConnection removes the connection from Vault structure and closes the Send channel
//
// return true if the Vault is empty after connection removal, false otherwise
//
// Send channel needs to be closed for WritePump to exit properly!
func (v *Vault) RemoveConnection(playerUUID string) bool {
	value, ok := v.Connections.Load(playerUUID)
	if !ok {
		return false
	}

	c := value.(*Connection)
	close(c.Send)         // close send channel
	err := c.conn.Close() // close connection
	if err != nil {
		return false
	}
	v.Connections.Delete(playerUUID) // remove connection

	// check if the vault is empty now
	isEmpty := true
	v.Connections.Range(func(k, v interface{}) bool {
		isEmpty = false
		return false
	})
	if isEmpty {
		return true
	}

	return false
}

// AddOrReplaceCell adds a VaultCell to the Cells map
func (v *Vault) AddOrReplaceCell(cell *pb.VaultCell) {
	key := fmt.Sprintf("%d,%d", cell.X, cell.Z)
	v.Cells.Store(key, cell)
}

// RemoveCell removes a VaultCell from the Cells map
func (v *Vault) RemoveCell(x, z int) {
	key := fmt.Sprintf("%d,%d", x, z)
	v.Cells.Delete(key)
}

// GetCell retrieves a VaultCell from the Cells map
func (v *Vault) GetCell(x, z int) (*pb.VaultCell, bool) {
	key := fmt.Sprintf("%d,%d", x, z)
	value, ok := v.Cells.Load(key)
	if !ok {
		return nil, false
	}
	return value.(*pb.VaultCell), true
}

// IterateCells runs a provided function on each cell in the vault
func (v *Vault) IterateCells(f func(key string, cell *pb.VaultCell)) { // iterate over cells
	v.Cells.Range(func(k, v interface{}) bool {
		key := k.(string)
		cell := v.(*pb.VaultCell)
		f(key, cell)
		return true
	})
}

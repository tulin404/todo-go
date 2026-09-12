package ipc

import (
	"encoding/json"
	"net"
)

// const 'Address' is the universal
const Address = "127.0.0.1:9876"

// 'Send' sends a received command to the TCP connection (IPC)
// *IPC is handled by TCP for multisystem functionality*
func Send(command Command) error {
	conn, err := net.Dial("tcp", Address)
	if err != nil {
		return err
	}
	defer conn.Close()

	return json.NewEncoder(conn).Encode(command)
}

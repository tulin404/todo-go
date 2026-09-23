package ipc

import (
	"encoding/json"
	"net"
)

// 'Listen' listen to the Unix socket
func Listen() (net.Listener, error) {
	return net.Listen("unix", socketPath)
}

// 'AcceptCommands' is a loop for accepting connection commands and redirecting them to the commands channel
func AcceptCommands(listener net.Listener, commands chan<- Command) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		var command Command

		err = json.NewDecoder(conn).Decode(&command)
		if err != nil {
			conn.Close()
			continue
		}

		commands <- command

		conn.Close()
	}
}

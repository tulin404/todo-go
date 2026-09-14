package ipc

import (
	"encoding/json"
	"net"
	"os"
	"path/filepath"
)

// 'getSocketPath' returns the unix socket path (IPC)
func getSocketPath() string {
	runtimeDir := os.Getenv("$XDG_RUNTIME_DIR")
	if runtimeDir != "" {
		return filepath.Join(runtimeDir, "todo.sock")
	}

	return "/tmp/todo.sock"
}

// const 'Path' is the universal path for the unix socket (IPC)
var socketPath = getSocketPath()

// 'Send' sends a received command to the socket connection (IPC)
func Send(command Command) error {
	conn, err := net.Dial("tcp", socketPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	return json.NewEncoder(conn).Encode(command)
}

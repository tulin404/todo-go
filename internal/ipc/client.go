package ipc

import (
	"encoding/json"
	"net"
	"os"
	"path/filepath"
)

// 'getSocketPath' returns the Unix socket path (IPC)
func getSocketPath() string {
	runtimeDir := os.Getenv("$XDG_RUNTIME_DIR")
	if runtimeDir != "" {
		return filepath.Join(runtimeDir, "todo.sock")
	}

	return "/tmp/todo.sock"
}

// var 'socketPath' is the universal path for the Unix socket (IPC)
var socketPath = getSocketPath()

// 'Send' sends a received command to the Unix socket connection (IPC)
func Send(command Command) error {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	return json.NewEncoder(conn).Encode(command)
}

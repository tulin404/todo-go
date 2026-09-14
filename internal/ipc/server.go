package ipc

import "net"

// 'Listen' listen to the Unix socket
func Listen() (net.Listener,error) {
	return net.Listen("unix", socketPath)
}

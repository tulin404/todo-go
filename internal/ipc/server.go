package ipc

import "net"

func Listen() (net.Listener,error) {
	return net.Listen("unix", socketPath)
}
// VAMOS MUDAR PARA SOCKET MESMO AO INVES DE CONEXOES TCP POR CONTA QUE SYSTEMD DA TRABALHO NO WINDOWS
func Receive()

package rtcp

import (
  "net"
)

type indirectServer struct {
  // Implement Server

  conn net.Conn

  status int
}

// func (this *indirectServer) Accept() (Conn, error) {
//   return this.l.Accept()
// }
//
// func (this *indirectServer) Close() error {
//   return this.l.Close()
// }

func NewInDirectServer(dphostport string, name string) (Server, error) {
  var server = &indirectServer{}
  var err error
  if server.conn, err = net.Dial("tcp", dphostport); err != nil {
    return nil, err
  }
  return nil, nil
}

package rtcp

/****
 *  Direct Server definition
 *  Requires nothing, but ip:port to listen. have to be publically accessible
 */

import (
  "net"
)

type directServer struct {
  // Implement Server
  l net.Listener
}

func (this *directServer) Accept() (Conn, error) {
  return this.l.Accept()
}

func (this *directServer) Close() error {
  return this.l.Close()
}

func NewDirectServer(hostport string) (Server, error) {
  if ln, err := net.Listen("tcp", hostport); err != nil {
    return nil, err
  } else {
    // Okay, return the TCP Server as it is.
    return &directServer{ln}, nil
  }
}


/******************************************************************************/

func NewDirectClient(hostport string) (Conn, error) {
  return net.Dial("tcp", hostport)
}

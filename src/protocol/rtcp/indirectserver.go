package rtcp

import (
  "net"
)

type indirectServer struct {
  // Implement Server

  conn net.Conn

  status int
}

func (this *indirectServer) Accept() (Conn, error) {
  return nil, nil
}

func (this *indirectServer) Close() error {
  return nil
}

func NewInDirectServer(dphostport string, name string) (Server, error) {
  var server = &indirectServer{}
  var err error
  if server.conn, err = net.Dial("tcp", dphostport); err != nil {
    LOG.Println("Fail to connect to DP:", err)
    return nil, err
  }

  if err = RMessageSend(&RMsgAloha{
       V: ProtocolVersion(),
       R: TypeRRoleIS,
       N: name,
     }, server.conn); err != nil {
    LOG.Println("Fail to send aloha to DP:", err)
    server.conn.Close()
    return nil, err
  }

  var recvBuffer = make([]byte, SafeMessageBufferSize, SafeMessageBufferSize)
  if msg, err := RMessageRcvX(server.conn, TypeRMsgAloha, recvBuffer[:]);
     err != nil {
    LOG.Println("Fail to receive aloha from DP:", err)
    server.conn.Close()
    return nil, err
  } else {
    // Check whether it is a valid DP
    var returnedAloha = msg.(*RMsgAloha)
    if returnedAloha.R != TypeRRoleDP || !isCompatible(returnedAloha.V) {
      LOG.Println("Target is not a good DP")
      server.conn.Close()
      return nil, err
    }
    // TODO record heartbeat
    LOG.Println("Connect to DP, aloha =", returnedAloha)
    return server, nil
  }

  return nil, nil
}

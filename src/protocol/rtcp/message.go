package rtcp

/****
 *  Defines common message types in the protocol.
 *  There're three roles in
 */

const (
  TypeRMsgAloha = 0
  TypeRMsgConnect = 1
  TypeRMsgAck = 2
  TypeRMsgData = 3
  TypeRMsgClose = 4
) // TypeRMsg

const (
  TypeRRoleDS = 0
  TypeRRoleDC = 1
  TypeRRoleIS = 2
  TypeRRoleIC = 3
  TypeRRoleDP = 4
  TypeRAbortVersion = 5
) // TypeRRole

type RMsgAloha struct {
  V int
  R int   // TypeRRole
  N string
}

type RMsgConnect struct {
  N string
}

type RMsgAck struct {
  X int
  R string
}

type RMsgData struct {
  X int
  C []byte
}

type RMsgClose struct {
  X int
  R string
}

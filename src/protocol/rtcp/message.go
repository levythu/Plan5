package rtcp

/****
 *  Defines common message types in the protocol.
 *  There're three roles in
 */

import (
  "io"
  "errors"
  "sync"
  "fmt"
  "bufio"
  "encoding/json"
)

const (
  MaxMessageLength = 4 * 1024
  SafeMessageBufferSize = MaxMessageLength + 10
)

const (
  TypeRMsgAloha = '0'
  TypeRMsgConnect = '1'
  TypeRMsgAck = '2'
  TypeRMsgData = '3'
  TypeRMsgClose = '4'
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
  HeartBeat int
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

/******************************************************************************/
// RMsg IO

func RMessageSend(msg interface{}, writer io.Writer) error {
  return RMessageSendX(msg, writer, nil)
}

func RMessageSendX(msg interface{}, writer io.Writer, lock *sync.Mutex) error {
  var typeToGo [1]byte
  switch msg.(type) {
  case *RMsgAloha:
    typeToGo[0] = TypeRMsgAloha
  case *RMsgConnect:
    typeToGo[0] = TypeRMsgConnect
  case *RMsgAck:
    typeToGo[0] = TypeRMsgAck
  case *RMsgData:
    typeToGo[0] = TypeRMsgData
  case *RMsgClose:
    typeToGo[0] = TypeRMsgClose
  default:
    return errors.New("msg must be a valid RMsg type.")
  }
  if jsonContent, err := json.Marshal(msg); err != nil {
    return err
  } else {
    if len(jsonContent) > MaxMessageLength {
      return errors.New("The message is too long. Caller needs to slice it.")
    }
    if lock != nil {
      lock.Lock()
      defer lock.Unlock()
    }
    if _, err := writer.Write(typeToGo[:]); err != nil {
      return err
    }
    if _, err := writer.Write([]byte("\n")); err != nil {
      return err
    }
    if _, err := writer.Write(jsonContent); err != nil {
      return err
    }
    if _, err := writer.Write([]byte("\n")); err != nil {
      return err
    }
  }
  return nil
}

func RMessageRcv(reader io.Reader) (interface{}, error) {
  return RMessageRcvX(reader, '*', nil)
}

// Assert = '*' for wildcard
func RMessageRcvX(
    reader io.Reader, assert byte, buf []byte) (interface{}, error) {
  if buf == nil {
    buf = make([]byte, SafeMessageBufferSize, SafeMessageBufferSize)
  }

  var scanner = bufio.NewScanner(reader)
  scanner.Buffer(buf, cap(buf))


  // First scan, get the message type
  scanner.Scan()
  var msgType = scanner.Bytes()
  if len(msgType) != 1 {
    return nil, errors.New("Broken protocol.")
  }

  scanner.Scan()
  var msgContent = scanner.Bytes()

  if assert != '*' && msgType[0] != assert {
    return nil, errors.New(
        fmt.Sprintf("Unmatched message type, expect %c, got %c",
                    assert, msgType[0]))
  }

  var res interface{}
  switch msgType[0] {
  case TypeRMsgAloha:
    res = &RMsgAloha{}
  case TypeRMsgConnect:
    res = &RMsgConnect{}
  case TypeRMsgAck:
    res = &RMsgAck{}
  case TypeRMsgData:
    res = &RMsgData{}
  case TypeRMsgClose:
    res = &RMsgClose{}
  default:
    return nil, errors.New("Receive invalid RMsg type.")
  }
  if err := json.Unmarshal(msgContent, res); err != nil {
    return nil, err
  }

  return res, nil
}

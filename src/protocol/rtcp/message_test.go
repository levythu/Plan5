package rtcp

import (
  "os"
  "bytes"
  "reflect"
  _ "fmt"
  "testing"
)

func TestMessageSend(t *testing.T) {
  var msg RMsgAloha
  if err := RMessageSend(&msg, os.Stdout); err != nil {
    t.Error("Error in sending the msg:", err)
  }
}

func TestMessageSendData(t *testing.T) {
  var msg RMsgData
  msg.C = []byte{0,1,2,6,22,55,12,71,88,12,20,255}
  if err := RMessageSend(&msg, os.Stdout); err != nil {
    t.Error("Error in sending the msg:", err)
  }
  msg.C = []byte("Huahua aichi baomihua! 你好。")
  if err := RMessageSend(&msg, os.Stdout); err != nil {
    t.Error("Error in sending the msg:", err)
  }
}

func TestMessageReceiveData(t *testing.T) {
  var b bytes.Buffer

  var msg RMsgData
  msg.C = []byte("Huahua aichi baomihua! 你好。")
  if err := RMessageSend(&msg, &b); err != nil {
    t.Error("Error in sending the msg:", err)
  }
  if res, err := RMessageRcv(&b); err != nil {
    t.Error("Error in receiving the msg:", err)
  } else {
    var aft = *(res.(*RMsgData))
    if !reflect.DeepEqual(msg, aft) {
      t.Error("Not the same data. Before:", msg, "After", aft)
    }
  }
}

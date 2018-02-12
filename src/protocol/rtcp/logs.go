package rtcp

import (
  "log"
  "os"
)

var LOG = log.New(os.Stderr, "rtcp: ", log.Lshortfile | log.LstdFlags)

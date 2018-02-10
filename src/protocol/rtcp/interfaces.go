package rtcp

type Conn interface {
  Close() error
  Read(b []byte) (int, error)
  Write(b []byte) (int, error)
}

type Server interface {
  Accept() (Conn, error)
  Close() error
}

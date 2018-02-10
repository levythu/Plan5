package rtcp

const _CURRENT_VERSION = 0x00071

func ProtocolVersion() int {
  return _CURRENT_VERSION
}

// Check whether target version number is compatible with myself.
func isCompatible(targetVersion int) bool {
  return true
}

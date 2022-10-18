package server

import protocol "github.com/tliron/glsp/protocol_3_16"

func uintToInt(v protocol.UInteger) int {
	return int(v) + 1
}

//func intToUInt(v int) protocol.UInteger {
//return protocol.UInteger(v) - 1
//}

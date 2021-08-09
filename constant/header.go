package constant

type (
	// Header is protocol header
	Header struct {
		SerialId  uint16  // serial id
		Version   Version // protocol version
		Timestamp int64   // timestamp
		Source    uint64  // source address
		Target    uint64  // target address
	}
)

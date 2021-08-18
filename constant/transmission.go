package constant

type (
	// TransmissionStateInfo is transmission state information
	TransmissionStateInfo struct {
		Flag      StateFlag // state flag
		Timestamp int64     // occurred timestamp
	}

	// TransmissionOperationInfo is transmission operation information
	TransmissionOperationInfo struct {
		Flag      OperationFlag // operation flag
		Operator  int           // operator
		Timestamp int64         // occurred timestamp
	}

	// TransmissionTimestamp is transmission timestamp
	TransmissionTimestamp struct {
		Timestamp int64
	}

	// TransmissionVersion is transmission version
	TransmissionVersion struct {
		Version Version
	}

	// TransmissionConfigure is transmission configure
	TransmissionConfigure struct {
		Configure string
	}
)

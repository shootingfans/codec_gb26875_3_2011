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

	// TransmissionTimestamp 用户传输装置系统时间
	TransmissionTimestamp struct {
		Timestamp int64
	}

	// TransmissionVersion 用户传输装置系统版本
	TransmissionVersion struct {
		Version Version
	}

	// TransmissionConfigure 用户传输装置配置信息
	TransmissionConfigure struct {
		Configure string
	}
)

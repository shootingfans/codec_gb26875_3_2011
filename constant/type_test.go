package constant

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateFlag_BitIndexBool(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		s    StateFlag
		args args
		want bool
	}{
		{name: "test bit index 0", s: 0xffff, args: args{index: 0}, want: false},
		{name: "test bit index 1", s: 0xfffe, args: args{index: 1}, want: false},
		{name: "test bit index 2", s: 0xfffd, args: args{index: 2}, want: false},
		{name: "test bit index 3", s: 0xfffb, args: args{index: 3}, want: false},
		{name: "test bit index 4", s: 0xfff7, args: args{index: 4}, want: false},
		{name: "test bit index 5", s: 0xffef, args: args{index: 5}, want: false},
		{name: "test bit index 6", s: 0xffdf, args: args{index: 6}, want: false},
		{name: "test bit index 7", s: 0xffbf, args: args{index: 7}, want: false},
		{name: "test bit index 8", s: 0xff7f, args: args{index: 8}, want: false},
		{name: "test bit index 9", s: 0xfeff, args: args{index: 9}, want: false},
		{name: "test bit index 10", s: 0xfdff, args: args{index: 10}, want: false},
		{name: "test bit index 11", s: 0xfbff, args: args{index: 11}, want: false},
		{name: "test bit index 12", s: 0xf7ff, args: args{index: 12}, want: false},
		{name: "test bit index 13", s: 0xefff, args: args{index: 13}, want: false},
		{name: "test bit index 14", s: 0xdfff, args: args{index: 14}, want: false},
		{name: "test bit index 15", s: 0xbfff, args: args{index: 15}, want: false},
		{name: "test bit index 16", s: 0x7fff, args: args{index: 16}, want: false},
		{name: "test bit index 17", s: 0xffff, args: args{index: 17}, want: false},
		{name: "test bit index 1", s: 0x0001, args: args{index: 1}, want: true},
		{name: "test bit index 2", s: 0x0002, args: args{index: 2}, want: true},
		{name: "test bit index 3", s: 0x0004, args: args{index: 3}, want: true},
		{name: "test bit index 4", s: 0x0008, args: args{index: 4}, want: true},
		{name: "test bit index 5", s: 0x0010, args: args{index: 5}, want: true},
		{name: "test bit index 6", s: 0x0020, args: args{index: 6}, want: true},
		{name: "test bit index 7", s: 0x0040, args: args{index: 7}, want: true},
		{name: "test bit index 8", s: 0x0080, args: args{index: 8}, want: true},
		{name: "test bit index 9", s: 0x0100, args: args{index: 9}, want: true},
		{name: "test bit index 10", s: 0x0200, args: args{index: 10}, want: true},
		{name: "test bit index 11", s: 0xf400, args: args{index: 11}, want: true},
		{name: "test bit index 12", s: 0x0800, args: args{index: 12}, want: true},
		{name: "test bit index 13", s: 0x1000, args: args{index: 13}, want: true},
		{name: "test bit index 14", s: 0x2000, args: args{index: 14}, want: true},
		{name: "test bit index 15", s: 0x4000, args: args{index: 15}, want: true},
		{name: "test bit index 16", s: 0x8000, args: args{index: 16}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, tt.s.BitIndexBool(tt.args.index), tt.want)
		})
	}
}

func TestStateFlag_Info(t *testing.T) {
	type args struct {
		mapper StateFlagBitMapper
	}
	tests := []struct {
		name     string
		s        StateFlag
		args     args
		wantInfo StateInfo
	}{
		{name: "test transmission 0x0000", s: 0x0000, args: args{mapper: StdTransmissionStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest}},
		{name: "test transmission 0x0001", s: 0x0001, args: args{mapper: StdTransmissionStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeNormal}},
		{name: "test transmission 0x0002", s: 0x0002, args: args{mapper: StdTransmissionStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfInAlarm}},
		{name: "test transmission 0x0004", s: 0x0004, args: args{mapper: StdTransmissionStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfInFault}},
		{name: "test transmission 0x0008", s: 0x0008, args: args{mapper: StdTransmissionStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfMainPowerFault}},
		{name: "test transmission 0x0010", s: 0x0010, args: args{mapper: StdTransmissionStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfBackupPowerFault}},
		{name: "test transmission 0x0020", s: 0x0020, args: args{mapper: StdTransmissionStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfCommunicationFault}},
		{name: "test transmission 0x0040", s: 0x0040, args: args{mapper: StdTransmissionStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfMonitorConnectionFault}},
		{name: "test transmission 0x0080", s: 0x0080, args: args{mapper: StdTransmissionStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest}},
		{name: "test controller 0x0000", s: 0x0000, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x0001", s: 0x0001, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeNormal, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x0002", s: 0x0002, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfInAlarm, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x0004", s: 0x0004, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfInFault, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x0008", s: 0x0008, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfInBlock, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x0010", s: 0x0010, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfInSupervise, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x0020", s: 0x0020, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfOpened, StateOfAutomaticRunning}},
		{name: "test controller 0x0040", s: 0x0040, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfFeedback, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x0080", s: 0x0080, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfDelay, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x0100", s: 0x0100, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfMainPowerFault, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x0200", s: 0x0200, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfBackupPowerFault, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x0400", s: 0x0400, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfBusFault, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x0800", s: 0x0800, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfClosed, StateOfManualRunning}},
		{name: "test controller 0x1000", s: 0x1000, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfConfigureChanged, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x2000", s: 0x2000, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfReset, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x4000", s: 0x4000, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test controller 0x8000", s: 0x8000, args: args{mapper: StdControllerStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfClosed, StateOfAutomaticRunning}},
		{name: "test equipment 0x0000", s: 0x0000, args: args{mapper: StdEquipmentStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfClosed}},
		{name: "test equipment 0x0001", s: 0x0001, args: args{mapper: StdEquipmentStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeNormal, StateOfClosed}},
		{name: "test equipment 0x0002", s: 0x0002, args: args{mapper: StdEquipmentStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfInAlarm, StateOfClosed}},
		{name: "test equipment 0x0004", s: 0x0004, args: args{mapper: StdEquipmentStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfInFault, StateOfClosed}},
		{name: "test equipment 0x0008", s: 0x0008, args: args{mapper: StdEquipmentStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfInBlock, StateOfClosed}},
		{name: "test equipment 0x0010", s: 0x0010, args: args{mapper: StdEquipmentStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfInSupervise, StateOfClosed}},
		{name: "test equipment 0x0020", s: 0x0020, args: args{mapper: StdEquipmentStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfOpened}},
		{name: "test equipment 0x0040", s: 0x0040, args: args{mapper: StdEquipmentStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfFeedback, StateOfClosed}},
		{name: "test equipment 0x0080", s: 0x0080, args: args{mapper: StdEquipmentStateFlagBitMapper}, wantInfo: StateInfo{StateOfRuntimeTest, StateOfDelay, StateOfClosed}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Info(tt.args.mapper)
			sort.Strings(got)
			sort.Strings(tt.wantInfo)
			assert.EqualValues(t, got, tt.wantInfo)
		})
	}
}

func TestOperationFlag_BitIndexBool(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		o    OperationFlag
		args args
		want bool
	}{
		{name: "test bit index 0 return false", o: 0xff, args: args{index: 0}, want: false},
		{name: "test bit index 1 return false", o: 0xfe, args: args{index: 1}, want: false},
		{name: "test bit index 2 return false", o: 0xfd, args: args{index: 2}, want: false},
		{name: "test bit index 3 return false", o: 0xfb, args: args{index: 3}, want: false},
		{name: "test bit index 4 return false", o: 0xf7, args: args{index: 4}, want: false},
		{name: "test bit index 5 return false", o: 0xef, args: args{index: 5}, want: false},
		{name: "test bit index 6 return false", o: 0xdf, args: args{index: 6}, want: false},
		{name: "test bit index 7 return false", o: 0xbf, args: args{index: 7}, want: false},
		{name: "test bit index 8 return false", o: 0x7f, args: args{index: 8}, want: false},
		{name: "test bit index 1 return true", o: 0x01, args: args{index: 1}, want: true},
		{name: "test bit index 2 return true", o: 0x02, args: args{index: 2}, want: true},
		{name: "test bit index 3 return true", o: 0x04, args: args{index: 3}, want: true},
		{name: "test bit index 4 return true", o: 0x08, args: args{index: 4}, want: true},
		{name: "test bit index 5 return true", o: 0x10, args: args{index: 5}, want: true},
		{name: "test bit index 6 return true", o: 0x20, args: args{index: 6}, want: true},
		{name: "test bit index 7 return true", o: 0x40, args: args{index: 7}, want: true},
		{name: "test bit index 8 return true", o: 0x80, args: args{index: 8}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, tt.o.BitIndexBool(tt.args.index), tt.want)
		})
	}
}

func TestOperationFlag_Info(t *testing.T) {
	type args struct {
		mapper OperationFlagBitMapper
	}
	tests := []struct {
		name     string
		o        OperationFlag
		args     args
		wantInfo OperationInfo
	}{
		{name: "test transmission 0x00", o: 0x00, args: args{mapper: StdTransmissionOperationFlagBitMapper}, wantInfo: nil},
		{name: "test transmission 0x02", o: 0x01, args: args{mapper: StdTransmissionOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfReset}},
		{name: "test transmission 0x02", o: 0x02, args: args{mapper: StdTransmissionOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfMute}},
		{name: "test transmission 0x04", o: 0x04, args: args{mapper: StdTransmissionOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfManualAlarm}},
		{name: "test transmission 0x08", o: 0x08, args: args{mapper: StdTransmissionOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfCancelAlarm}},
		{name: "test transmission 0x10", o: 0x10, args: args{mapper: StdTransmissionOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfSelfInspection}},
		{name: "test transmission 0x20", o: 0x20, args: args{mapper: StdTransmissionOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfInspectionReply}},
		{name: "test transmission 0x40", o: 0x40, args: args{mapper: StdTransmissionOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfTest}},
		{name: "test transmission 0x80", o: 0x80, args: args{mapper: StdTransmissionOperationFlagBitMapper}, wantInfo: nil},
		{name: "test controller 0x00", o: 0x00, args: args{mapper: StdControllerOperationFlagBitMapper}, wantInfo: nil},
		{name: "test controller 0x02", o: 0x01, args: args{mapper: StdControllerOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfReset}},
		{name: "test controller 0x02", o: 0x02, args: args{mapper: StdControllerOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfMute}},
		{name: "test controller 0x04", o: 0x04, args: args{mapper: StdControllerOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfManualAlarm}},
		{name: "test controller 0x08", o: 0x08, args: args{mapper: StdControllerOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfCancelAlarm}},
		{name: "test controller 0x10", o: 0x10, args: args{mapper: StdControllerOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfSelfInspection}},
		{name: "test controller 0x20", o: 0x20, args: args{mapper: StdControllerOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfConfirm}},
		{name: "test controller 0x40", o: 0x40, args: args{mapper: StdControllerOperationFlagBitMapper}, wantInfo: OperationInfo{OperationOfTest}},
		{name: "test controller 0x80", o: 0x80, args: args{mapper: StdControllerOperationFlagBitMapper}, wantInfo: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.o.Info(tt.args.mapper)
			sort.Strings(got)
			sort.Strings(tt.wantInfo)
			assert.EqualValues(t, got, tt.wantInfo)
		})
	}
}

func TestVersion_Major(t *testing.T) {
	tests := []struct {
		name string
		v    Version
		want int
	}{
		{name: "test 0x0101", v: 0x0101, want: 1},
		{name: "test 0x0501", v: 0x0501, want: 5},
		{name: "test 0x0001", v: 0x0001, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, tt.v.Major(), tt.want)
		})
	}
}

func TestVersion_Minor(t *testing.T) {
	tests := []struct {
		name string
		v    Version
		want int
	}{
		{name: "test 0x0101", v: 0x0101, want: 1},
		{name: "test 0x0504", v: 0x0504, want: 4},
		{name: "test 0x0400", v: 0x0400, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, tt.v.Minor(), tt.want)
		})
	}
}

func TestPacket_IsEmpty(t *testing.T) {
	for entry, want := range map[*Packet]bool{
		{}: true,
		{TransmissionStates: make([]TransmissionStateInfo, 5)}:         false,
		{TransmissionOperations: make([]TransmissionOperationInfo, 5)}: false,
		{TransmissionTimestamps: make([]TransmissionTimestamp, 5)}:     false,
		{TransmissionVersions: make([]TransmissionVersion, 5)}:         false,
		{TransmissionConfigures: make([]TransmissionConfigure, 5)}:     false,
		{ControllerStates: make([]ControllerStateInfo, 5)}:             false,
		{ControllerOperations: make([]ControllerOperationInfo, 5)}:     false,
		{ControllerParameters: make([]ControllerParameterInfo, 5)}:     false,
		{ControllerTimestamps: make([]ControllerTimestamp, 5)}:         false,
		{ControllerVersions: make([]ControllerVersion, 5)}:             false,
		{ControllerConfigures: make([]ControllerConfigure, 5)}:         false,
		{EquipmentStates: make([]EquipmentStateInfo, 5)}:               false,
		{EquipmentParameters: make([]EquipmentParameterInfo, 5)}:       false,
		{EquipmentConfigures: make([]EquipmentConfigure, 5)}:           false,
		{Others: make([]interface{}, 5)}:                               false,
	} {
		assert.Equal(t, entry.IsEmpty(), want)
	}
}

func TestVersion_String(t *testing.T) {
	tests := []struct {
		name string
		v    Version
		want string
	}{
		{name: "test for v1.1", v: Version(0x0101), want: "v1.1"},
		{name: "test for v1.3", v: Version(0x0103), want: "v1.3"},
		{name: "test for v2.4", v: Version(0x0204), want: "v2.4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.v.String(), tt.want)
		})
	}
}

func TestStateInfo_String(t *testing.T) {
	tests := []struct {
		name string
		s    StateInfo
		want string
	}{
		{name: "test for empty", s: StateInfo{}, want: ""},
		{name: "test for one", s: StateInfo{StateOfReset}, want: StateOfReset},
		{name: "test for more", s: StateInfo{StateOfRuntimeNormal, StateOfInAlarm}, want: StateOfRuntimeNormal + "," + StateOfInAlarm},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.s.String(), tt.want)
		})
	}
}

func TestOperationInfo_String(t *testing.T) {
	tests := []struct {
		name string
		o    OperationInfo
		want string
	}{
		{name: "test for empty", o: OperationInfo{}, want: ""},
		{name: "test for one", o: OperationInfo{OperationOfMute}, want: OperationOfMute},
		{name: "test for more", o: OperationInfo{OperationOfTest, OperationOfManualAlarm}, want: OperationOfTest + "," + OperationOfManualAlarm},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.o.String(), tt.want)
		})
	}
}

func TestStdStateName(t *testing.T) {
	for key, value := range stdStateNames {
		assert.Equal(t, StdStateName(key), value)
	}
	assert.Equal(t, "", StdStateName("not_exists"))
}

func TestStdOperationName(t *testing.T) {
	for key, value := range stdOperationNames {
		assert.Equal(t, StdOperationName(key), value)
	}
	assert.Equal(t, StdOperationName("not_exists"), "")
}

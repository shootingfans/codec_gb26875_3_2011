package constant

import (
	"strconv"
	"strings"
)

// Packet is packet of one request
type Packet struct {
	Header                 Header                      `json:"header,omitempty"`                  // request Header
	Action                 Action                      `json:"action,omitempty"`                  // request action
	ConnectionState        string                      `json:"connection_state,omitempty"`        // connect state, only set when connect or disconnect
	TransmissionStates     []TransmissionStateInfo     `json:"transmission_states,omitempty"`     // transmission states
	TransmissionOperations []TransmissionOperationInfo `json:"transmission_operations,omitempty"` // transmission operation information
	TransmissionTimestamps []TransmissionTimestamp     `json:"transmission_timestamps,omitempty"` // transmission time
	TransmissionVersions   []TransmissionVersion       `json:"transmission_versions,omitempty"`   // transmission version
	TransmissionConfigures []TransmissionConfigure     `json:"transmission_configures,omitempty"` // transmission configure
	ControllerStates       []ControllerStateInfo       `json:"controller_states,omitempty"`       // controller states
	ControllerOperations   []ControllerOperationInfo   `json:"controller_operations,omitempty"`   // controller operation information
	ControllerParameters   []ControllerParameterInfo   `json:"controller_parameters,omitempty"`   // controller parameters
	ControllerTimestamps   []ControllerTimestamp       `json:"controller_timestamps,omitempty"`   // controller timestamp
	ControllerVersions     []ControllerVersion         `json:"controller_versions,omitempty"`     // controller version
	ControllerConfigures   []ControllerConfigure       `json:"controller_configures,omitempty"`   // controller configure
	EquipmentStates        []EquipmentStateInfo        `json:"equipment_states,omitempty"`        // equipment states
	EquipmentParameters    []EquipmentParameterInfo    `json:"equipment_parameters,omitempty"`    // equipment parameters
	EquipmentConfigures    []EquipmentConfigure        `json:"equipment_configures,omitempty"`    // equipment configure
	Others                 []interface{}               `json:"others,omitempty"`                  // other custom information
	AppData                []byte                      `json:"app_data,omitempty"`                // app raw data
}

// IsEmpty is whether packet empty
func (p Packet) IsEmpty() bool {
	for _, call := range []func() int{
		func() int { return len(p.TransmissionStates) },
		func() int { return len(p.TransmissionOperations) },
		func() int { return len(p.TransmissionTimestamps) },
		func() int { return len(p.TransmissionVersions) },
		func() int { return len(p.TransmissionConfigures) },
		func() int { return len(p.ControllerStates) },
		func() int { return len(p.ControllerOperations) },
		func() int { return len(p.ControllerParameters) },
		func() int { return len(p.ControllerTimestamps) },
		func() int { return len(p.ControllerVersions) },
		func() int { return len(p.ControllerConfigures) },
		func() int { return len(p.EquipmentStates) },
		func() int { return len(p.EquipmentParameters) },
		func() int { return len(p.EquipmentConfigures) },
		func() int { return len(p.Others) },
	} {
		if call() > 0 {
			return false
		}
	}
	return true
}

type (
	// StateFlag is state flag
	StateFlag uint16
	// StateInfo is group state info
	StateInfo []string
	// OperationFlag is operation information flag
	OperationFlag uint8
	// OperationInfo is group operation information
	OperationInfo []string
	// BitValue is bit value selector 0 = false, 1 = true
	BitValue [2]string
	// Version is a version by major and minor, example: 1.1  3.2
	Version uint16
)

// StateFlagBitMapper is state mapper of bits
// The map key is bits position (1~16)
// The value is BitValue
type StateFlagBitMapper map[int]BitValue

// OperationFlagBitMapper is operation information mapper of bits
// The key is bits position (1~8)
// The value is BitValue
type OperationFlagBitMapper map[int]BitValue

// 定义了国标中状态
var (
	StateOfConnectionOpened       = "connection_opened"        // connection connect
	StateOfConnectionClosed       = "connection_closed"        // connection disconnect
	StateOfRuntimeNormal          = "runtime_normal"           // normal running
	StateOfRuntimeTest            = "runtime_test"             // test running
	StateOfInAlarm                = "alarming"                 // alarming
	StateOfInFault                = "fault"                    // faulting
	StateOfInBlock                = "blocking"                 // blocking
	StateOfInSupervise            = "supervise"                // supervise
	StateOfOpened                 = "opened"                   // opened
	StateOfClosed                 = "closed"                   // closed
	StateOfFeedback               = "feedback"                 // feedback
	StateOfDelay                  = "delay"                    // delay
	StateOfMainPowerFault         = "main_power_fault"         // main power fault
	StateOfBackupPowerFault       = "backup_power_fault"       // backup power fault
	StateOfBusFault               = "bus_fault"                // bus fault
	StateOfManualRunning          = "manual_running"           // manual running
	StateOfAutomaticRunning       = "automatic_running"        // automatic running
	StateOfConfigureChanged       = "configure_changed"        // configure changed
	StateOfReset                  = "reset"                    // reset
	StateOfCommunicationFault     = "communication_fault"      // communication fault
	StateOfMonitorConnectionFault = "monitor_connection_fault" // monitor connection fault
)

var stdStateNames = map[string]string{
	"connection_opened":        "connection_opened",
	"connection_closed":        "connection_closed",
	"runtime_normal":           "runtime_normal",
	"runtime_test":             "runtime_test",
	"alarming":                 "alarming",
	"fault":                    "fault",
	"blocking":                 "blocking",
	"supervise":                "supervise",
	"opened":                   "opened",
	"closed":                   "closed",
	"feedback":                 "feedback",
	"delay":                    "delay",
	"main_power_fault":         "main_power_fault",
	"backup_power_fault":       "backup_power_fault",
	"bus_fault":                "bus_fault",
	"manual_running":           "manual_running",
	"automatic_running":        "automatic_running",
	"configure_changed":        "configure_changed",
	"reset":                    "reset",
	"communication_fault":      "communication_fault",
	"monitor_connection_fault": "monitor_connection_fault",
}

// StdStateName is standard state name
func StdStateName(str string) string {
	if name, ok := stdStateNames[str]; ok {
		return name
	}
	return ""
}

// declare operation information
var (
	OperationOfReset           = "reset_action"           // reset operation
	OperationOfMute            = "mute_action"            // mute operation
	OperationOfManualAlarm     = "manual_alarm_action"    // manual alarm operation
	OperationOfCancelAlarm     = "cancel_alarm_action"    // cancel alarm operation
	OperationOfSelfInspection  = "self_inspection_action" // self inspection operation
	OperationOfInspectionReply = "inspection_reply"       // inspection reply operation
	OperationOfTest            = "test_action"            // test operation
	OperationOfConfirm         = "confirm_action"         // confirm operation
)

// StdOperationName is standard operation name
func StdOperationName(str string) string {
	if name, ok := stdOperationNames[str]; ok {
		return name
	}
	return ""
}

var stdOperationNames = map[string]string{
	"reset_action":           "reset_action",
	"mute_action":            "mute_action",
	"manual_alarm_action":    "manual_alarm_action",
	"cancel_alarm_action":    "cancel_alarm_action",
	"self_inspection_action": "self_inspection_action",
	"inspection_reply":       "inspection_reply",
	"test_action":            "test_action",
	"confirm_action":         "confirm_action",
}

// StdTransmissionStateFlagBitMapper is declare gb26875.3-2011 standard transmission states bits mapper
var StdTransmissionStateFlagBitMapper StateFlagBitMapper = map[int]BitValue{
	1: {StateOfRuntimeTest, StateOfRuntimeNormal},
	2: {"", StateOfInAlarm},
	3: {"", StateOfInFault},
	4: {"", StateOfMainPowerFault},
	5: {"", StateOfBackupPowerFault},
	6: {"", StateOfCommunicationFault},
	7: {"", StateOfMonitorConnectionFault},
}

// StdControllerStateFlagBitMapper is declare gb26875.3-2011 standard controller states bits mapper
var StdControllerStateFlagBitMapper StateFlagBitMapper = map[int]BitValue{
	1:  {StateOfRuntimeTest, StateOfRuntimeNormal},
	2:  {"", StateOfInAlarm},
	3:  {"", StateOfInFault},
	4:  {"", StateOfInBlock},
	5:  {"", StateOfInSupervise},
	6:  {StateOfClosed, StateOfOpened},
	7:  {"", StateOfFeedback},
	8:  {"", StateOfDelay},
	9:  {"", StateOfMainPowerFault},
	10: {"", StateOfBackupPowerFault},
	11: {"", StateOfBusFault},
	12: {StateOfAutomaticRunning, StateOfManualRunning},
	13: {"", StateOfConfigureChanged},
	14: {"", StateOfReset},
}

// StdEquipmentStateFlagBitMapper is declare gb26875.3-2011 standard equipment states bits mapper
var StdEquipmentStateFlagBitMapper StateFlagBitMapper = map[int]BitValue{
	1: {StateOfRuntimeTest, StateOfRuntimeNormal},
	2: {"", StateOfInAlarm},
	3: {"", StateOfInFault},
	4: {"", StateOfInBlock},
	5: {"", StateOfInSupervise},
	6: {StateOfClosed, StateOfOpened},
	7: {"", StateOfFeedback},
	8: {"", StateOfDelay},
	9: {"", StateOfMainPowerFault},
}

// StdTransmissionOperationFlagBitMapper is declare gb26875.3-2011 standard transmission operation bits mapper
var StdTransmissionOperationFlagBitMapper OperationFlagBitMapper = map[int]BitValue{
	1: {"", OperationOfReset},
	2: {"", OperationOfMute},
	3: {"", OperationOfManualAlarm},
	4: {"", OperationOfCancelAlarm},
	5: {"", OperationOfSelfInspection},
	6: {"", OperationOfInspectionReply},
	7: {"", OperationOfTest},
}

// StdControllerOperationFlagBitMapper is declare gb26875.3-2011 standard controller operation bits mapper
var StdControllerOperationFlagBitMapper OperationFlagBitMapper = map[int]BitValue{
	1: {"", OperationOfReset},
	2: {"", OperationOfMute},
	3: {"", OperationOfManualAlarm},
	4: {"", OperationOfCancelAlarm},
	5: {"", OperationOfSelfInspection},
	6: {"", OperationOfConfirm},
	7: {"", OperationOfTest},
}

var bits = []uint16{0xfffe, 0xfffd, 0xfffb, 0xfff7, 0xffef, 0xffdf, 0xffbf, 0xff7f, 0xfeff, 0xfdff, 0xfbff, 0xf7ff, 0xefff, 0xdfff, 0xbfff, 0x7fff}

func (bv BitValue) append(b bool, s *[]string) {
	v := bv[0]
	if b {
		v = bv[1]
	}
	if len(v) > 0 {
		*s = append(*s, v)
	}
}

// BitIndexBool is obtain bits index bool value
// index is between 1 to 16
func (s StateFlag) BitIndexBool(index int) bool {
	if index < 1 || index > 16 {
		return false
	}
	return bits[index-1]|uint16(s) == 0xffff
}

// Info is obtained state info by state bit mapper
func (s StateFlag) Info(mapper StateFlagBitMapper) (info StateInfo) {
	for index, args := range mapper {
		args.append(s.BitIndexBool(index), (*[]string)(&info))
	}
	return
}

// BitIndexBool is obtain bits index bool value
// index is between 1 to 8
func (o OperationFlag) BitIndexBool(index int) bool {
	if index < 1 || index > 8 {
		return false
	}
	return uint8(bits[index-1])|uint8(o) == 0xff
}

// Info obtained operation information by operation bit mapper
func (o OperationFlag) Info(mapper OperationFlagBitMapper) (info OperationInfo) {
	for index, args := range mapper {
		args.append(o.BitIndexBool(index), (*[]string)(&info))
	}
	return
}

// Major is major version
func (v Version) Major() int {
	return int(v >> 8)
}

// Minor is minor version
func (v Version) Minor() int {
	return int(uint8(v))
}

// String is implements fmt.Stringer
func (v Version) String() string {
	return "v" + strconv.Itoa(v.Major()) + "." + strconv.Itoa(v.Minor())
}

// String is implements fmt.Stringer
// Result is state join by ,
func (s StateInfo) String() string {
	return strings.Join(s, ",")
}

// String is implements fmt.Stringer
// Result is operation join by ,
func (o OperationInfo) String() string {
	return strings.Join(o, ",")
}

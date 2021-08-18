package constant

import "strconv"

// Action is packet action
type Action uint8

// Name is action name
func (ct Action) Name() string {
	if name, ok := ActionNames[ct]; ok {
		return name
	}
	return ""
}

// String is implements fmt.Stringer
func (ct Action) String() string {
	return "[" + strconv.Itoa(int(ct)) + "]" + ct.Name()
}

// define all actions
const (
	ControlAction  Action = iota + 1 // control action
	SendDataAction                   // send data action
	AckAction                        // ack action
	RequestAction                    // request action
	ResponseAction                   // response action
	RejectAction                     // reject action
)

// ActionNames is action names
var ActionNames = map[Action]string{
	ControlAction:  "Control",
	SendDataAction: "SendData",
	AckAction:      "Ack",
	RequestAction:  "Request",
	ResponseAction: "Response",
	RejectAction:   "Reject",
}

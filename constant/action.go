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
	ActionOfControl  Action = iota + 1 // control action
	ActionOfSendData                   // send data action
	ActionOfAck                        // ack action
	ActionOfRequest                    // request action
	ActionOfResponse                   // response action
	ActionOfReject                     // reject action
)

// ActionNames is action names
var ActionNames = map[Action]string{
	ActionOfControl:  "控制命令",
	ActionOfSendData: "发送数据",
	ActionOfAck:      "确认",
	ActionOfRequest:  "请求",
	ActionOfResponse: "响应",
	ActionOfReject:   "拒绝",
}

package constant

import "strconv"

type (
	// ControllerType is gb26875.3-2011 Controller type
	ControllerType uint
	// Controller is gb26875.3-2011 Controller
	Controller struct {
		Type ControllerType
		Addr int
	}
	// ControllerStateInfo is Controller States
	ControllerStateInfo struct {
		Ctrl      Controller
		Flag      StateFlag
		Timestamp int64
	}
	// ControllerParameterInfo is Controller parameter
	ControllerParameterInfo struct {
		Ctrl      Controller
		Info      ParameterInfo
		Timestamp int64
	}

	// ControllerOperationInfo is Controller operation information
	ControllerOperationInfo struct {
		Ctrl      Controller
		Flag      OperationFlag
		Operator  int
		Timestamp int64
	}

	// ControllerTimestamp is Controller time
	ControllerTimestamp struct {
		Ctrl      Controller
		Timestamp int64
	}

	// ControllerVersion is Controller version
	ControllerVersion struct {
		Ctrl    Controller
		Version Version
	}

	// ControllerConfigure is Controller configure
	ControllerConfigure struct {
		Ctrl      Controller
		Configure string
	}
)

// Name is Controller type name
func (ct ControllerType) Name() string {
	if name, ok := ControllerTypeNames[ct]; ok {
		return name
	}
	return ""
}

// String implements fmt.Stringer
func (ct ControllerType) String() string {
	return "[" + strconv.Itoa(int(ct)) + "]" + ct.Name()
}

const (
	ControllerTypeOfGeneral         ControllerType = iota // general controller
	ControllerTypeOfFireAlarmSystem                       // fire alarm controller
)

const (
	ControllerTypeOfFireLinkageSystem                                  ControllerType = 10 + iota // fire linkage controller
	ControllerTypeOfFireCockSystem                                                                // fire cock system
	ControllerTypeOfSprinklerSystem                                                               // sprinkler system
	ControllerTypeOfGasFireExtinguishingSystem                                                    // gas fire extinguishing system
	ControllerTypeOfPumpWaterSpraySystem                                                          // pump water spray system
	ControllerTypeOfPressureWaterSpraySystem                                                      // pressure water spray system
	ControllerTypeOfFoamFireExtinguishingSystem                                                   // foam fire extinguishing system
	ControllerTypeOfDryPowderFireExtinguishingSystem                                              // dry powder fire extinguishing system
	ControllerTypeOfSmokeExhaustSystem                                                            // smoke exhaust system
	ControllerTypeOfFireDoorAndShutterSystem                                                      // door and shutter system
	ControllerTypeOfFireLift                                                                      // fire lift
	ControllerTypeOfEmergencyBroadcast                                                            // emergency broadcast
	ControllerTypeOfFireEmergencyLightingAndEvacuationIndicationSystem                            // fire emergency lighting and evacuation indication system
	ControllerTypeOfFirePowerSupply                                                               // firepower supply
	ControllerTypeOfFireTelephone                                                                 // fire telephone

)

// ControllerTypeNames is declared all controller type names
var ControllerTypeNames = map[ControllerType]string{
	ControllerTypeOfGeneral:                                            "通用",
	ControllerTypeOfFireAlarmSystem:                                    "火灾报警系统",
	ControllerTypeOfFireLinkageSystem:                                  "消防联通控制器",
	ControllerTypeOfFireCockSystem:                                     "消火栓系统",
	ControllerTypeOfSprinklerSystem:                                    "自动喷水灭火系统",
	ControllerTypeOfGasFireExtinguishingSystem:                         "气体灭火系统",
	ControllerTypeOfPumpWaterSpraySystem:                               "水喷雾灭火系统(泵启动方式)",
	ControllerTypeOfPressureWaterSpraySystem:                           "水喷雾灭火系统(压力容器启动方式)",
	ControllerTypeOfFoamFireExtinguishingSystem:                        "泡沫灭火系统",
	ControllerTypeOfDryPowderFireExtinguishingSystem:                   "干粉灭火系统",
	ControllerTypeOfSmokeExhaustSystem:                                 "防烟排烟系统",
	ControllerTypeOfFireDoorAndShutterSystem:                           "防火门及卷帘系统",
	ControllerTypeOfFireLift:                                           "消防电梯",
	ControllerTypeOfEmergencyBroadcast:                                 "消防应急广播",
	ControllerTypeOfFireEmergencyLightingAndEvacuationIndicationSystem: "消防应急照明和疏散指示系统",
	ControllerTypeOfFirePowerSupply:                                    "消防电源",
	ControllerTypeOfFireTelephone:                                      "消防电话",
}

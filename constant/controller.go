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

// define controller types 0 ~ 1
const (
	GeneralControllerType         ControllerType = iota // general controller
	FireAlarmSystemControllerType                       // fire alarm controller
)

// define controller types 10 ~ 24
const (
	FireLinkageSystemControllerType                                  ControllerType = 10 + iota // fire linkage controller
	FireCockSystemControllerType                                                                // fire cock system
	SprinklerSystemControllerType                                                               // sprinkler system
	GasFireExtinguishingSystemControllerType                                                    // gas fire extinguishing system
	PumpWaterSpraySystemControllerType                                                          // pump water spray system
	PressureWaterSpraySystemControllerType                                                      // pressure water spray system
	FoamFireExtinguishingSystemControllerType                                                   // foam fire extinguishing system
	DryPowderFireExtinguishingSystemControllerType                                              // dry powder fire extinguishing system
	SmokeExhaustSystemControllerType                                                            // smoke exhaust system
	FireDoorAndShutterSystemControllerType                                                      // door and shutter system
	FireLiftControllerType                                                                      // fire lift
	EmergencyBroadcastControllerType                                                            // emergency broadcast
	FireEmergencyLightingAndEvacuationIndicationSystemControllerType                            // fire emergency lighting and evacuation indication system
	FirePowerSupplyControllerType                                                               // firepower supply
	FireTelephoneControllerType                                                                 // fire telephone
)

// ControllerTypeNames is declared all controller type names
var ControllerTypeNames = map[ControllerType]string{
	GeneralControllerType:                                            "General",
	FireAlarmSystemControllerType:                                    "FireAlarmSystem",
	FireLinkageSystemControllerType:                                  "FireLinkageSystem",
	FireCockSystemControllerType:                                     "FireCockSystem",
	SprinklerSystemControllerType:                                    "SprinklerSystem",
	GasFireExtinguishingSystemControllerType:                         "GasFireExtinguishingSystem",
	PumpWaterSpraySystemControllerType:                               "PumpWaterSpraySystem",
	PressureWaterSpraySystemControllerType:                           "PressureWaterSpraySystem",
	FoamFireExtinguishingSystemControllerType:                        "FoamFireExtinguishingSystem",
	DryPowderFireExtinguishingSystemControllerType:                   "DryPowderFireExtinguishingSystem",
	SmokeExhaustSystemControllerType:                                 "SmokeExhaustSystem",
	FireDoorAndShutterSystemControllerType:                           "FireDoorAndShutterSystem",
	FireLiftControllerType:                                           "FireLift",
	EmergencyBroadcastControllerType:                                 "EmergencyBroadcast",
	FireEmergencyLightingAndEvacuationIndicationSystemControllerType: "FireEmergencyLightingAndEvacuationIndicationSystem",
	FirePowerSupplyControllerType:                                    "FirePowerSupply",
	FireTelephoneControllerType:                                      "FireTelephone",
}

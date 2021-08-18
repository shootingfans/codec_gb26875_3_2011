package constant

import "strconv"

type (
	// EquipmentType is fire equipment type
	EquipmentType uint
	// EquipmentAddr is equipment address, in gb26875.3-2011 address is 4bits
	EquipmentAddr uint32
	// Equipment is fire equipment
	Equipment struct {
		Ctrl Controller
		Type EquipmentType
		Addr EquipmentAddr
	}
	// EquipmentStateInfo is fire state info
	EquipmentStateInfo struct {
		Equ         Equipment
		Flag        StateFlag
		Description string
		Timestamp   int64
	}
	// EquipmentParameterInfo is fire equipment parameter info
	EquipmentParameterInfo struct {
		Equ       Equipment
		Info      ParameterInfo
		Timestamp int64
	}

	// EquipmentConfigure is fire equipment configure
	EquipmentConfigure struct {
		Equ         Equipment
		Description string
	}
)

// Name is fire equipment type name
func (eq EquipmentType) Name() string {
	if name, ok := EquipmentTypeNames[eq]; ok {
		return name
	}
	return ""
}

// String implements fmt.Stringer
func (eq EquipmentType) String() string {
	return "[" + strconv.Itoa(int(eq)) + "]" + eq.Name()
}

// define 0 ~ 10 equipment type
const (
	GeneralEquipmentType          EquipmentType = iota // general
	FireAlarmControlEquipmentType                      // fire alarm control
)

// define other equipment type
const (
	FlammableGasDetectorEquipmentType              EquipmentType = 10 + iota // flammable gas detector
	PointTypeFlammableGasDetectorEquipmentType                               // point type flammable gas detector
	SelfContainedFlammableGasDetectorEquipmentType                           // self-contained flammable gas detector
	LineTypeFlammableGasDetectorEquipmentType                                // line type flammable gas detector
	_
	_
	ElectricalFireAlarmMonitorEquipmentType                     // electrical fire alarm monitor
	ResidualCurrentElectricalFireAlarmMonitorEquipmentType      // residual current electrical fire alarm monitor
	TemperatureMeasuringElectricalFireAlarmMonitorEquipmentType // temperature measuring electrical fire alarm monitor
	_
	_
	CircuitDetectorEquipmentType       // circuit detector
	FireDisplayPanelEquipmentType      // fire display panel
	ManualFireAlarmButtonEquipmentType // manual fire alarm button
	FireCockButtonEquipmentType        // fire cock button
	FireDetectorEquipmentType          // fire detector
	_
	_
	_
	_
	HeatFireDetectorEquipmentType           // heat fire detector
	PointTypeHeatFireDetectorEquipmentType  // point type heat fire detector
	SPointTypeHeatFireDetectorEquipmentType // S-point type heat fire detector
	RPointTypeHeatFireDetectorEquipmentType // R-point type heat fire detector
	LineTypeHeatFireDetectorEquipmentType   // line type heat fire detector
	SLineTypeHeatFireDetectorEquipmentType  // s line type heat fire detector
	RLineTypeHeatFireDetectorEquipmentType  // r line type heat fire detector
	OpticalHeatFireDetectorEquipmentType    // optical heat fire detector
	_
	_
	SmokeFireDetectorEquipmentType                          // smoke fire detector
	PointTypeIonSmokeFireDetectorEquipmentType              // point type ion smoke fire detector
	PointTypePhotoElectricitySmokeFireDetectorEquipmentType // point type photo electricity smoke fire detector
	PointTypeLightBeamSmokeFireDetectorEquipmentType        // point type light beam smoke fire detector
	AirBreathingTypeLightBeamSmokeFireDetectorEquipmentType // air breathing type light beam smoke fire detector
	_
	_
	_
	_
	_
	CombinationTypeFireDetectorEquipmentType                      // combination type fire detector
	CompoundSmokeAndTemperatureFireDetectorEquipmentType          // compound smoke and temperature fire detector
	CompoundPhotosensitiveAndTemperatureFireDetectorEquipmentType // compound photosensitive and temperature fire detector
	CompoundPhotosensitiveAndSmokeFireDetectorEquipmentType       // compound photosensitive and smoke fire detector
	_
	_
	_
	_
	_
	_
	_
	UltravioletFlameDetectorEquipmentType // equipment type of ultraviolet flame detector
	InfraredFlameDetectorEquipmentType    // equipment type of infrared flame detector
	_
	_
	_
	_
	_
	_
	OpticalFlameFireDetectorEquipmentType // optical flame fire detector
	_
	_
	_
	_
	GasDetectorEquipmentType // gas detector
	_
	_
	_
	ImageCameraModeFireDetectorEquipmentType // image camera mode fire detector
	AcousticFireDetectorEquipmentType        // acoustic fire detector
	_
	GasExtinguishingControllerEquipmentType            // gas extinguishing controller
	ElectricalFireControlDeviceEquipmentType           // electrical fire control device
	GraphicDisplayInFireControlRoomDeviceEquipmentType // graphic display in fire control room device
	ModuleEquipmentType                                // module
	InputModuleEquipmentType                           // input module
	OutputModuleEquipmentType                          // output module
	InputAndOutputModuleEquipmentType                  // input and output module
	RelayModuleEquipmentType                           // relay module
	_
	_
	FirePumpEquipmentType     // fire pump
	FireWaterBoxEquipmentType // fire water box
	_
	_
	SprayPumpEquipmentType          // spray pump
	WaterFlowIndicatorEquipmentType // water flow indicator
	SignalValveEquipmentType        // signal valve
	AlarmValveEquipmentType         // alarm valve
	PressureSwitchEquipmentType     // pressure switch
	_
	ValveActuatingDeviceEquipmentType       // valve actuating device
	FireDoorEquipmentType                   // fire door
	FireValveEquipmentType                  // fire valve
	VentilatingAirConditioningEquipmentType // ventilating air conditioning
	FoamConcentrateSupplyPumpEquipmentType  // foam concentrate supply pump
	PipeNetworkSolenoidValveEquipmentType   // pipe network solenoid valve
	_
	_
	_
	_
	SmokeControlAndExhaustFanEquipmentType // smoke control and exhaust fan
	_
	FireDamperEquipmentType                                   // fire damper
	AlwaysClosedAirOutletEquipmentType                        // always closed air outlet
	SmokeOutletEquipmentType                                  // smoke outlet
	ElectricallyControlledSmokeBlockVerticalWallEquipmentType // electrically controlled smoke block vertical wall
	FireShutterControllerEquipmentType                        // fire shutter controller
	FireDoorDetectorEquipmentType                             // fire door detector
	_
	_
	AlarmDeviceEquipmentType // alarm device
)

// EquipmentTypeNames is equipment type names
var EquipmentTypeNames = map[EquipmentType]string{
	GeneralEquipmentType:                                          "General",
	FireAlarmControlEquipmentType:                                 "FireAlarmControl",
	FlammableGasDetectorEquipmentType:                             "FlammableGasDetector",
	PointTypeFlammableGasDetectorEquipmentType:                    "PointTypeFlammableGasDetector",
	SelfContainedFlammableGasDetectorEquipmentType:                "SelfContainedFlammableGasDetector",
	LineTypeFlammableGasDetectorEquipmentType:                     "LineTypeFlammableGasDetector",
	ElectricalFireAlarmMonitorEquipmentType:                       "ElectricalFireAlarmMonitor",
	ResidualCurrentElectricalFireAlarmMonitorEquipmentType:        "ResidualCurrentElectricalFireAlarmMonitor",
	TemperatureMeasuringElectricalFireAlarmMonitorEquipmentType:   "TemperatureMeasuringElectricalFireAlarmMonitor",
	CircuitDetectorEquipmentType:                                  "CircuitDetector",
	FireDisplayPanelEquipmentType:                                 "FireDisplayPanel",
	ManualFireAlarmButtonEquipmentType:                            "ManualFireAlarmButton",
	FireCockButtonEquipmentType:                                   "FireCockButton",
	FireDetectorEquipmentType:                                     "FireDetector",
	HeatFireDetectorEquipmentType:                                 "HeatFireDetector",
	PointTypeHeatFireDetectorEquipmentType:                        "PointTypeHeatFireDetector",
	SPointTypeHeatFireDetectorEquipmentType:                       "SPointTypeHeatFireDetector",
	RPointTypeHeatFireDetectorEquipmentType:                       "RPointTypeHeatFireDetector",
	LineTypeHeatFireDetectorEquipmentType:                         "LineTypeHeatFireDetector",
	SLineTypeHeatFireDetectorEquipmentType:                        "SLineTypeHeatFireDetector",
	RLineTypeHeatFireDetectorEquipmentType:                        "RLineTypeHeatFireDetector",
	OpticalHeatFireDetectorEquipmentType:                          "OpticalHeatFireDetector",
	SmokeFireDetectorEquipmentType:                                "SmokeFireDetector",
	PointTypeIonSmokeFireDetectorEquipmentType:                    "PointTypeIonSmokeFireDetector",
	PointTypePhotoElectricitySmokeFireDetectorEquipmentType:       "PointTypePhotoElectricitySmokeFireDetector",
	PointTypeLightBeamSmokeFireDetectorEquipmentType:              "PointTypeLightBeamSmokeFireDetector",
	AirBreathingTypeLightBeamSmokeFireDetectorEquipmentType:       "AirBreathingTypeLightBeamSmokeFireDetector",
	CombinationTypeFireDetectorEquipmentType:                      "CombinationTypeFireDetector",
	CompoundSmokeAndTemperatureFireDetectorEquipmentType:          "CompoundSmokeAndTemperatureFireDetector",
	CompoundPhotosensitiveAndTemperatureFireDetectorEquipmentType: "CompoundPhotosensitiveAndTemperatureFireDetectorEquipmentType:",
	CompoundPhotosensitiveAndSmokeFireDetectorEquipmentType:       "CompoundPhotosensitiveAndSmokeFireDetector",
	UltravioletFlameDetectorEquipmentType:                         "UltravioletFlameDetector",
	InfraredFlameDetectorEquipmentType:                            "InfraredFlameDetector",
	OpticalFlameFireDetectorEquipmentType:                         "OpticalFlameFireDetector",
	GasDetectorEquipmentType:                                      "GasDetector",
	ImageCameraModeFireDetectorEquipmentType:                      "ImageCameraModeFireDetector",
	AcousticFireDetectorEquipmentType:                             "AcousticFireDetector",
	GasExtinguishingControllerEquipmentType:                       "GasExtinguishingController",
	ElectricalFireControlDeviceEquipmentType:                      "ElectricalFireControlDevice",
	GraphicDisplayInFireControlRoomDeviceEquipmentType:            "GraphicDisplayInFireControlRoomDevice",
	ModuleEquipmentType:                                           "Module",
	InputModuleEquipmentType:                                      "InputModule",
	OutputModuleEquipmentType:                                     "OutputModule",
	InputAndOutputModuleEquipmentType:                             "InputAndOutputModule",
	RelayModuleEquipmentType:                                      "RelayModule",
	FirePumpEquipmentType:                                         "FirePump",
	FireWaterBoxEquipmentType:                                     "FireWaterBox",
	SprayPumpEquipmentType:                                        "SprayPump",
	WaterFlowIndicatorEquipmentType:                               "WaterFlowIndicator",
	SignalValveEquipmentType:                                      "SignalValve",
	AlarmValveEquipmentType:                                       "AlarmValve",
	PressureSwitchEquipmentType:                                   "PressureSwitch",
	ValveActuatingDeviceEquipmentType:                             "ValveActuatingDevice",
	FireDoorEquipmentType:                                         "FireDoor",
	FireValveEquipmentType:                                        "FireValve",
	VentilatingAirConditioningEquipmentType:                       "VentilatingAirConditioning",
	FoamConcentrateSupplyPumpEquipmentType:                        "FoamConcentrateSupplyPump",
	PipeNetworkSolenoidValveEquipmentType:                         "PipeNetworkSolenoidValve",
	SmokeControlAndExhaustFanEquipmentType:                        "SmokeControlAndExhaustFan",
	FireDamperEquipmentType:                                       "FireDamper",
	AlwaysClosedAirOutletEquipmentType:                            "AlwaysClosedAirOutlet",
	SmokeOutletEquipmentType:                                      "SmokeOutlet",
	ElectricallyControlledSmokeBlockVerticalWallEquipmentType:     "ElectricallyControlledSmokeBlockVerticalWall",
	FireShutterControllerEquipmentType:                            "FireShutterController",
	FireDoorDetectorEquipmentType:                                 "FireDoorDetector",
	AlarmDeviceEquipmentType:                                      "AlarmDevice",
}

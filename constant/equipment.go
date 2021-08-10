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

const (
	EquipmentTypeOfGeneral          EquipmentType = iota // general
	EquipmentTypeOfFireAlarmControl                      // fire alarm control
)

const (
	EquipmentTypeOfFlammableGasDetector              EquipmentType = 10 + iota // flammable gas detector
	EquipmentTypeOfPointTypeFlammableGasDetector                               // point type flammable gas detector
	EquipmentTypeOfSelfContainedFlammableGasDetector                           // self-contained flammable gas detector
	EquipmentTypeOfLineTypeFlammableGasDetector                                // line type flammable gas detector
	_
	_
	EquipmentTypeOfElectricalFireAlarmMonitor                     // electrical fire alarm monitor
	EquipmentTypeOfResidualCurrentElectricalFireAlarmMonitor      // residual current electrical fire alarm monitor
	EquipmentTypeOfTemperatureMeasuringElectricalFireAlarmMonitor // temperature measuring electrical fire alarm monitor
	_
	_
	EquipmentTypeOfCircuitDetector       // circuit detector
	EquipmentTypeOfFireDisplayPanel      // fire display panel
	EquipmentTypeOfManualFireAlarmButton // manual fire alarm button
	EquipmentTypeOfFireCockButton        // fire cock button
	EquipmentTypeOfFireDetector          // fire detector
	_
	_
	_
	_
	EquipmentTypeOfHeatFireDetector           // heat fire detector
	EquipmentTypeOfPointTypeHeatFireDetector  // point type heat fire detector
	EquipmentTypeOfSPointTypeHeatFireDetector // S-point type heat fire detector
	EquipmentTypeOfRPointTypeHeatFireDetector // R-point type heat fire detector
	EquipmentTypeOfLineTypeHeatFireDetector   // line type heat fire detector
	EquipmentTypeOfSLineTypeHeatFireDetector  // s line type heat fire detector
	EquipmentTypeOfRLineTypeHeatFireDetector  // r line type heat fire detector
	EquipmentTypeOfOpticalHeatFireDetector    // optical heat fire detector
	_
	_
	EquipmentTypeOfSmokeFireDetector                          // smoke fire detector
	EquipmentTypeOfPointTypeIonSmokeFireDetector              // point type ion smoke fire detector
	EquipmentTypeOfPointTypePhotoElectricitySmokeFireDetector // point type photo electricity smoke fire detector
	EquipmentTypeOfPointTypeLightBeamSmokeFireDetector        // point type light beam smoke fire detector
	EquipmentTypeOfAirBreathingTypeLightBeamSmokeFireDetector // air breathing type light beam smoke fire detector
	_
	_
	_
	_
	_
	EquipmentTypeOfCombinationTypeFireDetector                      // combination type fire detector
	EquipmentTypeOfCompoundSmokeAndTemperatureFireDetector          // compound smoke and temperature fire detector
	EquipmentTypeOfCompoundPhotosensitiveAndTemperatureFireDetector // compound photosensitive and temperature fire detector
	EquipmentTypeOfCompoundPhotosensitiveAndSmokeFireDetector       // compound photosensitive and smoke fire detector
	_
	_
	_
	_
	_
	_
	_
	EquipmentTypeOfEquipmentTypeOfUltravioletFlameDetector // equipment type of ultraviolet flame detector
	EquipmentTypeOfEquipmentTypeOfInfraredFlameDetector    // equipment type of infrared flame detector
	_
	_
	_
	_
	_
	_
	EquipmentTypeOfOpticalFlameFireDetector // optical flame fire detector
	_
	_
	_
	_
	EquipmentTypeOfGasDetector // gas detector
	_
	_
	_
	EquipmentTypeOfImageCameraModeFireDetector // image camera mode fire detector
	EquipmentTypeOfAcousticFireDetector        // acoustic fire detector
	_
	EquipmentTypeOfGasExtinguishingController            // gas extinguishing controller
	EquipmentTypeOfElectricalFireControlDevice           // electrical fire control device
	EquipmentTypeOfGraphicDisplayInFireControlRoomDevice // graphic display in fire control room device
	EquipmentTypeOfModule                                // module
	EquipmentTypeOfInputModule                           // input module
	EquipmentTypeOfOutputModule                          // output module
	EquipmentTypeOfInputAndOutputModule                  // input and output module
	EquipmentTypeOfRelayModule                           // relay module
	_
	_
	EquipmentTypeOfFirePump     // fire pump
	EquipmentTypeOfFireWaterBox // fire water box
	_
	_
	EquipmentTypeOfSprayPump          // spray pump
	EquipmentTypeOfWaterFlowIndicator // water flow indicator
	EquipmentTypeOfSignalValve        // signal valve
	EquipmentTypeOfAlarmValve         // alarm valve
	EquipmentTypeOfPressureSwitch     // pressure switch
	_
	EquipmentTypeOfValveActuatingDevice       // valve actuating device
	EquipmentTypeOfFireDoor                   // fire door
	EquipmentTypeOfFireValve                  // fire valve
	EquipmentTypeOfVentilatingAirConditioning // ventilating air conditioning
	EquipmentTypeOfFoamConcentrateSupplyPump  // foam concentrate supply pump
	EquipmentTypeOfPipeNetworkSolenoidValve   // pipe network solenoid valve
	_
	_
	_
	_
	EquipmentTypeOfSmokeControlAndExhaustFan // smoke control and exhaust fan
	_
	EquipmentTypeOfFireDamper                                   // fire damper
	EquipmentTypeOfAlwaysClosedAirOutlet                        // always closed air outlet
	EquipmentTypeOfSmokeOutlet                                  // smoke outlet
	EquipmentTypeOfElectricallyControlledSmokeBlockVerticalWall // electrically controlled smoke block vertical wall
	EquipmentTypeOfFireShutterController                        // fire shutter controller
	EquipmentTypeOfFireDoorDetector                             // fire door detector
	_
	_
	EquipmentTypeOfAlarmDevice //报警装置
)

// EquipmentTypeNames 定义各部件名称
var EquipmentTypeNames = map[EquipmentType]string{
	EquipmentTypeOfGeneral:                                          "通用",
	EquipmentTypeOfFireAlarmControl:                                 "火灾报警控制器",
	EquipmentTypeOfFlammableGasDetector:                             "可燃气体探测器",
	EquipmentTypeOfPointTypeFlammableGasDetector:                    "点型可燃气体探测器",
	EquipmentTypeOfSelfContainedFlammableGasDetector:                "独立式可燃气体探测器",
	EquipmentTypeOfLineTypeFlammableGasDetector:                     "线型可燃气体探测器",
	EquipmentTypeOfElectricalFireAlarmMonitor:                       "电气火灾监控报警",
	EquipmentTypeOfResidualCurrentElectricalFireAlarmMonitor:        "剩余电流式电气火灾监控报警",
	EquipmentTypeOfTemperatureMeasuringElectricalFireAlarmMonitor:   "测温式电气火灾监控报警",
	EquipmentTypeOfCircuitDetector:                                  "回路探测",
	EquipmentTypeOfFireDisplayPanel:                                 "火灾显示盘",
	EquipmentTypeOfManualFireAlarmButton:                            "手动火灾报警按钮",
	EquipmentTypeOfFireCockButton:                                   "消火栓按钮",
	EquipmentTypeOfFireDetector:                                     "火灾探测器",
	EquipmentTypeOfHeatFireDetector:                                 "感温火灾探测器",
	EquipmentTypeOfPointTypeHeatFireDetector:                        "点型感温火灾探测器",
	EquipmentTypeOfSPointTypeHeatFireDetector:                       "点型感温火灾探测器(S型)",
	EquipmentTypeOfRPointTypeHeatFireDetector:                       "点型感温火灾探测器(R型)",
	EquipmentTypeOfLineTypeHeatFireDetector:                         "线型感温火灾探测器",
	EquipmentTypeOfSLineTypeHeatFireDetector:                        "线型感温火灾探测器(S型)",
	EquipmentTypeOfRLineTypeHeatFireDetector:                        "线型感温火灾探测器(R型)",
	EquipmentTypeOfOpticalHeatFireDetector:                          "光纤感温火灾探测器",
	EquipmentTypeOfSmokeFireDetector:                                "感烟火灾探测器",
	EquipmentTypeOfPointTypeIonSmokeFireDetector:                    "点型离子感烟火灾探测器",
	EquipmentTypeOfPointTypePhotoElectricitySmokeFireDetector:       "点型光电感烟火灾探测器",
	EquipmentTypeOfPointTypeLightBeamSmokeFireDetector:              "点型光束感烟火灾探测器",
	EquipmentTypeOfAirBreathingTypeLightBeamSmokeFireDetector:       "吸气式感烟火灾探测器",
	EquipmentTypeOfCombinationTypeFireDetector:                      "复合式火灾探测器",
	EquipmentTypeOfCompoundSmokeAndTemperatureFireDetector:          "复合式感烟感温火灾探测器",
	EquipmentTypeOfCompoundPhotosensitiveAndTemperatureFireDetector: "复合式感光感温火灾探测器",
	EquipmentTypeOfCompoundPhotosensitiveAndSmokeFireDetector:       "复合式感光感烟火灾探测器",
	EquipmentTypeOfEquipmentTypeOfUltravioletFlameDetector:          "紫外火焰探测器",
	EquipmentTypeOfEquipmentTypeOfInfraredFlameDetector:             "红外火焰探测器",
	EquipmentTypeOfOpticalFlameFireDetector:                         "感光火灾探测器",
	EquipmentTypeOfGasDetector:                                      "气体探测器",
	EquipmentTypeOfImageCameraModeFireDetector:                      "图像摄像方式火灾探测器",
	EquipmentTypeOfAcousticFireDetector:                             "感声火灾探测器",
	EquipmentTypeOfGasExtinguishingController:                       "气体灭火控制器",
	EquipmentTypeOfElectricalFireControlDevice:                      "消防电气控制装置",
	EquipmentTypeOfGraphicDisplayInFireControlRoomDevice:            "消防控制室图形显示装置",
	EquipmentTypeOfModule:                                           "模块",
	EquipmentTypeOfInputModule:                                      "输入模块",
	EquipmentTypeOfOutputModule:                                     "输出模块",
	EquipmentTypeOfInputAndOutputModule:                             "输入/输出模块",
	EquipmentTypeOfRelayModule:                                      "中继模块",
	EquipmentTypeOfFirePump:                                         "消防水泵",
	EquipmentTypeOfFireWaterBox:                                     "消防水箱",
	EquipmentTypeOfSprayPump:                                        "喷淋泵",
	EquipmentTypeOfWaterFlowIndicator:                               "水流指示器",
	EquipmentTypeOfSignalValve:                                      "信号阀",
	EquipmentTypeOfAlarmValve:                                       "报警阀",
	EquipmentTypeOfPressureSwitch:                                   "压力开关",
	EquipmentTypeOfValveActuatingDevice:                             "阀驱动装置",
	EquipmentTypeOfFireDoor:                                         "防火门",
	EquipmentTypeOfFireValve:                                        "防火阀",
	EquipmentTypeOfVentilatingAirConditioning:                       "通风空调",
	EquipmentTypeOfFoamConcentrateSupplyPump:                        "泡沫液泵",
	EquipmentTypeOfPipeNetworkSolenoidValve:                         "官网电磁阀",
	EquipmentTypeOfSmokeControlAndExhaustFan:                        "防烟排烟风机",
	EquipmentTypeOfFireDamper:                                       "排烟防火阀",
	EquipmentTypeOfAlwaysClosedAirOutlet:                            "常闭送风口",
	EquipmentTypeOfSmokeOutlet:                                      "排烟口",
	EquipmentTypeOfElectricallyControlledSmokeBlockVerticalWall:     "电控挡烟垂壁",
	EquipmentTypeOfFireShutterController:                            "防火卷帘控制器",
	EquipmentTypeOfFireDoorDetector:                                 "防火门监控器",
	EquipmentTypeOfAlarmDevice:                                      "报警装置",
}

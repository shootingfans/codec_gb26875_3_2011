package constant

import (
	"encoding/binary"
	"math"
	"strconv"

	"github.com/shootingfans/codec_gb26875_3_2011/utils"
)

type (
	// ParameterType is parameter type
	ParameterType uint
	// ParameterValue is parameter value
	ParameterValue int16
	// ParameterInfo is parameter information
	ParameterInfo struct {
		Type  ParameterType
		Value ParameterValue
	}
)

// NewParameterValue is created new parameter value
func NewParameterValue(b []byte) ParameterValue {
	return ParameterValue(utils.Bytes2Int16(b, binary.LittleEndian))
}

// Name is parameter type name
func (pt ParameterType) Name() string {
	if name, ok := ParameterTypeNames[pt]; ok {
		return name
	}
	return ""
}

// Unit is parameter unit description
func (pt ParameterType) Unit() string {
	if name, ok := ParameterTypeUnitNames[pt]; ok {
		return name
	}
	return ""
}

// String implements fmt.Stringer
func (pt ParameterType) String() string {
	return "[" + strconv.Itoa(int(pt)) + "]" + pt.Name()
}

// ToFloat is covert parameter value to float
func (pv ParameterValue) ToFloat(prec int) float64 {
	if prec == 0 {
		return float64(pv)
	}
	if prec < 0 {
		return float64(pv) * math.Pow10(-prec)
	}
	return float64(pv) / math.Pow10(prec)
}

// define parameter types
const (
	CounterParameterType          ParameterType = iota + 1 // event counter
	HeightParameterType                                    // height (0.01m)
	TemperatureParameterType                               // temperature (0.1℃)
	MPressureParameterType                                 // MPa (0.1MPa)
	KPressureParameterType                                 // kPa (0.1kpa)
	GasConcentrationParameterType                          // gas concentration (0.1%LEL)
	SecondsParameterType                                   // seconds (1s)
	VoltageParameterType                                   // voltage (0.1V)
	ElectricityParameterType                               // electricity (0.1A)
	FlowParameterType                                      // flow (0.1L/s)
	WindFlowParameterType                                  // wind flow (0.1m³/min)
	WindSpeedParameterType                                 // wind speed (m/s)
)

// ParameterTypeNames is parameter type name
var ParameterTypeNames = map[ParameterType]string{
	CounterParameterType:          "Counter",
	HeightParameterType:           "m (0.01m)",
	TemperatureParameterType:      "℃ (0.1℃)",
	MPressureParameterType:        "MPa (0.1MPa)",
	KPressureParameterType:        "kpa (0.1kpa)",
	GasConcentrationParameterType: "GasConcentration (0.1%LEL)",
	SecondsParameterType:          "Second (1s)",
	VoltageParameterType:          "Voltage (0.1V)",
	ElectricityParameterType:      "Electricity (0.1A)",
	FlowParameterType:             "Flow (0.1L/s)",
	WindFlowParameterType:         "WindFlow (0.1m³/min)",
	WindSpeedParameterType:        "WindSpeed (m/s)",
}

// ParameterTypeUnitNames is parameter unit description
var ParameterTypeUnitNames = map[ParameterType]string{
	CounterParameterType:          "次",
	HeightParameterType:           "m",
	TemperatureParameterType:      "℃",
	MPressureParameterType:        "MPa",
	KPressureParameterType:        "kPa",
	GasConcentrationParameterType: "%",
	SecondsParameterType:          "s",
	VoltageParameterType:          "V",
	ElectricityParameterType:      "A",
	FlowParameterType:             "L/s",
	WindFlowParameterType:         "m³/min",
	WindSpeedParameterType:        "m/s",
}

// ParameterInfoValue2StringHandler is convert func to value => string
var ParameterInfoValue2StringHandler = map[ParameterType]func(value ParameterValue) string{
	HeightParameterType:           parameterValue2Float(2),
	TemperatureParameterType:      parameterValue2Float(1),
	MPressureParameterType:        parameterValue2Float(1),
	KPressureParameterType:        parameterValue2Float(1),
	GasConcentrationParameterType: parameterValue2Float(1),
	VoltageParameterType:          parameterValue2Float(1),
	ElectricityParameterType:      parameterValue2Float(1),
	FlowParameterType:             parameterValue2Float(1),
	WindFlowParameterType:         parameterValue2Float(1),
}

func parameterValue2Float(prec int) func(value ParameterValue) string {
	return func(value ParameterValue) string {
		return strconv.FormatFloat(value.ToFloat(prec), 'g', 5, 64)
	}
}

// StringValue is value to string
func (pi ParameterInfo) StringValue() string {
	if handler, ok := ParameterInfoValue2StringHandler[pi.Type]; ok {
		return handler(pi.Value)
	}
	return strconv.Itoa(int(pi.Value))
}

// String implements fmt.Stringer
// format is value + ' ' + unit name
func (pi ParameterInfo) String() string {
	return pi.StringValue() + " " + pi.Type.Unit()
}

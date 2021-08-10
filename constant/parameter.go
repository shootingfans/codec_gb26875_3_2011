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

const (
	ParameterTypeOfCounter          ParameterType = iota + 1 // event counter
	ParameterTypeOfHeight                                    // height (0.01m)
	ParameterTypeOfTemperature                               // temperature (0.1℃)
	ParameterTypeOfMPressure                                 // MPa (0.1兆帕)
	ParameterTypeOfKPressure                                 // kPa (0.1千帕)
	ParameterTypeOfGasConcentration                          // gas concentration (0.1%LEL)
	ParameterTypeOfSeconds                                   // seconds (1s)
	ParameterTypeOfVoltage                                   // voltage (0.1V)
	ParameterTypeOfElectricity                               // electricity (0.1A)
	ParameterTypeOfFlow                                      // flow (0.1L/s)
	ParameterTypeOfWindFlow                                  // wind flow (0.1m³/min)
	ParameterTypeOfWindSpeed                                 // wind speed (m/s)
)

// ParameterTypeNames is parameter type name
var ParameterTypeNames = map[ParameterType]string{
	ParameterTypeOfCounter:          "事件计数",
	ParameterTypeOfHeight:           "高度 (0.01m)",
	ParameterTypeOfTemperature:      "温度(0.1℃)",
	ParameterTypeOfMPressure:        "压力(0.1兆帕)",
	ParameterTypeOfKPressure:        "压力(0.1千帕)",
	ParameterTypeOfGasConcentration: "气体浓度(0.1%LEL)",
	ParameterTypeOfSeconds:          "秒(1s)",
	ParameterTypeOfVoltage:          "电压(0.1V)",
	ParameterTypeOfElectricity:      "电流(0.1A)",
	ParameterTypeOfFlow:             "流量(0.1L/s)",
	ParameterTypeOfWindFlow:         "风量(0.1m³/min)",
	ParameterTypeOfWindSpeed:        "风速(m/s)",
}

// ParameterTypeUnitNames is parameter unit description
var ParameterTypeUnitNames = map[ParameterType]string{
	ParameterTypeOfCounter:          "次",
	ParameterTypeOfHeight:           "m",
	ParameterTypeOfTemperature:      "℃",
	ParameterTypeOfMPressure:        "MPa",
	ParameterTypeOfKPressure:        "kPa",
	ParameterTypeOfGasConcentration: "%",
	ParameterTypeOfSeconds:          "s",
	ParameterTypeOfVoltage:          "V",
	ParameterTypeOfElectricity:      "A",
	ParameterTypeOfFlow:             "L/s",
	ParameterTypeOfWindFlow:         "m³/min",
	ParameterTypeOfWindSpeed:        "m/s",
}

// ParameterInfoValue2StringHandler is convert func to value => string
var ParameterInfoValue2StringHandler = map[ParameterType]func(value ParameterValue) string{
	ParameterTypeOfHeight:           parameterValue2Float(2),
	ParameterTypeOfTemperature:      parameterValue2Float(1),
	ParameterTypeOfMPressure:        parameterValue2Float(1),
	ParameterTypeOfKPressure:        parameterValue2Float(1),
	ParameterTypeOfGasConcentration: parameterValue2Float(1),
	ParameterTypeOfVoltage:          parameterValue2Float(1),
	ParameterTypeOfElectricity:      parameterValue2Float(1),
	ParameterTypeOfFlow:             parameterValue2Float(1),
	ParameterTypeOfWindFlow:         parameterValue2Float(1),
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

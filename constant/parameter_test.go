package constant_test

import (
	"strconv"
	"testing"

	"github.com/shootingfans/codec_gb26875_3_2011/constant"

	"github.com/stretchr/testify/assert"
)

func TestParameterValue_ToFloat(t *testing.T) {
	assert.Equal(t, constant.ParameterValue(100).ToFloat(0), float64(100))
	assert.Equal(t, constant.ParameterValue(100).ToFloat(1), float64(10))
	assert.Equal(t, constant.ParameterValue(100).ToFloat(-1), float64(1000))
}

func TestParameterType_String(t *testing.T) {
	for k, v := range constant.ParameterTypeNames {
		assert.Equal(t, k.String(), "["+strconv.Itoa(int(k))+"]"+v)
		assert.Equal(t, k.Unit(), constant.ParameterTypeUnitNames[k])
	}
	assert.Equal(t, constant.ParameterType(200).Unit(), "")
	assert.Equal(t, constant.ParameterType(200).String(), "[200]")
}

func TestParameterInfo_String(t *testing.T) {
	testcases := map[constant.ParameterType]map[constant.ParameterValue]string{
		constant.CounterParameterType:          {1: "1 次", 10: "10 次", 35: "35 次"},
		constant.HeightParameterType:           {1: "0.01 m", 10: "0.1 m", 100: "1 m", 35: "0.35 m", 354: "3.54 m"},
		constant.TemperatureParameterType:      {1: "0.1 ℃", 10: "1 ℃", 100: "10 ℃", -131: "-13.1 ℃", 361: "36.1 ℃"},
		constant.MPressureParameterType:        {1: "0.1 MPa", 10: "1 MPa", 131: "13.1 MPa", 141: "14.1 MPa"},
		constant.KPressureParameterType:        {1: "0.1 kPa", 10: "1 kPa", 131: "13.1 kPa", 141: "14.1 kPa"},
		constant.GasConcentrationParameterType: {1: "0.1 %", 10: "1 %", 131: "13.1 %", 141: "14.1 %"},
		constant.SecondsParameterType:          {1: "1 s", 10: "10 s", 35: "35 s"},
		constant.VoltageParameterType:          {1: "0.1 V", 10: "1 V", 131: "13.1 V", 141: "14.1 V"},
		constant.ElectricityParameterType:      {1: "0.1 A", 10: "1 A", 131: "13.1 A", 141: "14.1 A"},
		constant.FlowParameterType:             {1: "0.1 L/s", 10: "1 L/s", 131: "13.1 L/s", 141: "14.1 L/s"},
		constant.WindFlowParameterType:         {1: "0.1 m³/min", 10: "1 m³/min", 131: "13.1 m³/min", 141: "14.1 m³/min"},
		constant.WindSpeedParameterType:        {1: "1 m/s", 10: "10 m/s", 35: "35 m/s"},
	}
	for tp, cases := range testcases {
		t.Run("test parameter type "+tp.String(), func(t *testing.T) {
			for va, str := range cases {
				assert.Equal(t, constant.ParameterInfo{Type: tp, Value: va}.String(), str)
			}
		})
	}
}

func TestNewParameterValue(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want constant.ParameterValue
	}{
		{name: "test 0xff00", args: args{[]byte{0xff, 0x00}}, want: constant.ParameterValue(255)},
		{name: "test 0xff01", args: args{[]byte{0xff, 0x01}}, want: constant.ParameterValue(511)},
		{name: "test 0xff8f", args: args{[]byte{0xff, 0x8f}}, want: constant.ParameterValue(-28673)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, constant.NewParameterValue(tt.args.b), tt.want)
		})
	}
}

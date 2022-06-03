package codec

import (
	"testing"
	"time"

	"github.com/shootingfans/codec_gb26875_3_2011/constant"
	"github.com/shootingfans/codec_gb26875_3_2011/utils"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	t.Run("test ack", func(t *testing.T) {
		b, err := Encode(&constant.Packet{
			Header: constant.Header{
				SerialId:  1,
				Version:   1,
				Timestamp: 1625187605,
				Source:    0x000102030405,
				Target:    0x060708090a0b,
			},
			Action: constant.AckAction,
		})
		assert.Nil(t, err)
		res := append(append([]byte{
			0x40, 0x40, 0x01, 0x00, 0x00, 0x01}, utils.Timestamp2Bytes(1625187605)...), 0x05, 0x04,
			0x03, 0x02, 0x01, 0x00, 0x0b, 0x0a, 0x09, 0x08, 0x07, 0x06, 0x00, 0x00, 0x03)
		res = append(res, uint8(utils.Sum(res[2:])), 0x23, 0x23)
		assert.EqualValues(t, b, res)
	})
	t.Run("test encode", func(t *testing.T) {
		b, err := Encode(&constant.Packet{
			Header: constant.Header{
				SerialId:  1,
				Version:   1,
				Timestamp: 1625187605,
				Source:    0x000102030405,
				Target:    0x060708090a0b,
			},
			Action:  constant.RequestAction,
			AppData: []byte{0x59, 0x01, 0x00},
		})
		assert.Nil(t, err)
		res := append(append([]byte{0x40, 0x40, 0x01, 0x00, 0x00, 0x01}, utils.Timestamp2Bytes(1625187605)...), 0x05, 0x04,
			0x03, 0x02, 0x01, 0x00, 0x0b, 0x0a, 0x09, 0x08, 0x07, 0x06, 0x03, 0x00, 0x04, 0x59,
			0x01, 0x00)
		res = append(res, uint8(utils.Sum(res[2:])), 0x23, 0x23)
		assert.EqualValues(t, b, res)
	})
}

func TestDecode(t *testing.T) {
	t.Run("test not enouth", func(t *testing.T) {
		p, n, err := Decode([]byte{0x40, 0x40, 0x00, 0x00, 0x01, 0x01, 0x18, 0x0d, 0x11, 0x16})
		assert.Nil(t, p)
		assert.Equal(t, n, 0)
		assert.ErrorIs(t, err, ErrPacketNotEnough)
	})
	t.Run("test invalid head flag", func(t *testing.T) {
		p, n, err := Decode([]byte{0x31, 0x40, 0x55, 0x31, 0x40, 0x55, 0x31, 0x40, 0x55, 0x31, 0x40, 0x55, 0x31, 0x40, 0x55, 0x31, 0x40, 0x55, 0x31, 0x40, 0x55, 0x31, 0x40, 0x55, 0x31, 0x40, 0x55, 0x31, 0x40, 0x55, 0x31, 0x40, 0x55, 0x31, 0x40, 0x55})
		assert.Nil(t, p)
		assert.Equal(t, n, 1)
		assert.ErrorIs(t, err, ErrPacketInvalid)
		assert.Contains(t, err.Error(), "head flag")
	})
	t.Run("test size not enough", func(t *testing.T) {
		p, n, err := Decode([]byte{0x40, 0x40, 0x00, 0x00, 0x01, 0x01, 0x18, 0x0d, 0x11, 0x16, 0x0a, 0x14, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x00, 0x02, 0x02, 0x01, 0x01, 0x03, 0x00, 0xd9, 0x00, 0x06, 0x00, 0x02, 0x00, 0xa3, 0xc1, 0xc7, 0xf8, 0xa3, 0xb1, 0xb2, 0xe3, 0xdf, 0xc8, 0xb2, 0xb8, 0xd7, 0xdf, 0xc0})
		assert.Nil(t, p)
		assert.Equal(t, n, 0)
		assert.ErrorIs(t, err, ErrPacketNotEnough)
	})
	t.Run("test invalid tail flag", func(t *testing.T) {
		p, n, err := Decode([]byte{0x40, 0x40, 0x00, 0x00, 0x01, 0x01, 0x18, 0x0d, 0x11, 0x16, 0x0a, 0x14, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x00, 0x02, 0x02, 0x01, 0x01, 0x03, 0x00, 0xd9, 0x00, 0x06, 0x00, 0x02, 0x00, 0xa3, 0xc1, 0xc7, 0xf8, 0xa3, 0xb1, 0xb2, 0xe3, 0xdf, 0xc8, 0xb2, 0xb8, 0xd7, 0xdf, 0xc0, 0xc8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x12, 0x13, 0x01, 0x08, 0x14, 0x53, 0x23, 0x13})
		assert.Nil(t, p)
		assert.Equal(t, n, 78)
		assert.ErrorIs(t, err, ErrPacketInvalid)
		assert.Contains(t, err.Error(), "tail flag")
	})
	t.Run("test check sum fail", func(t *testing.T) {
		p, n, err := Decode([]byte{0x40, 0x40, 0x00, 0x00, 0x01, 0x01, 0x18, 0x0d, 0x11, 0x16, 0x0a, 0x14, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x00, 0x02, 0x02, 0x01, 0x01, 0x03, 0x00, 0xd9, 0x00, 0x06, 0x00, 0x02, 0x00, 0xa3, 0xc1, 0xc7, 0xf8, 0xa3, 0xb1, 0xb2, 0xe3, 0xdf, 0xc8, 0xb2, 0xb8, 0xd7, 0xdf, 0xc0, 0xc8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x12, 0x13, 0x01, 0x08, 0x14, 0x43, 0x23, 0x23})
		assert.Nil(t, p)
		assert.Equal(t, n, 78)
		assert.ErrorIs(t, err, ErrPacketChecksumInvalid)
	})
	t.Run("test success", func(t *testing.T) {
		p, n, err := Decode([]byte{0x40, 0x40, 0x00, 0x00, 0x01, 0x01, 0x18, 0x0d, 0x11, 0x16, 0x0a, 0x14, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x30, 0x00, 0x02, 0x02, 0x01, 0x01, 0x03, 0x00, 0xd9, 0x00, 0x06, 0x00, 0x02, 0x00, 0xa3, 0xc1, 0xc7, 0xf8, 0xa3, 0xb1, 0xb2, 0xe3, 0xdf, 0xc8, 0xb2, 0xb8, 0xd7, 0xdf, 0xc0, 0xc8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x12, 0x13, 0x01, 0x08, 0x14, 0x68, 0x23, 0x23})
		assert.Equal(t, n, 78)
		assert.Nil(t, err)
		assert.EqualValues(t, p, &constant.Packet{
			Header: constant.Header{
				SerialId:  0,
				Version:   constant.Version(0x0101),
				Timestamp: utils.Bytes2Timestamp([]byte{0x18, 0x0d, 0x11, 0x16, 0x0a, 0x14}),
				Source:    0,
				Target:    0x010203040506,
			},
			Action: constant.SendDataAction,
			EquipmentStates: []constant.EquipmentStateInfo{
				{
					Equ: constant.Equipment{
						Ctrl: constant.Controller{
							Type: constant.FireAlarmSystemControllerType,
							Addr: 0x03,
						},
						Addr: 0x0600d9,
					},
					Flag:        0x02,
					Description: "Ａ区１层呷哺走廊",
					Timestamp:   utils.Bytes2Timestamp([]byte{0x30, 0x12, 0x13, 0x01, 0x08, 0x14}),
				},
			},
			AppData: []byte{0x02, 0x01, 0x01, 0x03, 0x00, 0xd9, 0x00, 0x06, 0x00, 0x02, 0x00, 0xa3, 0xc1, 0xc7, 0xf8, 0xa3, 0xb1, 0xb2, 0xe3, 0xdf, 0xc8, 0xb2, 0xb8, 0xd7, 0xdf, 0xc0, 0xc8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x12, 0x13, 0x01, 0x08, 0x14},
		})
	})
}

func TestDecodeAppData(t *testing.T) {
	t.Run("test not enough data", func(t *testing.T) {
		p := constant.Packet{AppData: []byte{0x01}, Action: constant.SendDataAction}
		DecodeAppData(&p)
		assert.EqualValues(t, p, constant.Packet{AppData: []byte{0x01}, Action: constant.SendDataAction})
	})
	t.Run("test none decoder type", func(t *testing.T) {
		p := constant.Packet{AppData: []byte{0xfe, 0x01}, Action: constant.SendDataAction}
		DecodeAppData(&p)
		assert.EqualValues(t, p, constant.Packet{AppData: []byte{0xfe, 0x01}, Action: constant.SendDataAction})
	})
	t.Run("test registry custom decoder", func(t *testing.T) {
		p := constant.Packet{AppData: []byte{0xfe, 0x01, 0x32}, Action: constant.SendDataAction}
		RegistryAppDecoder(constant.AppType(0xfe), AppDataDecoder(func(b []byte, packet *constant.Packet) {
			packet.Others = append(packet.Others, b[0])
		}))
		DecodeAppData(&p)
		assert.EqualValues(t, p.Others, []interface{}{uint8(0x32)})
	})
	t.Run("test decode upload system state", func(t *testing.T) {
		t.Run("test not enough", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x01, 0x01, 0x03, 0x04}}
			DecodeAppData(&p)
			assert.Empty(t, p.ControllerStates)
		})
		t.Run("test decode upload one system state", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x01, 0x01, 0x01, 0x02, 0x04, 0x02, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerStates, []constant.ControllerStateInfo{
				{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2}, Flag: constant.StateFlag(0x0204), Timestamp: utils.Bytes2Timestamp([]byte{0x04, 0x00, 0x09, 0x02, 0x07, 0x15})},
			})
		})
		t.Run("test decode upload more system state", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x01, 0x02, 0x01, 0x02, 0x04, 0x02, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15, 0x0d, 0x03, 0x05, 0x01, 0x05, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerStates, []constant.ControllerStateInfo{
				{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2}, Flag: constant.StateFlag(0x0204), Timestamp: utils.Bytes2Timestamp([]byte{0x04, 0x00, 0x09, 0x02, 0x07, 0x15})},
				{Ctrl: constant.Controller{Type: constant.GasFireExtinguishingSystemControllerType, Addr: 3}, Flag: constant.StateFlag(0x0105), Timestamp: utils.Bytes2Timestamp([]byte{0x05, 0x00, 0x09, 0x02, 0x07, 0x15})},
			})
		})
	})
	t.Run("test decode upload equipment state", func(t *testing.T) {
		t.Run("test not enough", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x02, 0x02, 0x03}}
			DecodeAppData(&p)
			assert.Empty(t, p.EquipmentStates)
		})
		t.Run("test decode upload one equipment state", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x02, 0x01, 0x01, 0x02, 0x79, 0x01, 0x02, 0x03, 0x04, 0x04, 0x02, 0xd5, 0xe2, 0xca, 0xc7, 0xd2, 0xbb, 0xb8, 0xf6, 0xd7, 0xaa, 0xc2, 0xeb, 0xb2, 0xe2, 0xca, 0xd4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.EquipmentStates, []constant.EquipmentStateInfo{
				{
					Equ: constant.Equipment{
						Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2},
						Type: constant.AlarmDeviceEquipmentType,
						Addr: constant.EquipmentAddr(0x04030201),
					}, Flag: constant.StateFlag(0x0204), Timestamp: utils.Bytes2Timestamp([]byte{0x04, 0x00, 0x09, 0x02, 0x07, 0x15}), Description: "这是一个转码测试",
				},
			})
		})
		t.Run("test decode upload more system state", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x02, 0x02,
				0x01, 0x02, 0x79, 0x01, 0x02, 0x03, 0x04, 0x04, 0x02, 0xd5, 0xe2, 0xca, 0xc7, 0xd2, 0xbb, 0xb8, 0xf6, 0xd7, 0xaa, 0xc2, 0xeb, 0xb2, 0xe2, 0xca, 0xd4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15,
				0x0d, 0x03, 0x79, 0x03, 0x04, 0x05, 0x06, 0x05, 0x01, 0xd5, 0xe2, 0xca, 0xc7, 0xd2, 0xbb, 0xb8, 0xf6, 0xd7, 0xaa, 0xc2, 0xeb, 0xb2, 0xe2, 0xca, 0xd4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.EquipmentStates, []constant.EquipmentStateInfo{
				{Equ: constant.Equipment{
					Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2},
					Type: constant.AlarmDeviceEquipmentType,
					Addr: constant.EquipmentAddr(0x04030201),
				}, Flag: constant.StateFlag(0x0204), Timestamp: utils.Bytes2Timestamp([]byte{0x04, 0x00, 0x09, 0x02, 0x07, 0x15}), Description: "这是一个转码测试"},
				{Equ: constant.Equipment{
					Ctrl: constant.Controller{Type: constant.GasFireExtinguishingSystemControllerType, Addr: 3},
					Type: constant.AlarmDeviceEquipmentType,
					Addr: constant.EquipmentAddr(0x06050403),
				}, Flag: constant.StateFlag(0x0105), Timestamp: utils.Bytes2Timestamp([]byte{0x05, 0x00, 0x09, 0x02, 0x07, 0x15}), Description: "这是一个转码测试"},
			})
		})
	})
	t.Run("test decode upload equipment parameter", func(t *testing.T) {
		t.Run("test not enough", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x03, 0x02, 0x03}}
			DecodeAppData(&p)
			assert.Empty(t, p.EquipmentParameters)
		})
		t.Run("test decode upload one equipment parameter", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x03, 0x01, 0x01, 0x02, 0x28, 0x01, 0x02, 0x03, 0x04, 0x03, 0x20, 0x03, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.EquipmentParameters, []constant.EquipmentParameterInfo{
				{
					Equ: constant.Equipment{
						Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2},
						Type: constant.SmokeFireDetectorEquipmentType,
						Addr: constant.EquipmentAddr(0x04030201),
					},
					Info: constant.ParameterInfo{
						Type:  constant.TemperatureParameterType,
						Value: constant.ParameterValue(0x0320),
					},
					Timestamp: utils.Bytes2Timestamp([]byte{0x04, 0x00, 0x09, 0x02, 0x07, 0x15}),
				},
			})
		})
		t.Run("test decode upload more equipment parameter", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x03, 0x02,
				0x01, 0x02, 0x28, 0x01, 0x02, 0x03, 0x04, 0x03, 0x20, 0x03, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15,
				0x0d, 0x03, 0x79, 0x03, 0x04, 0x05, 0x06, 0x02, 0x01, 0x50, 0x05, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.EquipmentParameters, []constant.EquipmentParameterInfo{
				{
					Equ: constant.Equipment{
						Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2},
						Type: constant.SmokeFireDetectorEquipmentType,
						Addr: constant.EquipmentAddr(0x04030201),
					},
					Info: constant.ParameterInfo{
						Type:  constant.TemperatureParameterType,
						Value: constant.ParameterValue(0x0320),
					},
					Timestamp: utils.Bytes2Timestamp([]byte{0x04, 0x00, 0x09, 0x02, 0x07, 0x15})},
				{
					Equ: constant.Equipment{
						Ctrl: constant.Controller{Type: constant.GasFireExtinguishingSystemControllerType, Addr: 3},
						Type: constant.AlarmDeviceEquipmentType,
						Addr: constant.EquipmentAddr(0x06050403),
					}, Info: constant.ParameterInfo{
						Type:  constant.HeightParameterType,
						Value: constant.ParameterValue(0x5001),
					},
					Timestamp: utils.Bytes2Timestamp([]byte{0x05, 0x00, 0x09, 0x02, 0x07, 0x15}),
				},
			})
		})
	})
	t.Run("test decode upload system operation", func(t *testing.T) {
		t.Run("test not enough", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x04, 0x02, 0x03}}
			DecodeAppData(&p)
			assert.Empty(t, p.ControllerOperations)
		})
		t.Run("test decode upload one equipment parameter", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x04, 0x01, 0x01, 0x02, 0x84, 0xfa, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerOperations, []constant.ControllerOperationInfo{
				{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2}, Flag: constant.OperationFlag(0x84), Operator: 0xfa, Timestamp: utils.Bytes2Timestamp([]byte{0x04, 0x00, 0x09, 0x02, 0x07, 0x15})},
			})
		})
		t.Run("test decode upload more equipment parameter", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x04, 0x02,
				0x01, 0x02, 0x84, 0xfa, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15,
				0x0d, 0x03, 0x73, 0xfb, 0x05, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerOperations, []constant.ControllerOperationInfo{
				{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2}, Flag: constant.OperationFlag(0x84), Operator: 0xfa, Timestamp: utils.Bytes2Timestamp([]byte{0x04, 0x00, 0x09, 0x02, 0x07, 0x15})},
				{Ctrl: constant.Controller{Type: constant.GasFireExtinguishingSystemControllerType, Addr: 3}, Flag: constant.OperationFlag(0x73), Operator: 0xfb, Timestamp: utils.Bytes2Timestamp([]byte{0x05, 0x00, 0x09, 0x02, 0x07, 0x15})},
			})
		})
	})
	t.Run("test decode upload system software version", func(t *testing.T) {
		t.Run("test not enough", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x05, 0x02}}
			DecodeAppData(&p)
			assert.Empty(t, p.ControllerVersions)
		})
		t.Run("test decode upload one system software version", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x05, 0x01, 0x01, 0x02, 0x03, 0x04}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerVersions, []constant.ControllerVersion{
				{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2}, Version: constant.Version(0x0304)},
			})
		})
		t.Run("test decode upload more system software version", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x05, 0x02, 0x01, 0x02, 0x03, 0x04, 0x0d, 0x03, 0x01, 0x01}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerVersions, []constant.ControllerVersion{
				{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2}, Version: constant.Version(0x0304)},
				{Ctrl: constant.Controller{Type: constant.GasFireExtinguishingSystemControllerType, Addr: 3}, Version: constant.Version(0x0101)},
			})
		})
	})
	t.Run("test decode upload system configure", func(t *testing.T) {
		t.Run("test not enough", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x06, 0x02}}
			DecodeAppData(&p)
			assert.Empty(t, p.ControllerConfigures)
		})
		t.Run("test invalid length", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x06, 0x01, 0x01, 0x01, 0x05, 0x01}}
			DecodeAppData(&p)
			assert.Empty(t, p.ControllerConfigures)
		})
		t.Run("test decode upload one system configure", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x06, 0x01, 0x01, 0x02, 0x1e, 0xd5, 0xe2, 0xca, 0xc7, 0xd2, 0xbb, 0xb8, 0xf6, 0xd7, 0xaa, 0xc2, 0xeb, 0xb2, 0xe2, 0xca, 0xd4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerConfigures, []constant.ControllerConfigure{
				{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2}, Configure: "这是一个转码测试"},
			})
		})
		t.Run("test decode upload more system configure", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x06, 0x02,
				0x01, 0x02, 0x1e, 0xd5, 0xe2, 0xca, 0xc7, 0xd2, 0xbb, 0xb8, 0xf6, 0xd7, 0xaa, 0xc2, 0xeb, 0xb2, 0xe2, 0xca, 0xd4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x0d, 0x03, 0x1e, 0xd5, 0xe2, 0xca, 0xc7, 0xd2, 0xbb, 0xb8, 0xf6, 0xd7, 0xaa, 0xc2, 0xeb, 0xb2, 0xe2, 0xca, 0xd4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerConfigures, []constant.ControllerConfigure{
				{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2}, Configure: "这是一个转码测试"},
				{Ctrl: constant.Controller{Type: constant.GasFireExtinguishingSystemControllerType, Addr: 3}, Configure: "这是一个转码测试"},
			})
		})
	})
	t.Run("test decode upload equipment configure", func(t *testing.T) {
		t.Run("test not enough", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x07, 0x02}}
			DecodeAppData(&p)
			assert.Empty(t, p.EquipmentConfigures)
		})
		t.Run("test invalid length", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x07, 0x01, 0x01, 0x01, 0x28, 0x01, 0x02, 0x03, 0x04, 0x01, 0x05, 0x01}}
			DecodeAppData(&p)
			assert.Empty(t, p.EquipmentConfigures)
		})
		t.Run("test decode upload one equipment configure", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x07, 0x01, 0x01, 0x02, 0x28, 0x01, 0x02, 0x03, 0x04, 0xd5, 0xe2, 0xca, 0xc7, 0xd2, 0xbb, 0xb8, 0xf6, 0xd7, 0xaa, 0xc2, 0xeb, 0xb2, 0xe2, 0xca, 0xd4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.EquipmentConfigures, []constant.EquipmentConfigure{
				{Equ: constant.Equipment{
					Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2},
					Type: constant.SmokeFireDetectorEquipmentType,
					Addr: constant.EquipmentAddr(0x04030201),
				}, Description: "这是一个转码测试"},
			})
		})
		t.Run("test decode upload more equipment configure", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x07, 0x02,
				0x01, 0x02, 0x28, 0x01, 0x02, 0x03, 0x04, 0xd5, 0xe2, 0xca, 0xc7, 0xd2, 0xbb, 0xb8, 0xf6, 0xd7, 0xaa, 0xc2, 0xeb, 0xb2, 0xe2, 0xca, 0xd4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x0d, 0x03, 0x55, 0x03, 0x04, 0x05, 0x06, 0xd5, 0xe2, 0xca, 0xc7, 0xd2, 0xbb, 0xb8, 0xf6, 0xd7, 0xaa, 0xc2, 0xeb, 0xb2, 0xe2, 0xca, 0xd4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.EquipmentConfigures, []constant.EquipmentConfigure{
				{Equ: constant.Equipment{
					Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2},
					Type: constant.SmokeFireDetectorEquipmentType,
					Addr: constant.EquipmentAddr(0x04030201),
				}, Description: "这是一个转码测试"},
				{Equ: constant.Equipment{
					Ctrl: constant.Controller{Type: constant.GasFireExtinguishingSystemControllerType, Addr: 3},
					Type: constant.InputModuleEquipmentType,
					Addr: constant.EquipmentAddr(0x06050403),
				}, Description: "这是一个转码测试"},
			})
		})
	})
	t.Run("test decode upload system time", func(t *testing.T) {
		t.Run("test not enough", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x08, 0x02}}
			DecodeAppData(&p)
			assert.Empty(t, p.ControllerTimestamps)
		})
		t.Run("test decode upload one system time", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x08, 0x01, 0x01, 0x02, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerTimestamps, []constant.ControllerTimestamp{
				{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 2}, Timestamp: utils.Bytes2Timestamp([]byte{0x04, 0x00, 0x09, 0x02, 0x07, 0x15})},
			})
		})
	})
	t.Run("test decode upload transmission state", func(t *testing.T) {
		t.Run("test not enough", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x15, 0x01}}
			DecodeAppData(&p)
			assert.Empty(t, p.TransmissionStates)
		})
		t.Run("test decode upload one transmission state", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x15, 0x01, 0x08, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.TransmissionStates, []constant.TransmissionStateInfo{
				{Flag: constant.StateFlag(0x08), Timestamp: utils.Bytes2Timestamp([]byte{0x04, 0x00, 0x09, 0x02, 0x07, 0x15})},
			})
		})
	})
	t.Run("test decode upload transmission operation", func(t *testing.T) {
		t.Run("test not enough", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x18, 0x01, 0x03}}
			DecodeAppData(&p)
			assert.Empty(t, p.TransmissionOperations)
		})
		t.Run("test decode upload one transmission operation", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x18, 0x01, 0x31, 0xfa, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.TransmissionOperations, []constant.TransmissionOperationInfo{
				{Flag: constant.OperationFlag(0x31), Operator: 0xfa, Timestamp: utils.Bytes2Timestamp([]byte{0x04, 0x00, 0x09, 0x02, 0x07, 0x15})},
			})
		})
		t.Run("test decode upload more transmission operation", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x18, 0x02,
				0x31, 0xfa, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15,
				0x73, 0xfb, 0x05, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.TransmissionOperations, []constant.TransmissionOperationInfo{
				{Flag: constant.OperationFlag(0x31), Operator: 0xfa, Timestamp: utils.Bytes2Timestamp([]byte{0x04, 0x00, 0x09, 0x02, 0x07, 0x15})},
				{Flag: constant.OperationFlag(0x73), Operator: 0xfb, Timestamp: utils.Bytes2Timestamp([]byte{0x05, 0x00, 0x09, 0x02, 0x07, 0x15})},
			})
		})
	})
	t.Run("test decode upload transmission software", func(t *testing.T) {
		t.Run("test not enough", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x19, 0x01}}
			DecodeAppData(&p)
			assert.Empty(t, p.TransmissionVersions)
		})
		t.Run("test decode upload one transmission software", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x19, 0x01, 0x03, 0x04}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.TransmissionVersions, []constant.TransmissionVersion{
				{Version: constant.Version(0x0304)},
			})
		})
	})
	t.Run("test decode upload transmission configure", func(t *testing.T) {
		t.Run("test not enough", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x1a, 0x01}}
			DecodeAppData(&p)
			assert.Empty(t, p.TransmissionConfigures)
		})
		t.Run("test invalid length", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x1a, 0x01, 0x1f, 0x01, 0x05, 0x01}}
			DecodeAppData(&p)
			assert.Empty(t, p.TransmissionConfigures)
		})
		t.Run("test decode upload one transmission configure", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x1a, 0x01, 0x1e, 0xd5, 0xe2, 0xca, 0xc7, 0xd2, 0xbb, 0xb8, 0xf6, 0xd7, 0xaa, 0xc2, 0xeb, 0xb2, 0xe2, 0xca, 0xd4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.TransmissionConfigures, []constant.TransmissionConfigure{
				{Configure: "这是一个转码测试"},
			})
		})
	})
	t.Run("test decode upload transmission time", func(t *testing.T) {
		t.Run("test not enough", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x1c, 0x01}}
			DecodeAppData(&p)
			assert.Empty(t, p.TransmissionTimestamps)
		})
		t.Run("test decode upload one transmission time", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x1c, 0x01, 0x05, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.TransmissionTimestamps, []constant.TransmissionTimestamp{{Timestamp: utils.Bytes2Timestamp([]byte{0x05, 0x00, 0x09, 0x02, 0x07, 0x15})}})
		})
	})
}

func TestNewQuerySystemStateAppData(t *testing.T) {
	type args struct {
		controllers []constant.Controller
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test empty", args: args{}, want: nil},
		{name: "test one controller", args: args{controllers: []constant.Controller{{Type: constant.FireAlarmSystemControllerType, Addr: 1}}}, want: []byte{0x3d, 0x01, 0x01, 0x01}},
		{name: "test more controllers", args: args{controllers: []constant.Controller{{Type: constant.FireAlarmSystemControllerType, Addr: 1}, {Type: constant.GasFireExtinguishingSystemControllerType, Addr: 2}}}, want: []byte{0x3d, 0x02, 0x01, 0x01, 0x0d, 0x02}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewQuerySystemStateAppData(tt.args.controllers...), tt.want)
		})
	}
}

func TestNewQueryEquipmentStateAppData(t *testing.T) {
	type args struct {
		equipments []constant.Equipment
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test empty", args: args{}, want: nil},
		{name: "test one equipment", args: args{equipments: []constant.Equipment{
			{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 1}, Type: constant.AlarmDeviceEquipmentType, Addr: 0x01110203},
		}}, want: []byte{0x3e, 0x01, 0x01, 0x01, 0x03, 0x02, 0x11, 0x01}},
		{name: "test more equipments", args: args{equipments: []constant.Equipment{
			{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 1}, Type: constant.AlarmDeviceEquipmentType, Addr: 0x01110203},
			{Ctrl: constant.Controller{Type: constant.GasFireExtinguishingSystemControllerType, Addr: 3}, Type: constant.SmokeFireDetectorEquipmentType, Addr: 0x12345678},
		}}, want: []byte{0x3e, 0x02, 0x01, 0x01, 0x03, 0x02, 0x11, 0x01, 0x0d, 0x03, 0x78, 0x56, 0x34, 0x12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewQueryEquipmentStateAppData(tt.args.equipments...), tt.want)
		})
	}
}

func TestNewQueryEquipmentParameterAppData(t *testing.T) {
	type args struct {
		equipments []constant.Equipment
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test empty", args: args{}, want: nil},
		{name: "test one equipment", args: args{equipments: []constant.Equipment{
			{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 1}, Type: constant.AlarmDeviceEquipmentType, Addr: 0x01110203},
		}}, want: []byte{0x3f, 0x01, 0x01, 0x01, 0x03, 0x02, 0x11, 0x01}},
		{name: "test more equipments", args: args{equipments: []constant.Equipment{
			{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 1}, Type: constant.AlarmDeviceEquipmentType, Addr: 0x01110203},
			{Ctrl: constant.Controller{Type: constant.GasFireExtinguishingSystemControllerType, Addr: 3}, Type: constant.SmokeFireDetectorEquipmentType, Addr: 0x12345678},
		}}, want: []byte{0x3f, 0x02, 0x01, 0x01, 0x03, 0x02, 0x11, 0x01, 0x0d, 0x03, 0x78, 0x56, 0x34, 0x12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewQueryEquipmentParameterAppData(tt.args.equipments...), tt.want)
		})
	}
}

func TestNewQuerySystemOperatingInformationAppData(t *testing.T) {
	type args struct {
		controller constant.Controller
		total      int
		startTime  time.Time
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test total 10 time start 2021-07-02 00:00:00", args: args{controller: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 0x01}, total: 10, startTime: time.Unix(1625155200, 0)},
			want: append([]byte{0x40, 0x01, 0x01, 0x01, 0x0a}, utils.Timestamp2Bytes(1625155200)...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewQuerySystemOperatingInformationAppData(tt.args.controller, tt.args.total, tt.args.startTime), tt.want)
		})
	}
}

func TestNewQuerySystemSoftwareVersionAppData(t *testing.T) {
	type args struct {
		controller constant.Controller
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test 1", args: args{controller: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 0x02}}, want: []byte{0x41, 0x01, 0x01, 0x02}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewQuerySystemSoftwareVersionAppData(tt.args.controller), tt.want)
		})
	}
}

func TestNewQuerySystemConfigureAppData(t *testing.T) {
	type args struct {
		controllers []constant.Controller
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test empty", args: args{}, want: nil},
		{name: "test one", args: args{controllers: []constant.Controller{
			{Type: constant.FireAlarmSystemControllerType, Addr: 0x03},
		}}, want: []byte{0x42, 0x01, 0x01, 0x03}},
		{name: "test more", args: args{controllers: []constant.Controller{
			{Type: constant.FireAlarmSystemControllerType, Addr: 0x03},
			{Type: constant.GasFireExtinguishingSystemControllerType, Addr: 0x02},
		}}, want: []byte{0x42, 0x02, 0x01, 0x03, 0x0d, 0x02}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewQuerySystemConfigureAppData(tt.args.controllers...), tt.want)
		})
	}
}

func TestNewQueryEquipmentConfigureAppData(t *testing.T) {
	type args struct {
		equipments []constant.Equipment
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test empty", args: args{}, want: nil},
		{name: "test one", args: args{equipments: []constant.Equipment{
			{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 0x01}, Addr: constant.EquipmentAddr(0x12345678), Type: constant.AlarmDeviceEquipmentType},
		}}, want: []byte{0x43, 0x01, 0x01, 0x01, 0x78, 0x56, 0x34, 0x12}},
		{name: "test more", args: args{equipments: []constant.Equipment{
			{Ctrl: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 0x01}, Addr: constant.EquipmentAddr(0x12345678), Type: constant.AlarmDeviceEquipmentType},
			{Ctrl: constant.Controller{Type: constant.GasFireExtinguishingSystemControllerType, Addr: 0x03}, Addr: constant.EquipmentAddr(0x01020304), Type: constant.GasDetectorEquipmentType},
		}}, want: []byte{0x43, 0x02, 0x01, 0x01, 0x78, 0x56, 0x34, 0x12, 0x0d, 0x03, 0x04, 0x03, 0x02, 0x01}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewQueryEquipmentConfigureAppData(tt.args.equipments...), tt.want)
		})
	}
}

func TestNewQuerySystemTimeAppData(t *testing.T) {
	type args struct {
		controller constant.Controller
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test", args: args{controller: constant.Controller{Type: constant.FireAlarmSystemControllerType, Addr: 0x03}}, want: []byte{0x44, 0x01, 0x01, 0x03}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewQuerySystemTimeAppData(tt.args.controller), tt.want)
		})
	}
}

func TestNewQueryTransmissionStateAppData(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{name: "test", want: []byte{0x51, 0x01, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewQueryTransmissionStateAppData(), tt.want)
		})
	}
}

func TestNewQueryTransmissionOperatingInformationAppData(t *testing.T) {
	type args struct {
		total     int
		startTime time.Time
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test total 16 startTime 2021-07-02 09:00:04", args: args{total: 16, startTime: time.Unix(1625187604, 0)}, want: append([]byte{0x54, 0x01, 0x10}, utils.Timestamp2Bytes(1625187604)...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewQueryTransmissionOperatingInformationAppData(tt.args.total, tt.args.startTime), tt.want)
		})
	}
}

func TestNewQueryTransmissionSoftwareVersionAppData(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{name: "test", want: []byte{0x55, 0x01, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewQueryTransmissionSoftwareVersionAppData(), tt.want)
		})
	}
}

func TestNewQueryTransmissionConfigureAppData(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{name: "test", want: []byte{0x56, 0x01, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewQueryTransmissionConfigureAppData(), tt.want)
		})
	}
}

func TestNewQueryTransmissionTimeAppData(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{name: "test", want: []byte{0x58, 0x01, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewQueryTransmissionTimeAppData(), tt.want)
		})
	}
}

func TestNewInitializeTransmissionAppData(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{name: "test", want: []byte{0x59, 0x01, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewInitializeTransmissionAppData(), tt.want)
		})
	}
}

func TestNewSyncTransmissionTimeAppData(t *testing.T) {
	type args struct {
		syncTime time.Time
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test sync to 2021-07-02 09:00:04", args: args{syncTime: time.Unix(1625187604, 0)}, want: append([]byte{0x5a, 0x01}, utils.Timestamp2Bytes(1625187604)...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewSyncTransmissionTimeAppData(tt.args.syncTime), tt.want)
		})
	}
}

func TestNewInspectSentriesAppData(t *testing.T) {
	type args struct {
		timeoutMinute int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test timeout 0 minute", args: args{timeoutMinute: 0}, want: []byte{0x5b, 0x01, 0x00}},
		{name: "test timeout 10 minute", args: args{timeoutMinute: 10}, want: []byte{0x5b, 0x01, 0x0a}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.EqualValues(t, NewInspectSentriesAppData(tt.args.timeoutMinute), tt.want)
		})
	}
}

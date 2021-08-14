package codec

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shootingfans/codec_gb26875_3_2011/constant"
)

func TestEncode(t *testing.T) {

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
				Timestamp: 1603358004,
				Source:    0,
				Target:    0x010203040506,
			},
			Action: constant.ActionOfSendData,
			EquipmentStates: []constant.EquipmentStateInfo{
				{
					Equ: constant.Equipment{
						Ctrl: constant.Controller{
							Type: constant.ControllerTypeOfFireAlarmSystem,
							Addr: 0x03,
						},
						Addr: 0x0600d9,
					},
					Flag:        0x02,
					Description: "Ａ区１层呷哺走廊",
					Timestamp:   1596280728,
				},
			},
			AppData: []byte{0x02, 0x01, 0x01, 0x03, 0x00, 0xd9, 0x00, 0x06, 0x00, 0x02, 0x00, 0xa3, 0xc1, 0xc7, 0xf8, 0xa3, 0xb1, 0xb2, 0xe3, 0xdf, 0xc8, 0xb2, 0xb8, 0xd7, 0xdf, 0xc0, 0xc8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x12, 0x13, 0x01, 0x08, 0x14},
		})
	})
}

func TestDecodeAppData(t *testing.T) {
	t.Run("test not enough data", func(t *testing.T) {
		p := constant.Packet{AppData: []byte{0x01}, Action: constant.ActionOfSendData}
		DecodeAppData(&p)
		assert.EqualValues(t, p, constant.Packet{AppData: []byte{0x01}, Action: constant.ActionOfSendData})
	})
	t.Run("test none decoder type", func(t *testing.T) {
		p := constant.Packet{AppData: []byte{0xfe, 0x01}, Action: constant.ActionOfSendData}
		DecodeAppData(&p)
		assert.EqualValues(t, p, constant.Packet{AppData: []byte{0xfe, 0x01}, Action: constant.ActionOfSendData})
	})
	t.Run("test registry custom decoder", func(t *testing.T) {
		p := constant.Packet{AppData: []byte{0xfe, 0x01, 0x32}, Action: constant.ActionOfSendData}
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
				{Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2}, Flag: constant.StateFlag(0x0204), Timestamp: 1625187604},
			})
		})
		t.Run("test decode upload more system state", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x01, 0x02, 0x01, 0x02, 0x04, 0x02, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15, 0x0d, 0x03, 0x05, 0x01, 0x05, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerStates, []constant.ControllerStateInfo{
				{Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2}, Flag: constant.StateFlag(0x0204), Timestamp: 1625187604},
				{Ctrl: constant.Controller{Type: constant.ControllerTypeOfGasFireExtinguishingSystem, Addr: 3}, Flag: constant.StateFlag(0x0105), Timestamp: 1625187605},
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
			p := constant.Packet{AppData: []byte{0x02, 0x01, 0x01, 0x02, 0x79, 0x01, 0x02, 0x03, 0x04, 0x04, 0x02, 0x37, 0xc2, 0xa5, 0x30, 0x32, 0xca, 0xd2, 0xbb, 0xe1, 0xd2, 0xe9, 0xca, 0xd2, 0xd1, 0xcc, 0xb8, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.EquipmentStates, []constant.EquipmentStateInfo{
				{
					Equ: constant.Equipment{
						Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2},
						Type: constant.EquipmentTypeOfAlarmDevice,
						Addr: constant.EquipmentAddr(0x04030201),
					}, Flag: constant.StateFlag(0x0204), Timestamp: 1625187604, Description: "7楼02室会议室烟感",
				},
			})
		})
		t.Run("test decode upload more system state", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x02, 0x02,
				0x01, 0x02, 0x79, 0x01, 0x02, 0x03, 0x04, 0x04, 0x02, 0x37, 0xc2, 0xa5, 0x30, 0x32, 0xca, 0xd2, 0xbb, 0xe1, 0xd2, 0xe9, 0xca, 0xd2, 0xd1, 0xcc, 0xb8, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15,
				0x0d, 0x03, 0x79, 0x03, 0x04, 0x05, 0x06, 0x05, 0x01, 0x37, 0xc2, 0xa5, 0x30, 0x32, 0xca, 0xd2, 0xbb, 0xe1, 0xd2, 0xe9, 0xca, 0xd2, 0xd1, 0xcc, 0xb8, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.EquipmentStates, []constant.EquipmentStateInfo{
				{Equ: constant.Equipment{
					Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2},
					Type: constant.EquipmentTypeOfAlarmDevice,
					Addr: constant.EquipmentAddr(0x04030201),
				}, Flag: constant.StateFlag(0x0204), Timestamp: 1625187604, Description: "7楼02室会议室烟感"},
				{Equ: constant.Equipment{
					Ctrl: constant.Controller{Type: constant.ControllerTypeOfGasFireExtinguishingSystem, Addr: 3},
					Type: constant.EquipmentTypeOfAlarmDevice,
					Addr: constant.EquipmentAddr(0x06050403),
				}, Flag: constant.StateFlag(0x0105), Timestamp: 1625187605, Description: "7楼02室会议室烟感"},
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
						Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2},
						Type: constant.EquipmentTypeOfSmokeFireDetector,
						Addr: constant.EquipmentAddr(0x04030201),
					},
					Info: constant.ParameterInfo{
						Type:  constant.ParameterTypeOfTemperature,
						Value: constant.ParameterValue(0x0320),
					},
					Timestamp: 1625187604,
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
						Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2},
						Type: constant.EquipmentTypeOfSmokeFireDetector,
						Addr: constant.EquipmentAddr(0x04030201),
					},
					Info: constant.ParameterInfo{
						Type:  constant.ParameterTypeOfTemperature,
						Value: constant.ParameterValue(0x0320),
					},
					Timestamp: 1625187604},
				{
					Equ: constant.Equipment{
						Ctrl: constant.Controller{Type: constant.ControllerTypeOfGasFireExtinguishingSystem, Addr: 3},
						Type: constant.EquipmentTypeOfAlarmDevice,
						Addr: constant.EquipmentAddr(0x06050403),
					}, Info: constant.ParameterInfo{
						Type:  constant.ParameterTypeOfHeight,
						Value: constant.ParameterValue(0x5001),
					},
					Timestamp: 1625187605,
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
				{Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2}, Flag: constant.OperationFlag(0x84), Operator: 0xfa, Timestamp: 1625187604},
			})
		})
		t.Run("test decode upload more equipment parameter", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x04, 0x02,
				0x01, 0x02, 0x84, 0xfa, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15,
				0x0d, 0x03, 0x73, 0xfb, 0x05, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerOperations, []constant.ControllerOperationInfo{
				{Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2}, Flag: constant.OperationFlag(0x84), Operator: 0xfa, Timestamp: 1625187604},
				{Ctrl: constant.Controller{Type: constant.ControllerTypeOfGasFireExtinguishingSystem, Addr: 3}, Flag: constant.OperationFlag(0x73), Operator: 0xfb, Timestamp: 1625187605},
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
				{Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2}, Version: constant.Version(0x0304)},
			})
		})
		t.Run("test decode upload more system software version", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x05, 0x02, 0x01, 0x02, 0x03, 0x04, 0x0d, 0x03, 0x01, 0x01}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerVersions, []constant.ControllerVersion{
				{Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2}, Version: constant.Version(0x0304)},
				{Ctrl: constant.Controller{Type: constant.ControllerTypeOfGasFireExtinguishingSystem, Addr: 3}, Version: constant.Version(0x0101)},
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
			p := constant.Packet{AppData: []byte{0x06, 0x01, 0x01, 0x02, 0x1e, 0x37, 0xc2, 0xa5, 0x30, 0x32, 0xca, 0xd2, 0xbb, 0xe1, 0xd2, 0xe9, 0xca, 0xd2, 0xd1, 0xcc, 0xb8, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerConfigures, []constant.ControllerConfigure{
				{Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2}, Configure: "7楼02室会议室烟感"},
			})
		})
		t.Run("test decode upload more system configure", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x06, 0x02,
				0x01, 0x02, 0x1e, 0x37, 0xc2, 0xa5, 0x30, 0x32, 0xca, 0xd2, 0xbb, 0xe1, 0xd2, 0xe9, 0xca, 0xd2, 0xd1, 0xcc, 0xb8, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x0d, 0x03, 0x1e, 0x37, 0xc2, 0xa5, 0x30, 0x32, 0xca, 0xd2, 0xbb, 0xe1, 0xd2, 0xe9, 0xca, 0xd2, 0xd1, 0xcc, 0xb8, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.ControllerConfigures, []constant.ControllerConfigure{
				{Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2}, Configure: "7楼02室会议室烟感"},
				{Ctrl: constant.Controller{Type: constant.ControllerTypeOfGasFireExtinguishingSystem, Addr: 3}, Configure: "7楼02室会议室烟感"},
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
			p := constant.Packet{AppData: []byte{0x07, 0x01, 0x01, 0x02, 0x28, 0x01, 0x02, 0x03, 0x04, 0x37, 0xc2, 0xa5, 0x30, 0x32, 0xca, 0xd2, 0xbb, 0xe1, 0xd2, 0xe9, 0xca, 0xd2, 0xd1, 0xcc, 0xb8, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.EquipmentConfigures, []constant.EquipmentConfigure{
				{Equ: constant.Equipment{
					Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2},
					Type: constant.EquipmentTypeOfSmokeFireDetector,
					Addr: constant.EquipmentAddr(0x04030201),
				}, Description: "7楼02室会议室烟感"},
			})
		})
		t.Run("test decode upload more equipment configure", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x07, 0x02,
				0x01, 0x02, 0x28, 0x01, 0x02, 0x03, 0x04, 0x37, 0xc2, 0xa5, 0x30, 0x32, 0xca, 0xd2, 0xbb, 0xe1, 0xd2, 0xe9, 0xca, 0xd2, 0xd1, 0xcc, 0xb8, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x0d, 0x03, 0x55, 0x03, 0x04, 0x05, 0x06, 0x37, 0xc2, 0xa5, 0x30, 0x32, 0xca, 0xd2, 0xbb, 0xe1, 0xd2, 0xe9, 0xca, 0xd2, 0xd1, 0xcc, 0xb8, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.EquipmentConfigures, []constant.EquipmentConfigure{
				{Equ: constant.Equipment{
					Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2},
					Type: constant.EquipmentTypeOfSmokeFireDetector,
					Addr: constant.EquipmentAddr(0x04030201),
				}, Description: "7楼02室会议室烟感"},
				{Equ: constant.Equipment{
					Ctrl: constant.Controller{Type: constant.ControllerTypeOfGasFireExtinguishingSystem, Addr: 3},
					Type: constant.EquipmentTypeOfInputModule,
					Addr: constant.EquipmentAddr(0x06050403),
				}, Description: "7楼02室会议室烟感"},
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
				{Ctrl: constant.Controller{Type: constant.ControllerTypeOfFireAlarmSystem, Addr: 2}, Timestamp: 1625187604},
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
				{Flag: constant.StateFlag(0x08), Timestamp: 1625187604},
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
				{Flag: constant.OperationFlag(0x31), Operator: 0xfa, Timestamp: 1625187604},
			})
		})
		t.Run("test decode upload more transmission operation", func(t *testing.T) {
			p := constant.Packet{AppData: []byte{0x18, 0x02,
				0x31, 0xfa, 0x04, 0x00, 0x09, 0x02, 0x07, 0x15,
				0x73, 0xfb, 0x05, 0x00, 0x09, 0x02, 0x07, 0x15}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.TransmissionOperations, []constant.TransmissionOperationInfo{
				{Flag: constant.OperationFlag(0x31), Operator: 0xfa, Timestamp: 1625187604},
				{Flag: constant.OperationFlag(0x73), Operator: 0xfb, Timestamp: 1625187605},
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
			p := constant.Packet{AppData: []byte{0x1a, 0x01, 0x1e, 0x37, 0xc2, 0xa5, 0x30, 0x32, 0xca, 0xd2, 0xbb, 0xe1, 0xd2, 0xe9, 0xca, 0xd2, 0xd1, 0xcc, 0xb8, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
			DecodeAppData(&p)
			assert.EqualValues(t, p.TransmissionConfigures, []constant.TransmissionConfigure{
				{Configure: "7楼02室会议室烟感"},
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
			assert.EqualValues(t, p.TransmissionTimestamps, []constant.TransmissionTimestamp{{Timestamp: 1625187605}})
		})
	})
}

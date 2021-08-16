package constant

import "strconv"

// AppType is application type
type AppType uint8

// Name is application name
func (ct AppType) Name() string {
	if name, ok := AppTypeNames[ct]; ok {
		return name
	}
	return ""
}

// String is implements fmt.Stringer
func (ct AppType) String() string {
	return "[" + strconv.Itoa(int(ct)) + "]" + ct.Name()
}

// define app types of 1 ~ 8
const (
	AppTypeOfUploadSystemState                AppType = iota + 1 // upload system state
	AppTypeOfUploadEquipmentState                                // upload equipment state
	AppTypeOfUploadEquipmentParameter                            // upload equipment parameter
	AppTypeOfUploadSystemOperatingInformation                    // upload system operating information
	AppTypeOfUploadSystemSoftwareVersion                         // upload system software version
	AppTypeOfUploadSystemConfigure                               // upload system configure
	AppTypeOfUploadEquipmentConfigure                            // upload equipment configure
	AppTypeOfUploadSystemTime                                    // upload system time
)

// define app types of 21 ~ 28
const (
	AppTypeOfUploadTransmissionState AppType = iota + 21 // upload transmission state
	_
	_
	AppTypeOfUploadTransmissionOperatingInformation // upload transmission operating information
	AppTypeOfUploadTransmissionSoftwareVersion      // upload transmission software version
	AppTypeOfUploadTransmissionConfigure            // upload transmission configure
	_
	AppTypeOfUploadTransmissionTime // upload transmission time
)

// define app types of 61 ~ 68
const (
	AppTypeOfQuerySystemState                AppType = iota + 61 // query system state
	AppTypeOfQueryEquipmentState                                 // query equipment state
	AppTypeOfQueryEquipmentParameter                             // query equipment parameter
	AppTypeOfQuerySystemOperatingInformation                     // query system operation information
	AppTypeOfQuerySystemSoftwareVersion                          // query system software version
	AppTypeOfQuerySystemConfigure                                // query system configure
	AppTypeOfQueryEquipmentConfigure                             // query equipment configure
	AppTypeOfQuerySystemTime                                     // query system time
)

// define app types of 81 ~ 91
const (
	AppTypeOfQueryTransmissionState AppType = iota + 81 // query transmission state
	_
	_
	AppTypeOfQueryTransmissionOperatingInformation // query transmission operation information
	AppTypeOfQueryTransmissionSoftwareVersion      // query transmission software version
	AppTypeOfQueryTransmissionConfigure            // query transmission configure
	_
	AppTypeOfQueryTransmissionTime  // query transmission time
	AppTypeOfInitializeTransmission // initialize transmission
	AppTypeOfSyncTransmissionTime   // sync transmission time
	AppTypeOfInspectSentries        // inspect sentries
)

// AppTypeNames is application names
var AppTypeNames = map[AppType]string{
	AppTypeOfUploadSystemState:                      "上传系统状态",
	AppTypeOfUploadEquipmentState:                   "上传部件状态",
	AppTypeOfUploadEquipmentParameter:               "上传部件参量",
	AppTypeOfUploadSystemOperatingInformation:       "上传操作信息",
	AppTypeOfUploadSystemSoftwareVersion:            "上传系统软件版本",
	AppTypeOfUploadSystemConfigure:                  "上传系统配置情况",
	AppTypeOfUploadEquipmentConfigure:               "上传部件配置情况",
	AppTypeOfUploadSystemTime:                       "上传系统时间",
	AppTypeOfUploadTransmissionState:                "上传传输装置状态",
	AppTypeOfUploadTransmissionOperatingInformation: "上传传输装置操作信息",
	AppTypeOfUploadTransmissionSoftwareVersion:      "上传传输装置软件版本",
	AppTypeOfUploadTransmissionConfigure:            "上传传输装置配置情况",
	AppTypeOfUploadTransmissionTime:                 "上传传输装置时间",
	AppTypeOfQuerySystemState:                       "查询系统状态",
	AppTypeOfQueryEquipmentState:                    "查询部件状态",
	AppTypeOfQueryEquipmentParameter:                "查询部件参量",
	AppTypeOfQuerySystemOperatingInformation:        "查询系统操作信息",
	AppTypeOfQuerySystemSoftwareVersion:             "查询系统软件版本",
	AppTypeOfQuerySystemConfigure:                   "查询系统配置情况",
	AppTypeOfQueryEquipmentConfigure:                "查询部件配置情况",
	AppTypeOfQuerySystemTime:                        "查询系统时间",
	AppTypeOfQueryTransmissionState:                 "查询传输装置状态",
	AppTypeOfQueryTransmissionOperatingInformation:  "查询传输装置的操作信息",
	AppTypeOfQueryTransmissionSoftwareVersion:       "查询传输装置软件版本",
	AppTypeOfQueryTransmissionConfigure:             "查询传输装置配置情况",
	AppTypeOfQueryTransmissionTime:                  "查询传输装置时间",
	AppTypeOfInitializeTransmission:                 "初始化传输装置",
	AppTypeOfSyncTransmissionTime:                   "同步传输装置时间",
	AppTypeOfInspectSentries:                        "查岗",
}

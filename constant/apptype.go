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
	UploadSystemStateAppType                AppType = iota + 1 // upload system state
	UploadEquipmentStateAppType                                // upload equipment state
	UploadEquipmentParameterAppType                            // upload equipment parameter
	UploadSystemOperatingInformationAppType                    // upload system operating information
	UploadSystemSoftwareVersionAppType                         // upload system software version
	UploadSystemConfigureAppType                               // upload system configure
	UploadEquipmentConfigureAppType                            // upload equipment configure
	UploadSystemTimeAppType                                    // upload system time
)

// define app types of 21 ~ 28
const (
	UploadTransmissionStateAppType AppType = iota + 21 // upload transmission state
	_
	_
	UploadTransmissionOperatingInformationAppType // upload transmission operating information
	UploadTransmissionSoftwareVersionAppType      // upload transmission software version
	UploadTransmissionConfigureAppType            // upload transmission configure
	_
	UploadTransmissionTimeAppType // upload transmission time
)

// define app types of 61 ~ 68
const (
	QuerySystemStateAppType                AppType = iota + 61 // query system state
	QueryEquipmentStateAppType                                 // query equipment state
	QueryEquipmentParameterAppType                             // query equipment parameter
	QuerySystemOperatingInformationAppType                     // query system operation information
	QuerySystemSoftwareVersionAppType                          // query system software version
	QuerySystemConfigureAppType                                // query system configure
	QueryEquipmentConfigureAppType                             // query equipment configure
	QuerySystemTimeAppType                                     // query system time
)

// define app types of 81 ~ 91
const (
	QueryTransmissionStateAppType AppType = iota + 81 // query transmission state
	_
	_
	QueryTransmissionOperatingInformationAppType // query transmission operation information
	QueryTransmissionSoftwareVersionAppType      // query transmission software version
	QueryTransmissionConfigureAppType            // query transmission configure
	_
	QueryTransmissionTimeAppType  // query transmission time
	InitializeTransmissionAppType // initialize transmission
	SyncTransmissionTimeAppType   // sync transmission time
	InspectSentriesAppType        // inspect sentries
)

// AppTypeNames is application names
var AppTypeNames = map[AppType]string{
	UploadSystemStateAppType:                      "UploadSystemState",
	UploadEquipmentStateAppType:                   "UploadEquipmentState",
	UploadEquipmentParameterAppType:               "UploadEquipmentParameter",
	UploadSystemOperatingInformationAppType:       "UploadSystemOperatingInformation",
	UploadSystemSoftwareVersionAppType:            "UploadSystemSoftwareVersion",
	UploadSystemConfigureAppType:                  "UploadSystemConfigure",
	UploadEquipmentConfigureAppType:               "UploadEquipmentConfigure",
	UploadSystemTimeAppType:                       "UploadSystemTime",
	UploadTransmissionStateAppType:                "UploadTransmissionState",
	UploadTransmissionOperatingInformationAppType: "UploadTransmissionOperatingInformation",
	UploadTransmissionSoftwareVersionAppType:      "UploadTransmissionSoftwareVersion",
	UploadTransmissionConfigureAppType:            "UploadTransmissionConfigure",
	UploadTransmissionTimeAppType:                 "UploadTransmissionTime",
	QuerySystemStateAppType:                       "QuerySystemState",
	QueryEquipmentStateAppType:                    "QueryEquipmentState",
	QueryEquipmentParameterAppType:                "QueryEquipmentParameter",
	QuerySystemOperatingInformationAppType:        "QuerySystemOperatingInformation",
	QuerySystemSoftwareVersionAppType:             "QuerySystemSoftwareVersion",
	QuerySystemConfigureAppType:                   "QuerySystemConfigure",
	QueryEquipmentConfigureAppType:                "QueryEquipmentConfigure",
	QuerySystemTimeAppType:                        "QuerySystemTime",
	QueryTransmissionStateAppType:                 "QueryTransmissionState",
	QueryTransmissionOperatingInformationAppType:  "QueryTransmissionOperatingInformation",
	QueryTransmissionSoftwareVersionAppType:       "QueryTransmissionSoftwareVersion",
	QueryTransmissionConfigureAppType:             "QueryTransmissionConfigure",
	QueryTransmissionTimeAppType:                  "QueryTransmissionTime",
	InitializeTransmissionAppType:                 "InitializeTransmission",
	SyncTransmissionTimeAppType:                   "SyncTransmissionTime",
	InspectSentriesAppType:                        "InspectSentries",
}

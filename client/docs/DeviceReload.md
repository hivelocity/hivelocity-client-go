# DeviceReload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicSshKeyIds** | **[]int32** | List of public ssh keys | [optional] [default to null]
**CustomIPXEScriptUrl** | **string** | Custom IPXE script URL. If device was provisioned was one before and this field is omitted, the original custom iPXE script is used. Otherwise specifying the script will override. | [optional] [default to null]
**ControlPanelId** | **int32** | The unique ID of the control panel product option to provision on this device. | [optional] [default to null]
**Body** | **string** | Describe any special requests for non instant devices. | [optional] [default to null]
**OperatingSystemId** | **int32** | The unique ID of the operating system product option to provision on this device. | [default to null]
**IgnitionId** | **int32** | The unique ID of an Ignition File for FlatcarOS provisions. | [optional] [default to null]
**Script** | **string** | A Cloud-Init script or a post-install script (Bash for Linux or Powershell for Windows). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



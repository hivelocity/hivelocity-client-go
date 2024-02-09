# BareMetalDeviceUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceReload** | **bool** | Whether to force a reload even if fields have not changed | [optional] [default to null]
**OsName** | **string** | The name of the Operating System to provision on this device. Must match name of an operating system product option. | [default to null]
**Hostname** | **string** | A FQDN for the device. For example: &#x60;example.hivelocity.net&#x60; | [default to null]
**CustomIPXEScriptContents** | **string** | Contents of iPXE script if not specifying URL. If both script URL and contents are not specified, the last iPXE script contents are used if OS selection requires an iPXE script. | [optional] [default to null]
**Tags** | **[]string** |  | [optional] [default to null]
**IgnitionId** | **int32** | The unique ID of an Ignition File for FlatcarOS provisions. | [optional] [default to null]
**PublicSshKeyId** | **int32** |  | [optional] [default to null]
**Script** | **string** | A Cloud-Init script or a post-install script (Bash for Linux or Powershell for Windows). | [optional] [default to null]
**CustomIPXEScriptURL** | **string** | URL to download custom iPXE script if not specifying contents in entirety. If both script URL and contents are not specified, the last iPXE script contents are used if OS selection requires an  iPXE script. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



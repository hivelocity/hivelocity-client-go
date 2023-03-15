# BareMetalDeviceCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | **string** | hourly|monthly|quarterly|semi-annually|annually|biennial|triennial | [optional] [default to null]
**IgnitionId** | **int32** | The unique ID of an Ignition File for FlatcarOS provisions. | [optional] [default to null]
**ProductId** | **int32** | The unique ID of the desired product to provision. | [default to null]
**BondingSupport** | **bool** | If true, ensures the provisioned device supports port bonding. If false, ensures no bonding support. | [optional] [default to null]
**Tags** | **[]string** | User specified values. | [optional] [default to null]
**VlanId** | **int32** |  | [optional] [default to null]
**ForceDeviceId** | **int32** | For users with reserved devices. The unique ID of the target device to provision. | [optional] [default to null]
**CustomIPXEScriptURL** | **string** | URL to download custom iPXE script if not supplying script in entirety | [optional] [default to null]
**CustomIPXEScriptContents** | **string** | Contents of iPXE script if not supplying URL | [optional] [default to null]
**PublicSshKeyId** | **int32** |  | [optional] [default to null]
**Hostname** | **string** | A FQDN for the device. For example: &#x60;example.hivelocity.net&#x60; | [default to null]
**LocationName** | **string** | A facility code. For example &#x60;NYC1&#x60;. | [default to null]
**OsName** | **string** | The name of the Operating System to provision on this device. Must match name of an operating system product option. | [default to null]
**Script** | **string** | A Cloud-Init script or a post-install script (Bash for Linux or Powershell for Windows). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



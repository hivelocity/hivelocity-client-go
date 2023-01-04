# BareMetalDevice

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | The user specified hostname for the device. Note: If the hostname is changed in the portal or on the device itself this value may not reflect the actual hostname on the device. | [optional] [default to null]
**PrimaryIp** | **string** | The first assigned public IP for accessing this device. | [optional] [default to null]
**Tags** | **[]string** | User specified values. | [optional] [default to null]
**CustomIPXEScriptURL** | **string** | URL of custom iPXE script used to provision device | [optional] [default to null]
**LocationName** | **string** | A facility code. For example &#x60;NYC1&#x60;. | [optional] [default to null]
**ServiceId** | **int32** | The unique ID of the service associated with this device. | [optional] [default to null]
**DeviceId** | **int32** | The unique ID of the device. | [optional] [default to null]
**ProductName** | **string** | The name of the product associated with this device. | [optional] [default to null]
**VlanId** | **int32** |  | [optional] [default to null]
**Period** | **string** | This device&#39;s service&#39;s billing period. | [optional] [default to null]
**PublicSshKeyId** | **int32** |  | [optional] [default to null]
**Script** | **string** | The post-install/cloud-init script used during this device&#39;s last provisioning. | [optional] [default to null]
**PowerStatus** | **string** | ON|OFF | [optional] [default to null]
**CustomIPXEScriptContents** | **string** | Contents of custom iPXE used to provision device | [optional] [default to null]
**OrderId** | **int32** | The unique ID of the order for this device. | [optional] [default to null]
**OsName** | **string** | The name of the operating system currently installed on this device. Note: If you manually reload your own OS over IPMI this value may not reflect the OS currently on your device. | [optional] [default to null]
**ProductId** | **int32** | The unique ID of the product associated with this device. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



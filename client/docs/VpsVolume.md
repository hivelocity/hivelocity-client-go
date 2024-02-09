# VpsVolume

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeId** | **string** | The unique VPS instance volume ID. | [optional] [default to null]
**Type_** | **string** | Disk type. Options are &#x60;ROOT&#x60; for root volumes and &#x60;DATADISK&#x60; for additional disks. | [optional] [default to null]
**Size** | **int32** | Disk size in bytes | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The created date/time of the VPS instance volume. | [optional] [default to null]
**DeviceId** | **int32** | Device ID the disk is attached to | [optional] [default to null]
**FacilityCode** | **string** | The facility code associated with the VPS Volume. For example: &#x60;TPA1&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



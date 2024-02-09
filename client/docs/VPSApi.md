# \VPSApi

All URIs are relative to *https://core.hivelocity.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVpsIdNetworkRemoveResource**](VPSApi.md#DeleteVpsIdNetworkRemoveResource) | **Delete** /vps/{deviceId}/network/{networkId} | Detach a VPS Network from a VPS Device
[**DeleteVpsNetworkIdResource**](VPSApi.md#DeleteVpsNetworkIdResource) | **Delete** /vps/network/{networkId} | Delete VPS Network
[**DeleteVpsSecondaryIpRemoveResource**](VPSApi.md#DeleteVpsSecondaryIpRemoveResource) | **Delete** /vps/{deviceId}/secondaryIp/{ipAssignmentId} | Remove a Secondary IP Address from VPS instance
[**DeleteVpsSnapshotIdResource**](VPSApi.md#DeleteVpsSnapshotIdResource) | **Delete** /vps/snapshot/{snapshotId} | Delete a VPS Volume Snapshot
[**DeleteVpsSnapshotScheduleIdResource**](VPSApi.md#DeleteVpsSnapshotScheduleIdResource) | **Delete** /vps/snapshotSchedule/{snapshotScheduleId} | Delete a VPS Snapshot Schedule
[**DeleteVpsVolumeIdResource**](VPSApi.md#DeleteVpsVolumeIdResource) | **Delete** /vps/volume/{volumeId} | Deletes a Volume
[**DeleteVpsisoIdResource**](VPSApi.md#DeleteVpsisoIdResource) | **Delete** /vps/iso/{isoId} | Delete an ISO
[**GetVpsAvailableNewVolumesResource**](VPSApi.md#GetVpsAvailableNewVolumesResource) | **Get** /vps/available-volume-sizes | Gets list of available sizes for new volumes
[**GetVpsAvailableUpgradeSizesResource**](VPSApi.md#GetVpsAvailableUpgradeSizesResource) | **Get** /vps/volume/{volumeId}/available-sizes | Get a volume&#39;s list of upgradable sizes
[**GetVpsIdEventsResource**](VPSApi.md#GetVpsIdEventsResource) | **Get** /vps/{deviceId}/events | Get VPS Instance Events
[**GetVpsIdResource**](VPSApi.md#GetVpsIdResource) | **Get** /vps/{deviceId} | Get VPS instance
[**GetVpsMetricsResource**](VPSApi.md#GetVpsMetricsResource) | **Get** /vps/{deviceId}/metrics | Get Metrics for a VPS instance
[**GetVpsNetworkIdResource**](VPSApi.md#GetVpsNetworkIdResource) | **Get** /vps/network/{networkId} | Get VPS Network
[**GetVpsNetworkResource**](VPSApi.md#GetVpsNetworkResource) | **Get** /vps/network | Get all VPS Networks
[**GetVpsResource**](VPSApi.md#GetVpsResource) | **Get** /vps/ | Get all VPS instances
[**GetVpsSecondaryIpAddResource**](VPSApi.md#GetVpsSecondaryIpAddResource) | **Get** /vps/{deviceId}/secondaryIp | Get all Secondary IP Addresses attached to VPS instance
[**GetVpsSelfMetadataResource**](VPSApi.md#GetVpsSelfMetadataResource) | **Get** /vps/self-metadata | 
[**GetVpsSnapshotIdResource**](VPSApi.md#GetVpsSnapshotIdResource) | **Get** /vps/snapshot/{snapshotId} | Get a VPS Volume Snapshot by ID
[**GetVpsSnapshotResource**](VPSApi.md#GetVpsSnapshotResource) | **Get** /vps/snapshot | Get all VPS Volume Snapshots available
[**GetVpsSnapshotScheduleIdResource**](VPSApi.md#GetVpsSnapshotScheduleIdResource) | **Get** /vps/snapshotSchedule/{snapshotScheduleId} | Get a VPS Snapshot Schedules
[**GetVpsSnapshotScheduleResource**](VPSApi.md#GetVpsSnapshotScheduleResource) | **Get** /vps/snapshotSchedule | Get all VPS Snapshot Schedules for a Facility or VPS Instance
[**GetVpsVolumeIdResource**](VPSApi.md#GetVpsVolumeIdResource) | **Get** /vps/volume/{volumeId} | Get a VPS Instance Volume
[**GetVpsVolumeResource**](VPSApi.md#GetVpsVolumeResource) | **Get** /vps/volume | Get all VPS Instance Volumes available
[**GetVpsisoIdResource**](VPSApi.md#GetVpsisoIdResource) | **Get** /vps/iso/{isoId} | Get the details of the VPS ISO ID specified
[**GetVpsisoResource**](VPSApi.md#GetVpsisoResource) | **Get** /vps/iso | Get all VPS ISOs available
[**PostVpsAttachIsoResource**](VPSApi.md#PostVpsAttachIsoResource) | **Post** /vps/{deviceId}/attachiso | Attach ISO to VPS instance
[**PostVpsAttachVolumeResource**](VPSApi.md#PostVpsAttachVolumeResource) | **Post** /vps/{deviceId}/volume/{volumeId}/attach | Attach Volume to VPS instance
[**PostVpsConsoleAccessResource**](VPSApi.md#PostVpsConsoleAccessResource) | **Post** /vps/{deviceId}/console | Create a one time use URL for console access to a VPS instance
[**PostVpsDetachIsoResource**](VPSApi.md#PostVpsDetachIsoResource) | **Post** /vps/{deviceId}/detachiso | Detach ISO from VPS instance
[**PostVpsDetachVolumeResource**](VPSApi.md#PostVpsDetachVolumeResource) | **Post** /vps/{deviceId}/volume/{volumeId}/detach | Detach Volume from VPS instance
[**PostVpsIdNetworkAddResource**](VPSApi.md#PostVpsIdNetworkAddResource) | **Post** /vps/{deviceId}/network | Attach VPS Network to a VPS Device
[**PostVpsNetworkResource**](VPSApi.md#PostVpsNetworkResource) | **Post** /vps/network | Create a new VPS Network
[**PostVpsResetPasswordResource**](VPSApi.md#PostVpsResetPasswordResource) | **Post** /vps/{deviceId}/reset-password | Request token to reset root password for VPS
[**PostVpsSecondaryIpAddResource**](VPSApi.md#PostVpsSecondaryIpAddResource) | **Post** /vps/{deviceId}/secondaryIp | Add a Secondary IP Address to VPS instance
[**PostVpsSnapshotIdResource**](VPSApi.md#PostVpsSnapshotIdResource) | **Post** /vps/snapshot/{snapshotId} | Restore a VPS Volume Snapshot
[**PostVpsSnapshotResource**](VPSApi.md#PostVpsSnapshotResource) | **Post** /vps/snapshot | Create a VPS Volume Snapshot
[**PostVpsSnapshotScheduleResource**](VPSApi.md#PostVpsSnapshotScheduleResource) | **Post** /vps/snapshotSchedule | Create a VPS Snapshot Schedule
[**PostVpsVolumeResource**](VPSApi.md#PostVpsVolumeResource) | **Post** /vps/volume | Creates a new Volume
[**PostVpsisoResource**](VPSApi.md#PostVpsisoResource) | **Post** /vps/iso | Create a new ISO
[**PutVpsIdResource**](VPSApi.md#PutVpsIdResource) | **Put** /vps/{deviceId} | Update VPS instance
[**PutVpsNetworkIdResource**](VPSApi.md#PutVpsNetworkIdResource) | **Put** /vps/network/{networkId} | Update VPS Network
[**PutVpsResetPasswordResource**](VPSApi.md#PutVpsResetPasswordResource) | **Put** /vps/{deviceId}/reset-password | Reset VPS root password, requires vps to be powered off
[**PutVpsVolumeIdResource**](VPSApi.md#PutVpsVolumeIdResource) | **Put** /vps/volume/{volumeId} | Resize a Volume


# **DeleteVpsIdNetworkRemoveResource**
> NetworkTaskDump DeleteVpsIdNetworkRemoveResource(ctx, deviceId, networkId, optional)
Detach a VPS Network from a VPS Device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **networkId** | **string**|  | 
 **optional** | ***VPSApiDeleteVpsIdNetworkRemoveResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiDeleteVpsIdNetworkRemoveResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVpsNetworkIdResource**
> NetworkTaskDump DeleteVpsNetworkIdResource(ctx, networkId, optional)
Delete VPS Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**|  | 
 **optional** | ***VPSApiDeleteVpsNetworkIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiDeleteVpsNetworkIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVpsSecondaryIpRemoveResource**
> NetworkTaskDump DeleteVpsSecondaryIpRemoveResource(ctx, deviceId, ipAssignmentId, optional)
Remove a Secondary IP Address from VPS instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **ipAssignmentId** | **int32**|  | 
 **optional** | ***VPSApiDeleteVpsSecondaryIpRemoveResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiDeleteVpsSecondaryIpRemoveResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVpsSnapshotIdResource**
> NetworkTaskDump DeleteVpsSnapshotIdResource(ctx, snapshotId, optional)
Delete a VPS Volume Snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**|  | 
 **optional** | ***VPSApiDeleteVpsSnapshotIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiDeleteVpsSnapshotIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **optional.Int32**| The unique client account ID. | 
 **facilityCode** | **optional.String**| The facility code associated with the VPS instance. For example: TPA1. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVpsSnapshotScheduleIdResource**
> DeleteVpsSnapshotScheduleIdResource(ctx, snapshotScheduleId, optional)
Delete a VPS Snapshot Schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotScheduleId** | **string**|  | 
 **optional** | ***VPSApiDeleteVpsSnapshotScheduleIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiDeleteVpsSnapshotScheduleIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **optional.Int32**| The unique client account ID. | 
 **facilityCode** | **optional.String**| The facility code associated with the VPS instance. For example: TPA1. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVpsVolumeIdResource**
> NetworkTaskDump DeleteVpsVolumeIdResource(ctx, volumeId, payload, optional)
Deletes a Volume

Warning: This will cancel the monthly billing service for this volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**|  | 
  **payload** | [**VpsVolumeDelete**](VpsVolumeDelete.md)|  | 
 **optional** | ***VPSApiDeleteVpsVolumeIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiDeleteVpsVolumeIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVpsisoIdResource**
> NetworkTaskDump DeleteVpsisoIdResource(ctx, isoId, optional)
Delete an ISO

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**|  | 
 **optional** | ***VPSApiDeleteVpsisoIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiDeleteVpsisoIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **facilityCode** | **optional.String**| The facility code where the ISOs are stored. | 
 **clientId** | **optional.Int32**| The unique client account ID. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsAvailableNewVolumesResource**
> VpsAvailableNewVolumeSizes GetVpsAvailableNewVolumesResource(ctx, optional)
Gets list of available sizes for new volumes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VPSApiGetVpsAvailableNewVolumesResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsAvailableNewVolumesResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**VpsAvailableNewVolumeSizes**](VPSAvailableNewVolumeSizes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsAvailableUpgradeSizesResource**
> VpsAvailableUpgradeSizes GetVpsAvailableUpgradeSizesResource(ctx, volumeId, optional)
Get a volume's list of upgradable sizes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**|  | 
 **optional** | ***VPSApiGetVpsAvailableUpgradeSizesResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsAvailableUpgradeSizesResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**VpsAvailableUpgradeSizes**](VPSAvailableUpgradeSizes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsIdEventsResource**
> []VpsEvent GetVpsIdEventsResource(ctx, deviceId, optional)
Get VPS Instance Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| Device ID of VPS instance | 
 **optional** | ***VPSApiGetVpsIdEventsResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsIdEventsResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]VpsEvent**](VPSEvent.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsIdResource**
> Vps GetVpsIdResource(ctx, deviceId, optional)
Get VPS instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| Device ID of VPS instance to View / Update | 
 **optional** | ***VPSApiGetVpsIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Vps**](VPS.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsMetricsResource**
> VpsMetricsResponse GetVpsMetricsResource(ctx, deviceId, metricType, optional)
Get Metrics for a VPS instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| Device ID of VPS instance | 
  **metricType** | **string**| Options: &#x60;CPU_USAGE&#x60;, &#x60;MEMORY_USAGE&#x60;, &#x60;DISK_RW&#x60;, &#x60;NETWORK_USAGE&#x60; | 
 **optional** | ***VPSApiGetVpsMetricsResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsMetricsResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **optional.Int32**| Desired Start Time for Chart Data. | 
 **endTime** | **optional.Int32**| Desired End Time for Chart Data. | 
 **imageWidth** | **optional.Int32**| Width to render metrics image in pixels. | 
 **imageHeight** | **optional.Int32**| Height to render metrics image in pixels. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**VpsMetricsResponse**](VPSMetricsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsNetworkIdResource**
> VpsNetwork GetVpsNetworkIdResource(ctx, networkId, optional)
Get VPS Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**|  | 
 **optional** | ***VPSApiGetVpsNetworkIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsNetworkIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **optional.Int32**| The unique client account ID. | 
 **facilityCode** | **optional.String**| The facility code where the VPS Network is located. For example: &#x60;TPA1&#x60;. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**VpsNetwork**](VPSNetwork.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsNetworkResource**
> []VpsNetwork GetVpsNetworkResource(ctx, optional)
Get all VPS Networks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VPSApiGetVpsNetworkResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsNetworkResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **optional.Int32**| The unique client account ID. | 
 **facilityCode** | **optional.String**| The facility code where the VPS Network is located. For example: &#x60;TPA1&#x60;. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]VpsNetwork**](VPSNetwork.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsResource**
> []Vps GetVpsResource(ctx, optional)
Get all VPS instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VPSApiGetVpsResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storageId** | **optional.String**| The storage server ID of the VPS instance. | 
 **clientId** | **optional.Int32**| The unique client account ID. | 
 **hostId** | **optional.String**| Host ID (hypervisor) of the VPS instance. | 
 **facilityCode** | **optional.String**| The facility code associated with the VPS instance. For example: TPA1. | 
 **networkId** | **optional.String**| The Network ID the VPS instance is attached to. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Vps**](VPS.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsSecondaryIpAddResource**
> []IpAssignment GetVpsSecondaryIpAddResource(ctx, deviceId, optional)
Get all Secondary IP Addresses attached to VPS instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
 **optional** | ***VPSApiGetVpsSecondaryIpAddResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsSecondaryIpAddResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]IpAssignment**](IPAssignment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsSelfMetadataResource**
> GetVpsSelfMetadataResource(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsSnapshotIdResource**
> VpsSnapshot GetVpsSnapshotIdResource(ctx, snapshotId, optional)
Get a VPS Volume Snapshot by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**|  | 
 **optional** | ***VPSApiGetVpsSnapshotIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsSnapshotIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **optional.Int32**| The unique client account ID. | 
 **facilityCode** | **optional.String**| The facility code associated with the VPS instance. For example: TPA1. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**VpsSnapshot**](VPSSnapshot.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsSnapshotResource**
> []VpsSnapshot GetVpsSnapshotResource(ctx, optional)
Get all VPS Volume Snapshots available

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VPSApiGetVpsSnapshotResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsSnapshotResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **optional.Int32**| The unique device ID of the VPS instance. | 
 **facilityCode** | **optional.String**| The facility code associated with the VPS instance. For example: TPA1. | 
 **clientId** | **optional.Int32**| The unique client account ID. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]VpsSnapshot**](VPSSnapshot.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsSnapshotScheduleIdResource**
> VpsSnapshotSchedule GetVpsSnapshotScheduleIdResource(ctx, snapshotScheduleId, optional)
Get a VPS Snapshot Schedules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotScheduleId** | **string**|  | 
 **optional** | ***VPSApiGetVpsSnapshotScheduleIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsSnapshotScheduleIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **optional.Int32**| The unique client account ID. | 
 **facilityCode** | **optional.String**| The facility code associated with the VPS instance. For example: TPA1. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**VpsSnapshotSchedule**](VPSSnapshotSchedule.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsSnapshotScheduleResource**
> []VpsSnapshotSchedule GetVpsSnapshotScheduleResource(ctx, optional)
Get all VPS Snapshot Schedules for a Facility or VPS Instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VPSApiGetVpsSnapshotScheduleResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsSnapshotScheduleResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **optional.Int32**| The unique device ID of the VPS instance. | 
 **facilityCode** | **optional.String**| The facility code associated with the VPS instance. For example: TPA1. | 
 **clientId** | **optional.Int32**| The unique client account ID. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]VpsSnapshotSchedule**](VPSSnapshotSchedule.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsVolumeIdResource**
> VpsVolume GetVpsVolumeIdResource(ctx, volumeId, optional)
Get a VPS Instance Volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**|  | 
 **optional** | ***VPSApiGetVpsVolumeIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsVolumeIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**VpsVolume**](VPSVolume.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsVolumeResource**
> []VpsVolume GetVpsVolumeResource(ctx, optional)
Get all VPS Instance Volumes available

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VPSApiGetVpsVolumeResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsVolumeResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **optional.Int32**| The unique device ID of the VPS instance. | 
 **facilityCode** | **optional.String**| The facility code associated with the VPS instance. For example: TPA1. | 
 **clientId** | **optional.Int32**| The unique client account ID. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]VpsVolume**](VPSVolume.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsisoIdResource**
> Vpsiso GetVpsisoIdResource(ctx, isoId, optional)
Get the details of the VPS ISO ID specified

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isoId** | **string**|  | 
 **optional** | ***VPSApiGetVpsisoIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsisoIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **facilityCode** | **optional.String**| The facility code where the ISOs are stored. | 
 **clientId** | **optional.Int32**| The unique client account ID. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Vpsiso**](VPSISO.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVpsisoResource**
> []Vpsiso GetVpsisoResource(ctx, optional)
Get all VPS ISOs available

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VPSApiGetVpsisoResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiGetVpsisoResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **facilityCode** | **optional.String**| The facility code where the ISOs are stored. | 
 **clientId** | **optional.Int32**| The unique client account ID. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Vpsiso**](VPSISO.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsAttachIsoResource**
> NetworkTaskDump PostVpsAttachIsoResource(ctx, deviceId, payload, optional)
Attach ISO to VPS instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| Device ID of VPS instance | 
  **payload** | [**VpsisoAttach**](VpsisoAttach.md)|  | 
 **optional** | ***VPSApiPostVpsAttachIsoResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPostVpsAttachIsoResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsAttachVolumeResource**
> NetworkTaskDump PostVpsAttachVolumeResource(ctx, volumeId, deviceId, optional)
Attach Volume to VPS instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| Volume ID to attach | 
  **deviceId** | **int32**| Device ID of VPS instance | 
 **optional** | ***VPSApiPostVpsAttachVolumeResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPostVpsAttachVolumeResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsConsoleAccessResource**
> VpsConsoleAccess PostVpsConsoleAccessResource(ctx, deviceId, optional)
Create a one time use URL for console access to a VPS instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| Device ID of VPS instance | 
 **optional** | ***VPSApiPostVpsConsoleAccessResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPostVpsConsoleAccessResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**VpsConsoleAccess**](VPSConsoleAccess.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsDetachIsoResource**
> NetworkTaskDump PostVpsDetachIsoResource(ctx, deviceId, optional)
Detach ISO from VPS instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| Device ID of VPS instance | 
 **optional** | ***VPSApiPostVpsDetachIsoResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPostVpsDetachIsoResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsDetachVolumeResource**
> NetworkTaskDump PostVpsDetachVolumeResource(ctx, volumeId, deviceId, optional)
Detach Volume from VPS instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| Volume ID to detach | 
  **deviceId** | **int32**| Device ID of VPS instance | 
 **optional** | ***VPSApiPostVpsDetachVolumeResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPostVpsDetachVolumeResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsIdNetworkAddResource**
> NetworkTaskDump PostVpsIdNetworkAddResource(ctx, deviceId, payload, optional)
Attach VPS Network to a VPS Device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
  **payload** | [**VpsInstanceNetworkAdd**](VpsInstanceNetworkAdd.md)|  | 
 **optional** | ***VPSApiPostVpsIdNetworkAddResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPostVpsIdNetworkAddResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsNetworkResource**
> VpsNetwork PostVpsNetworkResource(ctx, payload, optional)
Create a new VPS Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**VpsNetworkCreate**](VpsNetworkCreate.md)|  | 
 **optional** | ***VPSApiPostVpsNetworkResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPostVpsNetworkResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**VpsNetwork**](VPSNetwork.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsResetPasswordResource**
> PostVpsResetPasswordResource(ctx, deviceId)
Request token to reset root password for VPS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| Device ID of VPS instance | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsSecondaryIpAddResource**
> NetworkTaskDump PostVpsSecondaryIpAddResource(ctx, deviceId, optional)
Add a Secondary IP Address to VPS instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**|  | 
 **optional** | ***VPSApiPostVpsSecondaryIpAddResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPostVpsSecondaryIpAddResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsSnapshotIdResource**
> NetworkTaskDump PostVpsSnapshotIdResource(ctx, snapshotId, payload, optional)
Restore a VPS Volume Snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**|  | 
  **payload** | [**VpsSnapshotRestore**](VpsSnapshotRestore.md)|  | 
 **optional** | ***VPSApiPostVpsSnapshotIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPostVpsSnapshotIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsSnapshotResource**
> NetworkTaskDump PostVpsSnapshotResource(ctx, payload, optional)
Create a VPS Volume Snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**VpsSnapshotCreate**](VpsSnapshotCreate.md)|  | 
 **optional** | ***VPSApiPostVpsSnapshotResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPostVpsSnapshotResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsSnapshotScheduleResource**
> VpsSnapshotSchedule PostVpsSnapshotScheduleResource(ctx, payload, optional)
Create a VPS Snapshot Schedule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**VpsSnapshotScheduleCreate**](VpsSnapshotScheduleCreate.md)|  | 
 **optional** | ***VPSApiPostVpsSnapshotScheduleResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPostVpsSnapshotScheduleResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**VpsSnapshotSchedule**](VPSSnapshotSchedule.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsVolumeResource**
> NetworkTaskDump PostVpsVolumeResource(ctx, payload, optional)
Creates a new Volume

Warning: This endpoint incurs charges. The volume will be billed monthly regardless of attached device's billing cycle length.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**VpsVolumeCreate**](VpsVolumeCreate.md)|  | 
 **optional** | ***VPSApiPostVpsVolumeResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPostVpsVolumeResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostVpsisoResource**
> Vpsiso PostVpsisoResource(ctx, payload, optional)
Create a new ISO

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**VpsisoCreate**](VpsisoCreate.md)|  | 
 **optional** | ***VPSApiPostVpsisoResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPostVpsisoResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Vpsiso**](VPSISO.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutVpsIdResource**
> NetworkTaskDump PutVpsIdResource(ctx, deviceId, payload, optional)
Update VPS instance

This will incur additional charges if scaling up.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| Device ID of VPS instance to View / Update | 
  **payload** | [**VpsUpdate**](VpsUpdate.md)|  | 
 **optional** | ***VPSApiPutVpsIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPutVpsIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutVpsNetworkIdResource**
> NetworkTaskDump PutVpsNetworkIdResource(ctx, networkId, payload, optional)
Update VPS Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**|  | 
  **payload** | [**VpsNetworkUpdate**](VpsNetworkUpdate.md)|  | 
 **optional** | ***VPSApiPutVpsNetworkIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPutVpsNetworkIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutVpsResetPasswordResource**
> PutVpsResetPasswordResource(ctx, deviceId, payload, optional)
Reset VPS root password, requires vps to be powered off

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **int32**| Device ID of VPS instance | 
  **payload** | [**VpsResetRootPassword**](VpsResetRootPassword.md)|  | 
 **optional** | ***VPSApiPutVpsResetPasswordResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPutVpsResetPasswordResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutVpsVolumeIdResource**
> NetworkTaskDump PutVpsVolumeIdResource(ctx, volumeId, payload, optional)
Resize a Volume

Warning: This endpoint incurs charges.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**|  | 
  **payload** | [**VpsVolumeUpdate**](VpsVolumeUpdate.md)|  | 
 **optional** | ***VPSApiPutVpsVolumeIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VPSApiPutVpsVolumeIdResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**NetworkTaskDump**](NetworkTask-dump.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


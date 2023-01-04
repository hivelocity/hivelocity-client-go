# DeviceDump

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **int32** | The unique ID of the device. | [optional] [default to null]
**Name** | **string** | User given custom name. | [default to null]
**Status** | **string** | active|inactive | [optional] [default to null]
**DeviceType** | **string** | Generic description of device. Usually type and rack unit size. | [optional] [default to null]
**DeviceTypeGroup** | **string** | Generic group | [optional] [default to null]
**PowerStatus** | **interface{}** | ON|OFF | [optional] [default to null]
**HasCancellation** | **bool** | True if device has active cancellation request. | [optional] [default to null]
**IsManaged** | **bool** | True if device enrolled in managed services. | [optional] [default to null]
**IsReload** | **bool** | True if device currently reloading. | [optional] [default to null]
**MonitorsUp** | **int32** | # of passing device monitors | [optional] [default to null]
**MonitorsTotal** | **int32** | Total # device monitors | [optional] [default to null]
**ManagedAlertsTotal** | **int32** | # of managed service alerts. | [optional] [default to null]
**Ports** | **[]interface{}** | Device Ports (Network Interfaces). | [optional] [default to null]
**Hostname** | **string** | a fqdn for the device. for example: &#x60;example.hivelocity.net&#x60;. | [optional] [default to null]
**IpmiEnabled** | **bool** | True if device is accessible over IPMI by customer. | [optional] [default to null]
**DisplayedTags** | **[]interface{}** | List containing key/values of device info based on tag order. | [optional] [default to null]
**Tags** | **[]string** | List of all user set device tags. | [optional] [default to null]
**Location** | **interface{}** | Detailed information on the device location. | [optional] [default to null]
**NetworkAutomation** | **interface{}** | Network Automation status for device. | [optional] [default to null]
**PrimaryIp** | **string** | The first assigned public IP for accessing this device. | [optional] [default to null]
**IpmiAddress** | **interface{}** | IP address for IPMI connection. Requires you to whitelist your current IP or be on IPMI VPN. | [optional] [default to null]
**ServiceMonitors** | **[]string** |  | [optional] [default to null]
**BillingInfo** | **interface{}** | If set, detailed info on this device&#39;s billing method. Otherwise null. When null the accounts default billing info is used for payments. | [optional] [default to null]
**ServicePlan** | **int32** | The unique ID of the associated service. | [optional] [default to null]
**LastInvoiceId** | **int32** | The unique ID of the last invoice for this device. | [optional] [default to null]
**SelfProvisioning** | **bool** | True if instant server. | [optional] [default to null]
**Metadata** | **interface{}** | Additional metadata. | [optional] [default to null]
**SpsStatus** | **string** | BUILDING|IPMI_READY|PROVISIONABLE|RESERVED|WAIT_FOR_PXE|PROVISION_STARTED|PROVISION_WAIT_FOR_ADDONS|PROVISION_FINISHED|WAIT_TO_COMPLETE_ORDER|WAIT_TO_ASSIGN_SERVICE|WAIT_FOR_HARDWARE_SCAN|IN_USE|RELOADING|DEVICE_READY_TO_TEST|DEVICE_READY_TO_WIPE|DEVICE_READY_TO_UPGRADE_FIRMWARE|FAILED|CLEANUP_MOVE_TO_FAILED|IN_REVIEW | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



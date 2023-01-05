# IpAssignment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentId** | **int32** | Unique ID of this assignment | [optional] [default to null]
**ClientId** | **int32** | Unique ID of the assignment owner | [optional] [default to null]
**Subnet** | **string** | CIDR representado of this assignment. | [optional] [default to null]
**Netmask** | **string** | Netmask for this assignment. | [optional] [default to null]
**BroadcastIp** | **string** | Broadcast address for this subnet. | [optional] [default to null]
**GatewayIp** | **string** | Gateway address for this subnet. | [optional] [default to null]
**FirstUsableIp** | **string** | First address available for use by Devices on this subnet. | [optional] [default to null]
**LastUsableIp** | **string** | Last address available for use by Devices on this subnet. | [optional] [default to null]
**UsableIps** | **[]string** | A list of usable IP addresses on assignment. It is only filled for IPv4 subnets. | [optional] [default to null]
**FacilityCode** | **string** | Facility code of this assignment. | [optional] [default to null]
**Description** | **string** | Description for this assignment. | [optional] [default to null]
**DeviceId** | **int32** | The device receiving traffic from the assignment. | [optional] [default to null]
**PortId** | **int32** | The port receiving traffic from the assignment. | [optional] [default to null]
**VlanId** | **int32** | The VLAN receiving traffic from the assignment. | [optional] [default to null]
**NextHopIp** | **string** | Next Hop IP address, if this assignment is statically routed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



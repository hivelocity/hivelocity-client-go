# Vlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanId** | **int32** | Unique ID of the VLAN. | [optional] [default to null]
**PortIds** | **[]int32** | Unique IDs of ports or bonds. | [optional] [default to null]
**QInQ** | **bool** | If true, VLAN is configured in Q-in-Q Mode. Automation is not currently supported on Q-in-Q VLANs. | [optional] [default to null]
**FacilityCode** | **string** | For example: &#x60;NYC1&#x60;. | [optional] [default to null]
**Type_** | **string** | If &#x60;public&#x60;, this VLAN can have IPs assigned to become reachable from the internet. If &#x60;private&#x60;, this VLAN can not have IPs assigned and will never be reachabled from the internet. All VLANs are subject to traffic billing and overages, with the exception of private VLAN traffic on unbonded Devices. | [optional] [default to null]
**IpIds** | **[]int32** | Unique IDs of IP Assignments. | [optional] [default to null]
**Automated** | **bool** | If true, VLAN can be automated via API. If false, contact support to enable automation. | [optional] [default to null]
**VlanTag** | **int32** | The VLAN Tag id from the switch. Use this value when configuring your OS interfaces to use the VLAN. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



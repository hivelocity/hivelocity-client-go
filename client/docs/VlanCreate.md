# VlanCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FacilityCode** | **string** | The Facility where the VLAN will be created, for example: &#x60;NYC1&#x60;. Only Devices from the same Facility are allowed in the VLANs. | [default to null]
**Type_** | **string** | Choose &#x60;public&#x60; if you want to be able to assign IPs to the VLAN, making it reachable from the internet. Choose &#x60;private&#x60; if the VLAN should never be reachable from the internet. All VLANs are subject to traffic billing and overages, with the exception of private VLAN traffic on unbonded Devices. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



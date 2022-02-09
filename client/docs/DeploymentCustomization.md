# DeploymentCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatingSystem** | **string** | Operating System&#39;s Name or ID | 
**Hostnames** | **[]string** |  | 
**Quantity** | **int32** |  | [optional] 
**Options** | **[]int32** |  | [optional] 
**ProductId** | **int32** |  | 
**LocationCode** | **string** |  | [optional] [default to ]
**BillingPeriod** | **string** | must be one of [&#39;monthly&#39;, &#39;quarterly&#39;, &#39;semi-annually&#39;, &#39;annually&#39;, &#39;biennial&#39;, &#39;triennial&#39;, &#39;hourly&#39;] | [optional] [default to monthly]
**AdditionalNotes** | **[]string** |  | [optional] 
**ForceDeviceIds** | **[]int32** | Either deploy these Device IDs or fail | [optional] 
**PublicSshKeyId** | **int32** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



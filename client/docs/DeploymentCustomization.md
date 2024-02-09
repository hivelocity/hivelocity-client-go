# DeploymentCustomization

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingPeriod** | **string** | must be one of [&#39;monthly&#39;, &#39;quarterly&#39;, &#39;semi-annually&#39;, &#39;annually&#39;, &#39;biennial&#39;, &#39;triennial&#39;, &#39;hourly&#39;] | [optional] [default to null]
**OperatingSystem** | **string** | Operating System&#39;s Name or ID | [default to null]
**Options** | **[]int32** |  | [optional] [default to null]
**CustomIPXEScriptContents** | **string** | Contents of iPXE script if not supplying URL | [optional] [default to null]
**PublicSshKeyId** | **int32** | ID of SSH Key to use for deployment | [optional] [default to null]
**LocationCode** | **string** |  | [optional] [default to null]
**IgnitionIds** | **[]int32** | Specify Ignition file ID for CoreOS/Flatcar provisions | [optional] [default to null]
**CustomIPXEScriptURL** | **string** | URL to download custom iPXE script if not supplying script in entirety | [optional] [default to null]
**ProductId** | **int32** |  | [default to null]
**ForceDeviceIds** | **[]int32** | Either deploy these Device IDs or fail | [optional] [default to null]
**AdditionalNotes** | **[]string** |  | [optional] [default to null]
**Hostnames** | **[]string** |  | [default to null]
**Quantity** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



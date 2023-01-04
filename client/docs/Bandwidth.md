# Bandwidth

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | **interface{}** | Returns metadata about the query including the 95th percentile values of the returned data, legend for returned data arrays, and query args passed to the query. | [optional] [default to null]
**BandwidthData** | [**[][]float32**](array.md) | Returns an array where each item contains a dictionary of data and an identifier for each requested interface. The data itself is an array of arrays where each item is the measurementat the requested &#x60;step&#x60; and the index of each integer value corresponds with the legend returned in the &#x60;metadata&#x60;. | [optional] [default to null]
**Interface_** | **string** | The interface(s) displayed in the image. | [optional] [default to null]
**SwitchId** | **string** | The unique ID of the switch where bandwidth data is measured. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



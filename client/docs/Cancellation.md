# Cancellation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveMonths** | **int32** | The number of months the target device has been active. | [optional] [default to null]
**ClientId** | **int32** | Unique ID of a client. | [optional] [default to null]
**Id** | **int32** | Unique ID of the cancellation request. | [optional] [default to null]
**PlanId** | **int32** | Deprecated. | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Date of cancel request submission. | [optional] [default to null]
**ServiceId** | **int32** | Unique ID of a service. | [optional] [default to null]
**RequestDate** | [**time.Time**](time.Time.md) | The requested cancellation date, if specified. | [optional] [default to null]
**MonthlyPayment** | **int32** | The monthly cost of the cancelled device. | [optional] [default to null]
**DeviceId** | **int32** | Unique ID of a device. | [optional] [default to null]
**Cancelled** | **bool** | True if cancel completed. | [optional] [default to null]
**Name** | **string** | Name of user who submitted the request | [optional] [default to null]
**DeletedAt** | [**time.Time**](time.Time.md) | Date of cancellation completion. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



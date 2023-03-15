# NetworkTaskDump

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskCreated** | **interface{}** | Timestamp of task creation date. | [optional] [default to null]
**MetaData** | **string** | Key value arguments used to trigger the task. | [optional] [default to null]
**ClientId** | **int32** | Unique ID of the client triggering the task. | [optional] [default to null]
**Result** | **string** | Success|Pending|Failed. Null value also means Pending. | [optional] [default to null]
**TaskId** | **string** | Unique ID of the network task. | [optional] [default to null]
**DeviceId** | **int32** | Unique ID of the target device. | [optional] [default to null]
**TaskUpdated** | **interface{}** | Timestamp of most recent activity taken on the network task. This will update multiple times for multi-step tasks. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



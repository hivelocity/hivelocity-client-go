# VpsSnapshot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | **string** | The unique ID associated with the snapshot. | [optional] [default to null]
**SnapshotType** | **string** | Snapshot interval options: HOURLY, DAILY, WEEKLY, MONTHLY. | [optional] [default to null]
**Name** | **string** | The name of the VPS instance snapshot. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The created date/tie of the VPS instance snapshot. | [optional] [default to null]
**VolumeId** | **string** | The unique VPS instance volume ID the snapshot is taken of. | [optional] [default to null]
**Size** | **int32** | Snapshot size in bytes | [optional] [default to null]
**State** | **string** | The state of the snapshot ie. Scheduled, Started or Completed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# VpsSnapshotSchedule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weekday** | **int32** | The day of the week to run the snapshot process (1-7 with 1 being Monday). | [optional] [default to null]
**Hour** | **int32** | The hour of the day to run the snapshot process (0-23). | [optional] [default to null]
**Day** | **int32** | The day of the month to run the snapshot process (1-28). | [optional] [default to null]
**Timezone** | **string** | The timezone to use for the snapshot schedule in IANA format. Example: America/New_York. | [optional] [default to null]
**Minute** | **int32** | The minute of the hour to run the snapshot process (0-59). | [optional] [default to null]
**VolumeId** | **string** | The unique ID of the VPS instance volume to take a snapshot of. | [optional] [default to null]
**SnapshotScheduleId** | **string** | The unique ID associated with the snapshot schedule. | [optional] [default to null]
**IntervalType** | **string** | Snapshot interval options: HOURLY, DAILY, WEEKLY, MONTHLY. | [optional] [default to null]
**MaxSnapshots** | **int32** | The maximum number of snapshots to take within the interval period. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



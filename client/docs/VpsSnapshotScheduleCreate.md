# VpsSnapshotScheduleCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Minute** | **int32** | The minute of the hour to run the snapshot process (0-59). | [default to null]
**ClientId** | **int32** | The unique client account ID. | [optional] [default to null]
**VolumeId** | **string** | The unique ID of the VPS instance volume to take a snapshot of. | [default to null]
**MaxSnapshots** | **int32** | The maximum number of snapshots to retain at a time for this schedule. | [default to null]
**Day** | **int32** | The day of the month to run th snapshot process (1-28). This is required for Monthly. | [optional] [default to null]
**IntervalType** | **string** | Options: &#x60;HOURLY&#x60;, &#x60;DAILY&#x60;, &#x60;WEEKLY&#x60;, &#x60;MONTHLY&#x60; | [default to null]
**Timezone** | **string** | Timezone in IANA format. Example: &#x60;America/New_York&#x60; | [default to null]
**FacilityCode** | **string** | The facility code associated witht he VPS instance. For example:TPA1. | [default to null]
**Weekday** | **int32** | The day of the week to run the snapshot process (1-7 with 1 being Monday). This is required for Weekly. | [optional] [default to null]
**Hour** | **int32** | The hour of the day to run the snapshot process (0-23). | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



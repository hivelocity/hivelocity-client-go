# DeviceInitialCreds

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordReturnsUntil** | **int32** | Date password will expire from API. | [optional] [default to null]
**Password** | **string** | Password for initial ssh access. Will expire from api 7 days after initial provision. | [optional] [default to null]
**Port** | **int32** | Port for initial ssh access. | [optional] [default to null]
**LockerUrl** | **string** | Link to encrypted locker containing password for initial ssh access. Locker contents be expired from api 7 days after initial provision. | [optional] [default to null]
**User** | **string** | User for initial ssh access. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



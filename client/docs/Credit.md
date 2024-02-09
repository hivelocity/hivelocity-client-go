# Credit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID corresponding to a particular credit | [optional] [default to null]
**Amount** | **float32** | The amount of credit associated with the credit ID | [optional] [default to null]
**Balance** | **float32** | The amount of credit available for use associated with the credit ID | [optional] [default to null]
**ClientId** | **int32** | The client ID associated with the credit ID | [optional] [default to null]
**Created** | **int32** | The created date/time of the credit | [optional] [default to null]
**OrderId** | **int32** | The order ID associated with the credit ID | [optional] [default to null]
**Reason** | **string** | The reason code for the creation of the credit | [optional] [default to null]
**Status** | **string** | The status of the credit | [optional] [default to null]
**IsHourly** | **bool** | If credit is for hourly billing | [optional] [default to null]
**BillingInfoId** | **int32** | The Billing Info ID used to purchase the credits on this account | [optional] [default to null]
**ServiceIds** | **[]int32** | The list of service IDs associated with the credit ID | [optional] [default to null]
**AutoApply** | **int32** | Whether credits are automatically being applied to new orders or invoices. [0 &#x3D; Not Auto Apply, 1 &#x3D; Auto Apply New Invoices, 2 &#x3D; Auto Apply All Invoices] | [optional] [default to null]
**ExpirationDate** | **int32** | The expiration date of the credit | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



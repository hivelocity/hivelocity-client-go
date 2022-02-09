# \TicketApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTicketIdResource**](TicketApi.md#GetTicketIdResource) | **Get** /tickets/{ticketId} | Returns details of a specific ticket
[**GetTicketReplyResource**](TicketApi.md#GetTicketReplyResource) | **Get** /tickets/{ticketId}/reply | Returns a reply for a specific ticket
[**GetTicketResource**](TicketApi.md#GetTicketResource) | **Get** /tickets/ | Returns a list with all Tickets
[**GetTicketSearchResource**](TicketApi.md#GetTicketSearchResource) | **Get** /tickets/search | Return results of ticket search
[**PostTicketReplyResource**](TicketApi.md#PostTicketReplyResource) | **Post** /tickets/{ticketId}/reply | Creates reply for a specific Ticket
[**PostTicketResource**](TicketApi.md#PostTicketResource) | **Post** /tickets/ | Creates a new ticket
[**PutTicketIdResource**](TicketApi.md#PutTicketIdResource) | **Put** /tickets/{ticketId} | Updates a specific ticket



## GetTicketIdResource

> Ticket GetTicketIdResource(ctx, ticketId, optional)

Returns details of a specific ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32**| ticket database ID | 
 **optional** | ***GetTicketIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTicketIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicketReplyResource

> []TicketPost GetTicketReplyResource(ctx, ticketId, optional)

Returns a reply for a specific ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32**| ticket database ID | 
 **optional** | ***GetTicketReplyResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTicketReplyResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]TicketPost**](TicketPost.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicketResource

> []Ticket GetTicketResource(ctx, optional)

Returns a list with all Tickets

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTicketResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTicketResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Ticket**](Ticket.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicketSearchResource

> TicketSeachResult GetTicketSearchResource(ctx, optional)

Return results of ticket search

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTicketSearchResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTicketSearchResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int32**| Number of items per page. | [default to 10]
 **page** | **optional.Int32**| The page number of search. | [default to 1]
 **q** | **optional.String**| Content search. | 
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**TicketSeachResult**](TicketSeachResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTicketReplyResource

> TicketPost PostTicketReplyResource(ctx, ticketId, payload, optional)

Creates reply for a specific Ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32**| ticket database ID | 
**payload** | [**TicketCreateReply**](TicketCreateReply.md)|  | 
 **optional** | ***PostTicketReplyResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostTicketReplyResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**TicketPost**](TicketPost.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTicketResource

> Ticket PostTicketResource(ctx, payload, optional)

Creates a new ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**TicketCreate**](TicketCreate.md)|  | 
 **optional** | ***PostTicketResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostTicketResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTicketIdResource

> Ticket PutTicketIdResource(ctx, ticketId, payload, optional)

Updates a specific ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32**| ticket database ID | 
**payload** | [**TicketPut**](TicketPut.md)|  | 
 **optional** | ***PutTicketIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutTicketIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


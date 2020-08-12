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

> GetTicketIdResource(ctx, ticketId)

Returns details of a specific ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32**| ticket database ID | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicketReplyResource

> GetTicketReplyResource(ctx, ticketId)

Returns a reply for a specific ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32**| ticket database ID | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicketResource

> GetTicketResource(ctx, )

Returns a list with all Tickets

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicketSearchResource

> GetTicketSearchResource(ctx, optional)

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

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTicketReplyResource

> PostTicketReplyResource(ctx, ticketId, payload)

Creates reply for a specific Ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32**| ticket database ID | 
**payload** | [**TicketCreateReply**](TicketCreateReply.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTicketResource

> PostTicketResource(ctx, payload)

Creates a new ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**TicketCreate**](TicketCreate.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTicketIdResource

> PutTicketIdResource(ctx, ticketId, payload)

Updates a specific ticket

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int32**| ticket database ID | 
**payload** | [**TicketPut**](TicketPut.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


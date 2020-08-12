# \WebhookApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIdResource**](WebhookApi.md#DeleteIdResource) | **Delete** /webhooks/{webhookId} | Deletes a single webhook
[**GetIdResource**](WebhookApi.md#GetIdResource) | **Get** /webhooks/{webhookId} | Returns detailed information for a Single Webhook
[**GetWebhookEventResource**](WebhookApi.md#GetWebhookEventResource) | **Get** /webhooks/events | Returns all available Webhook Events
[**GetWebhookResource**](WebhookApi.md#GetWebhookResource) | **Get** /webhooks/ | Returns your active Webhooks
[**PostEventScriptActionTriggerResource**](WebhookApi.md#PostEventScriptActionTriggerResource) | **Post** /webhooks/trigger | Queues a webhook for the event script action that was triggered
[**PostWebhookResource**](WebhookApi.md#PostWebhookResource) | **Post** /webhooks/ | Create a new Webhook for a Webhook Event
[**PutIdResource**](WebhookApi.md#PutIdResource) | **Put** /webhooks/{webhookId} | Updates a Single Webhook


# **DeleteIdResource**
> DeleteIdResource(ctx, webhookId)
Deletes a single webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int32**| ID of Webhook to View / Update | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdResource**
> GetIdResource(ctx, webhookId)
Returns detailed information for a Single Webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int32**| ID of Webhook to View / Update | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhookEventResource**
> GetWebhookEventResource(ctx, )
Returns all available Webhook Events

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhookResource**
> GetWebhookResource(ctx, )
Returns your active Webhooks

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEventScriptActionTriggerResource**
> PostEventScriptActionTriggerResource(ctx, optional)
Queues a webhook for the event script action that was triggered

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhookApiPostEventScriptActionTriggerResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiPostEventScriptActionTriggerResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookName** | **optional.String**| The name of the webhook to trigger. | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWebhookResource**
> PostWebhookResource(ctx, payload)
Create a new Webhook for a Webhook Event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payload** | [**WebhookCreate**](WebhookCreate.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutIdResource**
> PutIdResource(ctx, webhookId, payload)
Updates a Single Webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int32**| ID of Webhook to View / Update | 
  **payload** | [**WebhookUpdate**](WebhookUpdate.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


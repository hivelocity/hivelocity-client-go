# \WebhookApi

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhookIdResource**](WebhookApi.md#DeleteWebhookIdResource) | **Delete** /webhooks/{webhookId} | Deletes a single webhook
[**GetWebhookEventResource**](WebhookApi.md#GetWebhookEventResource) | **Get** /webhooks/events | Returns all available Webhook Events
[**GetWebhookIdResource**](WebhookApi.md#GetWebhookIdResource) | **Get** /webhooks/{webhookId} | Returns detailed information for a Single Webhook
[**GetWebhookResource**](WebhookApi.md#GetWebhookResource) | **Get** /webhooks/ | Returns your active Webhooks
[**PostEventScriptActionTriggerResource**](WebhookApi.md#PostEventScriptActionTriggerResource) | **Post** /webhooks/trigger | Queues a webhook for the event script action that was triggered
[**PostWebhookResource**](WebhookApi.md#PostWebhookResource) | **Post** /webhooks/ | Create a new Webhook for a Webhook Event
[**PutWebhookIdResource**](WebhookApi.md#PutWebhookIdResource) | **Put** /webhooks/{webhookId} | Updates a Single Webhook



## DeleteWebhookIdResource

> DeleteWebhookIdResource(ctx, webhookId)

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

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookEventResource

> []WebhookEvent GetWebhookEventResource(ctx, optional)

Returns all available Webhook Events

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWebhookEventResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWebhookEventResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]WebhookEvent**](WebhookEvent.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookIdResource

> Webhook GetWebhookIdResource(ctx, webhookId, optional)

Returns detailed information for a Single Webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int32**| ID of Webhook to View / Update | 
 **optional** | ***GetWebhookIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWebhookIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookResource

> []Webhook GetWebhookResource(ctx, optional)

Returns your active Webhooks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWebhookResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWebhookResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**[]Webhook**](Webhook.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEventScriptActionTriggerResource

> PostEventScriptActionTriggerResource(ctx, optional)

Queues a webhook for the event script action that was triggered

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostEventScriptActionTriggerResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostEventScriptActionTriggerResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookName** | **optional.String**| The name of the webhook to trigger. | 

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


## PostWebhookResource

> Webhook PostWebhookResource(ctx, payload, optional)

Create a new Webhook for a Webhook Event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payload** | [**WebhookCreate**](WebhookCreate.md)|  | 
 **optional** | ***PostWebhookResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostWebhookResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWebhookIdResource

> Webhook PutWebhookIdResource(ctx, webhookId, payload, optional)

Updates a Single Webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int32**| ID of Webhook to View / Update | 
**payload** | [**WebhookUpdate**](WebhookUpdate.md)|  | 
 **optional** | ***PutWebhookIdResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutWebhookIdResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xFields** | **optional.String**| An optional fields mask | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// BandwidthApiService BandwidthApi service
type BandwidthApiService service

// PostDeviceIdBandwidthImageResourceOpts Optional parameters for the method 'PostDeviceIdBandwidthImageResource'
type PostDeviceIdBandwidthImageResourceOpts struct {
	Start optional.Int32
	End   optional.Int32
}

/*
PostDeviceIdBandwidthImageResource Returns RRDTool Graph based bandwidth in PNG format
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param deviceId ID of Device to View
 * @param period Preconfigured Time Periods for Graph Data
 * @param interface_ Network Interface to use for Graph Data
 * @param optional nil or *PostDeviceIdBandwidthImageResourceOpts - Optional Parameters:
 * @param "Start" (optional.Int32) -  Start Time of Custom Time Period. (Unix Epoch Time)
 * @param "End" (optional.Int32) -  End Time of Custom Time Period (Unix Epoch Time)
*/
func (a *BandwidthApiService) PostDeviceIdBandwidthImageResource(ctx _context.Context, deviceId int32, period string, interface_ string, localVarOptionals *PostDeviceIdBandwidthImageResourceOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/bandwidth/device/{deviceId}/image"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", _neturl.QueryEscape(parameterToString(deviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("period", parameterToString(period, ""))
	localVarQueryParams.Add("interface", parameterToString(interface_, ""))
	if localVarOptionals != nil && localVarOptionals.Start.IsSet() {
		localVarQueryParams.Add("start", parameterToString(localVarOptionals.Start.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.End.IsSet() {
		localVarQueryParams.Add("end", parameterToString(localVarOptionals.End.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-API-KEY"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// PostDeviceIdBandwidthResourceOpts Optional parameters for the method 'PostDeviceIdBandwidthResource'
type PostDeviceIdBandwidthResourceOpts struct {
	Historical optional.Bool
	Start      optional.Int32
	End        optional.Int32
}

/*
PostDeviceIdBandwidthResource Returns RRDTool Xport based bandwidth data in JSON format
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param deviceId ID of Device to View
 * @param period Preconfigured Time Periods for Graph Data
 * @param interface_ Network Interface to use for Graph Data
 * @param step Interval of Graph in Seconds
 * @param optional nil or *PostDeviceIdBandwidthResourceOpts - Optional Parameters:
 * @param "Historical" (optional.Bool) -  Include Historical Interface Data for Device for Resellers
 * @param "Start" (optional.Int32) -  Start Time of Custom Time Period. (Unix Epoch Time)
 * @param "End" (optional.Int32) -  End Time of Custom Time Period (Unix Epoch Time)
*/
func (a *BandwidthApiService) PostDeviceIdBandwidthResource(ctx _context.Context, deviceId int32, period string, interface_ string, step int32, localVarOptionals *PostDeviceIdBandwidthResourceOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/bandwidth/device/{deviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", _neturl.QueryEscape(parameterToString(deviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("period", parameterToString(period, ""))
	localVarQueryParams.Add("interface", parameterToString(interface_, ""))
	if localVarOptionals != nil && localVarOptionals.Historical.IsSet() {
		localVarQueryParams.Add("historical", parameterToString(localVarOptionals.Historical.Value(), ""))
	}
	localVarQueryParams.Add("step", parameterToString(step, ""))
	if localVarOptionals != nil && localVarOptionals.Start.IsSet() {
		localVarQueryParams.Add("start", parameterToString(localVarOptionals.Start.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.End.IsSet() {
		localVarQueryParams.Add("end", parameterToString(localVarOptionals.End.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-API-KEY"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// PostServiceIdBandwidthImageResourceOpts Optional parameters for the method 'PostServiceIdBandwidthImageResource'
type PostServiceIdBandwidthImageResourceOpts struct {
	Start optional.Int32
	End   optional.Int32
}

/*
PostServiceIdBandwidthImageResource Returns RRDTool Graph based bandwidth in PNG format
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceId ID of Service to View
 * @param period Preconfigured Time Periods for Graph Data
 * @param interface_ Network Interface to use for Graph Data
 * @param optional nil or *PostServiceIdBandwidthImageResourceOpts - Optional Parameters:
 * @param "Start" (optional.Int32) -  Start Time of Custom Time Period. (Unix Epoch Time)
 * @param "End" (optional.Int32) -  End Time of Custom Time Period (Unix Epoch Time)
*/
func (a *BandwidthApiService) PostServiceIdBandwidthImageResource(ctx _context.Context, serviceId int32, period string, interface_ string, localVarOptionals *PostServiceIdBandwidthImageResourceOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/bandwidth/service/{serviceId}/image"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", _neturl.QueryEscape(parameterToString(serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("period", parameterToString(period, ""))
	localVarQueryParams.Add("interface", parameterToString(interface_, ""))
	if localVarOptionals != nil && localVarOptionals.Start.IsSet() {
		localVarQueryParams.Add("start", parameterToString(localVarOptionals.Start.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.End.IsSet() {
		localVarQueryParams.Add("end", parameterToString(localVarOptionals.End.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-API-KEY"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// PostServiceIdBandwidthResourceOpts Optional parameters for the method 'PostServiceIdBandwidthResource'
type PostServiceIdBandwidthResourceOpts struct {
	Start optional.Int32
	End   optional.Int32
}

/*
PostServiceIdBandwidthResource Returns RRDTool Xport based bandwidth data in JSON format
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param serviceId ID of Service to View
 * @param period Preconfigured Time Periods for Graph Data
 * @param interface_ Network Interface to use for Graph Data
 * @param step Interval of Graph in Seconds
 * @param optional nil or *PostServiceIdBandwidthResourceOpts - Optional Parameters:
 * @param "Start" (optional.Int32) -  Start Time of Custom Time Period. (Unix Epoch Time)
 * @param "End" (optional.Int32) -  End Time of Custom Time Period (Unix Epoch Time)
*/
func (a *BandwidthApiService) PostServiceIdBandwidthResource(ctx _context.Context, serviceId int32, period string, interface_ string, step int32, localVarOptionals *PostServiceIdBandwidthResourceOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/bandwidth/service/{serviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", _neturl.QueryEscape(parameterToString(serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("period", parameterToString(period, ""))
	localVarQueryParams.Add("interface", parameterToString(interface_, ""))
	localVarQueryParams.Add("step", parameterToString(step, ""))
	if localVarOptionals != nil && localVarOptionals.Start.IsSet() {
		localVarQueryParams.Add("start", parameterToString(localVarOptionals.Start.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.End.IsSet() {
		localVarQueryParams.Add("end", parameterToString(localVarOptionals.End.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-API-KEY"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

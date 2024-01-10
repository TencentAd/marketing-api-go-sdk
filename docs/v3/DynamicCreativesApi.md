# TencentAds\DynamicCreativesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicCreativesAdd**](DynamicCreativesApi.md#DynamicCreativesAdd) | **Post** /dynamic_creatives/add | 创建动态创意
[**DynamicCreativesDelete**](DynamicCreativesApi.md#DynamicCreativesDelete) | **Post** /dynamic_creatives/delete | 删除广告创意
[**DynamicCreativesGet**](DynamicCreativesApi.md#DynamicCreativesGet) | **Get** /dynamic_creatives/get | 获取动态创意
[**DynamicCreativesUpdate**](DynamicCreativesApi.md#DynamicCreativesUpdate) | **Post** /dynamic_creatives/update | 更新创意


# **DynamicCreativesAdd**
> DynamicCreativesAddResponse DynamicCreativesAdd(ctx, data)
创建动态创意

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DynamicCreativesAddRequest**](DynamicCreativesAddRequest.md)|  | 

### Return type

[**DynamicCreativesAddResponse**](DynamicCreativesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicCreativesDelete**
> DynamicCreativesDeleteResponse DynamicCreativesDelete(ctx, data)
删除广告创意

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DynamicCreativesDeleteRequest**](DynamicCreativesDeleteRequest.md)|  | 

### Return type

[**DynamicCreativesDeleteResponse**](DynamicCreativesDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicCreativesGet**
> DynamicCreativesGetResponse DynamicCreativesGet(ctx, accountId, optional)
获取动态创意

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***DynamicCreativesApiDynamicCreativesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DynamicCreativesApiDynamicCreativesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)|  | 
 **isDeleted** | **optional.Bool**|  | 

### Return type

[**DynamicCreativesGetResponse**](DynamicCreativesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicCreativesUpdate**
> DynamicCreativesUpdateResponse DynamicCreativesUpdate(ctx, data)
更新创意

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DynamicCreativesUpdateRequest**](DynamicCreativesUpdateRequest.md)|  | 

### Return type

[**DynamicCreativesUpdateResponse**](DynamicCreativesUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


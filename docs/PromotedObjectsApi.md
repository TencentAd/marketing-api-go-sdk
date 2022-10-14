# TencentAds\PromotedObjectsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PromotedObjectsAdd**](PromotedObjectsApi.md#PromotedObjectsAdd) | **Post** /promoted_objects/add | 创建推广目标
[**PromotedObjectsAuthorize**](PromotedObjectsApi.md#PromotedObjectsAuthorize) | **Post** /promoted_objects/authorize | 推广目标授权接口
[**PromotedObjectsDelete**](PromotedObjectsApi.md#PromotedObjectsDelete) | **Post** /promoted_objects/delete | 删除推广目标
[**PromotedObjectsGet**](PromotedObjectsApi.md#PromotedObjectsGet) | **Get** /promoted_objects/get | 获取推广目标
[**PromotedObjectsUpdate**](PromotedObjectsApi.md#PromotedObjectsUpdate) | **Post** /promoted_objects/update | 更新推广目标


# **PromotedObjectsAdd**
> PromotedObjectsAddResponse PromotedObjectsAdd(ctx, data)
创建推广目标

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**PromotedObjectsAddRequest**](PromotedObjectsAddRequest.md)|  | 

### Return type

[**PromotedObjectsAddResponse**](PromotedObjectsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PromotedObjectsAuthorize**
> PromotedObjectsAuthorizeResponse PromotedObjectsAuthorize(ctx, data)
推广目标授权接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**PromotedObjectsAuthorizeRequest**](PromotedObjectsAuthorizeRequest.md)|  | 

### Return type

[**PromotedObjectsAuthorizeResponse**](PromotedObjectsAuthorizeResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PromotedObjectsDelete**
> PromotedObjectsDeleteResponse PromotedObjectsDelete(ctx, data)
删除推广目标

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**PromotedObjectsDeleteRequest**](PromotedObjectsDeleteRequest.md)|  | 

### Return type

[**PromotedObjectsDeleteResponse**](PromotedObjectsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PromotedObjectsGet**
> PromotedObjectsGetResponse PromotedObjectsGet(ctx, accountId, optional)
获取推广目标

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***PromotedObjectsApiPromotedObjectsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PromotedObjectsApiPromotedObjectsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**PromotedObjectsGetResponse**](PromotedObjectsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PromotedObjectsUpdate**
> PromotedObjectsUpdateResponse PromotedObjectsUpdate(ctx, data)
更新推广目标

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**PromotedObjectsUpdateRequest**](PromotedObjectsUpdateRequest.md)|  | 

### Return type

[**PromotedObjectsUpdateResponse**](PromotedObjectsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


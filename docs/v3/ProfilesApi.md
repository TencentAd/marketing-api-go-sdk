# TencentAds\ProfilesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProfilesAdd**](ProfilesApi.md#ProfilesAdd) | **Post** /profiles/add | 创建朋友圈头像昵称跳转页
[**ProfilesDelete**](ProfilesApi.md#ProfilesDelete) | **Post** /profiles/delete | 删除朋友圈头像昵称跳转页
[**ProfilesGet**](ProfilesApi.md#ProfilesGet) | **Get** /profiles/get | 获取朋友圈头像昵称跳转页


# **ProfilesAdd**
> ProfilesAddResponse ProfilesAdd(ctx, data)
创建朋友圈头像昵称跳转页

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProfilesAddRequest**](ProfilesAddRequest.md)|  | 

### Return type

[**ProfilesAddResponse**](ProfilesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProfilesDelete**
> ProfilesDeleteResponse ProfilesDelete(ctx, data)
删除朋友圈头像昵称跳转页

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProfilesDeleteRequest**](ProfilesDeleteRequest.md)|  | 

### Return type

[**ProfilesDeleteResponse**](ProfilesDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProfilesGet**
> ProfilesGetResponse ProfilesGet(ctx, accountId, optional)
获取朋友圈头像昵称跳转页

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***ProfilesApiProfilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProfilesApiProfilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProfilesGetResponse**](ProfilesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


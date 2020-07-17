# TencentAds\UserPropertySetsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserPropertySetsAdd**](UserPropertySetsApi.md#UserPropertySetsAdd) | **Post** /user_property_sets/add | 创建用户属性数据源
[**UserPropertySetsGet**](UserPropertySetsApi.md#UserPropertySetsGet) | **Get** /user_property_sets/get | 获取用户属性数据源
[**UserPropertySetsUpdate**](UserPropertySetsApi.md#UserPropertySetsUpdate) | **Post** /user_property_sets/update | 更新用户属性数据源


# **UserPropertySetsAdd**
> UserPropertySetsAddResponse UserPropertySetsAdd(ctx, data)
创建用户属性数据源

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**UserPropertySetsAddRequest**](UserPropertySetsAddRequest.md)|  | 

### Return type

[**UserPropertySetsAddResponse**](UserPropertySetsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPropertySetsGet**
> UserPropertySetsGetResponse UserPropertySetsGet(ctx, accountId, optional)
获取用户属性数据源

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***UserPropertySetsApiUserPropertySetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserPropertySetsApiUserPropertySetsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userPropertySetId** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**UserPropertySetsGetResponse**](UserPropertySetsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPropertySetsUpdate**
> UserPropertySetsUpdateResponse UserPropertySetsUpdate(ctx, data)
更新用户属性数据源

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**UserPropertySetsUpdateRequest**](UserPropertySetsUpdateRequest.md)|  | 

### Return type

[**UserPropertySetsUpdateResponse**](UserPropertySetsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


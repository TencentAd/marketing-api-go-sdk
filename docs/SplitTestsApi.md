# TencentAds\SplitTestsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SplitTestsAdd**](SplitTestsApi.md#SplitTestsAdd) | **Post** /split_tests/add | 创建拆分对比实验
[**SplitTestsDelete**](SplitTestsApi.md#SplitTestsDelete) | **Post** /split_tests/delete | 删除拆分对比实验
[**SplitTestsGet**](SplitTestsApi.md#SplitTestsGet) | **Get** /split_tests/get | 获取拆分对比实验
[**SplitTestsUpdate**](SplitTestsApi.md#SplitTestsUpdate) | **Post** /split_tests/update | 更新拆分对比实验


# **SplitTestsAdd**
> SplitTestsAddResponse SplitTestsAdd(ctx, data)
创建拆分对比实验

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SplitTestsAddRequest**](SplitTestsAddRequest.md)|  | 

### Return type

[**SplitTestsAddResponse**](SplitTestsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SplitTestsDelete**
> SplitTestsDeleteResponse SplitTestsDelete(ctx, data)
删除拆分对比实验

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SplitTestsDeleteRequest**](SplitTestsDeleteRequest.md)|  | 

### Return type

[**SplitTestsDeleteResponse**](SplitTestsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SplitTestsGet**
> SplitTestsGetResponse SplitTestsGet(ctx, accountId, optional)
获取拆分对比实验

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***SplitTestsApiSplitTestsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SplitTestsApiSplitTestsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**SplitTestsGetResponse**](SplitTestsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SplitTestsUpdate**
> SplitTestsUpdateResponse SplitTestsUpdate(ctx, data)
更新拆分对比实验

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SplitTestsUpdateRequest**](SplitTestsUpdateRequest.md)|  | 

### Return type

[**SplitTestsUpdateResponse**](SplitTestsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


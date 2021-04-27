# TencentAds\PlayablePagesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlayablePagesAdd**](PlayablePagesApi.md#PlayablePagesAdd) | **Post** /playable_pages/add | 添加互动推广页（待废弃,详见公告）
[**PlayablePagesGet**](PlayablePagesApi.md#PlayablePagesGet) | **Get** /playable_pages/get | 获取互动推广页


# **PlayablePagesAdd**
> PlayablePagesAddResponse PlayablePagesAdd(ctx, accountId, playablePageName, materialFile)
添加互动推广页（待废弃,详见公告）

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **playablePageName** | **string**|  | 
  **materialFile** | ***os.File**|  | 

### Return type

[**PlayablePagesAddResponse**](PlayablePagesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayablePagesGet**
> PlayablePagesGetResponse PlayablePagesGet(ctx, accountId, optional)
获取互动推广页

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***PlayablePagesApiPlayablePagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayablePagesApiPlayablePagesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**PlayablePagesGetResponse**](PlayablePagesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


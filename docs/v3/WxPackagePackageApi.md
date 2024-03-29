# TencentAds\WxPackagePackageApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WxPackagePackageAdd**](WxPackagePackageApi.md#WxPackagePackageAdd) | **Post** /wx_package_package/add | 添加蹊径号码包
[**WxPackagePackageGet**](WxPackagePackageApi.md#WxPackagePackageGet) | **Get** /wx_package_package/get | 获取蹊径号码包列表
[**WxPackagePackageUpdate**](WxPackagePackageApi.md#WxPackagePackageUpdate) | **Post** /wx_package_package/update | 更新蹊径号码包


# **WxPackagePackageAdd**
> WxPackagePackageAddResponse WxPackagePackageAdd(ctx, data)
添加蹊径号码包

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WxPackagePackageAddRequest**](WxPackagePackageAddRequest.md)|  | 

### Return type

[**WxPackagePackageAddResponse**](WxPackagePackageAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WxPackagePackageGet**
> WxPackagePackageGetResponse WxPackagePackageGet(ctx, accountId, optional)
获取蹊径号码包列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WxPackagePackageApiWxPackagePackageGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WxPackagePackageApiWxPackagePackageGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int64**|  | 
 **pageIndex** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WxPackagePackageGetResponse**](WxPackagePackageGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WxPackagePackageUpdate**
> WxPackagePackageUpdateResponse WxPackagePackageUpdate(ctx, data)
更新蹊径号码包

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WxPackagePackageUpdateRequest**](WxPackagePackageUpdateRequest.md)|  | 

### Return type

[**WxPackagePackageUpdateResponse**](WxPackagePackageUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


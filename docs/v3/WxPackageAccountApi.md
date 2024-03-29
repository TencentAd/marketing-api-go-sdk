# TencentAds\WxPackageAccountApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WxPackageAccountGet**](WxPackageAccountApi.md#WxPackageAccountGet) | **Get** /wx_package_account/get | 获取蹊径微信号列表
[**WxPackageAccountUpdate**](WxPackageAccountApi.md#WxPackageAccountUpdate) | **Post** /wx_package_account/update | 更新蹊径微信号


# **WxPackageAccountGet**
> WxPackageAccountGetResponse WxPackageAccountGet(ctx, accountId, optional)
获取蹊径微信号列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WxPackageAccountApiWxPackageAccountGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WxPackageAccountApiWxPackageAccountGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int64**|  | 
 **pageIndex** | **optional.Int64**|  | 
 **beginTime** | **optional.String**|  | 
 **endTime** | **optional.String**|  | 
 **keyword** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WxPackageAccountGetResponse**](WxPackageAccountGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WxPackageAccountUpdate**
> WxPackageAccountUpdateResponse WxPackageAccountUpdate(ctx, accountId, wechatId, optional)
更新蹊径微信号

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **wechatId** | **int64**|  | 
 **optional** | ***WxPackageAccountApiWxPackageAccountUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WxPackageAccountApiWxPackageAccountUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nickName** | **optional.String**|  | 
 **file** | **optional.Interface of *os.File**|  | 
 **enableFlag** | **optional.Int64**|  | 

### Return type

[**WxPackageAccountUpdateResponse**](WxPackageAccountUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


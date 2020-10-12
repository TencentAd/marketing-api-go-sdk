# TencentAds\WechatQualificationsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatQualificationsAdd**](WechatQualificationsApi.md#WechatQualificationsAdd) | **Post** /wechat_qualifications/add | 添加附近推商家资质
[**WechatQualificationsDelete**](WechatQualificationsApi.md#WechatQualificationsDelete) | **Post** /wechat_qualifications/delete | 删除附近推商家资质
[**WechatQualificationsGet**](WechatQualificationsApi.md#WechatQualificationsGet) | **Get** /wechat_qualifications/get | 查询附近推商家资质信息


# **WechatQualificationsAdd**
> WechatQualificationsAddResponse WechatQualificationsAdd(ctx, accountId, wechatQualificationType, wechatQualificationFile, wechatQualificationFileSignature)
添加附近推商家资质

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **wechatQualificationType** | **string**|  | 
  **wechatQualificationFile** | ***os.File**|  | 
  **wechatQualificationFileSignature** | **string**|  | 

### Return type

[**WechatQualificationsAddResponse**](WechatQualificationsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatQualificationsDelete**
> WechatQualificationsDeleteResponse WechatQualificationsDelete(ctx, data)
删除附近推商家资质

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatQualificationsDeleteRequest**](WechatQualificationsDeleteRequest.md)|  | 

### Return type

[**WechatQualificationsDeleteResponse**](WechatQualificationsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatQualificationsGet**
> WechatQualificationsGetResponse WechatQualificationsGet(ctx, accountId, optional)
查询附近推商家资质信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WechatQualificationsApiWechatQualificationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatQualificationsApiWechatQualificationsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatQualificationsGetResponse**](WechatQualificationsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


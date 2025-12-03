# TencentAds\OnlinePreviewQrcodeApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OnlinePreviewQrcodeGet**](OnlinePreviewQrcodeApi.md#OnlinePreviewQrcodeGet) | **Get** /online_preview_qrcode/get | 获取在线预览二维码


# **OnlinePreviewQrcodeGet**
> OnlinePreviewQrcodeGetResponse OnlinePreviewQrcodeGet(ctx, accountId, dynamicCreativeId, optional)
获取在线预览二维码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **dynamicCreativeId** | **int64**|  | 
 **optional** | ***OnlinePreviewQrcodeApiOnlinePreviewQrcodeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OnlinePreviewQrcodeApiOnlinePreviewQrcodeGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userIdType** | **optional.String**|  | 
 **previewCreativeComponents** | [**optional.Interface of PreviewCreativeComponents**](PreviewCreativeComponents.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**OnlinePreviewQrcodeGetResponse**](OnlinePreviewQrcodeGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


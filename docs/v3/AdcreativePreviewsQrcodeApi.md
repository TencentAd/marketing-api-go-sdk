# TencentAds\AdcreativePreviewsQrcodeApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdcreativePreviewsQrcodeGet**](AdcreativePreviewsQrcodeApi.md#AdcreativePreviewsQrcodeGet) | **Get** /adcreative_previews_qrcode/get | 获取广告预览二维码


# **AdcreativePreviewsQrcodeGet**
> AdcreativePreviewsQrcodeGetResponse AdcreativePreviewsQrcodeGet(ctx, accountId, adgroupId, optional)
获取广告预览二维码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **adgroupId** | **int64**|  | 
 **optional** | ***AdcreativePreviewsQrcodeApiAdcreativePreviewsQrcodeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdcreativePreviewsQrcodeApiAdcreativePreviewsQrcodeGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AdcreativePreviewsQrcodeGetResponse**](AdcreativePreviewsQrcodeGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


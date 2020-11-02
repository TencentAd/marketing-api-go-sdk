# TencentAds\CreativevideoLabelApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreativevideoLabelAdd**](CreativevideoLabelApi.md#CreativevideoLabelAdd) | **Post** /creativevideo_label/add | 获取视频创意标签


# **CreativevideoLabelAdd**
> CreativevideoLabelAddResponse CreativevideoLabelAdd(ctx, accountId, signature, filename, videofile)
获取视频创意标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **signature** | **string**|  | 
  **filename** | **string**|  | 
  **videofile** | ***os.File**|  | 

### Return type

[**CreativevideoLabelAddResponse**](CreativevideoLabelAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


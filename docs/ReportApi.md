# TencentAds\ReportApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportConversionsPredict**](ReportApi.md#ReportConversionsPredict) | **Post** /report/conversions_predict | 获取当日转化效果预估数据
[**ReportVideoFrame**](ReportApi.md#ReportVideoFrame) | **Post** /report/video_frame | 视频流失分析接口


# **ReportConversionsPredict**
> ReportConversionsPredictResponse ReportConversionsPredict(ctx, data)
获取当日转化效果预估数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ReportConversionsPredictRequest**](ReportConversionsPredictRequest.md)|  | 

### Return type

[**ReportConversionsPredictResponse**](ReportConversionsPredictResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReportVideoFrame**
> ReportVideoFrameResponse ReportVideoFrame(ctx, data)
视频流失分析接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ReportVideoFrameRequest**](ReportVideoFrameRequest.md)|  | 

### Return type

[**ReportVideoFrameResponse**](ReportVideoFrameResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


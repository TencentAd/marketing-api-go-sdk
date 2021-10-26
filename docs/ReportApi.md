# TencentAds\ReportApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportAdUnion**](ReportApi.md#ReportAdUnion) | **Post** /report/ad_union | 联盟广告位报表接口
[**ReportConversionsPredict**](ReportApi.md#ReportConversionsPredict) | **Post** /report/conversions_predict | 获取当日转化效果预估数据
[**ReportLandingPage**](ReportApi.md#ReportLandingPage) | **Post** /report/landing_page | 落地页报表数据接口
[**ReportVideoFrame**](ReportApi.md#ReportVideoFrame) | **Post** /report/video_frame | 视频流失分析接口


# **ReportAdUnion**
> ReportAdUnionResponse ReportAdUnion(ctx, data)
联盟广告位报表接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ReportAdUnionRequest**](ReportAdUnionRequest.md)|  | 

### Return type

[**ReportAdUnionResponse**](ReportAdUnionResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **ReportLandingPage**
> ReportLandingPageResponse ReportLandingPage(ctx, data)
落地页报表数据接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ReportLandingPageRequest**](ReportLandingPageRequest.md)|  | 

### Return type

[**ReportLandingPageResponse**](ReportLandingPageResponse.md)

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


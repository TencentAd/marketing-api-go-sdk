# TencentAds\ReportApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportConversionsPredict**](ReportApi.md#ReportConversionsPredict) | **Post** /report/conversions_predict | 获取当日转化效果预估数据
[**ReportJdCreativeTemplateHourlyReport**](ReportApi.md#ReportJdCreativeTemplateHourlyReport) | **Post** /report/jd_creative_template_hourly_report | 获取京东创意形式小时报表
[**ReportJdOfflineReportFile**](ReportApi.md#ReportJdOfflineReportFile) | **Get** /report/jd_offline_report_file | 获取京东离线报表文件
[**ReportJdOfflineReportStatus**](ReportApi.md#ReportJdOfflineReportStatus) | **Post** /report/jd_offline_report_status | 获取京东离线报表文件状态
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

# **ReportJdCreativeTemplateHourlyReport**
> ReportJdCreativeTemplateHourlyReportResponse ReportJdCreativeTemplateHourlyReport(ctx, data)
获取京东创意形式小时报表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ReportJdCreativeTemplateHourlyReportRequest**](ReportJdCreativeTemplateHourlyReportRequest.md)|  | 

### Return type

[**ReportJdCreativeTemplateHourlyReportResponse**](ReportJdCreativeTemplateHourlyReportResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReportJdOfflineReportFile**
> ReportJdOfflineReportFileResponse ReportJdOfflineReportFile(ctx, accountId, task, date, optional)
获取京东离线报表文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **task** | **string**|  | 
  **date** | **string**|  | 
 **optional** | ***ReportApiReportJdOfflineReportFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportApiReportJdOfflineReportFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **hour** | **optional.String**|  | 
 **timeFrame** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ReportJdOfflineReportFileResponse**](ReportJdOfflineReportFileResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReportJdOfflineReportStatus**
> ReportJdOfflineReportStatusResponse ReportJdOfflineReportStatus(ctx, data)
获取京东离线报表文件状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ReportJdOfflineReportStatusRequest**](ReportJdOfflineReportStatusRequest.md)|  | 

### Return type

[**ReportJdOfflineReportStatusResponse**](ReportJdOfflineReportStatusResponse.md)

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


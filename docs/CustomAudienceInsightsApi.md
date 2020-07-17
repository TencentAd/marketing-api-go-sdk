# TencentAds\CustomAudienceInsightsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomAudienceInsightsGet**](CustomAudienceInsightsApi.md#CustomAudienceInsightsGet) | **Get** /custom_audience_insights/get | 人群洞察分析


# **CustomAudienceInsightsGet**
> CustomAudienceInsightsGetResponse CustomAudienceInsightsGet(ctx, accountId, audienceId, dimensionType, optional)
人群洞察分析

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **audienceId** | **int64**|  | 
  **dimensionType** | [**[]string**](string.md)|  | 
 **optional** | ***CustomAudienceInsightsApiCustomAudienceInsightsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomAudienceInsightsApiCustomAudienceInsightsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CustomAudienceInsightsGetResponse**](CustomAudienceInsightsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


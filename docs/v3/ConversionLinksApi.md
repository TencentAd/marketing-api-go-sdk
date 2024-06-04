# TencentAds\ConversionLinksApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConversionLinksGet**](ConversionLinksApi.md#ConversionLinksGet) | **Get** /conversion_links/get | 获取营销链路模板


# **ConversionLinksGet**
> ConversionLinksGetResponse ConversionLinksGet(ctx, accountId, secondCategoryType, optional)
获取营销链路模板

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **secondCategoryType** | **string**|  | 
 **optional** | ***ConversionLinksApiConversionLinksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversionLinksApiConversionLinksGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ConversionLinksGetResponse**](ConversionLinksGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


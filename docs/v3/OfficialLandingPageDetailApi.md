# TencentAds\OfficialLandingPageDetailApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficialLandingPageDetailGet**](OfficialLandingPageDetailApi.md#OfficialLandingPageDetailGet) | **Get** /official_landing_page_detail/get | 官方落地页-获取落地页详情


# **OfficialLandingPageDetailGet**
> OfficialLandingPageDetailGetResponse OfficialLandingPageDetailGet(ctx, accountId, pageId, optional)
官方落地页-获取落地页详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **pageId** | **int64**|  | 
 **optional** | ***OfficialLandingPageDetailApiOfficialLandingPageDetailGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficialLandingPageDetailApiOfficialLandingPageDetailGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **protoVersion** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**OfficialLandingPageDetailGetResponse**](OfficialLandingPageDetailGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


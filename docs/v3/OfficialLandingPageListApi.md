# TencentAds\OfficialLandingPageListApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficialLandingPageListGet**](OfficialLandingPageListApi.md#OfficialLandingPageListGet) | **Get** /official_landing_page_list/get | 官方落地页-获取落地页列表


# **OfficialLandingPageListGet**
> OfficialLandingPageListGetResponse OfficialLandingPageListGet(ctx, accountId, optional)
官方落地页-获取落地页列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***OfficialLandingPageListApiOfficialLandingPageListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OfficialLandingPageListApiOfficialLandingPageListGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**OfficialLandingPageListGetResponse**](OfficialLandingPageListGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\AgencyApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgencyGet**](AgencyApi.md#AgencyGet) | **Get** /agency/get | 查询腾讯广告服务商信息


# **AgencyGet**
> AgencyGetResponse AgencyGet(ctx, fields, page, pageSize, optional)
查询腾讯广告服务商信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fields** | [**[]string**](string.md)|  | 
  **page** | **int64**|  | 
  **pageSize** | **int64**|  | 
 **optional** | ***AgencyApiAgencyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgencyApiAgencyGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accountId** | **optional.Int64**|  | 

### Return type

[**AgencyGetResponse**](AgencyGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


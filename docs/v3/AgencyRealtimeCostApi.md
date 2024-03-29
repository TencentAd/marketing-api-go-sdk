# TencentAds\AgencyRealtimeCostApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgencyRealtimeCostGet**](AgencyRealtimeCostApi.md#AgencyRealtimeCostGet) | **Get** /agency_realtime_cost/get | 服务商当日分账户实时消耗


# **AgencyRealtimeCostGet**
> AgencyRealtimeCostGetResponse AgencyRealtimeCostGet(ctx, accountId, optional)
服务商当日分账户实时消耗

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***AgencyRealtimeCostApiAgencyRealtimeCostGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgencyRealtimeCostApiAgencyRealtimeCostGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AgencyRealtimeCostGetResponse**](AgencyRealtimeCostGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


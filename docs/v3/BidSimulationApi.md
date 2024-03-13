# TencentAds\BidSimulationApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BidSimulationGet**](BidSimulationApi.md#BidSimulationGet) | **Get** /bid_simulation/get | 出价模拟预估


# **BidSimulationGet**
> BidSimulationGetResponse BidSimulationGet(ctx, accountId, optional)
出价模拟预估

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***BidSimulationApiBidSimulationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BidSimulationApiBidSimulationGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reqType** | **optional.Int64**|  | 
 **optimizationGoalLevel** | **optional.Int64**|  | 
 **adgroupId** | **optional.Int64**|  | 
 **bidList** | [**optional.Interface of []int64**](int64.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**BidSimulationGetResponse**](BidSimulationGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


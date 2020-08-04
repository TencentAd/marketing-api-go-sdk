# TencentAds\BidSimulationApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BidSimulationGet**](BidSimulationApi.md#BidSimulationGet) | **Post** /bid_simulation/get | 获取出价模拟信息


# **BidSimulationGet**
> BidSimulationGetResponse BidSimulationGet(ctx, data)
获取出价模拟信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BidSimulationGetRequest**](BidSimulationGetRequest.md)|  | 

### Return type

[**BidSimulationGetResponse**](BidSimulationGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


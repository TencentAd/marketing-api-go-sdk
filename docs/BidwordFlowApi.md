# TencentAds\BidwordFlowApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BidwordFlowGet**](BidwordFlowApi.md#BidwordFlowGet) | **Post** /bidword_flow/get | 查询关键词流量接口


# **BidwordFlowGet**
> BidwordFlowGetResponse BidwordFlowGet(ctx, data)
查询关键词流量接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BidwordFlowGetRequest**](BidwordFlowGetRequest.md)|  | 

### Return type

[**BidwordFlowGetResponse**](BidwordFlowGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


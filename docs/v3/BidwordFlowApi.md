# TencentAds\BidwordFlowApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BidwordFlowGet**](BidwordFlowApi.md#BidwordFlowGet) | **Get** /bidword_flow/get | 查询关键词流量接口


# **BidwordFlowGet**
> BidwordFlowGetResponse BidwordFlowGet(ctx, accountId, bidwordList, optional)
查询关键词流量接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **bidwordList** | [**[]string**](string.md)|  | 
 **optional** | ***BidwordFlowApiBidwordFlowGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BidwordFlowApiBidwordFlowGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orderBy** | [**optional.Interface of []OrderByStructInfo**](OrderByStructInfo.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**BidwordFlowGetResponse**](BidwordFlowGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


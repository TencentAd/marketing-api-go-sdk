# TencentAds\AgencyInnerTransferApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgencyInnerTransferAdd**](AgencyInnerTransferApi.md#AgencyInnerTransferAdd) | **Post** /agency_inner_transfer/add | 服务商内部划账


# **AgencyInnerTransferAdd**
> AgencyInnerTransferAddResponse AgencyInnerTransferAdd(ctx, data)
服务商内部划账

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AgencyInnerTransferAddRequest**](AgencyInnerTransferAddRequest.md)|  | 

### Return type

[**AgencyInnerTransferAddResponse**](AgencyInnerTransferAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


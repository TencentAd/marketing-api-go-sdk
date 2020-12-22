# TencentAds\AgencyPeerTransferApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgencyPeerTransferAdd**](AgencyPeerTransferApi.md#AgencyPeerTransferAdd) | **Post** /agency_peer_transfer/add | 服务商之间划账


# **AgencyPeerTransferAdd**
> AgencyPeerTransferAddResponse AgencyPeerTransferAdd(ctx, data)
服务商之间划账

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AgencyPeerTransferAddRequest**](AgencyPeerTransferAddRequest.md)|  | 

### Return type

[**AgencyPeerTransferAddResponse**](AgencyPeerTransferAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\RtaApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RtaInfo**](RtaApi.md#RtaInfo) | **Post** /rta/info | 基本信息查询


# **RtaInfo**
> RtaInfoResponse RtaInfo(ctx, data)
基本信息查询

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RtaInfoRequest**](RtaInfoRequest.md)|  | 

### Return type

[**RtaInfoResponse**](RtaInfoResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


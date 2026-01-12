# TencentAds\PunishDetailApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PunishDetailGet**](PunishDetailApi.md#PunishDetailGet) | **Post** /punish_detail/get | 获取计量处罚明细


# **PunishDetailGet**
> PunishDetailGetResponse PunishDetailGet(ctx, data)
获取计量处罚明细

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**PunishDetailGetRequest**](PunishDetailGetRequest.md)|  | 

### Return type

[**PunishDetailGetResponse**](PunishDetailGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


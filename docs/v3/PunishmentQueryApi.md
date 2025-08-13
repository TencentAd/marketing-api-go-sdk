# TencentAds\PunishmentQueryApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PunishmentQueryGet**](PunishmentQueryApi.md#PunishmentQueryGet) | **Post** /punishment_query/get | 获取违规处罚列表


# **PunishmentQueryGet**
> PunishmentQueryGetResponse PunishmentQueryGet(ctx, data)
获取违规处罚列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**PunishmentQueryGetRequest**](PunishmentQueryGetRequest.md)|  | 

### Return type

[**PunishmentQueryGetResponse**](PunishmentQueryGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


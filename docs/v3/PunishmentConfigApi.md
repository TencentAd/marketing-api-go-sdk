# TencentAds\PunishmentConfigApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PunishmentConfigGet**](PunishmentConfigApi.md#PunishmentConfigGet) | **Get** /punishment_config/get | 获取处罚系统配置


# **PunishmentConfigGet**
> PunishmentConfigGetResponse PunishmentConfigGet(ctx, optional)
获取处罚系统配置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PunishmentConfigApiPunishmentConfigGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PunishmentConfigApiPunishmentConfigGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**PunishmentConfigGetResponse**](PunishmentConfigGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


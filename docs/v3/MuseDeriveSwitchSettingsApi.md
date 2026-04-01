# TencentAds\MuseDeriveSwitchSettingsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MuseDeriveSwitchSettingsGet**](MuseDeriveSwitchSettingsApi.md#MuseDeriveSwitchSettingsGet) | **Get** /muse_derive_switch_settings/get | 创意增强配置


# **MuseDeriveSwitchSettingsGet**
> MuseDeriveSwitchSettingsGetResponse MuseDeriveSwitchSettingsGet(ctx, accountId, marketingTargetType, marketingCarrierType, optional)
创意增强配置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **marketingTargetType** | **string**|  | 
  **marketingCarrierType** | **string**|  | 
 **optional** | ***MuseDeriveSwitchSettingsApiMuseDeriveSwitchSettingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MuseDeriveSwitchSettingsApiMuseDeriveSwitchSettingsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **automaticSiteEnabled** | **optional.Bool**|  | 
 **siteSet** | [**optional.Interface of []string**](string.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**MuseDeriveSwitchSettingsGetResponse**](MuseDeriveSwitchSettingsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\ConversionLinkAssetAvailableApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConversionLinkAssetAvailableGet**](ConversionLinkAssetAvailableApi.md#ConversionLinkAssetAvailableGet) | **Get** /conversion_link_asset_available/get | 获取可投放链路列表


# **ConversionLinkAssetAvailableGet**
> ConversionLinkAssetAvailableGetResponse ConversionLinkAssetAvailableGet(ctx, accountId, marketingGoal, marketingTargetType, optional)
获取可投放链路列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **marketingGoal** | **string**|  | 
  **marketingTargetType** | **string**|  | 
 **optional** | ***ConversionLinkAssetAvailableApiConversionLinkAssetAvailableGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversionLinkAssetAvailableApiConversionLinkAssetAvailableGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **marketingAssetOuterId** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ConversionLinkAssetAvailableGetResponse**](ConversionLinkAssetAvailableGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\MarketingTargetAssetDetailApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketingTargetAssetDetailGet**](MarketingTargetAssetDetailApi.md#MarketingTargetAssetDetailGet) | **Get** /marketing_target_asset_detail/get | 获取推广内容资产详情


# **MarketingTargetAssetDetailGet**
> MarketingTargetAssetDetailGetResponse MarketingTargetAssetDetailGet(ctx, accountId, marketingAssetId, marketingTargetType, optional)
获取推广内容资产详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **marketingAssetId** | **int64**|  | 
  **marketingTargetType** | **string**|  | 
 **optional** | ***MarketingTargetAssetDetailApiMarketingTargetAssetDetailGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketingTargetAssetDetailApiMarketingTargetAssetDetailGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**MarketingTargetAssetDetailGetResponse**](MarketingTargetAssetDetailGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


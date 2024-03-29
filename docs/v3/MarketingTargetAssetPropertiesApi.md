# TencentAds\MarketingTargetAssetPropertiesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketingTargetAssetPropertiesGet**](MarketingTargetAssetPropertiesApi.md#MarketingTargetAssetPropertiesGet) | **Get** /marketing_target_asset_properties/get | 获取可用的推广内容资产属性


# **MarketingTargetAssetPropertiesGet**
> MarketingTargetAssetPropertiesGetResponse MarketingTargetAssetPropertiesGet(ctx, organizationId, marketingTargetType, optional)
获取可用的推广内容资产属性

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationId** | **int64**|  | 
  **marketingTargetType** | **string**|  | 
 **optional** | ***MarketingTargetAssetPropertiesApiMarketingTargetAssetPropertiesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketingTargetAssetPropertiesApiMarketingTargetAssetPropertiesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **marketingAssetType** | **optional.String**|  | 
 **marketingAssetCategory** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**MarketingTargetAssetPropertiesGetResponse**](MarketingTargetAssetPropertiesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


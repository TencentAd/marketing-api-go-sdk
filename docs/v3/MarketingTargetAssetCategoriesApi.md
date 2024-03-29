# TencentAds\MarketingTargetAssetCategoriesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketingTargetAssetCategoriesGet**](MarketingTargetAssetCategoriesApi.md#MarketingTargetAssetCategoriesGet) | **Get** /marketing_target_asset_categories/get | 获取可用的推广内容资产类目


# **MarketingTargetAssetCategoriesGet**
> MarketingTargetAssetCategoriesGetResponse MarketingTargetAssetCategoriesGet(ctx, organizationId, marketingTargetType, optional)
获取可用的推广内容资产类目

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationId** | **int64**|  | 
  **marketingTargetType** | **string**|  | 
 **optional** | ***MarketingTargetAssetCategoriesApiMarketingTargetAssetCategoriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketingTargetAssetCategoriesApiMarketingTargetAssetCategoriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **marketingAssetType** | **optional.String**|  | 
 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**MarketingTargetAssetCategoriesGetResponse**](MarketingTargetAssetCategoriesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


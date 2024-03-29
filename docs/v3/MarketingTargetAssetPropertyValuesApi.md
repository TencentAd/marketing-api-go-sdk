# TencentAds\MarketingTargetAssetPropertyValuesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketingTargetAssetPropertyValuesGet**](MarketingTargetAssetPropertyValuesApi.md#MarketingTargetAssetPropertyValuesGet) | **Get** /marketing_target_asset_property_values/get | 获取可用的推广内容资产属性值


# **MarketingTargetAssetPropertyValuesGet**
> MarketingTargetAssetPropertyValuesGetResponse MarketingTargetAssetPropertyValuesGet(ctx, organizationId, marketingTargetType, optional)
获取可用的推广内容资产属性值

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationId** | **int64**|  | 
  **marketingTargetType** | **string**|  | 
 **optional** | ***MarketingTargetAssetPropertyValuesApiMarketingTargetAssetPropertyValuesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketingTargetAssetPropertyValuesApiMarketingTargetAssetPropertyValuesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **marketingAssetType** | **optional.String**|  | 
 **marketingAssetCategory** | **optional.String**|  | 
 **propertyName** | **optional.String**|  | 
 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**MarketingTargetAssetPropertyValuesGetResponse**](MarketingTargetAssetPropertyValuesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


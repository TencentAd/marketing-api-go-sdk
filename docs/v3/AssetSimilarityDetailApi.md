# TencentAds\AssetSimilarityDetailApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetSimilarityDetailGet**](AssetSimilarityDetailApi.md#AssetSimilarityDetailGet) | **Get** /asset_similarity_detail/get | 获取创意资产相似度检测详情
[**AssetSimilarityDetailUpdate**](AssetSimilarityDetailApi.md#AssetSimilarityDetailUpdate) | **Post** /asset_similarity_detail/update | 更新创意资产相似度


# **AssetSimilarityDetailGet**
> AssetSimilarityDetailGetResponse AssetSimilarityDetailGet(ctx, optional)
获取创意资产相似度检测详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetSimilarityDetailApiAssetSimilarityDetailGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetSimilarityDetailApiAssetSimilarityDetailGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **optional.Int64**|  | 
 **organizationId** | **optional.Int64**|  | 
 **creativeAssetId** | **optional.String**|  | 
 **creativeAssetType** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AssetSimilarityDetailGetResponse**](AssetSimilarityDetailGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetSimilarityDetailUpdate**
> AssetSimilarityDetailUpdateResponse AssetSimilarityDetailUpdate(ctx, data)
更新创意资产相似度

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetSimilarityDetailUpdateRequest**](AssetSimilarityDetailUpdateRequest.md)|  | 

### Return type

[**AssetSimilarityDetailUpdateResponse**](AssetSimilarityDetailUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


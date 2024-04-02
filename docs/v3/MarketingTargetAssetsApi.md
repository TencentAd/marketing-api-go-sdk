# TencentAds\MarketingTargetAssetsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketingTargetAssetsAdd**](MarketingTargetAssetsApi.md#MarketingTargetAssetsAdd) | **Post** /marketing_target_assets/add | 创建推广内容资产
[**MarketingTargetAssetsDelete**](MarketingTargetAssetsApi.md#MarketingTargetAssetsDelete) | **Post** /marketing_target_assets/delete | 删除推广内容资产
[**MarketingTargetAssetsGet**](MarketingTargetAssetsApi.md#MarketingTargetAssetsGet) | **Get** /marketing_target_assets/get | 获取可投放推广内容资产列表
[**MarketingTargetAssetsUpdate**](MarketingTargetAssetsApi.md#MarketingTargetAssetsUpdate) | **Post** /marketing_target_assets/update | 更新推广内容资产


# **MarketingTargetAssetsAdd**
> MarketingTargetAssetsAddResponse MarketingTargetAssetsAdd(ctx, data)
创建推广内容资产

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MarketingTargetAssetsAddRequest**](MarketingTargetAssetsAddRequest.md)|  | 

### Return type

[**MarketingTargetAssetsAddResponse**](MarketingTargetAssetsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarketingTargetAssetsDelete**
> MarketingTargetAssetsDeleteResponse MarketingTargetAssetsDelete(ctx, data)
删除推广内容资产

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MarketingTargetAssetsDeleteRequest**](MarketingTargetAssetsDeleteRequest.md)|  | 

### Return type

[**MarketingTargetAssetsDeleteResponse**](MarketingTargetAssetsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarketingTargetAssetsGet**
> MarketingTargetAssetsGetResponse MarketingTargetAssetsGet(ctx, accountId, marketingTargetType, optional)
获取可投放推广内容资产列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **marketingTargetType** | **string**|  | 
 **optional** | ***MarketingTargetAssetsApiMarketingTargetAssetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketingTargetAssetsApiMarketingTargetAssetsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**MarketingTargetAssetsGetResponse**](MarketingTargetAssetsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarketingTargetAssetsUpdate**
> MarketingTargetAssetsUpdateResponse MarketingTargetAssetsUpdate(ctx, data)
更新推广内容资产

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MarketingTargetAssetsUpdateRequest**](MarketingTargetAssetsUpdateRequest.md)|  | 

### Return type

[**MarketingTargetAssetsUpdateResponse**](MarketingTargetAssetsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


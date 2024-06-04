# TencentAds\ConversionLinkAssetsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConversionLinkAssetsAdd**](ConversionLinkAssetsApi.md#ConversionLinkAssetsAdd) | **Post** /conversion_link_assets/add | 创建营销链路
[**ConversionLinkAssetsGet**](ConversionLinkAssetsApi.md#ConversionLinkAssetsGet) | **Get** /conversion_link_assets/get | 获取营销链路
[**ConversionLinkAssetsUpdate**](ConversionLinkAssetsApi.md#ConversionLinkAssetsUpdate) | **Post** /conversion_link_assets/update | 更新营销链路


# **ConversionLinkAssetsAdd**
> ConversionLinkAssetsAddResponse ConversionLinkAssetsAdd(ctx, data)
创建营销链路

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ConversionLinkAssetsAddRequest**](ConversionLinkAssetsAddRequest.md)|  | 

### Return type

[**ConversionLinkAssetsAddResponse**](ConversionLinkAssetsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversionLinkAssetsGet**
> ConversionLinkAssetsGetResponse ConversionLinkAssetsGet(ctx, accountId, filtering, optional)
获取营销链路

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **filtering** | [**[]FilteringStruct**](FilteringStruct.md)|  | 
 **optional** | ***ConversionLinkAssetsApiConversionLinkAssetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversionLinkAssetsApiConversionLinkAssetsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **key** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ConversionLinkAssetsGetResponse**](ConversionLinkAssetsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversionLinkAssetsUpdate**
> ConversionLinkAssetsUpdateResponse ConversionLinkAssetsUpdate(ctx, data)
更新营销链路

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ConversionLinkAssetsUpdateRequest**](ConversionLinkAssetsUpdateRequest.md)|  | 

### Return type

[**ConversionLinkAssetsUpdateResponse**](ConversionLinkAssetsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


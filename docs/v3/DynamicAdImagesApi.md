# TencentAds\DynamicAdImagesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicAdImagesAdd**](DynamicAdImagesApi.md#DynamicAdImagesAdd) | **Post** /dynamic_ad_images/add | 创建用于广告投放的动态广告图片
[**DynamicAdImagesGet**](DynamicAdImagesApi.md#DynamicAdImagesGet) | **Get** /dynamic_ad_images/get | 获取动态广告图片信息


# **DynamicAdImagesAdd**
> DynamicAdImagesAddResponse DynamicAdImagesAdd(ctx, data)
创建用于广告投放的动态广告图片

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DynamicAdImagesAddRequest**](DynamicAdImagesAddRequest.md)|  | 

### Return type

[**DynamicAdImagesAddResponse**](DynamicAdImagesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicAdImagesGet**
> DynamicAdImagesGetResponse DynamicAdImagesGet(ctx, accountId, optional)
获取动态广告图片信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***DynamicAdImagesApiDynamicAdImagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DynamicAdImagesApiDynamicAdImagesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**DynamicAdImagesGetResponse**](DynamicAdImagesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


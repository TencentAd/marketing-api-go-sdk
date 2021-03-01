# TencentAds\AndroidUnionChannelPackagesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AndroidUnionChannelPackagesAdd**](AndroidUnionChannelPackagesApi.md#AndroidUnionChannelPackagesAdd) | **Post** /android_union_channel_packages/add | 添加广告渠道包
[**AndroidUnionChannelPackagesGet**](AndroidUnionChannelPackagesApi.md#AndroidUnionChannelPackagesGet) | **Get** /android_union_channel_packages/get | 获取广告渠道包
[**AndroidUnionChannelPackagesUpdate**](AndroidUnionChannelPackagesApi.md#AndroidUnionChannelPackagesUpdate) | **Post** /android_union_channel_packages/update | 更新广告渠道包


# **AndroidUnionChannelPackagesAdd**
> AndroidUnionChannelPackagesAddResponse AndroidUnionChannelPackagesAdd(ctx, data)
添加广告渠道包

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AndroidUnionChannelPackagesAddRequest**](AndroidUnionChannelPackagesAddRequest.md)|  | 

### Return type

[**AndroidUnionChannelPackagesAddResponse**](AndroidUnionChannelPackagesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AndroidUnionChannelPackagesGet**
> AndroidUnionChannelPackagesGetResponse AndroidUnionChannelPackagesGet(ctx, accountId, androidUnionAppId, optional)
获取广告渠道包

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **androidUnionAppId** | **int64**|  | 
 **optional** | ***AndroidUnionChannelPackagesApiAndroidUnionChannelPackagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AndroidUnionChannelPackagesApiAndroidUnionChannelPackagesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AndroidUnionChannelPackagesGetResponse**](AndroidUnionChannelPackagesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AndroidUnionChannelPackagesUpdate**
> AndroidUnionChannelPackagesUpdateResponse AndroidUnionChannelPackagesUpdate(ctx, data)
更新广告渠道包

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AndroidUnionChannelPackagesUpdateRequest**](AndroidUnionChannelPackagesUpdateRequest.md)|  | 

### Return type

[**AndroidUnionChannelPackagesUpdateResponse**](AndroidUnionChannelPackagesUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


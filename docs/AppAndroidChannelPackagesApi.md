# TencentAds\AppAndroidChannelPackagesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppAndroidChannelPackagesGet**](AppAndroidChannelPackagesApi.md#AppAndroidChannelPackagesGet) | **Get** /app_android_channel_packages/get | 获取Android渠道包


# **AppAndroidChannelPackagesGet**
> AppAndroidChannelPackagesGetResponse AppAndroidChannelPackagesGet(ctx, accountId, promotedObjectId, promotedObjectType, optional)
获取Android渠道包

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **promotedObjectId** | **string**|  | 
  **promotedObjectType** | **string**|  | 
 **optional** | ***AppAndroidChannelPackagesApiAppAndroidChannelPackagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppAndroidChannelPackagesApiAppAndroidChannelPackagesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AppAndroidChannelPackagesGetResponse**](AppAndroidChannelPackagesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


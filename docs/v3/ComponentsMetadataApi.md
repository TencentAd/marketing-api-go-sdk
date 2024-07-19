# TencentAds\ComponentsMetadataApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentsMetadataGet**](ComponentsMetadataApi.md#ComponentsMetadataGet) | **Get** /components_metadata/get | 查询创意组件元数据


# **ComponentsMetadataGet**
> ComponentsMetadataGetResponse ComponentsMetadataGet(ctx, accountId, optional)
查询创意组件元数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***ComponentsMetadataApiComponentsMetadataGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentsMetadataApiComponentsMetadataGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []GetMetadataFilteringStruct**](GetMetadataFilteringStruct.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ComponentsMetadataGetResponse**](ComponentsMetadataGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\UnionPositionPackagesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UnionPositionPackagesAdd**](UnionPositionPackagesApi.md#UnionPositionPackagesAdd) | **Post** /union_position_packages/add | 联盟流量包模块
[**UnionPositionPackagesDelete**](UnionPositionPackagesApi.md#UnionPositionPackagesDelete) | **Post** /union_position_packages/delete | 联盟流量包删除接口
[**UnionPositionPackagesGet**](UnionPositionPackagesApi.md#UnionPositionPackagesGet) | **Get** /union_position_packages/get | 联盟流量包获取接口
[**UnionPositionPackagesUpdate**](UnionPositionPackagesApi.md#UnionPositionPackagesUpdate) | **Post** /union_position_packages/update | 联盟流量包模块


# **UnionPositionPackagesAdd**
> UnionPositionPackagesAddResponse UnionPositionPackagesAdd(ctx, data)
联盟流量包模块

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**UnionPositionPackagesAddRequest**](UnionPositionPackagesAddRequest.md)|  | 

### Return type

[**UnionPositionPackagesAddResponse**](UnionPositionPackagesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnionPositionPackagesDelete**
> UnionPositionPackagesDeleteResponse UnionPositionPackagesDelete(ctx, data)
联盟流量包删除接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**UnionPositionPackagesDeleteRequest**](UnionPositionPackagesDeleteRequest.md)|  | 

### Return type

[**UnionPositionPackagesDeleteResponse**](UnionPositionPackagesDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnionPositionPackagesGet**
> UnionPositionPackagesGetResponse UnionPositionPackagesGet(ctx, accountId, optional)
联盟流量包获取接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***UnionPositionPackagesApiUnionPositionPackagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UnionPositionPackagesApiUnionPositionPackagesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**UnionPositionPackagesGetResponse**](UnionPositionPackagesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnionPositionPackagesUpdate**
> UnionPositionPackagesUpdateResponse UnionPositionPackagesUpdate(ctx, data)
联盟流量包模块

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**UnionPositionPackagesUpdateRequest**](UnionPositionPackagesUpdateRequest.md)|  | 

### Return type

[**UnionPositionPackagesUpdateResponse**](UnionPositionPackagesUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


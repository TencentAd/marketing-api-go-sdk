# TencentAds\ExtendPackageApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtendPackageAdd**](ExtendPackageApi.md#ExtendPackageAdd) | **Post** /extend_package/add | 创建应用分包
[**ExtendPackageGet**](ExtendPackageApi.md#ExtendPackageGet) | **Get** /extend_package/get | 查询应用分包列表
[**ExtendPackageUpdate**](ExtendPackageApi.md#ExtendPackageUpdate) | **Post** /extend_package/update | 更新应用子包版本


# **ExtendPackageAdd**
> ExtendPackageAddResponse ExtendPackageAdd(ctx, data)
创建应用分包

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ExtendPackageAddRequest**](ExtendPackageAddRequest.md)|  | 

### Return type

[**ExtendPackageAddResponse**](ExtendPackageAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtendPackageGet**
> ExtendPackageGetResponse ExtendPackageGet(ctx, accountId, packageId, optional)
查询应用分包列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **packageId** | **int64**|  | 
 **optional** | ***ExtendPackageApiExtendPackageGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtendPackageApiExtendPackageGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ExtendPackageGetResponse**](ExtendPackageGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtendPackageUpdate**
> ExtendPackageUpdateResponse ExtendPackageUpdate(ctx, data)
更新应用子包版本

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ExtendPackageUpdateRequest**](ExtendPackageUpdateRequest.md)|  | 

### Return type

[**ExtendPackageUpdateResponse**](ExtendPackageUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


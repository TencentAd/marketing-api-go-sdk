# TencentAds\LocalStorePackagesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalStorePackagesAdd**](LocalStorePackagesApi.md#LocalStorePackagesAdd) | **Post** /local_store_packages/add | 创建门店包
[**LocalStorePackagesDelete**](LocalStorePackagesApi.md#LocalStorePackagesDelete) | **Post** /local_store_packages/delete | 删除门店包
[**LocalStorePackagesGet**](LocalStorePackagesApi.md#LocalStorePackagesGet) | **Get** /local_store_packages/get | 查询门店包信息
[**LocalStorePackagesUpdate**](LocalStorePackagesApi.md#LocalStorePackagesUpdate) | **Post** /local_store_packages/update | 更新门店包信息


# **LocalStorePackagesAdd**
> LocalStorePackagesAddResponse LocalStorePackagesAdd(ctx, data)
创建门店包

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalStorePackagesAddRequest**](LocalStorePackagesAddRequest.md)|  | 

### Return type

[**LocalStorePackagesAddResponse**](LocalStorePackagesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalStorePackagesDelete**
> LocalStorePackagesDeleteResponse LocalStorePackagesDelete(ctx, data)
删除门店包

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalStorePackagesDeleteRequest**](LocalStorePackagesDeleteRequest.md)|  | 

### Return type

[**LocalStorePackagesDeleteResponse**](LocalStorePackagesDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalStorePackagesGet**
> LocalStorePackagesGetResponse LocalStorePackagesGet(ctx, accountId, optional)
查询门店包信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***LocalStorePackagesApiLocalStorePackagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalStorePackagesApiLocalStorePackagesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**LocalStorePackagesGetResponse**](LocalStorePackagesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalStorePackagesUpdate**
> LocalStorePackagesUpdateResponse LocalStorePackagesUpdate(ctx, data)
更新门店包信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalStorePackagesUpdateRequest**](LocalStorePackagesUpdateRequest.md)|  | 

### Return type

[**LocalStorePackagesUpdateResponse**](LocalStorePackagesUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


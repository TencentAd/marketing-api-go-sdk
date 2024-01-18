# TencentAds\LocalStoresApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalStoresAdd**](LocalStoresApi.md#LocalStoresAdd) | **Post** /local_stores/add | 批量录入门店
[**LocalStoresDelete**](LocalStoresApi.md#LocalStoresDelete) | **Post** /local_stores/delete | 批量删除门店信息
[**LocalStoresGet**](LocalStoresApi.md#LocalStoresGet) | **Get** /local_stores/get | 查询门店信息
[**LocalStoresUpdate**](LocalStoresApi.md#LocalStoresUpdate) | **Post** /local_stores/update | 更新门店信息


# **LocalStoresAdd**
> LocalStoresAddResponse LocalStoresAdd(ctx, data)
批量录入门店

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalStoresAddRequest**](LocalStoresAddRequest.md)|  | 

### Return type

[**LocalStoresAddResponse**](LocalStoresAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalStoresDelete**
> LocalStoresDeleteResponse LocalStoresDelete(ctx, data)
批量删除门店信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalStoresDeleteRequest**](LocalStoresDeleteRequest.md)|  | 

### Return type

[**LocalStoresDeleteResponse**](LocalStoresDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalStoresGet**
> LocalStoresGetResponse LocalStoresGet(ctx, accountId, optional)
查询门店信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***LocalStoresApiLocalStoresGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalStoresApiLocalStoresGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**LocalStoresGetResponse**](LocalStoresGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalStoresUpdate**
> LocalStoresUpdateResponse LocalStoresUpdate(ctx, data)
更新门店信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalStoresUpdateRequest**](LocalStoresUpdateRequest.md)|  | 

### Return type

[**LocalStoresUpdateResponse**](LocalStoresUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\ComponentsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentsAdd**](ComponentsApi.md#ComponentsAdd) | **Post** /components/add | 创建创意组件
[**ComponentsDelete**](ComponentsApi.md#ComponentsDelete) | **Post** /components/delete | 删除创意组件
[**ComponentsGet**](ComponentsApi.md#ComponentsGet) | **Get** /components/get | 获取创意组件


# **ComponentsAdd**
> ComponentsAddResponse ComponentsAdd(ctx, data)
创建创意组件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ComponentsAddRequest**](ComponentsAddRequest.md)|  | 

### Return type

[**ComponentsAddResponse**](ComponentsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComponentsDelete**
> ComponentsDeleteResponse ComponentsDelete(ctx, data)
删除创意组件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ComponentsDeleteRequest**](ComponentsDeleteRequest.md)|  | 

### Return type

[**ComponentsDeleteResponse**](ComponentsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComponentsGet**
> ComponentsGetResponse ComponentsGet(ctx, optional)
获取创意组件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ComponentsApiComponentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentsApiComponentsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **optional.Int64**|  | 
 **organizationId** | **optional.Int64**|  | 
 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **isDeleted** | **optional.Bool**|  | 
 **fields** | [**optional.Interface of []string**](string.md)|  | 
 **componentIdFilteringMode** | **optional.String**|  | 

### Return type

[**ComponentsGetResponse**](ComponentsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


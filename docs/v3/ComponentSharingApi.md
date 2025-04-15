# TencentAds\ComponentSharingApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentSharingAdd**](ComponentSharingApi.md#ComponentSharingAdd) | **Post** /component_sharing/add | 共享创意组件
[**ComponentSharingGet**](ComponentSharingApi.md#ComponentSharingGet) | **Get** /component_sharing/get | 查询创意组件共享信息
[**ComponentSharingUpdate**](ComponentSharingApi.md#ComponentSharingUpdate) | **Post** /component_sharing/update | 修改创意组件共享


# **ComponentSharingAdd**
> ComponentSharingAddResponse ComponentSharingAdd(ctx, data)
共享创意组件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ComponentSharingAddRequest**](ComponentSharingAddRequest.md)|  | 

### Return type

[**ComponentSharingAddResponse**](ComponentSharingAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComponentSharingGet**
> ComponentSharingGetResponse ComponentSharingGet(ctx, organizationId, optional)
查询创意组件共享信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationId** | **int64**|  | 
 **optional** | ***ComponentSharingApiComponentSharingGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentSharingApiComponentSharingGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentId** | **optional.Int64**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **isDeleted** | **optional.Bool**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ComponentSharingGetResponse**](ComponentSharingGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComponentSharingUpdate**
> ComponentSharingUpdateResponse ComponentSharingUpdate(ctx, data)
修改创意组件共享

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ComponentSharingUpdateRequest**](ComponentSharingUpdateRequest.md)|  | 

### Return type

[**ComponentSharingUpdateResponse**](ComponentSharingUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


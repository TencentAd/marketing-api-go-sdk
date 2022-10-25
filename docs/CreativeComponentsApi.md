# TencentAds\CreativeComponentsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreativeComponentsAdd**](CreativeComponentsApi.md#CreativeComponentsAdd) | **Post** /creative_components/add | 创建创意组件
[**CreativeComponentsDelete**](CreativeComponentsApi.md#CreativeComponentsDelete) | **Post** /creative_components/delete | 删除创意组件
[**CreativeComponentsGet**](CreativeComponentsApi.md#CreativeComponentsGet) | **Get** /creative_components/get | 查询创意组件信息
[**CreativeComponentsUpdate**](CreativeComponentsApi.md#CreativeComponentsUpdate) | **Post** /creative_components/update | 更新创意组件 
[**CreativeComponentsUpdateStatus**](CreativeComponentsApi.md#CreativeComponentsUpdateStatus) | **Post** /creative_components/update_status | 更新创意组件状态


# **CreativeComponentsAdd**
> CreativeComponentsAddResponse CreativeComponentsAdd(ctx, data)
创建创意组件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CreativeComponentsAddRequest**](CreativeComponentsAddRequest.md)|  | 

### Return type

[**CreativeComponentsAddResponse**](CreativeComponentsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreativeComponentsDelete**
> CreativeComponentsDeleteResponse CreativeComponentsDelete(ctx, data)
删除创意组件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CreativeComponentsDeleteRequest**](CreativeComponentsDeleteRequest.md)|  | 

### Return type

[**CreativeComponentsDeleteResponse**](CreativeComponentsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreativeComponentsGet**
> CreativeComponentsGetResponse CreativeComponentsGet(ctx, accountId, optional)
查询创意组件信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***CreativeComponentsApiCreativeComponentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreativeComponentsApiCreativeComponentsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CreativeComponentsGetResponse**](CreativeComponentsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreativeComponentsUpdate**
> CreativeComponentsUpdateResponse CreativeComponentsUpdate(ctx, data)
更新创意组件 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CreativeComponentsUpdateRequest**](CreativeComponentsUpdateRequest.md)|  | 

### Return type

[**CreativeComponentsUpdateResponse**](CreativeComponentsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreativeComponentsUpdateStatus**
> CreativeComponentsUpdateStatusResponse CreativeComponentsUpdateStatus(ctx, data)
更新创意组件状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CreativeComponentsUpdateStatusRequest**](CreativeComponentsUpdateStatusRequest.md)|  | 

### Return type

[**CreativeComponentsUpdateStatusResponse**](CreativeComponentsUpdateStatusResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


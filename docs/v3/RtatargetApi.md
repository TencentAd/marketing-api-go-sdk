# TencentAds\RtatargetApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RtatargetAdd**](RtatargetApi.md#RtatargetAdd) | **Post** /rtatarget/add | 新增RTA策略
[**RtatargetDelete**](RtatargetApi.md#RtatargetDelete) | **Post** /rtatarget/delete | 删除RTA策略
[**RtatargetGet**](RtatargetApi.md#RtatargetGet) | **Post** /rtatarget/get | 获取RTA策略


# **RtatargetAdd**
> RtatargetAddResponse RtatargetAdd(ctx, data)
新增RTA策略

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RtatargetAddRequest**](RtatargetAddRequest.md)|  | 

### Return type

[**RtatargetAddResponse**](RtatargetAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RtatargetDelete**
> RtatargetDeleteResponse RtatargetDelete(ctx, data)
删除RTA策略

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RtatargetDeleteRequest**](RtatargetDeleteRequest.md)|  | 

### Return type

[**RtatargetDeleteResponse**](RtatargetDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RtatargetGet**
> RtatargetGetResponse RtatargetGet(ctx, data)
获取RTA策略

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RtatargetGetRequest**](RtatargetGetRequest.md)|  | 

### Return type

[**RtatargetGetResponse**](RtatargetGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


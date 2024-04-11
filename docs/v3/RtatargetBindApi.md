# TencentAds\RtatargetBindApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RtatargetBindAdd**](RtatargetBindApi.md#RtatargetBindAdd) | **Post** /rtatarget_bind/add | 新增策略绑定
[**RtatargetBindDelete**](RtatargetBindApi.md#RtatargetBindDelete) | **Post** /rtatarget_bind/delete | 解除策略绑定
[**RtatargetBindGet**](RtatargetBindApi.md#RtatargetBindGet) | **Post** /rtatarget_bind/get | 查询已绑定ID列表


# **RtatargetBindAdd**
> RtatargetBindAddResponse RtatargetBindAdd(ctx, data)
新增策略绑定

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RtatargetBindAddRequest**](RtatargetBindAddRequest.md)|  | 

### Return type

[**RtatargetBindAddResponse**](RtatargetBindAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RtatargetBindDelete**
> RtatargetBindDeleteResponse RtatargetBindDelete(ctx, data)
解除策略绑定

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RtatargetBindDeleteRequest**](RtatargetBindDeleteRequest.md)|  | 

### Return type

[**RtatargetBindDeleteResponse**](RtatargetBindDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RtatargetBindGet**
> RtatargetBindGetResponse RtatargetBindGet(ctx, data)
查询已绑定ID列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**RtatargetBindGetRequest**](RtatargetBindGetRequest.md)|  | 

### Return type

[**RtatargetBindGetResponse**](RtatargetBindGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


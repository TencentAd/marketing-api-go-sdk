# TencentAds\ShopApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShopAdd**](ShopApi.md#ShopAdd) | **Post** /shop/add | 创建广告
[**ShopDelete**](ShopApi.md#ShopDelete) | **Post** /shop/delete | 删除广告
[**ShopGet**](ShopApi.md#ShopGet) | **Post** /shop/get | 拉取广告列表
[**ShopUpdate**](ShopApi.md#ShopUpdate) | **Post** /shop/update | 更新广告


# **ShopAdd**
> ShopAddResponse ShopAdd(ctx, data)
创建广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ShopAddRequest**](ShopAddRequest.md)|  | 

### Return type

[**ShopAddResponse**](ShopAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShopDelete**
> ShopDeleteResponse ShopDelete(ctx, data)
删除广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ShopDeleteRequest**](ShopDeleteRequest.md)|  | 

### Return type

[**ShopDeleteResponse**](ShopDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShopGet**
> ShopGetResponse ShopGet(ctx, data)
拉取广告列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ShopGetRequest**](ShopGetRequest.md)|  | 

### Return type

[**ShopGetResponse**](ShopGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShopUpdate**
> ShopUpdateResponse ShopUpdate(ctx, data)
更新广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ShopUpdateRequest**](ShopUpdateRequest.md)|  | 

### Return type

[**ShopUpdateResponse**](ShopUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


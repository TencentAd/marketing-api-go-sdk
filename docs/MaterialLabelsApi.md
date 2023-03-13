# TencentAds\MaterialLabelsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MaterialLabelsAdd**](MaterialLabelsApi.md#MaterialLabelsAdd) | **Post** /material_labels/add | 素材标签保存
[**MaterialLabelsGet**](MaterialLabelsApi.md#MaterialLabelsGet) | **Post** /material_labels/get | 素材标签列表


# **MaterialLabelsAdd**
> MaterialLabelsAddResponse MaterialLabelsAdd(ctx, data)
素材标签保存

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MaterialLabelsAddRequest**](MaterialLabelsAddRequest.md)|  | 

### Return type

[**MaterialLabelsAddResponse**](MaterialLabelsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MaterialLabelsGet**
> MaterialLabelsGetResponse MaterialLabelsGet(ctx, data)
素材标签列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MaterialLabelsGetRequest**](MaterialLabelsGetRequest.md)|  | 

### Return type

[**MaterialLabelsGetResponse**](MaterialLabelsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


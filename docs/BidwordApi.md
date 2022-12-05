# TencentAds\BidwordApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BidwordAdd**](BidwordApi.md#BidwordAdd) | **Post** /bidword/add | 创建关键词
[**BidwordDelete**](BidwordApi.md#BidwordDelete) | **Post** /bidword/delete | 删除关键词
[**BidwordGet**](BidwordApi.md#BidwordGet) | **Post** /bidword/get | 获取关键词
[**BidwordUpdate**](BidwordApi.md#BidwordUpdate) | **Post** /bidword/update | 更新关键词


# **BidwordAdd**
> BidwordAddResponse BidwordAdd(ctx, data)
创建关键词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BidwordAddRequest**](BidwordAddRequest.md)|  | 

### Return type

[**BidwordAddResponse**](BidwordAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BidwordDelete**
> BidwordDeleteResponse BidwordDelete(ctx, data)
删除关键词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BidwordDeleteRequest**](BidwordDeleteRequest.md)|  | 

### Return type

[**BidwordDeleteResponse**](BidwordDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BidwordGet**
> BidwordGetResponse BidwordGet(ctx, data)
获取关键词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BidwordGetRequest**](BidwordGetRequest.md)|  | 

### Return type

[**BidwordGetResponse**](BidwordGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BidwordUpdate**
> BidwordUpdateResponse BidwordUpdate(ctx, data)
更新关键词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BidwordUpdateRequest**](BidwordUpdateRequest.md)|  | 

### Return type

[**BidwordUpdateResponse**](BidwordUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


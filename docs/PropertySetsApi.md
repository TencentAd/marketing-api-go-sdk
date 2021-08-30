# TencentAds\PropertySetsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PropertySetsAdd**](PropertySetsApi.md#PropertySetsAdd) | **Post** /property_sets/add | 创建属性数据源
[**PropertySetsGet**](PropertySetsApi.md#PropertySetsGet) | **Get** /property_sets/get | 获取属性数据源


# **PropertySetsAdd**
> PropertySetsAddResponse PropertySetsAdd(ctx, data)
创建属性数据源

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**PropertySetsAddRequest**](PropertySetsAddRequest.md)|  | 

### Return type

[**PropertySetsAddResponse**](PropertySetsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PropertySetsGet**
> PropertySetsGetResponse PropertySetsGet(ctx, accountId, optional)
获取属性数据源

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***PropertySetsApiPropertySetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PropertySetsApiPropertySetsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertySetId** | **optional.Int64**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**PropertySetsGetResponse**](PropertySetsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\PropertySetSchemasApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PropertySetSchemasAdd**](PropertySetSchemasApi.md#PropertySetSchemasAdd) | **Post** /property_set_schemas/add | 创建属性数据源schema
[**PropertySetSchemasGet**](PropertySetSchemasApi.md#PropertySetSchemasGet) | **Get** /property_set_schemas/get | 获取属性数据源Schema
[**PropertySetSchemasUpdate**](PropertySetSchemasApi.md#PropertySetSchemasUpdate) | **Post** /property_set_schemas/update | 更新属性数据源schema


# **PropertySetSchemasAdd**
> PropertySetSchemasAddResponse PropertySetSchemasAdd(ctx, data)
创建属性数据源schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**PropertySetSchemasAddRequest**](PropertySetSchemasAddRequest.md)|  | 

### Return type

[**PropertySetSchemasAddResponse**](PropertySetSchemasAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PropertySetSchemasGet**
> PropertySetSchemasGetResponse PropertySetSchemasGet(ctx, accountId, propertySetId, optional)
获取属性数据源Schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **propertySetId** | **int64**|  | 
 **optional** | ***PropertySetSchemasApiPropertySetSchemasGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PropertySetSchemasApiPropertySetSchemasGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**PropertySetSchemasGetResponse**](PropertySetSchemasGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PropertySetSchemasUpdate**
> PropertySetSchemasUpdateResponse PropertySetSchemasUpdate(ctx, data)
更新属性数据源schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**PropertySetSchemasUpdateRequest**](PropertySetSchemasUpdateRequest.md)|  | 

### Return type

[**PropertySetSchemasUpdateResponse**](PropertySetSchemasUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


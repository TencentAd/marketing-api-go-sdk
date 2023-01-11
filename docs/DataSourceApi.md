# TencentAds\DataSourceApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataSourceAdd**](DataSourceApi.md#DataSourceAdd) | **Post** /data_source/add | 添加数据源
[**DataSourceGet**](DataSourceApi.md#DataSourceGet) | **Get** /data_source/get | 查询数据源


# **DataSourceAdd**
> DataSourceAddResponse DataSourceAdd(ctx, data)
添加数据源

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DataSourceAddRequest**](DataSourceAddRequest.md)|  | 

### Return type

[**DataSourceAddResponse**](DataSourceAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataSourceGet**
> DataSourceGetResponse DataSourceGet(ctx, accountId, optional)
查询数据源

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***DataSourceApiDataSourceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataSourceApiDataSourceGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataSourceId** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**DataSourceGetResponse**](DataSourceGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


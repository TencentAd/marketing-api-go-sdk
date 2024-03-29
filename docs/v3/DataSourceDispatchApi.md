# TencentAds\DataSourceDispatchApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataSourceDispatchGet**](DataSourceDispatchApi.md#DataSourceDispatchGet) | **Get** /data_source_dispatch/get | 数据源分发关系获取


# **DataSourceDispatchGet**
> DataSourceDispatchGetResponse DataSourceDispatchGet(ctx, accountId, optional)
数据源分发关系获取

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***DataSourceDispatchApiDataSourceDispatchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataSourceDispatchApiDataSourceDispatchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userActionSetId** | **optional.Int64**|  | 
 **type_** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **scenes** | [**optional.Interface of []string**](string.md)|  | 
 **switchType** | **optional.String**|  | 
 **accessWay** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**DataSourceDispatchGetResponse**](DataSourceDispatchGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


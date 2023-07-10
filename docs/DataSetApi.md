# TencentAds\DataSetApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataSetAdd**](DataSetApi.md#DataSetAdd) | **Post** /data_set/add | 添加数据集
[**DataSetGet**](DataSetApi.md#DataSetGet) | **Get** /data_set/get | 获取数据集


# **DataSetAdd**
> DataSetAddResponse DataSetAdd(ctx, data)
添加数据集

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**DataSetAddRequest**](DataSetAddRequest.md)|  | 

### Return type

[**DataSetAddResponse**](DataSetAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataSetGet**
> DataSetGetResponse DataSetGet(ctx, accountId, optional)
获取数据集

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***DataSetApiDataSetGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataSetApiDataSetGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userActionSetId** | **optional.Int64**|  | 
 **dataSetId** | **optional.Int64**|  | 
 **envType** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**DataSetGetResponse**](DataSetGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\AdgroupNegativewordsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdgroupNegativewordsAdd**](AdgroupNegativewordsApi.md#AdgroupNegativewordsAdd) | **Post** /adgroup_negativewords/add | 新增广告组否定词
[**AdgroupNegativewordsGet**](AdgroupNegativewordsApi.md#AdgroupNegativewordsGet) | **Get** /adgroup_negativewords/get | 查询广告组否定词
[**AdgroupNegativewordsUpdate**](AdgroupNegativewordsApi.md#AdgroupNegativewordsUpdate) | **Post** /adgroup_negativewords/update | 修改广告组否定词


# **AdgroupNegativewordsAdd**
> AdgroupNegativewordsAddResponse AdgroupNegativewordsAdd(ctx, data)
新增广告组否定词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdgroupNegativewordsAddRequest**](AdgroupNegativewordsAddRequest.md)|  | 

### Return type

[**AdgroupNegativewordsAddResponse**](AdgroupNegativewordsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdgroupNegativewordsGet**
> AdgroupNegativewordsGetResponse AdgroupNegativewordsGet(ctx, accountId, adgroupIds, optional)
查询广告组否定词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **adgroupIds** | [**[]int64**](int64.md)|  | 
 **optional** | ***AdgroupNegativewordsApiAdgroupNegativewordsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdgroupNegativewordsApiAdgroupNegativewordsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AdgroupNegativewordsGetResponse**](AdgroupNegativewordsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdgroupNegativewordsUpdate**
> AdgroupNegativewordsUpdateResponse AdgroupNegativewordsUpdate(ctx, data)
修改广告组否定词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdgroupNegativewordsUpdateRequest**](AdgroupNegativewordsUpdateRequest.md)|  | 

### Return type

[**AdgroupNegativewordsUpdateResponse**](AdgroupNegativewordsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


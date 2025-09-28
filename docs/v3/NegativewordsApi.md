# TencentAds\NegativewordsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NegativewordsAdd**](NegativewordsApi.md#NegativewordsAdd) | **Post** /negativewords/add | 新增否定词，可在广告、创意上新增
[**NegativewordsGet**](NegativewordsApi.md#NegativewordsGet) | **Get** /negativewords/get | 查询否定词，可查询广告、创意的否词
[**NegativewordsUpdate**](NegativewordsApi.md#NegativewordsUpdate) | **Post** /negativewords/update | 更新否定词，可在广告、创意上更新


# **NegativewordsAdd**
> NegativewordsAddResponse NegativewordsAdd(ctx, data)
新增否定词，可在广告、创意上新增

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**NegativewordsAddRequest**](NegativewordsAddRequest.md)|  | 

### Return type

[**NegativewordsAddResponse**](NegativewordsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NegativewordsGet**
> NegativewordsGetResponse NegativewordsGet(ctx, accountId, optional)
查询否定词，可查询广告、创意的否词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***NegativewordsApiNegativewordsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegativewordsApiNegativewordsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adgroupIds** | [**optional.Interface of []int64**](int64.md)|  | 
 **dynamicCreativeIds** | [**optional.Interface of []int64**](int64.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**NegativewordsGetResponse**](NegativewordsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NegativewordsUpdate**
> NegativewordsUpdateResponse NegativewordsUpdate(ctx, data)
更新否定词，可在广告、创意上更新

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**NegativewordsUpdateRequest**](NegativewordsUpdateRequest.md)|  | 

### Return type

[**NegativewordsUpdateResponse**](NegativewordsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


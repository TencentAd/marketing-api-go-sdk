# TencentAds\AdsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdsAdd**](AdsApi.md#AdsAdd) | **Post** /ads/add | 创建广告
[**AdsDelete**](AdsApi.md#AdsDelete) | **Post** /ads/delete | 删除广告
[**AdsGet**](AdsApi.md#AdsGet) | **Get** /ads/get | 获取广告
[**AdsUpdate**](AdsApi.md#AdsUpdate) | **Post** /ads/update | 更新广告
[**AdsUpdateConfiguredStatus**](AdsApi.md#AdsUpdateConfiguredStatus) | **Post** /ads/update_configured_status | 更新广告状态


# **AdsAdd**
> AdsAddResponse AdsAdd(ctx, data)
创建广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdsAddRequest**](AdsAddRequest.md)|  | 

### Return type

[**AdsAddResponse**](AdsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdsDelete**
> AdsDeleteResponse AdsDelete(ctx, data)
删除广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdsDeleteRequest**](AdsDeleteRequest.md)|  | 

### Return type

[**AdsDeleteResponse**](AdsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdsGet**
> AdsGetResponse AdsGet(ctx, accountId, optional)
获取广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***AdsApiAdsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdsApiAdsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **isDeleted** | **optional.Bool**|  | 
 **weixinOfficialAccountsUpgradeEnabled** | **optional.Bool**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AdsGetResponse**](AdsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdsUpdate**
> AdsUpdateResponse AdsUpdate(ctx, data)
更新广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdsUpdateRequest**](AdsUpdateRequest.md)|  | 

### Return type

[**AdsUpdateResponse**](AdsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdsUpdateConfiguredStatus**
> AdsUpdateConfiguredStatusResponse AdsUpdateConfiguredStatus(ctx, data)
更新广告状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdsUpdateConfiguredStatusRequest**](AdsUpdateConfiguredStatusRequest.md)|  | 

### Return type

[**AdsUpdateConfiguredStatusResponse**](AdsUpdateConfiguredStatusResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


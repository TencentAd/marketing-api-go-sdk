# TencentAds\AdcreativesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdcreativesAdd**](AdcreativesApi.md#AdcreativesAdd) | **Post** /adcreatives/add | 创建广告创意
[**AdcreativesDelete**](AdcreativesApi.md#AdcreativesDelete) | **Post** /adcreatives/delete | 删除广告创意
[**AdcreativesGet**](AdcreativesApi.md#AdcreativesGet) | **Get** /adcreatives/get | 获取广告创意
[**AdcreativesUpdate**](AdcreativesApi.md#AdcreativesUpdate) | **Post** /adcreatives/update | 更新广告创意


# **AdcreativesAdd**
> AdcreativesAddResponse AdcreativesAdd(ctx, data)
创建广告创意

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdcreativesAddRequest**](AdcreativesAddRequest.md)|  | 

### Return type

[**AdcreativesAddResponse**](AdcreativesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdcreativesDelete**
> AdcreativesDeleteResponse AdcreativesDelete(ctx, data)
删除广告创意

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdcreativesDeleteRequest**](AdcreativesDeleteRequest.md)|  | 

### Return type

[**AdcreativesDeleteResponse**](AdcreativesDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdcreativesGet**
> AdcreativesGetResponse AdcreativesGet(ctx, accountId, optional)
获取广告创意

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***AdcreativesApiAdcreativesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdcreativesApiAdcreativesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **isDeleted** | **optional.Bool**|  | 
 **linkPageTypeCompatible** | **optional.Bool**|  | 
 **weixinOfficialAccountsUpgradeEnabled** | **optional.Bool**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AdcreativesGetResponse**](AdcreativesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdcreativesUpdate**
> AdcreativesUpdateResponse AdcreativesUpdate(ctx, data)
更新广告创意

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdcreativesUpdateRequest**](AdcreativesUpdateRequest.md)|  | 

### Return type

[**AdcreativesUpdateResponse**](AdcreativesUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


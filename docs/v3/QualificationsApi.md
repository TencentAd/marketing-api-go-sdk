# TencentAds\QualificationsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QualificationsAdd**](QualificationsApi.md#QualificationsAdd) | **Post** /qualifications/add | 创建广告主资质
[**QualificationsDelete**](QualificationsApi.md#QualificationsDelete) | **Post** /qualifications/delete | 删除广告主资质
[**QualificationsGet**](QualificationsApi.md#QualificationsGet) | **Get** /qualifications/get | 获取广告主资质
[**QualificationsUpdate**](QualificationsApi.md#QualificationsUpdate) | **Post** /qualifications/update | 更新广告主资质


# **QualificationsAdd**
> QualificationsAddResponse QualificationsAdd(ctx, data)
创建广告主资质

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**QualificationsAddRequest**](QualificationsAddRequest.md)|  | 

### Return type

[**QualificationsAddResponse**](QualificationsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QualificationsDelete**
> QualificationsDeleteResponse QualificationsDelete(ctx, data)
删除广告主资质

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**QualificationsDeleteRequest**](QualificationsDeleteRequest.md)|  | 

### Return type

[**QualificationsDeleteResponse**](QualificationsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QualificationsGet**
> QualificationsGetResponse QualificationsGet(ctx, accountId, qualificationType, optional)
获取广告主资质

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **qualificationType** | **string**|  | 
 **optional** | ***QualificationsApiQualificationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualificationsApiQualificationsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**QualificationsGetResponse**](QualificationsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QualificationsUpdate**
> QualificationsUpdateResponse QualificationsUpdate(ctx, data)
更新广告主资质

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**QualificationsUpdateRequest**](QualificationsUpdateRequest.md)|  | 

### Return type

[**QualificationsUpdateResponse**](QualificationsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


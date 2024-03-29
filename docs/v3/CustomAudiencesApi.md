# TencentAds\CustomAudiencesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomAudiencesAdd**](CustomAudiencesApi.md#CustomAudiencesAdd) | **Post** /custom_audiences/add | 创建客户人群
[**CustomAudiencesDelete**](CustomAudiencesApi.md#CustomAudiencesDelete) | **Post** /custom_audiences/delete | 删除客户人群
[**CustomAudiencesGet**](CustomAudiencesApi.md#CustomAudiencesGet) | **Get** /custom_audiences/get | 获取客户人群
[**CustomAudiencesUpdate**](CustomAudiencesApi.md#CustomAudiencesUpdate) | **Post** /custom_audiences/update | 更新客户人群


# **CustomAudiencesAdd**
> CustomAudiencesAddResponse CustomAudiencesAdd(ctx, data)
创建客户人群

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CustomAudiencesAddRequest**](CustomAudiencesAddRequest.md)|  | 

### Return type

[**CustomAudiencesAddResponse**](CustomAudiencesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomAudiencesDelete**
> CustomAudiencesDeleteResponse CustomAudiencesDelete(ctx, data)
删除客户人群

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CustomAudiencesDeleteRequest**](CustomAudiencesDeleteRequest.md)|  | 

### Return type

[**CustomAudiencesDeleteResponse**](CustomAudiencesDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomAudiencesGet**
> CustomAudiencesGetResponse CustomAudiencesGet(ctx, accountId, optional)
获取客户人群

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***CustomAudiencesApiCustomAudiencesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomAudiencesApiCustomAudiencesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **audienceId** | **optional.Int64**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CustomAudiencesGetResponse**](CustomAudiencesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomAudiencesUpdate**
> CustomAudiencesUpdateResponse CustomAudiencesUpdate(ctx, data)
更新客户人群

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CustomAudiencesUpdateRequest**](CustomAudiencesUpdateRequest.md)|  | 

### Return type

[**CustomAudiencesUpdateResponse**](CustomAudiencesUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


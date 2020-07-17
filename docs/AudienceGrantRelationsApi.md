# TencentAds\AudienceGrantRelationsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AudienceGrantRelationsAdd**](AudienceGrantRelationsApi.md#AudienceGrantRelationsAdd) | **Post** /audience_grant_relations/add | 添加人群授权
[**AudienceGrantRelationsGet**](AudienceGrantRelationsApi.md#AudienceGrantRelationsGet) | **Get** /audience_grant_relations/get | 获取人群授权信息


# **AudienceGrantRelationsAdd**
> AudienceGrantRelationsAddResponse AudienceGrantRelationsAdd(ctx, data)
添加人群授权

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AudienceGrantRelationsAddRequest**](AudienceGrantRelationsAddRequest.md)|  | 

### Return type

[**AudienceGrantRelationsAddResponse**](AudienceGrantRelationsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AudienceGrantRelationsGet**
> AudienceGrantRelationsGetResponse AudienceGrantRelationsGet(ctx, accountId, optional)
获取人群授权信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***AudienceGrantRelationsApiAudienceGrantRelationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AudienceGrantRelationsApiAudienceGrantRelationsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AudienceGrantRelationsGetResponse**](AudienceGrantRelationsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


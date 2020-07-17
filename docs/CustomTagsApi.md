# TencentAds\CustomTagsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomTagsAdd**](CustomTagsApi.md#CustomTagsAdd) | **Post** /custom_tags/add | 创建客户标签
[**CustomTagsDelete**](CustomTagsApi.md#CustomTagsDelete) | **Post** /custom_tags/delete | 删除客户标签
[**CustomTagsGet**](CustomTagsApi.md#CustomTagsGet) | **Get** /custom_tags/get | 获取客户标签
[**CustomTagsUpdate**](CustomTagsApi.md#CustomTagsUpdate) | **Post** /custom_tags/update | 更新客户标签


# **CustomTagsAdd**
> CustomTagsAddResponse CustomTagsAdd(ctx, data)
创建客户标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CustomTagsAddRequest**](CustomTagsAddRequest.md)|  | 

### Return type

[**CustomTagsAddResponse**](CustomTagsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomTagsDelete**
> CustomTagsDeleteResponse CustomTagsDelete(ctx, data)
删除客户标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CustomTagsDeleteRequest**](CustomTagsDeleteRequest.md)|  | 

### Return type

[**CustomTagsDeleteResponse**](CustomTagsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomTagsGet**
> CustomTagsGetResponse CustomTagsGet(ctx, accountId, optional)
获取客户标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***CustomTagsApiCustomTagsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomTagsApiCustomTagsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentTagId** | **optional.Int64**| 父节点Id，parent_tag_id，tag_code，tag_id只能有一个存在 | 
 **tagId** | **optional.Int64**| 标签id，parent_tag_id，tag_code，tag_id只能有一个存在 | 
 **tagCode** | **optional.String**| 广告主对标签在自己系统里的编码，parent_tag_id，tag_code，tag_id只能有一个存在 | 
 **platform** | **optional.String**| 数据应用，不填写默认为DMP | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CustomTagsGetResponse**](CustomTagsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomTagsUpdate**
> CustomTagsUpdateResponse CustomTagsUpdate(ctx, data)
更新客户标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CustomTagsUpdateRequest**](CustomTagsUpdateRequest.md)|  | 

### Return type

[**CustomTagsUpdateResponse**](CustomTagsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


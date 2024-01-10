# TencentAds\MaterialLabelsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MaterialLabelsAdd**](MaterialLabelsApi.md#MaterialLabelsAdd) | **Post** /material_labels/add | 创建素材标签
[**MaterialLabelsBind**](MaterialLabelsApi.md#MaterialLabelsBind) | **Post** /material_labels/bind | 绑定标签素材关联关系
[**MaterialLabelsDelete**](MaterialLabelsApi.md#MaterialLabelsDelete) | **Post** /material_labels/delete | 删除素材标签
[**MaterialLabelsGet**](MaterialLabelsApi.md#MaterialLabelsGet) | **Get** /material_labels/get | 获取素材标签列表
[**MaterialLabelsUpdate**](MaterialLabelsApi.md#MaterialLabelsUpdate) | **Post** /material_labels/update | 更新素材标签


# **MaterialLabelsAdd**
> MaterialLabelsAddResponse MaterialLabelsAdd(ctx, data)
创建素材标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MaterialLabelsAddRequest**](MaterialLabelsAddRequest.md)|  | 

### Return type

[**MaterialLabelsAddResponse**](MaterialLabelsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MaterialLabelsBind**
> MaterialLabelsBindResponse MaterialLabelsBind(ctx, data)
绑定标签素材关联关系

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MaterialLabelsBindRequest**](MaterialLabelsBindRequest.md)|  | 

### Return type

[**MaterialLabelsBindResponse**](MaterialLabelsBindResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MaterialLabelsDelete**
> MaterialLabelsDeleteResponse MaterialLabelsDelete(ctx, data)
删除素材标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MaterialLabelsDeleteRequest**](MaterialLabelsDeleteRequest.md)|  | 

### Return type

[**MaterialLabelsDeleteResponse**](MaterialLabelsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MaterialLabelsGet**
> MaterialLabelsGetResponse MaterialLabelsGet(ctx, accountId, optional)
获取素材标签列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***MaterialLabelsApiMaterialLabelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MaterialLabelsApiMaterialLabelsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labelId** | **optional.Int64**|  | 
 **labelName** | **optional.String**|  | 
 **firstLabelLevelIdList** | [**optional.Interface of []int64**](int64.md)| 一级标签类目ID列表 | 
 **secondLabelLevelIdList** | [**optional.Interface of []int64**](int64.md)| 二级标签类目ID列表 | 
 **needCount** | **optional.Bool**|  | 
 **businessScenario** | **optional.String**|  | 
 **orderBy** | [**optional.Interface of []OrderByStruct**](OrderByStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**MaterialLabelsGetResponse**](MaterialLabelsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MaterialLabelsUpdate**
> MaterialLabelsUpdateResponse MaterialLabelsUpdate(ctx, data)
更新素材标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MaterialLabelsUpdateRequest**](MaterialLabelsUpdateRequest.md)|  | 

### Return type

[**MaterialLabelsUpdateResponse**](MaterialLabelsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


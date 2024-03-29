# TencentAds\MaterialDcatagApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MaterialDcatagAdd**](MaterialDcatagApi.md#MaterialDcatagAdd) | **Post** /material_dcatag/add | 素材DCA标签绑定新增
[**MaterialDcatagGet**](MaterialDcatagApi.md#MaterialDcatagGet) | **Get** /material_dcatag/get | 素材DCA标签绑定查询


# **MaterialDcatagAdd**
> MaterialDcatagAddResponse MaterialDcatagAdd(ctx, data)
素材DCA标签绑定新增

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MaterialDcatagAddRequest**](MaterialDcatagAddRequest.md)|  | 

### Return type

[**MaterialDcatagAddResponse**](MaterialDcatagAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MaterialDcatagGet**
> MaterialDcatagGetResponse MaterialDcatagGet(ctx, accountId, optional)
素材DCA标签绑定查询

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***MaterialDcatagApiMaterialDcatagGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MaterialDcatagApiMaterialDcatagGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imageIdList** | [**optional.Interface of []int64**](int64.md)|  | 
 **mediaIdList** | [**optional.Interface of []int64**](int64.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**MaterialDcatagGetResponse**](MaterialDcatagGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


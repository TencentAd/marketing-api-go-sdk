# TencentAds\BusinessManagerRelationsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BusinessManagerRelationsGet**](BusinessManagerRelationsApi.md#BusinessManagerRelationsGet) | **Get** /business_manager_relations/get | 查询商务管家账号下广告主信息


# **BusinessManagerRelationsGet**
> BusinessManagerRelationsGetResponse BusinessManagerRelationsGet(ctx, optional)
查询商务管家账号下广告主信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BusinessManagerRelationsApiBusinessManagerRelationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessManagerRelationsApiBusinessManagerRelationsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **advertiserType** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**BusinessManagerRelationsGetResponse**](BusinessManagerRelationsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


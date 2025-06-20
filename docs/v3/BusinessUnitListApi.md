# TencentAds\BusinessUnitListApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BusinessUnitListGet**](BusinessUnitListApi.md#BusinessUnitListGet) | **Get** /business_unit_list/get | 查询业务单元列表


# **BusinessUnitListGet**
> BusinessUnitListGetResponse BusinessUnitListGet(ctx, page, pageSize, optional)
查询业务单元列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int64**|  | 
  **pageSize** | **int64**|  | 
 **optional** | ***BusinessUnitListApiBusinessUnitListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessUnitListApiBusinessUnitListGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**BusinessUnitListGetResponse**](BusinessUnitListGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


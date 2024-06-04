# TencentAds\AgencyBusinessUnitListApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgencyBusinessUnitListGet**](AgencyBusinessUnitListApi.md#AgencyBusinessUnitListGet) | **Get** /agency_business_unit_list/get | 查询服务商业务单元列表


# **AgencyBusinessUnitListGet**
> AgencyBusinessUnitListGetResponse AgencyBusinessUnitListGet(ctx, page, pageSize, optional)
查询服务商业务单元列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int64**|  | 
  **pageSize** | **int64**|  | 
 **optional** | ***AgencyBusinessUnitListApiAgencyBusinessUnitListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgencyBusinessUnitListApiAgencyBusinessUnitListGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **organizationId** | **optional.Int64**|  | 
 **businessUnitName** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AgencyBusinessUnitListGetResponse**](AgencyBusinessUnitListGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


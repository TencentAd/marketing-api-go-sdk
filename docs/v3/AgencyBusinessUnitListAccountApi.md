# TencentAds\AgencyBusinessUnitListAccountApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgencyBusinessUnitListAccountGet**](AgencyBusinessUnitListAccountApi.md#AgencyBusinessUnitListAccountGet) | **Get** /agency_business_unit_list_account/get | 查询服务商业务单元下的广告主账号列表


# **AgencyBusinessUnitListAccountGet**
> AgencyBusinessUnitListAccountGetResponse AgencyBusinessUnitListAccountGet(ctx, organizationId, page, pageSize, optional)
查询服务商业务单元下的广告主账号列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationId** | **int64**|  | 
  **page** | **int64**|  | 
  **pageSize** | **int64**|  | 
 **optional** | ***AgencyBusinessUnitListAccountApiAgencyBusinessUnitListAccountGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgencyBusinessUnitListAccountApiAgencyBusinessUnitListAccountGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AgencyBusinessUnitListAccountGetResponse**](AgencyBusinessUnitListAccountGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


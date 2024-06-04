# TencentAds\AgencyBusinessUnitApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgencyBusinessUnitAdd**](AgencyBusinessUnitApi.md#AgencyBusinessUnitAdd) | **Post** /agency_business_unit/add | 创建服务商业务单元
[**AgencyBusinessUnitUpdate**](AgencyBusinessUnitApi.md#AgencyBusinessUnitUpdate) | **Post** /agency_business_unit/update | 修改服务商业务单元


# **AgencyBusinessUnitAdd**
> AgencyBusinessUnitAddResponse AgencyBusinessUnitAdd(ctx, data)
创建服务商业务单元

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AgencyBusinessUnitAddRequest**](AgencyBusinessUnitAddRequest.md)|  | 

### Return type

[**AgencyBusinessUnitAddResponse**](AgencyBusinessUnitAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgencyBusinessUnitUpdate**
> AgencyBusinessUnitUpdateResponse AgencyBusinessUnitUpdate(ctx, data)
修改服务商业务单元

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AgencyBusinessUnitUpdateRequest**](AgencyBusinessUnitUpdateRequest.md)|  | 

### Return type

[**AgencyBusinessUnitUpdateResponse**](AgencyBusinessUnitUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


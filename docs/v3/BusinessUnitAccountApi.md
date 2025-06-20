# TencentAds\BusinessUnitAccountApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BusinessUnitAccountGet**](BusinessUnitAccountApi.md#BusinessUnitAccountGet) | **Post** /business_unit_account/get | 查询广告主账号所属的业务单元
[**BusinessUnitAccountUpdate**](BusinessUnitAccountApi.md#BusinessUnitAccountUpdate) | **Post** /business_unit_account/update | 修改客户业务单元帐户


# **BusinessUnitAccountGet**
> BusinessUnitAccountGetResponse BusinessUnitAccountGet(ctx, data)
查询广告主账号所属的业务单元

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BusinessUnitAccountGetRequest**](BusinessUnitAccountGetRequest.md)|  | 

### Return type

[**BusinessUnitAccountGetResponse**](BusinessUnitAccountGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BusinessUnitAccountUpdate**
> BusinessUnitAccountUpdateResponse BusinessUnitAccountUpdate(ctx, data)
修改客户业务单元帐户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BusinessUnitAccountUpdateRequest**](BusinessUnitAccountUpdateRequest.md)|  | 

### Return type

[**BusinessUnitAccountUpdateResponse**](BusinessUnitAccountUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


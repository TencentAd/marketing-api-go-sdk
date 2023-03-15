# TencentAds\LeadsCallRecordsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadsCallRecordsGet**](LeadsCallRecordsApi.md#LeadsCallRecordsGet) | **Post** /leads_call_records/get | 获取一个账号下的全部通话结果


# **LeadsCallRecordsGet**
> LeadsCallRecordsGetResponse LeadsCallRecordsGet(ctx, data)
获取一个账号下的全部通话结果

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LeadsCallRecordsGetRequest**](LeadsCallRecordsGetRequest.md)|  | 

### Return type

[**LeadsCallRecordsGetResponse**](LeadsCallRecordsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\EstimationApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EstimationGet**](EstimationApi.md#EstimationGet) | **Post** /estimation/get | 预估覆盖人数


# **EstimationGet**
> EstimationGetResponse EstimationGet(ctx, data)
预估覆盖人数

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**EstimationGetRequest**](EstimationGetRequest.md)|  | 

### Return type

[**EstimationGetResponse**](EstimationGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


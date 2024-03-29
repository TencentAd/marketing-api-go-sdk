# TencentAds\CustomAudienceEstimationsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomAudienceEstimationsGet**](CustomAudienceEstimationsApi.md#CustomAudienceEstimationsGet) | **Get** /custom_audience_estimations/get | 人群覆盖数预估


# **CustomAudienceEstimationsGet**
> CustomAudienceEstimationsGetResponse CustomAudienceEstimationsGet(ctx, accountId, type_, audienceSpec, optional)
人群覆盖数预估

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **type_** | **string**|  | 
  **audienceSpec** | [**EstimationAudienceSpec**](EstimationAudienceSpec.md)|  | 
 **optional** | ***CustomAudienceEstimationsApiCustomAudienceEstimationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomAudienceEstimationsApiCustomAudienceEstimationsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CustomAudienceEstimationsGetResponse**](CustomAudienceEstimationsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\AdcreativesRelatedCapabilityApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdcreativesRelatedCapabilityGet**](AdcreativesRelatedCapabilityApi.md#AdcreativesRelatedCapabilityGet) | **Get** /adcreatives_related_capability/get | 检查广告创意是否可以修改


# **AdcreativesRelatedCapabilityGet**
> AdcreativesRelatedCapabilityGetResponse AdcreativesRelatedCapabilityGet(ctx, accountId, adId, optional)
检查广告创意是否可以修改

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **adId** | **int64**|  | 
 **optional** | ***AdcreativesRelatedCapabilityApiAdcreativesRelatedCapabilityGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdcreativesRelatedCapabilityApiAdcreativesRelatedCapabilityGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AdcreativesRelatedCapabilityGetResponse**](AdcreativesRelatedCapabilityGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


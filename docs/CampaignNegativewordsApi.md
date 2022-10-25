# TencentAds\CampaignNegativewordsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CampaignNegativewordsAdd**](CampaignNegativewordsApi.md#CampaignNegativewordsAdd) | **Post** /campaign_negativewords/add | 新增推广计划否定词
[**CampaignNegativewordsGet**](CampaignNegativewordsApi.md#CampaignNegativewordsGet) | **Post** /campaign_negativewords/get | 查询推广计划否定词
[**CampaignNegativewordsUpdate**](CampaignNegativewordsApi.md#CampaignNegativewordsUpdate) | **Post** /campaign_negativewords/update | 修改推广计划否定词


# **CampaignNegativewordsAdd**
> CampaignNegativewordsAddResponse CampaignNegativewordsAdd(ctx, data)
新增推广计划否定词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CampaignNegativewordsAddRequest**](CampaignNegativewordsAddRequest.md)|  | 

### Return type

[**CampaignNegativewordsAddResponse**](CampaignNegativewordsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignNegativewordsGet**
> CampaignNegativewordsGetResponse CampaignNegativewordsGet(ctx, data)
查询推广计划否定词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CampaignNegativewordsGetRequest**](CampaignNegativewordsGetRequest.md)|  | 

### Return type

[**CampaignNegativewordsGetResponse**](CampaignNegativewordsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignNegativewordsUpdate**
> CampaignNegativewordsUpdateResponse CampaignNegativewordsUpdate(ctx, data)
修改推广计划否定词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CampaignNegativewordsUpdateRequest**](CampaignNegativewordsUpdateRequest.md)|  | 

### Return type

[**CampaignNegativewordsUpdateResponse**](CampaignNegativewordsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


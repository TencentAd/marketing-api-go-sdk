# TencentAds\AdvertiserDailyBudgetApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdvertiserDailyBudgetGet**](AdvertiserDailyBudgetApi.md#AdvertiserDailyBudgetGet) | **Get** /advertiser_daily_budget/get | 获取竞价广告账户日预算


# **AdvertiserDailyBudgetGet**
> AdvertiserDailyBudgetGetResponse AdvertiserDailyBudgetGet(ctx, accountId, fields)
获取竞价广告账户日预算

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **fields** | [**[]string**](string.md)|  | 

### Return type

[**AdvertiserDailyBudgetGetResponse**](AdvertiserDailyBudgetGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


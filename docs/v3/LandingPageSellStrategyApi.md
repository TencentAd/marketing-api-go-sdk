# TencentAds\LandingPageSellStrategyApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LandingPageSellStrategyAdd**](LandingPageSellStrategyApi.md#LandingPageSellStrategyAdd) | **Post** /landing_page_sell_strategy/add | 短剧售卖策略-创建售卖策略
[**LandingPageSellStrategyGet**](LandingPageSellStrategyApi.md#LandingPageSellStrategyGet) | **Get** /landing_page_sell_strategy/get | 短剧售卖策略-获取售卖策略列表


# **LandingPageSellStrategyAdd**
> LandingPageSellStrategyAddResponse LandingPageSellStrategyAdd(ctx, data)
短剧售卖策略-创建售卖策略

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LandingPageSellStrategyAddRequest**](LandingPageSellStrategyAddRequest.md)|  | 

### Return type

[**LandingPageSellStrategyAddResponse**](LandingPageSellStrategyAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LandingPageSellStrategyGet**
> LandingPageSellStrategyGetResponse LandingPageSellStrategyGet(ctx, accountId, optional)
短剧售卖策略-获取售卖策略列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***LandingPageSellStrategyApiLandingPageSellStrategyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LandingPageSellStrategyApiLandingPageSellStrategyGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **strategyId** | **optional.Int64**|  | 
 **strategyStatus** | **optional.Int64**|  | 
 **strategyName** | **optional.String**|  | 
 **fullStrategyName** | **optional.String**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**LandingPageSellStrategyGetResponse**](LandingPageSellStrategyGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


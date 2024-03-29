# TencentAds\AdvertiserApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdvertiserAdd**](AdvertiserApi.md#AdvertiserAdd) | **Post** /advertiser/add | 添加腾讯广告服务商子客
[**AdvertiserGet**](AdvertiserApi.md#AdvertiserGet) | **Get** /advertiser/get | 查询腾讯广告广告主信息
[**AdvertiserUpdate**](AdvertiserApi.md#AdvertiserUpdate) | **Post** /advertiser/update | 更新腾讯广告广告主信息
[**AdvertiserUpdateDailyBudget**](AdvertiserApi.md#AdvertiserUpdateDailyBudget) | **Post** /advertiser/update_daily_budget | 更新广告主日限额信息


# **AdvertiserAdd**
> AdvertiserAddResponse AdvertiserAdd(ctx, data)
添加腾讯广告服务商子客

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdvertiserAddRequest**](AdvertiserAddRequest.md)|  | 

### Return type

[**AdvertiserAddResponse**](AdvertiserAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdvertiserGet**
> AdvertiserGetResponse AdvertiserGet(ctx, optional)
查询腾讯广告广告主信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdvertiserApiAdvertiserGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdvertiserApiAdvertiserGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agencyId** | **optional.Int64**|  | 
 **accountId** | **optional.Int64**|  | 
 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)|  | 
 **paginationMode** | **optional.String**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **cursor** | **optional.Int64**|  | 

### Return type

[**AdvertiserGetResponse**](AdvertiserGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdvertiserUpdate**
> AdvertiserUpdateResponse AdvertiserUpdate(ctx, data)
更新腾讯广告广告主信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdvertiserUpdateRequest**](AdvertiserUpdateRequest.md)|  | 

### Return type

[**AdvertiserUpdateResponse**](AdvertiserUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdvertiserUpdateDailyBudget**
> AdvertiserUpdateDailyBudgetResponse AdvertiserUpdateDailyBudget(ctx, data)
更新广告主日限额信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdvertiserUpdateDailyBudgetRequest**](AdvertiserUpdateDailyBudgetRequest.md)|  | 

### Return type

[**AdvertiserUpdateDailyBudgetResponse**](AdvertiserUpdateDailyBudgetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


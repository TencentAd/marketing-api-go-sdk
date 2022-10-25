# TencentAds\CampaignsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CampaignsAdd**](CampaignsApi.md#CampaignsAdd) | **Post** /campaigns/add | 创建推广计划
[**CampaignsAddNegativeword**](CampaignsApi.md#CampaignsAddNegativeword) | **Post** /campaigns/add_negativeword | 新增推广计划否定词
[**CampaignsDelete**](CampaignsApi.md#CampaignsDelete) | **Post** /campaigns/delete | 删除推广计划
[**CampaignsGet**](CampaignsApi.md#CampaignsGet) | **Get** /campaigns/get | 获取推广计划
[**CampaignsGetNegativeword**](CampaignsApi.md#CampaignsGetNegativeword) | **Post** /campaigns/get_negativeword | 查询推广计划否定词
[**CampaignsUpdate**](CampaignsApi.md#CampaignsUpdate) | **Post** /campaigns/update | 更新推广计划
[**CampaignsUpdateConfiguredStatus**](CampaignsApi.md#CampaignsUpdateConfiguredStatus) | **Post** /campaigns/update_configured_status | 更新推广计划状态
[**CampaignsUpdateDailyBudget**](CampaignsApi.md#CampaignsUpdateDailyBudget) | **Post** /campaigns/update_daily_budget | 更新推广计划日限额信息
[**CampaignsUpdateNegativeword**](CampaignsApi.md#CampaignsUpdateNegativeword) | **Post** /campaigns/update_negativeword | 修改推广计划否定词


# **CampaignsAdd**
> CampaignsAddResponse CampaignsAdd(ctx, data)
创建推广计划

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CampaignsAddRequest**](CampaignsAddRequest.md)|  | 

### Return type

[**CampaignsAddResponse**](CampaignsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignsAddNegativeword**
> CampaignsAddNegativewordResponse CampaignsAddNegativeword(ctx, data)
新增推广计划否定词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CampaignsAddNegativewordRequest**](CampaignsAddNegativewordRequest.md)|  | 

### Return type

[**CampaignsAddNegativewordResponse**](CampaignsAddNegativewordResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignsDelete**
> CampaignsDeleteResponse CampaignsDelete(ctx, data)
删除推广计划

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CampaignsDeleteRequest**](CampaignsDeleteRequest.md)|  | 

### Return type

[**CampaignsDeleteResponse**](CampaignsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignsGet**
> CampaignsGetResponse CampaignsGet(ctx, accountId, optional)
获取推广计划

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***CampaignsApiCampaignsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiCampaignsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **isDeleted** | **optional.Bool**|  | 
 **weixinOfficialAccountsUpgradeEnabled** | **optional.Bool**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CampaignsGetResponse**](CampaignsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignsGetNegativeword**
> CampaignsGetNegativewordResponse CampaignsGetNegativeword(ctx, data)
查询推广计划否定词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CampaignsGetNegativewordRequest**](CampaignsGetNegativewordRequest.md)|  | 

### Return type

[**CampaignsGetNegativewordResponse**](CampaignsGetNegativewordResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignsUpdate**
> CampaignsUpdateResponse CampaignsUpdate(ctx, data)
更新推广计划

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CampaignsUpdateRequest**](CampaignsUpdateRequest.md)|  | 

### Return type

[**CampaignsUpdateResponse**](CampaignsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignsUpdateConfiguredStatus**
> CampaignsUpdateConfiguredStatusResponse CampaignsUpdateConfiguredStatus(ctx, data)
更新推广计划状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CampaignsUpdateConfiguredStatusRequest**](CampaignsUpdateConfiguredStatusRequest.md)|  | 

### Return type

[**CampaignsUpdateConfiguredStatusResponse**](CampaignsUpdateConfiguredStatusResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignsUpdateDailyBudget**
> CampaignsUpdateDailyBudgetResponse CampaignsUpdateDailyBudget(ctx, data)
更新推广计划日限额信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CampaignsUpdateDailyBudgetRequest**](CampaignsUpdateDailyBudgetRequest.md)|  | 

### Return type

[**CampaignsUpdateDailyBudgetResponse**](CampaignsUpdateDailyBudgetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignsUpdateNegativeword**
> CampaignsUpdateNegativewordResponse CampaignsUpdateNegativeword(ctx, data)
修改推广计划否定词

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CampaignsUpdateNegativewordRequest**](CampaignsUpdateNegativewordRequest.md)|  | 

### Return type

[**CampaignsUpdateNegativewordResponse**](CampaignsUpdateNegativewordResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


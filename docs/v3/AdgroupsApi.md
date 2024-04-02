# TencentAds\AdgroupsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdgroupsAdd**](AdgroupsApi.md#AdgroupsAdd) | **Post** /adgroups/add | 创建广告组
[**AdgroupsDelete**](AdgroupsApi.md#AdgroupsDelete) | **Post** /adgroups/delete | 删除广告组
[**AdgroupsGet**](AdgroupsApi.md#AdgroupsGet) | **Get** /adgroups/get | 获取广告组
[**AdgroupsUpdate**](AdgroupsApi.md#AdgroupsUpdate) | **Post** /adgroups/update | 更新广告组
[**AdgroupsUpdateBidAmount**](AdgroupsApi.md#AdgroupsUpdateBidAmount) | **Post** /adgroups/update_bid_amount | 批量修改广告组出价
[**AdgroupsUpdateConfiguredStatus**](AdgroupsApi.md#AdgroupsUpdateConfiguredStatus) | **Post** /adgroups/update_configured_status | 批量修改广告组开启/暂停状态
[**AdgroupsUpdateDailyBudget**](AdgroupsApi.md#AdgroupsUpdateDailyBudget) | **Post** /adgroups/update_daily_budget | 批量修改广告组日限额
[**AdgroupsUpdateDatetime**](AdgroupsApi.md#AdgroupsUpdateDatetime) | **Post** /adgroups/update_datetime | 批量修改广告组投放起止时间


# **AdgroupsAdd**
> AdgroupsAddResponse AdgroupsAdd(ctx, data)
创建广告组

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdgroupsAddRequest**](AdgroupsAddRequest.md)|  | 

### Return type

[**AdgroupsAddResponse**](AdgroupsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdgroupsDelete**
> AdgroupsDeleteResponse AdgroupsDelete(ctx, data)
删除广告组

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdgroupsDeleteRequest**](AdgroupsDeleteRequest.md)|  | 

### Return type

[**AdgroupsDeleteResponse**](AdgroupsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdgroupsGet**
> AdgroupsGetResponse AdgroupsGet(ctx, accountId, optional)
获取广告组

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***AdgroupsApiAdgroupsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdgroupsApiAdgroupsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **isDeleted** | **optional.Bool**|  | 
 **fields** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**AdgroupsGetResponse**](AdgroupsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdgroupsUpdate**
> AdgroupsUpdateResponse AdgroupsUpdate(ctx, data)
更新广告组

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdgroupsUpdateRequest**](AdgroupsUpdateRequest.md)|  | 

### Return type

[**AdgroupsUpdateResponse**](AdgroupsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdgroupsUpdateBidAmount**
> AdgroupsUpdateBidAmountResponse AdgroupsUpdateBidAmount(ctx, data)
批量修改广告组出价

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdgroupsUpdateBidAmountRequest**](AdgroupsUpdateBidAmountRequest.md)|  | 

### Return type

[**AdgroupsUpdateBidAmountResponse**](AdgroupsUpdateBidAmountResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdgroupsUpdateConfiguredStatus**
> AdgroupsUpdateConfiguredStatusResponse AdgroupsUpdateConfiguredStatus(ctx, data)
批量修改广告组开启/暂停状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdgroupsUpdateConfiguredStatusRequest**](AdgroupsUpdateConfiguredStatusRequest.md)|  | 

### Return type

[**AdgroupsUpdateConfiguredStatusResponse**](AdgroupsUpdateConfiguredStatusResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdgroupsUpdateDailyBudget**
> AdgroupsUpdateDailyBudgetResponse AdgroupsUpdateDailyBudget(ctx, data)
批量修改广告组日限额

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdgroupsUpdateDailyBudgetRequest**](AdgroupsUpdateDailyBudgetRequest.md)|  | 

### Return type

[**AdgroupsUpdateDailyBudgetResponse**](AdgroupsUpdateDailyBudgetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdgroupsUpdateDatetime**
> AdgroupsUpdateDatetimeResponse AdgroupsUpdateDatetime(ctx, data)
批量修改广告组投放起止时间

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdgroupsUpdateDatetimeRequest**](AdgroupsUpdateDatetimeRequest.md)|  | 

### Return type

[**AdgroupsUpdateDatetimeResponse**](AdgroupsUpdateDatetimeResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


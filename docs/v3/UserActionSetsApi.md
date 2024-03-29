# TencentAds\UserActionSetsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserActionSetsAdd**](UserActionSetsApi.md#UserActionSetsAdd) | **Post** /user_action_sets/add | 创建用户行为数据源
[**UserActionSetsGet**](UserActionSetsApi.md#UserActionSetsGet) | **Get** /user_action_sets/get | 获取用户行为数据源


# **UserActionSetsAdd**
> UserActionSetsAddResponse UserActionSetsAdd(ctx, data)
创建用户行为数据源

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**UserActionSetsAddRequest**](UserActionSetsAddRequest.md)|  | 

### Return type

[**UserActionSetsAddResponse**](UserActionSetsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserActionSetsGet**
> UserActionSetsGetResponse UserActionSetsGet(ctx, accountId, optional)
获取用户行为数据源

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***UserActionSetsApiUserActionSetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserActionSetsApiUserActionSetsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userActionSetId** | **optional.Int64**|  | 
 **type_** | [**optional.Interface of []string**](string.md)|  | 
 **mobileAppId** | **optional.Int64**|  | 
 **wechatAppId** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **includePermission** | **optional.Bool**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**UserActionSetsGetResponse**](UserActionSetsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\GameFeatureV6Api

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameFeatureV6Add**](GameFeatureV6Api.md#GameFeatureV6Add) | **Post** /game_feature_v6/add | 新增游戏App特征V6
[**GameFeatureV6Get**](GameFeatureV6Api.md#GameFeatureV6Get) | **Get** /game_feature_v6/get | 获取游戏App特征V6


# **GameFeatureV6Add**
> GameFeatureV6AddResponse GameFeatureV6Add(ctx, data)
新增游戏App特征V6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**GameFeatureV6AddRequest**](GameFeatureV6AddRequest.md)|  | 

### Return type

[**GameFeatureV6AddResponse**](GameFeatureV6AddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GameFeatureV6Get**
> GameFeatureV6GetResponse GameFeatureV6Get(ctx, accountId, marketingTargetType, marketingTargetDetailId, optional)
获取游戏App特征V6

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **marketingTargetType** | **string**|  | 
  **marketingTargetDetailId** | **string**|  | 
 **optional** | ***GameFeatureV6ApiGameFeatureV6GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GameFeatureV6ApiGameFeatureV6GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**GameFeatureV6GetResponse**](GameFeatureV6GetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


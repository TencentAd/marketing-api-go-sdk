# TencentAds\OptimizationGoalPermissionsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OptimizationGoalPermissionsGet**](OptimizationGoalPermissionsApi.md#OptimizationGoalPermissionsGet) | **Get** /optimization_goal_permissions/get | 查询优化目标权限


# **OptimizationGoalPermissionsGet**
> OptimizationGoalPermissionsGetResponse OptimizationGoalPermissionsGet(ctx, accountId, siteSet, marketingGoal, marketingSubGoal, marketingCarrierType, marketingTargetType, optional)
查询优化目标权限

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **siteSet** | [**[]string**](string.md)|  | 
  **marketingGoal** | **string**|  | 
  **marketingSubGoal** | **string**|  | 
  **marketingCarrierType** | **string**|  | 
  **marketingTargetType** | **string**|  | 
 **optional** | ***OptimizationGoalPermissionsApiOptimizationGoalPermissionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OptimizationGoalPermissionsApiOptimizationGoalPermissionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **bidMode** | **optional.String**|  | 
 **marketingCarrierDetail** | [**optional.Interface of MarketingCarrierDetail**](MarketingCarrierDetail.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**OptimizationGoalPermissionsGetResponse**](OptimizationGoalPermissionsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


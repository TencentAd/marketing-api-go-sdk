# TencentAds\JointBudgetRulesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JointBudgetRulesAdd**](JointBudgetRulesApi.md#JointBudgetRulesAdd) | **Post** /joint_budget_rules/add | 创建联合预算规则
[**JointBudgetRulesGet**](JointBudgetRulesApi.md#JointBudgetRulesGet) | **Get** /joint_budget_rules/get | 获取联合预算
[**JointBudgetRulesUpdate**](JointBudgetRulesApi.md#JointBudgetRulesUpdate) | **Post** /joint_budget_rules/update | 更新联合预算规则


# **JointBudgetRulesAdd**
> JointBudgetRulesAddResponse JointBudgetRulesAdd(ctx, data)
创建联合预算规则

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**JointBudgetRulesAddRequest**](JointBudgetRulesAddRequest.md)|  | 

### Return type

[**JointBudgetRulesAddResponse**](JointBudgetRulesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JointBudgetRulesGet**
> JointBudgetRulesGetResponse JointBudgetRulesGet(ctx, accountId, optional)
获取联合预算

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***JointBudgetRulesApiJointBudgetRulesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a JointBudgetRulesApiJointBudgetRulesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**JointBudgetRulesGetResponse**](JointBudgetRulesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JointBudgetRulesUpdate**
> JointBudgetRulesUpdateResponse JointBudgetRulesUpdate(ctx, data)
更新联合预算规则

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**JointBudgetRulesUpdateRequest**](JointBudgetRulesUpdateRequest.md)|  | 

### Return type

[**JointBudgetRulesUpdateResponse**](JointBudgetRulesUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


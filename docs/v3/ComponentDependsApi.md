# TencentAds\ComponentDependsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentDependsGet**](ComponentDependsApi.md#ComponentDependsGet) | **Get** /component_depends/get | 查询创意组件字段选项对于其他组件的依赖信息


# **ComponentDependsGet**
> ComponentDependsGetResponse ComponentDependsGet(ctx, accountId, marketingGoal, marketingTargetType, marketingCarrierType, deliveryMode, optional)
查询创意组件字段选项对于其他组件的依赖信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **marketingGoal** | **string**|  | 
  **marketingTargetType** | **string**|  | 
  **marketingCarrierType** | **string**|  | 
  **deliveryMode** | **string**|  | 
 **optional** | ***ComponentDependsApiComponentDependsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentDependsApiComponentDependsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **marketingSubGoal** | **optional.String**|  | 
 **automaticSiteEnabled** | **optional.Bool**|  | 
 **siteSet** | [**optional.Interface of []string**](string.md)|  | 
 **dynamicCreativeType** | **optional.String**|  | 
 **creativeTemplateId** | **optional.Int64**|  | 
 **componentType** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ComponentDependsGetResponse**](ComponentDependsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


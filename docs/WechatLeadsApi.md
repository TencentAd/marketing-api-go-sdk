# TencentAds\WechatLeadsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatLeadsGet**](WechatLeadsApi.md#WechatLeadsGet) | **Get** /wechat_leads/get | 获取销售线索(待废弃)


# **WechatLeadsGet**
> WechatLeadsGetResponse WechatLeadsGet(ctx, optional)
获取销售线索(待废弃)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WechatLeadsApiWechatLeadsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatLeadsApiWechatLeadsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateRange** | [**optional.Interface of DateRange**](DateRange.md)|  | 
 **timeRange** | [**optional.Interface of TimeRange**](TimeRange.md)|  | 
 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatLeadsGetResponse**](WechatLeadsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


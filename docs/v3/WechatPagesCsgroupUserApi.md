# TencentAds\WechatPagesCsgroupUserApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatPagesCsgroupUserGet**](WechatPagesCsgroupUserApi.md#WechatPagesCsgroupUserGet) | **Get** /wechat_pages_csgroup_user/get | 获取企业微信组件客服列表


# **WechatPagesCsgroupUserGet**
> WechatPagesCsgroupUserGetResponse WechatPagesCsgroupUserGet(ctx, accountId, corpId, optional)
获取企业微信组件客服列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **corpId** | **string**|  | 
 **optional** | ***WechatPagesCsgroupUserApiWechatPagesCsgroupUserGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatPagesCsgroupUserApiWechatPagesCsgroupUserGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **departmentId** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatPagesCsgroupUserGetResponse**](WechatPagesCsgroupUserGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\WechatPagesGrantinfoApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatPagesGrantinfoGet**](WechatPagesGrantinfoApi.md#WechatPagesGrantinfoGet) | **Get** /wechat_pages_grantinfo/get | 获取原生页授权方信息


# **WechatPagesGrantinfoGet**
> WechatPagesGrantinfoGetResponse WechatPagesGrantinfoGet(ctx, accountId, optional)
获取原生页授权方信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WechatPagesGrantinfoApiWechatPagesGrantinfoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatPagesGrantinfoApiWechatPagesGrantinfoGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchKey** | **optional.String**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatPagesGrantinfoGetResponse**](WechatPagesGrantinfoGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


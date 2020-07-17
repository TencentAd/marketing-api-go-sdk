# TencentAds\AuthorizationApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizationWechatBind**](AuthorizationApi.md#AuthorizationWechatBind) | **Get** /authorization/wechat_bind | 绑定微信公众号(待废弃)


# **AuthorizationWechatBind**
> string AuthorizationWechatBind(ctx, accessToken, redirectUri, optional)
绑定微信公众号(待废弃)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accessToken** | **string**|  | 
  **redirectUri** | **string**|  | 
 **optional** | ***AuthorizationApiAuthorizationWechatBindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationApiAuthorizationWechatBindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **optional.Int64**| 需绑定公众号的广告主 id，有操作权限的帐号 id | 
 **wechatAccountId** | **optional.String**| 微信公众号id，用于判断扫描的公众号与请求时的wechat_account_id是否一致 | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


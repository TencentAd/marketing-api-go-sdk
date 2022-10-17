# TencentAds\WeixinOfficialAccountsUpgradeStatusApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WeixinOfficialAccountsUpgradeStatusGet**](WeixinOfficialAccountsUpgradeStatusApi.md#WeixinOfficialAccountsUpgradeStatusGet) | **Get** /weixin_official_accounts_upgrade_status/get | 获取微信公众号账户升级状态


# **WeixinOfficialAccountsUpgradeStatusGet**
> WeixinOfficialAccountsUpgradeStatusGetResponse WeixinOfficialAccountsUpgradeStatusGet(ctx, accountId, optional)
获取微信公众号账户升级状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WeixinOfficialAccountsUpgradeStatusApiWeixinOfficialAccountsUpgradeStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WeixinOfficialAccountsUpgradeStatusApiWeixinOfficialAccountsUpgradeStatusGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WeixinOfficialAccountsUpgradeStatusGetResponse**](WeixinOfficialAccountsUpgradeStatusGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


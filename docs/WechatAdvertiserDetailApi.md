# TencentAds\WechatAdvertiserDetailApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatAdvertiserDetailGet**](WechatAdvertiserDetailApi.md#WechatAdvertiserDetailGet) | **Get** /wechat_advertiser_detail/get | 【待废弃，请使用新接口wechat_advertiser_specification/get】查询微信公众平台广告主详情信息


# **WechatAdvertiserDetailGet**
> WechatAdvertiserDetailGetResponse WechatAdvertiserDetailGet(ctx, optional)
【待废弃，请使用新接口wechat_advertiser_specification/get】查询微信公众平台广告主详情信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WechatAdvertiserDetailApiWechatAdvertiserDetailGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatAdvertiserDetailApiWechatAdvertiserDetailGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **optional.Int64**|  | 
 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatAdvertiserDetailGetResponse**](WechatAdvertiserDetailGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\WechatShopActivityListApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatShopActivityListGet**](WechatShopActivityListApi.md#WechatShopActivityListGet) | **Get** /wechat_shop_activity_list/get | 查询微信小店活动列表


# **WechatShopActivityListGet**
> WechatShopActivityListGetResponse WechatShopActivityListGet(ctx, accountId, wechatShopId, wechatShopProductId, optional)
查询微信小店活动列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **wechatShopId** | **string**|  | 
  **wechatShopProductId** | **string**|  | 
 **optional** | ***WechatShopActivityListApiWechatShopActivityListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatShopActivityListApiWechatShopActivityListGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatShopActivityListGetResponse**](WechatShopActivityListGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


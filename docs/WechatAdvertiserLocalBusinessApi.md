# TencentAds\WechatAdvertiserLocalBusinessApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatAdvertiserLocalBusinessAdd**](WechatAdvertiserLocalBusinessApi.md#WechatAdvertiserLocalBusinessAdd) | **Post** /wechat_advertiser_local_business/add | 附近推商家开户
[**WechatAdvertiserLocalBusinessGet**](WechatAdvertiserLocalBusinessApi.md#WechatAdvertiserLocalBusinessGet) | **Get** /wechat_advertiser_local_business/get | 查询附近推商家信息
[**WechatAdvertiserLocalBusinessUpdate**](WechatAdvertiserLocalBusinessApi.md#WechatAdvertiserLocalBusinessUpdate) | **Post** /wechat_advertiser_local_business/update | 更新附近推商家信息


# **WechatAdvertiserLocalBusinessAdd**
> WechatAdvertiserLocalBusinessAddResponse WechatAdvertiserLocalBusinessAdd(ctx, headImage, name, description, contactPerson, contactPersonMobile, contactPersonCardId, contactPersonTele, corporation, corporationLicence, businessContent, industryId, businessId)
附近推商家开户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **headImage** | ***os.File**|  | 
  **name** | **string**|  | 
  **description** | **string**|  | 
  **contactPerson** | **string**|  | 
  **contactPersonMobile** | **string**|  | 
  **contactPersonCardId** | **string**|  | 
  **contactPersonTele** | **string**|  | 
  **corporation** | **string**|  | 
  **corporationLicence** | **string**|  | 
  **businessContent** | **string**|  | 
  **industryId** | **int64**|  | 
  **businessId** | **string**|  | 

### Return type

[**WechatAdvertiserLocalBusinessAddResponse**](WechatAdvertiserLocalBusinessAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatAdvertiserLocalBusinessGet**
> WechatAdvertiserLocalBusinessGetResponse WechatAdvertiserLocalBusinessGet(ctx, accountId, optional)
查询附近推商家信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WechatAdvertiserLocalBusinessApiWechatAdvertiserLocalBusinessGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatAdvertiserLocalBusinessApiWechatAdvertiserLocalBusinessGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatAdvertiserLocalBusinessGetResponse**](WechatAdvertiserLocalBusinessGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatAdvertiserLocalBusinessUpdate**
> WechatAdvertiserLocalBusinessUpdateResponse WechatAdvertiserLocalBusinessUpdate(ctx, headImage, name, description, contactPerson, contactPersonMobile, contactPersonCardId, contactPersonTele, corporation, corporationLicence, businessContent, industryId, accountId)
更新附近推商家信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **headImage** | ***os.File**|  | 
  **name** | **string**|  | 
  **description** | **string**|  | 
  **contactPerson** | **string**|  | 
  **contactPersonMobile** | **string**|  | 
  **contactPersonCardId** | **string**|  | 
  **contactPersonTele** | **string**|  | 
  **corporation** | **string**|  | 
  **corporationLicence** | **string**|  | 
  **businessContent** | **string**|  | 
  **industryId** | **int64**|  | 
  **accountId** | **int64**|  | 

### Return type

[**WechatAdvertiserLocalBusinessUpdateResponse**](WechatAdvertiserLocalBusinessUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\XijingPageListApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**XijingPageListGet**](XijingPageListApi.md#XijingPageListGet) | **Get** /xijing_page_list/get | 蹊径-获取落地页列表


# **XijingPageListGet**
> XijingPageListGetResponse XijingPageListGet(ctx, accountId, optional)
蹊径-获取落地页列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***XijingPageListApiXijingPageListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a XijingPageListApiXijingPageListGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageId** | **optional.Int64**|  | 
 **pageServiceId** | **optional.String**|  | 
 **pageName** | **optional.String**|  | 
 **pageType** | [**optional.Interface of []string**](string.md)|  | 
 **pageLastModifyStartTime** | **optional.String**|  | 
 **pageLastModifyEndTime** | **optional.String**|  | 
 **pageSize** | **optional.Int64**|  | 
 **pageIndex** | **optional.Int64**|  | 
 **pagePublishStatus** | [**optional.Interface of []string**](string.md)|  | 
 **pageStatus** | [**optional.Interface of []string**](string.md)|  | 
 **pageSource** | **optional.String**|  | 
 **pageOwnerId** | **optional.Int64**|  | 
 **appId** | **optional.Int64**|  | 
 **appType** | **optional.String**|  | 
 **queryType** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**XijingPageListGetResponse**](XijingPageListGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


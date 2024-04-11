# TencentAds\XijingPageInteractiveApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**XijingPageInteractiveAdd**](XijingPageInteractiveApi.md#XijingPageInteractiveAdd) | **Post** /xijing_page_interactive/add | 蹊径-创建互动落地页


# **XijingPageInteractiveAdd**
> XijingPageInteractiveAddResponse XijingPageInteractiveAdd(ctx, accountId, isAutoSubmit, pageType, interactivePageType, pageTitle, pageName, mobileAppId, optional)
蹊径-创建互动落地页

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **isAutoSubmit** | **int64**|  | 
  **pageType** | **string**|  | 
  **interactivePageType** | **string**|  | 
  **pageTitle** | **string**|  | 
  **pageName** | **string**|  | 
  **mobileAppId** | **string**|  | 
 **optional** | ***XijingPageInteractiveApiXijingPageInteractiveAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a XijingPageInteractiveApiXijingPageInteractiveAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **file** | **optional.Interface of *os.File**|  | 
 **transformType** | **optional.String**|  | 
 **pageConfig** | **optional.String**|  | 

### Return type

[**XijingPageInteractiveAddResponse**](XijingPageInteractiveAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


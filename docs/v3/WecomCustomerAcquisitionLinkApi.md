# TencentAds\WecomCustomerAcquisitionLinkApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WecomCustomerAcquisitionLinkAdd**](WecomCustomerAcquisitionLinkApi.md#WecomCustomerAcquisitionLinkAdd) | **Post** /wecom_customer_acquisition_link/add | 创建获客链接
[**WecomCustomerAcquisitionLinkGet**](WecomCustomerAcquisitionLinkApi.md#WecomCustomerAcquisitionLinkGet) | **Get** /wecom_customer_acquisition_link/get | 查询获客链接列表
[**WecomCustomerAcquisitionLinkUpdate**](WecomCustomerAcquisitionLinkApi.md#WecomCustomerAcquisitionLinkUpdate) | **Post** /wecom_customer_acquisition_link/update | 更新获客链接


# **WecomCustomerAcquisitionLinkAdd**
> WecomCustomerAcquisitionLinkAddResponse WecomCustomerAcquisitionLinkAdd(ctx, data)
创建获客链接

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WecomCustomerAcquisitionLinkAddRequest**](WecomCustomerAcquisitionLinkAddRequest.md)|  | 

### Return type

[**WecomCustomerAcquisitionLinkAddResponse**](WecomCustomerAcquisitionLinkAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WecomCustomerAcquisitionLinkGet**
> WecomCustomerAcquisitionLinkGetResponse WecomCustomerAcquisitionLinkGet(ctx, accountId, corpId, optional)
查询获客链接列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **corpId** | **string**|  | 
 **optional** | ***WecomCustomerAcquisitionLinkApiWecomCustomerAcquisitionLinkGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WecomCustomerAcquisitionLinkApiWecomCustomerAcquisitionLinkGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **linkId** | **optional.String**| 指定获客链接 ID，传入时只返回该链接信息，不走分页 | 
 **cursor** | **optional.String**| 分页游标，首次请求不传 | 
 **limit** | **optional.Int64**| 单页大小，默认 10 | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WecomCustomerAcquisitionLinkGetResponse**](WecomCustomerAcquisitionLinkGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WecomCustomerAcquisitionLinkUpdate**
> WecomCustomerAcquisitionLinkUpdateResponse WecomCustomerAcquisitionLinkUpdate(ctx, data)
更新获客链接

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WecomCustomerAcquisitionLinkUpdateRequest**](WecomCustomerAcquisitionLinkUpdateRequest.md)|  | 

### Return type

[**WecomCustomerAcquisitionLinkUpdateResponse**](WecomCustomerAcquisitionLinkUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


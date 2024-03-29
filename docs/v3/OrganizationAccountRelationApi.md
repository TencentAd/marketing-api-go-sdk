# TencentAds\OrganizationAccountRelationApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationAccountRelationGet**](OrganizationAccountRelationApi.md#OrganizationAccountRelationGet) | **Get** /organization_account_relation/get | 查询组织下广告账号信息


# **OrganizationAccountRelationGet**
> OrganizationAccountRelationGetResponse OrganizationAccountRelationGet(ctx, paginationMode, optional)
查询组织下广告账号信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **paginationMode** | **string**|  | 
 **optional** | ***OrganizationAccountRelationApiOrganizationAccountRelationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationAccountRelationApiOrganizationAccountRelationGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.Int64**|  | 
 **advertiserType** | **optional.String**|  | 
 **cursor** | **optional.Int64**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**OrganizationAccountRelationGetResponse**](OrganizationAccountRelationGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


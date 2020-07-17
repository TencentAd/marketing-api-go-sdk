# TencentAds\AssetPermissionsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetPermissionsAdd**](AssetPermissionsApi.md#AssetPermissionsAdd) | **Post** /asset_permissions/add | 资产权限授予
[**AssetPermissionsDelete**](AssetPermissionsApi.md#AssetPermissionsDelete) | **Post** /asset_permissions/delete | 资产权限回收


# **AssetPermissionsAdd**
> AssetPermissionsAddResponse AssetPermissionsAdd(ctx, data)
资产权限授予

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionsAddRequest**](AssetPermissionsAddRequest.md)|  | 

### Return type

[**AssetPermissionsAddResponse**](AssetPermissionsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetPermissionsDelete**
> AssetPermissionsDeleteResponse AssetPermissionsDelete(ctx, data)
资产权限回收

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPermissionsDeleteRequest**](AssetPermissionsDeleteRequest.md)|  | 

### Return type

[**AssetPermissionsDeleteResponse**](AssetPermissionsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


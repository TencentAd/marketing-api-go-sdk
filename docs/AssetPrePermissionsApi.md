# TencentAds\AssetPrePermissionsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetPrePermissionsGet**](AssetPrePermissionsApi.md#AssetPrePermissionsGet) | **Get** /asset_pre_permissions/get | 获取待确认接收授权列表接口
[**AssetPrePermissionsUpdate**](AssetPrePermissionsApi.md#AssetPrePermissionsUpdate) | **Post** /asset_pre_permissions/update | 资产预授权确认


# **AssetPrePermissionsGet**
> AssetPrePermissionsGetResponse AssetPrePermissionsGet(ctx, accountId, assetType, optional)
获取待确认接收授权列表接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **assetType** | **string**|  | 
 **optional** | ***AssetPrePermissionsApiAssetPrePermissionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetPrePermissionsApiAssetPrePermissionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assetId** | **optional.Int64**|  | 
 **assetName** | **optional.String**|  | 
 **pathType** | **optional.String**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AssetPrePermissionsGetResponse**](AssetPrePermissionsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetPrePermissionsUpdate**
> AssetPrePermissionsUpdateResponse AssetPrePermissionsUpdate(ctx, data)
资产预授权确认

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AssetPrePermissionsUpdateRequest**](AssetPrePermissionsUpdateRequest.md)|  | 

### Return type

[**AssetPrePermissionsUpdateResponse**](AssetPrePermissionsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


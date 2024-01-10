# TencentAds\OauthApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OauthAuthorize**](OauthApi.md#OauthAuthorize) | **Get** /oauth/authorize | 获取Authorization Code
[**OauthToken**](OauthApi.md#OauthToken) | **Get** /oauth/token | 通过Authorization Code获取Access Token或刷新Access Token


# **OauthAuthorize**
> string OauthAuthorize(ctx, clientId, redirectUri, optional)
获取Authorization Code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **int64**|  | 
  **redirectUri** | **string**|  | 
 **optional** | ***OauthApiOauthAuthorizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthApiOauthAuthorizeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **state** | **optional.String**|  | 
 **scope** | **optional.String**|  | 
 **accountType** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OauthToken**
> OauthTokenResponse OauthToken(ctx, clientId, clientSecret, grantType, optional)
通过Authorization Code获取Access Token或刷新Access Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **int64**|  | 
  **clientSecret** | **string**|  | 
  **grantType** | **string**|  | 
 **optional** | ***OauthApiOauthTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthApiOauthTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorizationCode** | **optional.String**|  | 
 **refreshToken** | **optional.String**| 当 grant_type&#x3D;refresh_token 时必填； | 
 **redirectUri** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**OauthTokenResponse**](OauthTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


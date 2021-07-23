# TencentAds\VideomakerVideocapturesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VideomakerVideocapturesAdd**](VideomakerVideocapturesApi.md#VideomakerVideocapturesAdd) | **Post** /videomaker_videocaptures/add | 生成视频封面图


# **VideomakerVideocapturesAdd**
> VideomakerVideocapturesAddResponse VideomakerVideocapturesAdd(ctx, accountId, optional)
生成视频封面图

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***VideomakerVideocapturesApiVideomakerVideocapturesAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VideomakerVideocapturesApiVideomakerVideocapturesAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **videoId** | **optional.String**|  | 
 **videoFile** | **optional.Interface of *os.File**|  | 
 **signature** | **optional.String**|  | 
 **number** | **optional.Int64**|  | 
 **returnImageIds** | **optional.Bool**|  | 

### Return type

[**VideomakerVideocapturesAddResponse**](VideomakerVideocapturesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# TencentAds\ChannelsCommentApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsCommentAdd**](ChannelsCommentApi.md#ChannelsCommentAdd) | **Post** /channels_comment/add | 回复视频号评论
[**ChannelsCommentDelete**](ChannelsCommentApi.md#ChannelsCommentDelete) | **Post** /channels_comment/delete | 删除视频号评论


# **ChannelsCommentAdd**
> ChannelsCommentAddResponse ChannelsCommentAdd(ctx, data)
回复视频号评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ChannelsCommentAddRequest**](ChannelsCommentAddRequest.md)|  | 

### Return type

[**ChannelsCommentAddResponse**](ChannelsCommentAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsCommentDelete**
> ChannelsCommentDeleteResponse ChannelsCommentDelete(ctx, data)
删除视频号评论

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ChannelsCommentDeleteRequest**](ChannelsCommentDeleteRequest.md)|  | 

### Return type

[**ChannelsCommentDeleteResponse**](ChannelsCommentDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


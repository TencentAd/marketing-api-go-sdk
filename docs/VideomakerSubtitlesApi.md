# TencentAds\VideomakerSubtitlesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VideomakerSubtitlesAdd**](VideomakerSubtitlesApi.md#VideomakerSubtitlesAdd) | **Post** /videomaker_subtitles/add | 创建视频加字幕任务


# **VideomakerSubtitlesAdd**
> VideomakerSubtitlesAddResponse VideomakerSubtitlesAdd(ctx, accountId, optional)
创建视频加字幕任务

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***VideomakerSubtitlesApiVideomakerSubtitlesAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VideomakerSubtitlesApiVideomakerSubtitlesAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **videoId** | **optional.String**|  | 
 **videoFile** | **optional.Interface of *os.File**|  | 
 **signature** | **optional.String**|  | 
 **onlySubtitleFile** | **optional.Bool**|  | 

### Return type

[**VideomakerSubtitlesAddResponse**](VideomakerSubtitlesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


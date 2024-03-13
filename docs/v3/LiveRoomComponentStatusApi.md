# TencentAds\LiveRoomComponentStatusApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LiveRoomComponentStatusUpdate**](LiveRoomComponentStatusApi.md#LiveRoomComponentStatusUpdate) | **Post** /live_room_component_status/update | 更新直播间组件状态


# **LiveRoomComponentStatusUpdate**
> LiveRoomComponentStatusUpdateResponse LiveRoomComponentStatusUpdate(ctx, data)
更新直播间组件状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LiveRoomComponentStatusUpdateRequest**](LiveRoomComponentStatusUpdateRequest.md)|  | 

### Return type

[**LiveRoomComponentStatusUpdateResponse**](LiveRoomComponentStatusUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


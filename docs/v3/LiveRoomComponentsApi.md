# TencentAds\LiveRoomComponentsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LiveRoomComponentsAdd**](LiveRoomComponentsApi.md#LiveRoomComponentsAdd) | **Post** /live_room_components/add | 创建直播间组件
[**LiveRoomComponentsDelete**](LiveRoomComponentsApi.md#LiveRoomComponentsDelete) | **Post** /live_room_components/delete | 删除直播间组件
[**LiveRoomComponentsGet**](LiveRoomComponentsApi.md#LiveRoomComponentsGet) | **Get** /live_room_components/get | 查询直播间组件信息
[**LiveRoomComponentsUpdate**](LiveRoomComponentsApi.md#LiveRoomComponentsUpdate) | **Post** /live_room_components/update | 更新直播间组件 


# **LiveRoomComponentsAdd**
> LiveRoomComponentsAddResponse LiveRoomComponentsAdd(ctx, data)
创建直播间组件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LiveRoomComponentsAddRequest**](LiveRoomComponentsAddRequest.md)|  | 

### Return type

[**LiveRoomComponentsAddResponse**](LiveRoomComponentsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LiveRoomComponentsDelete**
> LiveRoomComponentsDeleteResponse LiveRoomComponentsDelete(ctx, data)
删除直播间组件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LiveRoomComponentsDeleteRequest**](LiveRoomComponentsDeleteRequest.md)|  | 

### Return type

[**LiveRoomComponentsDeleteResponse**](LiveRoomComponentsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LiveRoomComponentsGet**
> LiveRoomComponentsGetResponse LiveRoomComponentsGet(ctx, accountId, optional)
查询直播间组件信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***LiveRoomComponentsApiLiveRoomComponentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LiveRoomComponentsApiLiveRoomComponentsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**LiveRoomComponentsGetResponse**](LiveRoomComponentsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LiveRoomComponentsUpdate**
> LiveRoomComponentsUpdateResponse LiveRoomComponentsUpdate(ctx, data)
更新直播间组件 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LiveRoomComponentsUpdateRequest**](LiveRoomComponentsUpdateRequest.md)|  | 

### Return type

[**LiveRoomComponentsUpdateResponse**](LiveRoomComponentsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


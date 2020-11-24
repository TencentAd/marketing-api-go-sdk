# TencentAds\VideomakerAutoadjustmentsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VideomakerAutoadjustmentsAdd**](VideomakerAutoadjustmentsApi.md#VideomakerAutoadjustmentsAdd) | **Post** /videomaker_autoadjustments/add | 创建智能调整任务


# **VideomakerAutoadjustmentsAdd**
> VideomakerAutoadjustmentsAddResponse VideomakerAutoadjustmentsAdd(ctx, accountId, adjustmentType, optional)
创建智能调整任务

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **adjustmentType** | **string**|  | 
 **optional** | ***VideomakerAutoadjustmentsApiVideomakerAutoadjustmentsAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VideomakerAutoadjustmentsApiVideomakerAutoadjustmentsAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **videoId** | **optional.String**|  | 
 **videoFile** | **optional.Interface of *os.File**|  | 
 **signature** | **optional.String**|  | 
 **smartAdjustment** | [**optional.Interface of SmartAdjustment**](SmartAdjustment.md)|  | 
 **manualAdjustment** | [**optional.Interface of ManualAdjustment**](ManualAdjustment.md)|  | 

### Return type

[**VideomakerAutoadjustmentsAddResponse**](VideomakerAutoadjustmentsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


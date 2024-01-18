# TencentAds\KeywordRecommendApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeywordRecommendGet**](KeywordRecommendApi.md#KeywordRecommendGet) | **Get** /keyword_recommend/get | 获取关键词推荐结果


# **KeywordRecommendGet**
> KeywordRecommendGetResponse KeywordRecommendGet(ctx, siteSets, recommendCategory, accountId, systemIndustryId, optional)
获取关键词推荐结果

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteSets** | [**[]string**](string.md)|  | 
  **recommendCategory** | **string**|  | 
  **accountId** | **int64**|  | 
  **systemIndustryId** | **int64**|  | 
 **optional** | ***KeywordRecommendApiKeywordRecommendGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KeywordRecommendApiKeywordRecommendGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **queryWord** | [**optional.Interface of []string**](string.md)|  | 
 **businessPointId** | **optional.String**|  | 
 **adgroupId** | **optional.Int64**|  | 
 **campaignId** | **optional.Int64**|  | 
 **includeWord** | [**optional.Interface of []string**](string.md)|  | 
 **excludeWord** | [**optional.Interface of []string**](string.md)|  | 
 **filterAdWord** | **optional.Bool**|  | 
 **filterAccountWord** | **optional.Bool**|  | 
 **recommendReasons** | [**optional.Interface of []string**](string.md)|  | 
 **province** | [**optional.Interface of []int64**](int64.md)|  | 
 **city** | [**optional.Interface of []int64**](int64.md)|  | 
 **orderBy** | [**optional.Interface of []OrderByStructInfo**](OrderByStructInfo.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**KeywordRecommendGetResponse**](KeywordRecommendGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


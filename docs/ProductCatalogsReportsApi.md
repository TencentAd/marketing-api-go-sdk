# TencentAds\ProductCatalogsReportsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductCatalogsReportsGet**](ProductCatalogsReportsApi.md#ProductCatalogsReportsGet) | **Get** /product_catalogs_reports/get | 获取商品库报表(待废弃)


# **ProductCatalogsReportsGet**
> ProductCatalogsReportsGetResponse ProductCatalogsReportsGet(ctx, accountId, productCatalogId, dateRange, optional)
获取商品库报表(待废弃)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **productCatalogId** | **int64**|  | 
  **dateRange** | [**ReportDateRange**](ReportDateRange.md)|  | 
 **optional** | ***ProductCatalogsReportsApiProductCatalogsReportsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductCatalogsReportsApiProductCatalogsReportsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **groupBy** | [**optional.Interface of []string**](string.md)|  | 
 **orderBy** | [**optional.Interface of []OrderByStruct**](OrderByStruct.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProductCatalogsReportsGetResponse**](ProductCatalogsReportsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


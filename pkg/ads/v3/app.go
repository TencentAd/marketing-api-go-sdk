/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ads

import (
	"github.com/tencentad/marketing-api-go-sdk/pkg/api/v3"
)

func (c *SDKClient) AccountVersion() *api.AccountVersionApiService {
	return c.Client.AccountVersionApi
}

func (c *SDKClient) AdLabel() *api.AdLabelApiService {
	return c.Client.AdLabelApi
}

func (c *SDKClient) AdParam() *api.AdParamApiService {
	return c.Client.AdParamApi
}

func (c *SDKClient) AdUnionReports() *api.AdUnionReportsApiService {
	return c.Client.AdUnionReportsApi
}

func (c *SDKClient) AdcreativePreviews() *api.AdcreativePreviewsApiService {
	return c.Client.AdcreativePreviewsApi
}

func (c *SDKClient) AdcreativePreviewsQrcode() *api.AdcreativePreviewsQrcodeApiService {
	return c.Client.AdcreativePreviewsQrcodeApi
}

func (c *SDKClient) AdgroupNegativewords() *api.AdgroupNegativewordsApiService {
	return c.Client.AdgroupNegativewordsApi
}

func (c *SDKClient) Adgroups() *api.AdgroupsApiService {
	return c.Client.AdgroupsApi
}

func (c *SDKClient) Advertiser() *api.AdvertiserApiService {
	return c.Client.AdvertiserApi
}

func (c *SDKClient) AdvertiserDailyBudget() *api.AdvertiserDailyBudgetApiService {
	return c.Client.AdvertiserDailyBudgetApi
}

func (c *SDKClient) Agency() *api.AgencyApiService {
	return c.Client.AgencyApi
}

func (c *SDKClient) AgencyBusinessUnit() *api.AgencyBusinessUnitApiService {
	return c.Client.AgencyBusinessUnitApi
}

func (c *SDKClient) AgencyBusinessUnitList() *api.AgencyBusinessUnitListApiService {
	return c.Client.AgencyBusinessUnitListApi
}

func (c *SDKClient) AgencyBusinessUnitListAccount() *api.AgencyBusinessUnitListAccountApiService {
	return c.Client.AgencyBusinessUnitListAccountApi
}

func (c *SDKClient) AgencyBusinessUnitListByAccount() *api.AgencyBusinessUnitListByAccountApiService {
	return c.Client.AgencyBusinessUnitListByAccountApi
}

func (c *SDKClient) AgencyRealtimeCost() *api.AgencyRealtimeCostApiService {
	return c.Client.AgencyRealtimeCostApi
}

func (c *SDKClient) AndroidChannel() *api.AndroidChannelApiService {
	return c.Client.AndroidChannelApi
}

func (c *SDKClient) AsyncReportFiles() *api.AsyncReportFilesApiService {
	return c.Client.AsyncReportFilesApi
}

func (c *SDKClient) AsyncReports() *api.AsyncReportsApiService {
	return c.Client.AsyncReportsApi
}

func (c *SDKClient) AsyncTasks() *api.AsyncTasksApiService {
	return c.Client.AsyncTasksApi
}

func (c *SDKClient) AudienceGrantRelations() *api.AudienceGrantRelationsApiService {
	return c.Client.AudienceGrantRelationsApi
}

func (c *SDKClient) Barrage() *api.BarrageApiService {
	return c.Client.BarrageApi
}

func (c *SDKClient) BarrageRecommend() *api.BarrageRecommendApiService {
	return c.Client.BarrageRecommendApi
}

func (c *SDKClient) BatchAsyncRequestSpecification() *api.BatchAsyncRequestSpecificationApiService {
	return c.Client.BatchAsyncRequestSpecificationApi
}

func (c *SDKClient) BatchAsyncRequests() *api.BatchAsyncRequestsApiService {
	return c.Client.BatchAsyncRequestsApi
}

func (c *SDKClient) BatchRequests() *api.BatchRequestsApiService {
	return c.Client.BatchRequestsApi
}

func (c *SDKClient) BidSimulation() *api.BidSimulationApiService {
	return c.Client.BidSimulationApi
}

func (c *SDKClient) Bidword() *api.BidwordApiService {
	return c.Client.BidwordApi
}

func (c *SDKClient) BidwordFlow() *api.BidwordFlowApiService {
	return c.Client.BidwordFlowApi
}

func (c *SDKClient) Brand() *api.BrandApiService {
	return c.Client.BrandApi
}

func (c *SDKClient) BusinessPoint() *api.BusinessPointApiService {
	return c.Client.BusinessPointApi
}

func (c *SDKClient) Categories() *api.CategoriesApiService {
	return c.Client.CategoriesApi
}

func (c *SDKClient) CategoriesAttribute() *api.CategoriesAttributeApiService {
	return c.Client.CategoriesAttributeApi
}

func (c *SDKClient) ChannelsComment() *api.ChannelsCommentApiService {
	return c.Client.ChannelsCommentApi
}

func (c *SDKClient) ChannelsFinderobject() *api.ChannelsFinderobjectApiService {
	return c.Client.ChannelsFinderobjectApi
}

func (c *SDKClient) ChannelsLivenoticeinfo() *api.ChannelsLivenoticeinfoApiService {
	return c.Client.ChannelsLivenoticeinfoApi
}

func (c *SDKClient) ChannelsUserpageobjects() *api.ChannelsUserpageobjectsApiService {
	return c.Client.ChannelsUserpageobjectsApi
}

func (c *SDKClient) CommentList() *api.CommentListApiService {
	return c.Client.CommentListApi
}

func (c *SDKClient) ComponentElementUrgeReview() *api.ComponentElementUrgeReviewApiService {
	return c.Client.ComponentElementUrgeReviewApi
}

func (c *SDKClient) ComponentReviewResults() *api.ComponentReviewResultsApiService {
	return c.Client.ComponentReviewResultsApi
}

func (c *SDKClient) ComponentSharing() *api.ComponentSharingApiService {
	return c.Client.ComponentSharingApi
}

func (c *SDKClient) Components() *api.ComponentsApiService {
	return c.Client.ComponentsApi
}

func (c *SDKClient) ComponentsMetadata() *api.ComponentsMetadataApiService {
	return c.Client.ComponentsMetadataApi
}

func (c *SDKClient) ConversionLinkAssetAvailable() *api.ConversionLinkAssetAvailableApiService {
	return c.Client.ConversionLinkAssetAvailableApi
}

func (c *SDKClient) ConversionLinkAssets() *api.ConversionLinkAssetsApiService {
	return c.Client.ConversionLinkAssetsApi
}

func (c *SDKClient) ConversionLinks() *api.ConversionLinksApiService {
	return c.Client.ConversionLinksApi
}

func (c *SDKClient) Conversions() *api.ConversionsApiService {
	return c.Client.ConversionsApi
}

func (c *SDKClient) CreativeTemplate() *api.CreativeTemplateApiService {
	return c.Client.CreativeTemplateApi
}

func (c *SDKClient) CreativeTemplateList() *api.CreativeTemplateListApiService {
	return c.Client.CreativeTemplateListApi
}

func (c *SDKClient) CreativeTemplatePreviews() *api.CreativeTemplatePreviewsApiService {
	return c.Client.CreativeTemplatePreviewsApi
}

func (c *SDKClient) CreativetoolsText() *api.CreativetoolsTextApiService {
	return c.Client.CreativetoolsTextApi
}

func (c *SDKClient) CustomAudienceEstimations() *api.CustomAudienceEstimationsApiService {
	return c.Client.CustomAudienceEstimationsApi
}

func (c *SDKClient) CustomAudienceFiles() *api.CustomAudienceFilesApiService {
	return c.Client.CustomAudienceFilesApi
}

func (c *SDKClient) CustomAudiences() *api.CustomAudiencesApiService {
	return c.Client.CustomAudiencesApi
}

func (c *SDKClient) DailyBalanceReport() *api.DailyBalanceReportApiService {
	return c.Client.DailyBalanceReportApi
}

func (c *SDKClient) DailyReports() *api.DailyReportsApiService {
	return c.Client.DailyReportsApi
}

func (c *SDKClient) DataSourceDispatch() *api.DataSourceDispatchApiService {
	return c.Client.DataSourceDispatchApi
}

func (c *SDKClient) DynamicAdImageTemplates() *api.DynamicAdImageTemplatesApiService {
	return c.Client.DynamicAdImageTemplatesApi
}

func (c *SDKClient) DynamicAdImages() *api.DynamicAdImagesApiService {
	return c.Client.DynamicAdImagesApi
}

func (c *SDKClient) DynamicAdVideo() *api.DynamicAdVideoApiService {
	return c.Client.DynamicAdVideoApi
}

func (c *SDKClient) DynamicAdVideoTemplates() *api.DynamicAdVideoTemplatesApiService {
	return c.Client.DynamicAdVideoTemplatesApi
}

func (c *SDKClient) DynamicCreativeReviewResults() *api.DynamicCreativeReviewResultsApiService {
	return c.Client.DynamicCreativeReviewResultsApi
}

func (c *SDKClient) DynamicCreatives() *api.DynamicCreativesApiService {
	return c.Client.DynamicCreativesApi
}

func (c *SDKClient) EcommerceOrder() *api.EcommerceOrderApiService {
	return c.Client.EcommerceOrderApi
}

func (c *SDKClient) ElementAppealQuota() *api.ElementAppealQuotaApiService {
	return c.Client.ElementAppealQuotaApi
}

func (c *SDKClient) ElementAppealReview() *api.ElementAppealReviewApiService {
	return c.Client.ElementAppealReviewApi
}

func (c *SDKClient) Estimation() *api.EstimationApiService {
	return c.Client.EstimationApi
}

func (c *SDKClient) ExtendPackage() *api.ExtendPackageApiService {
	return c.Client.ExtendPackageApi
}

func (c *SDKClient) FinderAdObjectList() *api.FinderAdObjectListApiService {
	return c.Client.FinderAdObjectListApi
}

func (c *SDKClient) FundStatementsDetailed() *api.FundStatementsDetailedApiService {
	return c.Client.FundStatementsDetailedApi
}

func (c *SDKClient) FundTransfer() *api.FundTransferApiService {
	return c.Client.FundTransferApi
}

func (c *SDKClient) Funds() *api.FundsApiService {
	return c.Client.FundsApi
}

func (c *SDKClient) GameFeature() *api.GameFeatureApiService {
	return c.Client.GameFeatureApi
}

func (c *SDKClient) GameFeatureTags() *api.GameFeatureTagsApiService {
	return c.Client.GameFeatureTagsApi
}

func (c *SDKClient) GetWxGameAppGiftPack() *api.GetWxGameAppGiftPackApiService {
	return c.Client.GetWxGameAppGiftPackApi
}

func (c *SDKClient) HourlyReports() *api.HourlyReportsApiService {
	return c.Client.HourlyReportsApi
}

func (c *SDKClient) ImageProcessing() *api.ImageProcessingApiService {
	return c.Client.ImageProcessingApi
}

func (c *SDKClient) Images() *api.ImagesApiService {
	return c.Client.ImagesApi
}

func (c *SDKClient) JointBudgetRules() *api.JointBudgetRulesApiService {
	return c.Client.JointBudgetRulesApi
}

func (c *SDKClient) KeywordRecommend() *api.KeywordRecommendApiService {
	return c.Client.KeywordRecommendApi
}

func (c *SDKClient) Labels() *api.LabelsApiService {
	return c.Client.LabelsApi
}

func (c *SDKClient) LandingPageSellStrategy() *api.LandingPageSellStrategyApiService {
	return c.Client.LandingPageSellStrategyApi
}

func (c *SDKClient) Leads() *api.LeadsApiService {
	return c.Client.LeadsApi
}

func (c *SDKClient) LeadsActionTypeReport() *api.LeadsActionTypeReportApiService {
	return c.Client.LeadsActionTypeReportApi
}

func (c *SDKClient) LeadsCallRecord() *api.LeadsCallRecordApiService {
	return c.Client.LeadsCallRecordApi
}

func (c *SDKClient) LeadsCallRecords() *api.LeadsCallRecordsApiService {
	return c.Client.LeadsCallRecordsApi
}

func (c *SDKClient) LeadsCallVirtualNumber() *api.LeadsCallVirtualNumberApiService {
	return c.Client.LeadsCallVirtualNumberApi
}

func (c *SDKClient) LeadsClaim() *api.LeadsClaimApiService {
	return c.Client.LeadsClaimApi
}

func (c *SDKClient) LeadsInvalidPay() *api.LeadsInvalidPayApiService {
	return c.Client.LeadsInvalidPayApi
}

func (c *SDKClient) LeadsList() *api.LeadsListApiService {
	return c.Client.LeadsListApi
}

func (c *SDKClient) LeadsStatus() *api.LeadsStatusApiService {
	return c.Client.LeadsStatusApi
}

func (c *SDKClient) LeadsVoipCall() *api.LeadsVoipCallApiService {
	return c.Client.LeadsVoipCallApi
}

func (c *SDKClient) LeadsVoipCallToken() *api.LeadsVoipCallTokenApiService {
	return c.Client.LeadsVoipCallTokenApi
}

func (c *SDKClient) LiveRoomComponentStatus() *api.LiveRoomComponentStatusApiService {
	return c.Client.LiveRoomComponentStatusApi
}

func (c *SDKClient) LiveRoomComponents() *api.LiveRoomComponentsApiService {
	return c.Client.LiveRoomComponentsApi
}

func (c *SDKClient) LocalStorePackages() *api.LocalStorePackagesApiService {
	return c.Client.LocalStorePackagesApi
}

func (c *SDKClient) LocalStores() *api.LocalStoresApiService {
	return c.Client.LocalStoresApi
}

func (c *SDKClient) LocalStoresAddressParsingResult() *api.LocalStoresAddressParsingResultApiService {
	return c.Client.LocalStoresAddressParsingResultApi
}

func (c *SDKClient) LocalStoresCategories() *api.LocalStoresCategoriesApiService {
	return c.Client.LocalStoresCategoriesApi
}

func (c *SDKClient) LocalStoresSearchInfo() *api.LocalStoresSearchInfoApiService {
	return c.Client.LocalStoresSearchInfoApi
}

func (c *SDKClient) LocalStoresWxpayMerchants() *api.LocalStoresWxpayMerchantsApiService {
	return c.Client.LocalStoresWxpayMerchantsApi
}

func (c *SDKClient) MarketingRules() *api.MarketingRulesApiService {
	return c.Client.MarketingRulesApi
}

func (c *SDKClient) MarketingTargetAssetCategories() *api.MarketingTargetAssetCategoriesApiService {
	return c.Client.MarketingTargetAssetCategoriesApi
}

func (c *SDKClient) MarketingTargetAssetDetail() *api.MarketingTargetAssetDetailApiService {
	return c.Client.MarketingTargetAssetDetailApi
}

func (c *SDKClient) MarketingTargetAssetProperties() *api.MarketingTargetAssetPropertiesApiService {
	return c.Client.MarketingTargetAssetPropertiesApi
}

func (c *SDKClient) MarketingTargetAssetPropertyValues() *api.MarketingTargetAssetPropertyValuesApiService {
	return c.Client.MarketingTargetAssetPropertyValuesApi
}

func (c *SDKClient) MarketingTargetAssets() *api.MarketingTargetAssetsApiService {
	return c.Client.MarketingTargetAssetsApi
}

func (c *SDKClient) MarketingTargetTypes() *api.MarketingTargetTypesApiService {
	return c.Client.MarketingTargetTypesApi
}

func (c *SDKClient) MaterialDcaset() *api.MaterialDcasetApiService {
	return c.Client.MaterialDcasetApi
}

func (c *SDKClient) MaterialDcatag() *api.MaterialDcatagApiService {
	return c.Client.MaterialDcatagApi
}

func (c *SDKClient) MaterialLabels() *api.MaterialLabelsApiService {
	return c.Client.MaterialLabelsApi
}

func (c *SDKClient) MergeFundTypeDailyBalanceReport() *api.MergeFundTypeDailyBalanceReportApiService {
	return c.Client.MergeFundTypeDailyBalanceReportApi
}

func (c *SDKClient) MergeFundTypeFundStatementsDetailed() *api.MergeFundTypeFundStatementsDetailedApiService {
	return c.Client.MergeFundTypeFundStatementsDetailedApi
}

func (c *SDKClient) MergeFundTypeFunds() *api.MergeFundTypeFundsApiService {
	return c.Client.MergeFundTypeFundsApi
}

func (c *SDKClient) MergeFundTypeSubcustomerTransfer() *api.MergeFundTypeSubcustomerTransferApiService {
	return c.Client.MergeFundTypeSubcustomerTransferApi
}

func (c *SDKClient) MuseAiMaterial() *api.MuseAiMaterialApiService {
	return c.Client.MuseAiMaterialApi
}

func (c *SDKClient) MuseAiTask() *api.MuseAiTaskApiService {
	return c.Client.MuseAiTaskApi
}

func (c *SDKClient) MuseAiUgc() *api.MuseAiUgcApiService {
	return c.Client.MuseAiUgcApi
}

func (c *SDKClient) Oauth() *api.OauthApiService {
	return c.Client.OauthApi
}

func (c *SDKClient) ObjectCommentFlag() *api.ObjectCommentFlagApiService {
	return c.Client.ObjectCommentFlagApi
}

func (c *SDKClient) OperationLogList() *api.OperationLogListApiService {
	return c.Client.OperationLogListApi
}

func (c *SDKClient) OptimizationGoalPermissions() *api.OptimizationGoalPermissionsApiService {
	return c.Client.OptimizationGoalPermissionsApi
}

func (c *SDKClient) OrganizationAccountRelation() *api.OrganizationAccountRelationApiService {
	return c.Client.OrganizationAccountRelationApi
}

func (c *SDKClient) Pages() *api.PagesApiService {
	return c.Client.PagesApi
}

func (c *SDKClient) ProductCatalogs() *api.ProductCatalogsApiService {
	return c.Client.ProductCatalogsApi
}

func (c *SDKClient) ProductCategoriesList() *api.ProductCategoriesListApiService {
	return c.Client.ProductCategoriesListApi
}

func (c *SDKClient) ProductItems() *api.ProductItemsApiService {
	return c.Client.ProductItemsApi
}

func (c *SDKClient) ProductItemsDetail() *api.ProductItemsDetailApiService {
	return c.Client.ProductItemsDetailApi
}

func (c *SDKClient) ProductItemsVerticals() *api.ProductItemsVerticalsApiService {
	return c.Client.ProductItemsVerticalsApi
}

func (c *SDKClient) ProductSeries() *api.ProductSeriesApiService {
	return c.Client.ProductSeriesApi
}

func (c *SDKClient) ProductsSystemStatus() *api.ProductsSystemStatusApiService {
	return c.Client.ProductsSystemStatusApi
}

func (c *SDKClient) Profiles() *api.ProfilesApiService {
	return c.Client.ProfilesApi
}

func (c *SDKClient) Programmed() *api.ProgrammedApiService {
	return c.Client.ProgrammedApi
}

func (c *SDKClient) ProgrammedMaterialMappings() *api.ProgrammedMaterialMappingsApiService {
	return c.Client.ProgrammedMaterialMappingsApi
}

func (c *SDKClient) ProgrammedTemplate() *api.ProgrammedTemplateApiService {
	return c.Client.ProgrammedTemplateApi
}

func (c *SDKClient) QualificationImages() *api.QualificationImagesApiService {
	return c.Client.QualificationImagesApi
}

func (c *SDKClient) QualificationStructure() *api.QualificationStructureApiService {
	return c.Client.QualificationStructureApi
}

func (c *SDKClient) Qualifications() *api.QualificationsApiService {
	return c.Client.QualificationsApi
}

func (c *SDKClient) RealtimeCost() *api.RealtimeCostApiService {
	return c.Client.RealtimeCostApi
}

func (c *SDKClient) ReviewElementPrereviewResults() *api.ReviewElementPrereviewResultsApiService {
	return c.Client.ReviewElementPrereviewResultsApi
}

func (c *SDKClient) RtaInfo() *api.RtaInfoApiService {
	return c.Client.RtaInfoApi
}

func (c *SDKClient) Rtaexp() *api.RtaexpApiService {
	return c.Client.RtaexpApi
}

func (c *SDKClient) RtaexpDataRoi() *api.RtaexpDataRoiApiService {
	return c.Client.RtaexpDataRoiApi
}

func (c *SDKClient) RtaexpDspTagData() *api.RtaexpDspTagDataApiService {
	return c.Client.RtaexpDspTagDataApi
}

func (c *SDKClient) Rtatarget() *api.RtatargetApiService {
	return c.Client.RtatargetApi
}

func (c *SDKClient) RtatargetBind() *api.RtatargetBindApiService {
	return c.Client.RtatargetBindApi
}

func (c *SDKClient) SceneSpecTags() *api.SceneSpecTagsApiService {
	return c.Client.SceneSpecTagsApi
}

func (c *SDKClient) SubcustomerTransfer() *api.SubcustomerTransferApiService {
	return c.Client.SubcustomerTransferApi
}

func (c *SDKClient) TargetingTagReports() *api.TargetingTagReportsApiService {
	return c.Client.TargetingTagReportsApi
}

func (c *SDKClient) TargetingTags() *api.TargetingTagsApiService {
	return c.Client.TargetingTagsApi
}

func (c *SDKClient) TargetingTagsUv() *api.TargetingTagsUvApiService {
	return c.Client.TargetingTagsUvApi
}

func (c *SDKClient) Targetings() *api.TargetingsApiService {
	return c.Client.TargetingsApi
}

func (c *SDKClient) UnionPositionPackages() *api.UnionPositionPackagesApiService {
	return c.Client.UnionPositionPackagesApi
}

func (c *SDKClient) UserActionSetReports() *api.UserActionSetReportsApiService {
	return c.Client.UserActionSetReportsApi
}

func (c *SDKClient) UserActionSets() *api.UserActionSetsApiService {
	return c.Client.UserActionSetsApi
}

func (c *SDKClient) UserActions() *api.UserActionsApiService {
	return c.Client.UserActionsApi
}

func (c *SDKClient) VideoChannelDealerData() *api.VideoChannelDealerDataApiService {
	return c.Client.VideoChannelDealerDataApi
}

func (c *SDKClient) VideoChannelFansData() *api.VideoChannelFansDataApiService {
	return c.Client.VideoChannelFansDataApi
}

func (c *SDKClient) VideoChannelLeadsData() *api.VideoChannelLeadsDataApiService {
	return c.Client.VideoChannelLeadsDataApi
}

func (c *SDKClient) VideoChannelLiveData() *api.VideoChannelLiveDataApiService {
	return c.Client.VideoChannelLiveDataApi
}

func (c *SDKClient) Videos() *api.VideosApiService {
	return c.Client.VideosApi
}

func (c *SDKClient) Wallet() *api.WalletApiService {
	return c.Client.WalletApi
}

func (c *SDKClient) WechatChannelsAccounts() *api.WechatChannelsAccountsApiService {
	return c.Client.WechatChannelsAccountsApi
}

func (c *SDKClient) WechatChannelsAdAccount() *api.WechatChannelsAdAccountApiService {
	return c.Client.WechatChannelsAdAccountApi
}

func (c *SDKClient) WechatChannelsAdAccountCertificationFile() *api.WechatChannelsAdAccountCertificationFileApiService {
	return c.Client.WechatChannelsAdAccountCertificationFileApi
}

func (c *SDKClient) WechatChannelsAdAccountValidation() *api.WechatChannelsAdAccountValidationApiService {
	return c.Client.WechatChannelsAdAccountValidationApi
}

func (c *SDKClient) WechatChannelsAdAccountWechatBinding() *api.WechatChannelsAdAccountWechatBindingApiService {
	return c.Client.WechatChannelsAdAccountWechatBindingApi
}

func (c *SDKClient) WechatChannelsAuthorization() *api.WechatChannelsAuthorizationApiService {
	return c.Client.WechatChannelsAuthorizationApi
}

func (c *SDKClient) WechatOfficialAccounts() *api.WechatOfficialAccountsApiService {
	return c.Client.WechatOfficialAccountsApi
}

func (c *SDKClient) WechatPages() *api.WechatPagesApiService {
	return c.Client.WechatPagesApi
}

func (c *SDKClient) WechatPagesCsgroupStatus() *api.WechatPagesCsgroupStatusApiService {
	return c.Client.WechatPagesCsgroupStatusApi
}

func (c *SDKClient) WechatPagesCsgroupUser() *api.WechatPagesCsgroupUserApiService {
	return c.Client.WechatPagesCsgroupUserApi
}

func (c *SDKClient) WechatPagesCsgrouplist() *api.WechatPagesCsgrouplistApiService {
	return c.Client.WechatPagesCsgrouplistApi
}

func (c *SDKClient) WechatPagesCustom() *api.WechatPagesCustomApiService {
	return c.Client.WechatPagesCustomApi
}

func (c *SDKClient) WechatPagesGrantinfo() *api.WechatPagesGrantinfoApiService {
	return c.Client.WechatPagesGrantinfoApi
}

func (c *SDKClient) Wildcards() *api.WildcardsApiService {
	return c.Client.WildcardsApi
}

func (c *SDKClient) WxGamePlayablePage() *api.WxGamePlayablePageApiService {
	return c.Client.WxGamePlayablePageApi
}

func (c *SDKClient) WxPackageAccount() *api.WxPackageAccountApiService {
	return c.Client.WxPackageAccountApi
}

func (c *SDKClient) WxPackagePackage() *api.WxPackagePackageApiService {
	return c.Client.WxPackagePackageApi
}

func (c *SDKClient) XijingComplexTemplate() *api.XijingComplexTemplateApiService {
	return c.Client.XijingComplexTemplateApi
}

func (c *SDKClient) XijingPage() *api.XijingPageApiService {
	return c.Client.XijingPageApi
}

func (c *SDKClient) XijingPageByComponents() *api.XijingPageByComponentsApiService {
	return c.Client.XijingPageByComponentsApi
}

func (c *SDKClient) XijingPageInteractive() *api.XijingPageInteractiveApiService {
	return c.Client.XijingPageInteractiveApi
}

func (c *SDKClient) XijingPageList() *api.XijingPageListApiService {
	return c.Client.XijingPageListApi
}

func (c *SDKClient) XijingTemplate() *api.XijingTemplateApiService {
	return c.Client.XijingTemplateApi
}

func (c *SDKClient) XijingTemplateList() *api.XijingTemplateListApiService {
	return c.Client.XijingTemplateListApi
}

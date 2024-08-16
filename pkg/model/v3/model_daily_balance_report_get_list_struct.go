/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type DailyBalanceReportGetListStruct struct {
	AccountId       *int64         `json:"account_id,omitempty"`
	FundType        AccountTypeMap `json:"fund_type,omitempty"`
	Time            *int64         `json:"time,omitempty"`
	Deposit         *int64         `json:"deposit,omitempty"`
	Paid            *int64         `json:"paid,omitempty"`
	TransIn         *int64         `json:"trans_in,omitempty"`
	TransOut        *int64         `json:"trans_out,omitempty"`
	CreditModify    *int64         `json:"credit_modify,omitempty"`
	Balance         *int64         `json:"balance,omitempty"`
	PreauthBalance  *int64         `json:"preauth_balance,omitempty"`
	PreauthOutPay   *int64         `json:"preauth_out_pay,omitempty"`
	PreauthInRefund *int64         `json:"preauth_in_refund,omitempty"`
	AcctOutPay      *int64         `json:"acct_out_pay,omitempty"`
	AcctOutPayShare *int64         `json:"acct_out_pay_share,omitempty"`
	ShareOutPay     *int64         `json:"share_out_pay,omitempty"`
}

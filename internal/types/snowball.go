package types

type Suggest struct {
	Code int `json:"code"`
	Data []struct {
		Code      string `json:"code"`
		Label     string `json:"label"`
		Query     string `json:"query"`
		State     int    `json:"state"`
		StockType int    `json:"stock_type"`
		Type      int    `json:"type"`
	} `json:"data"`
	Message string `json:"message"`
	Meta    struct {
		Count       int   `json:"count"`
		Feedback    int   `json:"feedback"`
		HasNextPage bool  `json:"has_next_page"`
		MaxPage     int   `json:"maxPage"`
		Page        int   `json:"page"`
		QueryID     int64 `json:"query_id"`
		Size        int   `json:"size"`
	} `json:"meta"`
	Success bool `json:"success"`
}

type Kline struct {
	Data struct {
		Symbol string      `json:"symbol"`
		Column []string    `json:"column"`
		Item   [][]float64 `json:"item"`
	} `json:"data"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

/*
	type Indicator struct {
		Data struct {
			QuoteName      string `json:"quote_name"`
			CurrencyName   string `json:"currency_name"`
			OrgType        int    `json:"org_type"`
			LastReportName string `json:"last_report_name"`
			Statuses       any    `json:"statuses"`
			Currency       string `json:"currency"`
			List           []struct {
				ReportDate                  int64     `json:"report_date,omitempty"`
				ReportName                  string    `json:"report_name,omitempty"`
				Ctime                       int64     `json:"ctime,omitempty"`
				AvgRoe                      []float64 `json:"avg_roe,omitempty"`
				NpPerShare                  []float64 `json:"np_per_share,omitempty"`
				OperateCashFlowPs           []float64 `json:"operate_cash_flow_ps,omitempty"`
				BasicEps                    []float64 `json:"basic_eps,omitempty"`
				CapitalReserve              []float64 `json:"capital_reserve,omitempty"`
				UndistriProfitPs            []float64 `json:"undistri_profit_ps,omitempty"`
				NetInterestOfTotalAssets    []float64 `json:"net_interest_of_total_assets,omitempty"`
				NetSellingRate              []float64 `json:"net_selling_rate,omitempty"`
				GrossSellingRate            []float64 `json:"gross_selling_rate,omitempty"`
				TotalRevenue                []float64 `json:"total_revenue,omitempty"`
				OperatingIncomeYoy          []float64 `json:"operating_income_yoy,omitempty"`
				NetProfitAtsopc             []float64 `json:"net_profit_atsopc,omitempty"`
				NetProfitAtsopcYoy          []float64 `json:"net_profit_atsopc_yoy,omitempty"`
				NetProfitAfterNrgalAtsolc   []float64 `json:"net_profit_after_nrgal_atsolc,omitempty"`
				NpAtsopcNrgalYoy            []float64 `json:"np_atsopc_nrgal_yoy,omitempty"`
				OreDlt                      []float64 `json:"ore_dlt,omitempty"`
				Rop                         []float64 `json:"rop,omitempty"`
				AssetLiabRatio              []float64 `json:"asset_liab_ratio,omitempty"`
				CurrentRatio                []float64 `json:"current_ratio,omitempty"`
				QuickRatio                  []float64 `json:"quick_ratio,omitempty"`
				EquityMultiplier            []float64 `json:"equity_multiplier,omitempty"`
				EquityRatio                 []float64 `json:"equity_ratio,omitempty"`
				HolderEquity                []float64 `json:"holder_equity,omitempty"`
				NcfFromOaToTotalLiab        []float64 `json:"ncf_from_oa_to_total_liab,omitempty"`
				InventoryTurnoverDays       []float64 `json:"inventory_turnover_days,omitempty"`
				ReceivableTurnoverDays      []float64 `json:"receivable_turnover_days,omitempty"`
				AccountsPayableTurnoverDays []float64 `json:"accounts_payable_turnover_days,omitempty"`
				CashCycle                   []float64 `json:"cash_cycle,omitempty"`
				OperatingCycle              []float64 `json:"operating_cycle,omitempty"`
				TotalCapitalTurnover        []float64 `json:"total_capital_turnover,omitempty"`
				InventoryTurnover           []float64 `json:"inventory_turnover,omitempty"`
				AccountReceivableTurnover   []float64 `json:"account_receivable_turnover,omitempty"`
				AccountsPayableTurnover     []float64 `json:"accounts_payable_turnover,omitempty"`
				CurrentAssetTurnoverRate    []float64 `json:"current_asset_turnover_rate,omitempty"`
				FixedAssetTurnoverRatio     []float64 `json:"fixed_asset_turnover_ratio,omitempty"`
			} `json:"list"`
		} `json:"data"`
		ErrorCode        int    `json:"error_code"`
		ErrorDescription string `json:"error_description"`
	}
*/
type Indicator struct {
	Data struct {
		QuoteName      string                   `json:"quote_name"`
		CurrencyName   string                   `json:"currency_name"`
		OrgType        int                      `json:"org_type"`
		LastReportName string                   `json:"last_report_name"`
		Statuses       any                      `json:"statuses"`
		Currency       string                   `json:"currency"`
		List           []map[string]interface{} `json:"list"`
	} `json:"data"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

/*
	type Income struct {
		Data struct {
			QuoteName      string `json:"quote_name,omitempty"`
			CurrencyName   string `json:"currency_name,omitempty"`
			OrgType        int    `json:"org_type,omitempty"`
			LastReportName string `json:"last_report_name,omitempty"`
			Statuses       any    `json:"statuses,omitempty"`
			Currency       string `json:"currency,omitempty"`
			List           []struct {
				ReportDate                  int64     `json:"report_date,omitempty"`
				ReportName                  string    `json:"report_name,omitempty"`
				Ctime                       int64     `json:"ctime,omitempty"`
				NetProfit                   []float64 `json:"net_profit,omitempty"`
				NetProfitAtsopc             []float64 `json:"net_profit_atsopc,omitempty"`
				TotalRevenue                []float64 `json:"total_revenue,omitempty"`
				Op                          []float64 `json:"op,omitempty"`
				IncomeFromChgInFv           []float64 `json:"income_from_chg_in_fv,omitempty"`
				InvestIncomesFromRr         []float64 `json:"invest_incomes_from_rr,omitempty"`
				InvestIncome                []float64 `json:"invest_income,omitempty"`
				ExchgGain                   []any     `json:"exchg_gain,omitempty"`
				OperatingTaxesAndSurcharge  []float64 `json:"operating_taxes_and_surcharge,omitempty"`
				AssetImpairmentLoss         []float64 `json:"asset_impairment_loss,omitempty"`
				NonOperatingIncome          []float64 `json:"non_operating_income,omitempty"`
				NonOperatingPayout          []float64 `json:"non_operating_payout,omitempty"`
				ProfitTotalAmt              []float64 `json:"profit_total_amt,omitempty"`
				MinorityGal                 []any     `json:"minority_gal,omitempty"`
				BasicEps                    []float64 `json:"basic_eps,omitempty"`
				DltEarningsPerShare         []float64 `json:"dlt_earnings_per_share,omitempty"`
				OthrCompreIncomeAtoopc      []float64 `json:"othr_compre_income_atoopc,omitempty"`
				OthrCompreIncomeAtms        []any     `json:"othr_compre_income_atms,omitempty"`
				TotalCompreIncome           []float64 `json:"total_compre_income,omitempty"`
				TotalCompreIncomeAtsopc     []float64 `json:"total_compre_income_atsopc,omitempty"`
				TotalCompreIncomeAtms       []any     `json:"total_compre_income_atms,omitempty"`
				OthrCompreIncome            []float64 `json:"othr_compre_income,omitempty"`
				NetProfitAfterNrgalAtsolc   []float64 `json:"net_profit_after_nrgal_atsolc,omitempty"`
				IncomeTaxExpenses           []float64 `json:"income_tax_expenses,omitempty"`
				CreditImpairmentLoss        []float64 `json:"credit_impairment_loss,omitempty"`
				Revenue                     []float64 `json:"revenue,omitempty"`
				OperatingCosts              []float64 `json:"operating_costs,omitempty"`
				OperatingCost               []float64 `json:"operating_cost,omitempty"`
				SalesFee                    []float64 `json:"sales_fee,omitempty"`
				ManageFee                   []float64 `json:"manage_fee,omitempty"`
				FinancingExpenses           []float64 `json:"financing_expenses,omitempty"`
				RadCost                     []float64 `json:"rad_cost,omitempty"`
				FinanceCostInterestFee      []float64 `json:"finance_cost_interest_fee,omitempty"`
				FinanceCostInterestIncome   []float64 `json:"finance_cost_interest_income,omitempty"`
				AssetDisposalIncome         []float64 `json:"asset_disposal_income,omitempty"`
				OtherIncome                 []float64 `json:"other_income,omitempty"`
				NoncurrentAssetsDisposeGain []any     `json:"noncurrent_assets_dispose_gain,omitempty"`
				NoncurrentAssetDisposalLoss []any     `json:"noncurrent_asset_disposal_loss,omitempty"`
				NetProfitBi                 []any     `json:"net_profit_bi,omitempty"`
				ContinousOperatingNp        []float64 `json:"continous_operating_np,omitempty"`
			} `json:"list,omitempty"`
		} `json:"data,omitempty"`
		ErrorCode        int    `json:"error_code,omitempty"`
		ErrorDescription string `json:"error_description,omitempty"`
	}
*/
type Income struct {
	Data struct {
		QuoteName      string                   `json:"quote_name,omitempty"`
		CurrencyName   string                   `json:"currency_name,omitempty"`
		OrgType        int                      `json:"org_type,omitempty"`
		LastReportName string                   `json:"last_report_name,omitempty"`
		Statuses       any                      `json:"statuses,omitempty"`
		Currency       string                   `json:"currency,omitempty"`
		List           []map[string]interface{} `json:"list,omitempty"`
	} `json:"data,omitempty"`
	ErrorCode        int    `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

/*
	type Balance struct {
		Data struct {
			QuoteName      string `json:"quote_name,omitempty"`
			CurrencyName   string `json:"currency_name,omitempty"`
			OrgType        int    `json:"org_type,omitempty"`
			LastReportName string `json:"last_report_name,omitempty"`
			Statuses       any    `json:"statuses,omitempty"`
			Currency       string `json:"currency,omitempty"`
			List           []struct {
				ReportDate                 int64     `json:"report_date,omitempty"`
				ReportName                 string    `json:"report_name,omitempty"`
				Ctime                      any       `json:"ctime,omitempty"`
				TotalAssets                []float64 `json:"total_assets,omitempty"`
				TotalLiab                  []float64 `json:"total_liab,omitempty"`
				AssetLiabRatio             []float64 `json:"asset_liab_ratio,omitempty"`
				TotalQuityAtsopc           []float64 `json:"total_quity_atsopc,omitempty"`
				TradableFnnclAssets        []float64 `json:"tradable_fnncl_assets,omitempty"`
				InterestReceivable         []any     `json:"interest_receivable,omitempty"`
				SaleableFinacialAssets     []any     `json:"saleable_finacial_assets,omitempty"`
				HeldToMaturityInvest       []any     `json:"held_to_maturity_invest,omitempty"`
				FixedAsset                 []any     `json:"fixed_asset,omitempty"`
				IntangibleAssets           []float64 `json:"intangible_assets,omitempty"`
				ConstructionInProcess      []any     `json:"construction_in_process,omitempty"`
				DtAssets                   []float64 `json:"dt_assets,omitempty"`
				TradableFnnclLiab          []any     `json:"tradable_fnncl_liab,omitempty"`
				PayrollPayable             []float64 `json:"payroll_payable,omitempty"`
				TaxPayable                 []float64 `json:"tax_payable,omitempty"`
				EstimatedLiab              []any     `json:"estimated_liab,omitempty"`
				DtLiab                     []float64 `json:"dt_liab,omitempty"`
				BondPayable                []float64 `json:"bond_payable,omitempty"`
				Shares                     []float64 `json:"shares,omitempty"`
				CapitalReserve             []float64 `json:"capital_reserve,omitempty"`
				EarnedSurplus              []float64 `json:"earned_surplus,omitempty"`
				UndstrbtdProfit            []float64 `json:"undstrbtd_profit,omitempty"`
				MinorityEquity             []any     `json:"minority_equity,omitempty"`
				TotalHoldersEquity         []float64 `json:"total_holders_equity,omitempty"`
				TotalLiabAndHoldersEquity  []float64 `json:"total_liab_and_holders_equity,omitempty"`
				LtEquityInvest             []float64 `json:"lt_equity_invest,omitempty"`
				DerivativeFnnclLiab        []any     `json:"derivative_fnncl_liab,omitempty"`
				GeneralRiskProvision       []any     `json:"general_risk_provision,omitempty"`
				FrgnCurrencyConvertDiff    []any     `json:"frgn_currency_convert_diff,omitempty"`
				Goodwill                   []any     `json:"goodwill,omitempty"`
				InvestProperty             []any     `json:"invest_property,omitempty"`
				InterestPayable            []any     `json:"interest_payable,omitempty"`
				TreasuryStock              []float64 `json:"treasury_stock,omitempty"`
				OthrCompreIncome           []float64 `json:"othr_compre_income,omitempty"`
				OthrEquityInstruments      []float64 `json:"othr_equity_instruments,omitempty"`
				CurrencyFunds              []float64 `json:"currency_funds,omitempty"`
				BillsReceivable            []float64 `json:"bills_receivable,omitempty"`
				AccountReceivable          []float64 `json:"account_receivable,omitempty"`
				PrePayment                 []float64 `json:"pre_payment,omitempty"`
				DividendReceivable         []any     `json:"dividend_receivable,omitempty"`
				OthrReceivables            []any     `json:"othr_receivables,omitempty"`
				Inventory                  []float64 `json:"inventory,omitempty"`
				NcaDueWithinOneYear        []float64 `json:"nca_due_within_one_year,omitempty"`
				OthrCurrentAssets          []float64 `json:"othr_current_assets,omitempty"`
				CurrentAssetsSi            []any     `json:"current_assets_si,omitempty"`
				TotalCurrentAssets         []float64 `json:"total_current_assets,omitempty"`
				LtReceivable               []float64 `json:"lt_receivable,omitempty"`
				DevExpenditure             []any     `json:"dev_expenditure,omitempty"`
				LtDeferredExpense          []float64 `json:"lt_deferred_expense,omitempty"`
				OthrNoncurrentAssets       []float64 `json:"othr_noncurrent_assets,omitempty"`
				NoncurrentAssetsSi         []any     `json:"noncurrent_assets_si,omitempty"`
				TotalNoncurrentAssets      []float64 `json:"total_noncurrent_assets,omitempty"`
				StLoan                     []any     `json:"st_loan,omitempty"`
				BillPayable                []float64 `json:"bill_payable,omitempty"`
				AccountsPayable            []float64 `json:"accounts_payable,omitempty"`
				PreReceivable              []any     `json:"pre_receivable,omitempty"`
				DividendPayable            []any     `json:"dividend_payable,omitempty"`
				OthrPayables               []any     `json:"othr_payables,omitempty"`
				NoncurrentLiabDueIn1Y      []float64 `json:"noncurrent_liab_due_in1y,omitempty"`
				CurrentLiabSi              []any     `json:"current_liab_si,omitempty"`
				TotalCurrentLiab           []float64 `json:"total_current_liab,omitempty"`
				LtLoan                     []any     `json:"lt_loan,omitempty"`
				LtPayable                  []any     `json:"lt_payable,omitempty"`
				SpecialPayable             []any     `json:"special_payable,omitempty"`
				OthrNonCurrentLiab         []float64 `json:"othr_non_current_liab,omitempty"`
				NoncurrentLiabSi           []any     `json:"noncurrent_liab_si,omitempty"`
				TotalNoncurrentLiab        []float64 `json:"total_noncurrent_liab,omitempty"`
				SalableFinancialAssets     []any     `json:"salable_financial_assets,omitempty"`
				OthrCurrentLiab            []float64 `json:"othr_current_liab,omitempty"`
				ArAndBr                    []float64 `json:"ar_and_br,omitempty"`
				ContractualAssets          []any     `json:"contractual_assets,omitempty"`
				BpAndAp                    []float64 `json:"bp_and_ap,omitempty"`
				ContractLiabilities        []float64 `json:"contract_liabilities,omitempty"`
				ToSaleAsset                []any     `json:"to_sale_asset,omitempty"`
				OtherEqInsInvest           []float64 `json:"other_eq_ins_invest,omitempty"`
				OtherIlliquidFnnclAssets   []float64 `json:"other_illiquid_fnncl_assets,omitempty"`
				FixedAssetSum              []float64 `json:"fixed_asset_sum,omitempty"`
				FixedAssetsDisposal        []any     `json:"fixed_assets_disposal,omitempty"`
				ConstructionInProcessSum   []float64 `json:"construction_in_process_sum,omitempty"`
				ProjectGoodsAndMaterial    []any     `json:"project_goods_and_material,omitempty"`
				ProductiveBiologicalAssets []any     `json:"productive_biological_assets,omitempty"`
				OilAndGasAsset             []any     `json:"oil_and_gas_asset,omitempty"`
				ToSaleDebt                 []any     `json:"to_sale_debt,omitempty"`
				LtPayableSum               []any     `json:"lt_payable_sum,omitempty"`
				NoncurrentLiabDi           []float64 `json:"noncurrent_liab_di,omitempty"`
				PerpetualBond              []any     `json:"perpetual_bond,omitempty"`
				SpecialReserve             []any     `json:"special_reserve,omitempty"`
			} `json:"list,omitempty"`
		} `json:"data,omitempty"`
		ErrorCode        int    `json:"error_code,omitempty"`
		ErrorDescription string `json:"error_description,omitempty"`
	}
*/
type Balance struct {
	Data struct {
		QuoteName      string                   `json:"quote_name,omitempty"`
		CurrencyName   string                   `json:"currency_name,omitempty"`
		OrgType        int                      `json:"org_type,omitempty"`
		LastReportName string                   `json:"last_report_name,omitempty"`
		Statuses       any                      `json:"statuses,omitempty"`
		Currency       string                   `json:"currency,omitempty"`
		List           []map[string]interface{} `json:"list,omitempty"`
	} `json:"data,omitempty"`
	ErrorCode        int    `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

/*
type CashFlow struct {
	Data struct {
		QuoteName      string `json:"quote_name,omitempty"`
		CurrencyName   string `json:"currency_name,omitempty"`
		OrgType        int    `json:"org_type,omitempty"`
		LastReportName string `json:"last_report_name,omitempty"`
		Statuses       any    `json:"statuses,omitempty"`
		Currency       string `json:"currency,omitempty"`
		List           []struct {
			ReportDate                 int64     `json:"report_date,omitempty"`
			ReportName                 string    `json:"report_name,omitempty"`
			Ctime                      any       `json:"ctime,omitempty"`
			NcfFromOa                  []float64 `json:"ncf_from_oa,omitempty"`
			NcfFromIa                  []float64 `json:"ncf_from_ia,omitempty"`
			NcfFromFa                  []float64 `json:"ncf_from_fa,omitempty"`
			CashReceivedOfOthrOa       []float64 `json:"cash_received_of_othr_oa,omitempty"`
			SubTotalOfCiFromOa         []float64 `json:"sub_total_of_ci_from_oa,omitempty"`
			CashPaidToEmployeeEtc      []float64 `json:"cash_paid_to_employee_etc,omitempty"`
			PaymentsOfAllTaxes         []float64 `json:"payments_of_all_taxes,omitempty"`
			OthrcashPaidRelatingToOa   []float64 `json:"othrcash_paid_relating_to_oa,omitempty"`
			SubTotalOfCosFromOa        []float64 `json:"sub_total_of_cos_from_oa,omitempty"`
			CashReceivedOfDspslInvest  []float64 `json:"cash_received_of_dspsl_invest,omitempty"`
			InvestIncomeCashReceived   []float64 `json:"invest_income_cash_received,omitempty"`
			NetCashOfDisposalAssets    []float64 `json:"net_cash_of_disposal_assets,omitempty"`
			NetCashOfDisposalBranch    []any     `json:"net_cash_of_disposal_branch,omitempty"`
			CashReceivedOfOthrIa       []any     `json:"cash_received_of_othr_ia,omitempty"`
			SubTotalOfCiFromIa         []float64 `json:"sub_total_of_ci_from_ia,omitempty"`
			InvestPaidCash             []float64 `json:"invest_paid_cash,omitempty"`
			CashPaidForAssets          []float64 `json:"cash_paid_for_assets,omitempty"`
			OthrcashPaidRelatingToIa   []any     `json:"othrcash_paid_relating_to_ia,omitempty"`
			SubTotalOfCosFromIa        []float64 `json:"sub_total_of_cos_from_ia,omitempty"`
			CashReceivedOfAbsorbInvest []float64 `json:"cash_received_of_absorb_invest,omitempty"`
			CashReceivedFromInvestor   []any     `json:"cash_received_from_investor,omitempty"`
			CashReceivedFromBondIssue  []any     `json:"cash_received_from_bond_issue,omitempty"`
			CashReceivedOfBorrowing    []any     `json:"cash_received_of_borrowing,omitempty"`
			CashReceivedOfOthrFa       []any     `json:"cash_received_of_othr_fa,omitempty"`
			SubTotalOfCiFromFa         []float64 `json:"sub_total_of_ci_from_fa,omitempty"`
			CashPayForDebt             []float64 `json:"cash_pay_for_debt,omitempty"`
			CashPaidOfDistribution     []float64 `json:"cash_paid_of_distribution,omitempty"`
			BranchPaidToMinorityHolder []any     `json:"branch_paid_to_minority_holder,omitempty"`
			OthrcashPaidRelatingToFa   []float64 `json:"othrcash_paid_relating_to_fa,omitempty"`
			SubTotalOfCosFromFa        []float64 `json:"sub_total_of_cos_from_fa,omitempty"`
			EffectOfExchangeChgOnCce   []float64 `json:"effect_of_exchange_chg_on_cce,omitempty"`
			NetIncreaseInCce           []float64 `json:"net_increase_in_cce,omitempty"`
			InitialBalanceOfCce        []float64 `json:"initial_balance_of_cce,omitempty"`
			FinalBalanceOfCce          []float64 `json:"final_balance_of_cce,omitempty"`
			CashReceivedOfSalesService []float64 `json:"cash_received_of_sales_service,omitempty"`
			RefundOfTaxAndLevies       []float64 `json:"refund_of_tax_and_levies,omitempty"`
			GoodsBuyAndServiceCashPay  []float64 `json:"goods_buy_and_service_cash_pay,omitempty"`
			NetCashAmtFromBranch       []any     `json:"net_cash_amt_from_branch,omitempty"`
		} `json:"list,omitempty"`
	} `json:"data,omitempty"`
	ErrorCode        int    `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}
*/

type CashFlow struct {
	Data struct {
		QuoteName      string                   `json:"quote_name,omitempty"`
		CurrencyName   string                   `json:"currency_name,omitempty"`
		OrgType        int                      `json:"org_type,omitempty"`
		LastReportName string                   `json:"last_report_name,omitempty"`
		Statuses       any                      `json:"statuses,omitempty"`
		Currency       string                   `json:"currency,omitempty"`
		List           []map[string]interface{} `json:"list,omitempty"`
	} `json:"data,omitempty"`
	ErrorCode        int    `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

type Quote struct {
	Data struct {
		Market struct {
			StatusID              int    `json:"status_id,omitempty"`
			Region                string `json:"region,omitempty"`
			Status                string `json:"status,omitempty"`
			TimeZone              string `json:"time_zone,omitempty"`
			TimeZoneDesc          any    `json:"time_zone_desc,omitempty"`
			DelayTag              int    `json:"delay_tag,omitempty"`
			DowngradeNightSession bool   `json:"downgrade_night_session,omitempty"`
		} `json:"market,omitempty"`
		Quote struct {
			CurrentExt               float64 `json:"current_ext,omitempty"`
			Symbol                   string  `json:"symbol,omitempty"`
			VolumeExt                float64 `json:"volume_ext,omitempty"`
			High52W                  float64 `json:"high52w,omitempty"`
			Delayed                  int     `json:"delayed,omitempty"`
			Type                     int     `json:"type,omitempty"`
			TickSize                 float64 `json:"tick_size,omitempty"`
			FloatShares              float64 `json:"float_shares,omitempty"`
			LimitDown                float64 `json:"limit_down,omitempty"`
			NoProfit                 string  `json:"no_profit,omitempty"`
			High                     float64 `json:"high,omitempty"`
			FloatMarketCapital       float64 `json:"float_market_capital,omitempty"`
			TimestampExt             int64   `json:"timestamp_ext,omitempty"`
			LotSize                  float64 `json:"lot_size,omitempty"`
			LockSet                  any     `json:"lock_set,omitempty"`
			WeightedVotingRights     string  `json:"weighted_voting_rights,omitempty"`
			Chg                      float64 `json:"chg,omitempty"`
			Eps                      float64 `json:"eps,omitempty"`
			LastClose                float64 `json:"last_close,omitempty"`
			ProfitFour               float64 `json:"profit_four,omitempty"`
			Volume                   float64 `json:"volume,omitempty"`
			VolumeRatio              float64 `json:"volume_ratio,omitempty"`
			ProfitForecast           float64 `json:"profit_forecast,omitempty"`
			TurnoverRate             float64 `json:"turnover_rate,omitempty"`
			Low52W                   float64 `json:"low52w,omitempty"`
			Name                     string  `json:"name,omitempty"`
			Exchange                 string  `json:"exchange,omitempty"`
			PeForecast               float64 `json:"pe_forecast,omitempty"`
			TotalShares              float64 `json:"total_shares,omitempty"`
			Status                   float64 `json:"status,omitempty"`
			IsVieDesc                string  `json:"is_vie_desc,omitempty"`
			SecurityStatus           any     `json:"security_status,omitempty"`
			Code                     string  `json:"code,omitempty"`
			GoodwillInNetAssets      any     `json:"goodwill_in_net_assets,omitempty"`
			AvgPrice                 float64 `json:"avg_price,omitempty"`
			Percent                  float64 `json:"percent,omitempty"`
			WeightedVotingRightsDesc string  `json:"weighted_voting_rights_desc,omitempty"`
			Amplitude                float64 `json:"amplitude,omitempty"`
			Current                  float64 `json:"current,omitempty"`
			IsVie                    string  `json:"is_vie,omitempty"`
			CurrentYearPercent       float64 `json:"current_year_percent,omitempty"`
			IssueDate                int64   `json:"issue_date,omitempty"`
			SubType                  string  `json:"sub_type,omitempty"`
			Low                      float64 `json:"low,omitempty"`
			IsRegistrationDesc       string  `json:"is_registration_desc,omitempty"`
			NoProfitDesc             string  `json:"no_profit_desc,omitempty"`
			MarketCapital            float64 `json:"market_capital,omitempty"`
			Dividend                 float64 `json:"dividend,omitempty"`
			DividendYield            float64 `json:"dividend_yield,omitempty"`
			Currency                 string  `json:"currency,omitempty"`
			Navps                    float64 `json:"navps,omitempty"`
			Profit                   float64 `json:"profit,omitempty"`
			Timestamp                int64   `json:"timestamp,omitempty"`
			PeLyr                    float64 `json:"pe_lyr,omitempty"`
			Amount                   float64 `json:"amount,omitempty"`
			PledgeRatio              any     `json:"pledge_ratio,omitempty"`
			TradedAmountExt          float64 `json:"traded_amount_ext,omitempty"`
			IsRegistration           string  `json:"is_registration,omitempty"`
			Pb                       float64 `json:"pb,omitempty"`
			LimitUp                  float64 `json:"limit_up,omitempty"`
			PeTtm                    float64 `json:"pe_ttm,omitempty"`
			Time                     int64   `json:"time,omitempty"`
			Open                     float64 `json:"open,omitempty"`
		} `json:"quote,omitempty"`
		Others struct {
			PankouRatio float64 `json:"pankou_ratio,omitempty"`
			CybSwitch   bool    `json:"cyb_switch,omitempty"`
		} `json:"others,omitempty"`
		Tags []struct {
			Description string `json:"description,omitempty"`
			Value       int    `json:"value,omitempty"`
		} `json:"tags,omitempty"`
	} `json:"data,omitempty"`
	ErrorCode        int    `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package studiosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/clear-street/studio-sdk-go/internal/apijson"
	"github.com/clear-street/studio-sdk-go/internal/requestconfig"
	"github.com/clear-street/studio-sdk-go/option"
)

// EntityService contains methods and other services that help with interacting
// with the studio-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityService] method instead.
type EntityService struct {
	Options               []option.RequestOption
	PnlSummaries          *EntityPnlSummaryService
	RegtMargins           *EntityRegtMarginService
	PortfolioMargins      *EntityPortfolioMarginService
	RegtMarginSimulations *EntityRegtMarginSimulationService
}

// NewEntityService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEntityService(opts ...option.RequestOption) (r *EntityService) {
	r = &EntityService{}
	r.Options = opts
	r.PnlSummaries = NewEntityPnlSummaryService(opts...)
	r.RegtMargins = NewEntityRegtMarginService(opts...)
	r.PortfolioMargins = NewEntityPortfolioMarginService(opts...)
	r.RegtMarginSimulations = NewEntityRegtMarginSimulationService(opts...)
	return
}

// Get an entity by its ID.
func (r *EntityService) Get(ctx context.Context, entityID string, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all available entities.
func (r *EntityService) List(ctx context.Context, opts ...option.RequestOption) (res *EntityListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "entities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Entity struct {
	ClientCode string `json:"client_code,required"`
	// Entity ID for the legal entity.
	EntityID  string     `json:"entity_id,required"`
	LegalName string     `json:"legal_name"`
	JSON      entityJSON `json:"-"`
}

// entityJSON contains the JSON metadata for the struct [Entity]
type entityJSON struct {
	ClientCode  apijson.Field
	EntityID    apijson.Field
	LegalName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Entity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityJSON) RawJSON() string {
	return r.raw
}

type PnlSummary struct {
	// Profit and loss from intraday trading activities.
	DayPnl float64 `json:"day_pnl,required"`
	// Entity ID for the legal entity.
	EntityID string `json:"entity_id,required"`
	// Net value of instruments held in the portfolio.
	Equity float64 `json:"equity,required"`
	// Absolute market value of long and short market values.
	GrossMarketValue float64 `json:"gross_market_value,required"`
	// Market value of securities positioned long.
	LongMarketValue float64 `json:"long_market_value,required"`
	// Market value net of long and short market values.
	NetMarketValue float64 `json:"net_market_value,required"`
	// `total_pnl + total_fees`
	NetPnl float64 `json:"net_pnl,required"`
	// Profit and loss from previous trading date.
	OvernightPnl float64 `json:"overnight_pnl,required"`
	// Profit and loss realized from position closing trading activity
	RealizedPnl float64 `json:"realized_pnl,required"`
	// Market value of securities positioned short.
	ShortMarketValue float64 `json:"short_market_value,required"`
	// Net value of instruments held in the portfolio at the start of a trading day.
	SodEquity float64 `json:"sod_equity,required"`
	// Absolute market value at the start of a trading day.
	SodGrossMarketValue float64 `json:"sod_gross_market_value,required"`
	// Market value of securities positioned long at the start of a trading day.
	SodLongMarketValue float64 `json:"sod_long_market_value,required"`
	// Market value of securities positioned short at the start of a trading day.
	SodShortMarketValue float64 `json:"sod_short_market_value,required"`
	// Milliseconds since epoch.
	Timestamp int64 `json:"timestamp,required"`
	// Total fees incurred from trading activities.
	TotalFees float64 `json:"total_fees,required"`
	// `realized_pnl + unrealized_pnl`
	TotalPnl float64 `json:"total_pnl,required"`
	// Profit and loss from market changes.
	UnrealizedPnl float64        `json:"unrealized_pnl,required"`
	JSON          pnlSummaryJSON `json:"-"`
}

// pnlSummaryJSON contains the JSON metadata for the struct [PnlSummary]
type pnlSummaryJSON struct {
	DayPnl              apijson.Field
	EntityID            apijson.Field
	Equity              apijson.Field
	GrossMarketValue    apijson.Field
	LongMarketValue     apijson.Field
	NetMarketValue      apijson.Field
	NetPnl              apijson.Field
	OvernightPnl        apijson.Field
	RealizedPnl         apijson.Field
	ShortMarketValue    apijson.Field
	SodEquity           apijson.Field
	SodGrossMarketValue apijson.Field
	SodLongMarketValue  apijson.Field
	SodShortMarketValue apijson.Field
	Timestamp           apijson.Field
	TotalFees           apijson.Field
	TotalPnl            apijson.Field
	UnrealizedPnl       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PnlSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pnlSummaryJSON) RawJSON() string {
	return r.raw
}

type PortfolioMargin struct {
	// Sum of add-on margin requirements. Formula:
	// `liquidity_add_on + concentration_add_on + discretionary_requirement`
	AddOnRequirement float64 `json:"add_on_requirement"`
	// The percentage add-on margin requirements in terms of total house requirement.
	// Formula: `add_on_requirement / house_requirement`
	AddOnRequirementPercent float64 `json:"add_on_requirement_percent"`
	// A component margin requirement that captures risk based on gross exposure to
	// total equity.
	ConcentrationAddOn float64 `json:"concentration_add_on"`
	// The percentage concentration add-on margin requirements in terms of total house
	// requirement. Formula: `concentration_add_on / house_requirement`
	ConcentrationAddOnPercent float64 `json:"concentration_add_on_percent"`
	// A component margin requirement that captures miscellaneous risk factors.
	DiscretionaryRequirement float64 `json:"discretionary_requirement"`
	// The percentage discretionary margin requirements in terms of total house
	// requirement Formula: `discretionary_requirement / house_requirement`
	DiscretionaryRequirementPercent float64 `json:"discretionary_requirement_percent"`
	// The maring amount by taking the difference between total equity and the
	// effective requirement. A negative number reflects an effective margin deficit.
	EffeciveExcess float64 `json:"effecive_excess"`
	// The enforced margin requirement in effect.
	EffectiveRequirement float64 `json:"effective_requirement"`
	// Portfolio margin groups
	Groups []PortfolioMarginGroup `json:"groups"`
	// The margin amount by taking the difference between total equity and the house
	// requirement. A negative number reflects a house margin deficit.
	HouseExcess float64 `json:"house_excess"`
	// Margin requirements based on Clear Street's house margin methodology.
	HouseRequirement float64 `json:"house_requirement"`
	// A component margin requirement that captures risk based on liquidity, Market
	// Cap, and Average Daily Volume factors.
	LiquidityAddOn float64 `json:"liquidity_add_on"`
	// The percentage liquidity add-on margin requirements in terms of total house
	// requirement. Formula: `liquidity_add_on / house_requirement`
	LiquidityAddOnPercent float64 `json:"liquidity_add_on_percent"`
	// Sum of market values across all positions.
	NetMarketValue float64 `json:"net_market_value"`
	// A component margin requirement that captures risk for security instruments that
	// are not margin eligible.
	NonMarginableRequirement float64 `json:"non_marginable_requirement"`
	// The percentage non-marginable requirement in terms of total house requirement
	// Formula: `non_marginable_requirement / house_requirement`
	NonMarginableRequirementPercent float64 `json:"non_marginable_requirement_percent"`
	// A component margin requirement that captures base-case risk under house margin
	// methodology.
	RiskBasedRequirement float64 `json:"risk_based_requirement"`
	// The percentage risk_base margin requirement in terms of total house requirement
	// Formula: `risk_based_requirement / house_requirement`
	RiskBasedRequirementPercent float64 `json:"risk_based_requirement_percent"`
	// Timestamp of when this margin was calculated.
	Timestamp int64 `json:"timestamp"`
	// A component margin requirement that captures risk based on vega.
	VegaRequirement float64 `json:"vega_requirement"`
	// Unique identifier for this margin calculation.
	Version string              `json:"version"`
	JSON    portfolioMarginJSON `json:"-"`
}

// portfolioMarginJSON contains the JSON metadata for the struct [PortfolioMargin]
type portfolioMarginJSON struct {
	AddOnRequirement                apijson.Field
	AddOnRequirementPercent         apijson.Field
	ConcentrationAddOn              apijson.Field
	ConcentrationAddOnPercent       apijson.Field
	DiscretionaryRequirement        apijson.Field
	DiscretionaryRequirementPercent apijson.Field
	EffeciveExcess                  apijson.Field
	EffectiveRequirement            apijson.Field
	Groups                          apijson.Field
	HouseExcess                     apijson.Field
	HouseRequirement                apijson.Field
	LiquidityAddOn                  apijson.Field
	LiquidityAddOnPercent           apijson.Field
	NetMarketValue                  apijson.Field
	NonMarginableRequirement        apijson.Field
	NonMarginableRequirementPercent apijson.Field
	RiskBasedRequirement            apijson.Field
	RiskBasedRequirementPercent     apijson.Field
	Timestamp                       apijson.Field
	VegaRequirement                 apijson.Field
	Version                         apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *PortfolioMargin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r portfolioMarginJSON) RawJSON() string {
	return r.raw
}

type PortfolioMarginGroup struct {
	// The enforced margin requirement in effect for the group.
	EffectiveRequirement float64 `json:"effective_requirement,required"`
	// The percentage effective margin requirement in terms of the group market value.
	// Formula: `(effective_requirement / net_market_value)`
	MarginPercent float64 `json:"margin_percent,required"`
	// The percentage effective margin requirement in terms of the total effective
	// requirement. Formula: `(effective_requirement / sum(effective_requirement))`
	MarginPercentContribution float64 `json:"margin_percent_contribution,required"`
	// The aggregated market value of all instruments for the group.
	MarketValue float64 `json:"market_value,required"`
	// The percentage market value of the group in terms of the total net_market_value
	// of all positions. Formula: `(market_value / net_market_value)`
	MarketValuePercent float64 `json:"market_value_percent,required"`
	// A list of securities that comprise this group.
	Members []PortfolioMarginGroupsMember `json:"members,required"`
	// Unique name of the group, typically the symbol of the underlier.
	Name string `json:"name,required"`
	// A component margin requirement that captures risk for the group based on gross
	// exposure to total equity
	ConcentrationRequirement float64 `json:"concentration_requirement"`
	// A component margin requirement that captures miscellaneous risk factors for the
	// group.
	DiscretionaryRequirement float64 `json:"discretionary_requirement"`
	// A component margin requirement that captures risk for the group based on
	// liquidity, Market Cap, and Average Daily Volume factors.
	LiquidityRequirement float64 `json:"liquidity_requirement"`
	// A component margin requirement that captures risk for the group that are not
	// margin eligible.
	NonMarginableRequirement float64 `json:"non_marginable_requirement"`
	// Margin requirements based on OCC TIMS regulatory margin methodology
	RegulatoryRequirement float64 `json:"regulatory_requirement"`
	// A component margin requirement that captures base-case risk for the group under
	// house margin methodology
	RiskBasedRequirement float64 `json:"risk_based_requirement"`
	// Maps shock scenarios to their resulting pnl.
	Shocks map[string]float64 `json:"shocks"`
	// Margin requirements based on value-at-risk over any 5-day period in a 2 year
	// historic lookback
	VarRequirement float64                  `json:"var_requirement"`
	JSON           portfolioMarginGroupJSON `json:"-"`
}

// portfolioMarginGroupJSON contains the JSON metadata for the struct
// [PortfolioMarginGroup]
type portfolioMarginGroupJSON struct {
	EffectiveRequirement      apijson.Field
	MarginPercent             apijson.Field
	MarginPercentContribution apijson.Field
	MarketValue               apijson.Field
	MarketValuePercent        apijson.Field
	Members                   apijson.Field
	Name                      apijson.Field
	ConcentrationRequirement  apijson.Field
	DiscretionaryRequirement  apijson.Field
	LiquidityRequirement      apijson.Field
	NonMarginableRequirement  apijson.Field
	RegulatoryRequirement     apijson.Field
	RiskBasedRequirement      apijson.Field
	Shocks                    apijson.Field
	VarRequirement            apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *PortfolioMarginGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r portfolioMarginGroupJSON) RawJSON() string {
	return r.raw
}

type PortfolioMarginGroupsMember struct {
	// The asset class of the symbol.
	AssetClass PortfolioMarginGroupsMembersAssetClass `json:"asset_class"`
	// Market value of the instrument.
	MarketValue float64 `json:"market_value"`
	// The percentage market value of the instrument in terms of the total
	// `net_market_value` of all positions held. Formula:
	// `market_value / net_market_value`
	MarketValuePercent float64 `json:"market_value_percent"`
	// The quantity held for this instrument.
	Quantity string `json:"quantity"`
	// Maps shock scenarios to their resulting pnl.
	Shocks map[string]float64 `json:"shocks"`
	// The symbol for the instrument.
	Symbol string                          `json:"symbol"`
	JSON   portfolioMarginGroupsMemberJSON `json:"-"`
}

// portfolioMarginGroupsMemberJSON contains the JSON metadata for the struct
// [PortfolioMarginGroupsMember]
type portfolioMarginGroupsMemberJSON struct {
	AssetClass         apijson.Field
	MarketValue        apijson.Field
	MarketValuePercent apijson.Field
	Quantity           apijson.Field
	Shocks             apijson.Field
	Symbol             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PortfolioMarginGroupsMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r portfolioMarginGroupsMemberJSON) RawJSON() string {
	return r.raw
}

// The asset class of the symbol.
type PortfolioMarginGroupsMembersAssetClass string

const (
	PortfolioMarginGroupsMembersAssetClassOther  PortfolioMarginGroupsMembersAssetClass = "other"
	PortfolioMarginGroupsMembersAssetClassEquity PortfolioMarginGroupsMembersAssetClass = "equity"
	PortfolioMarginGroupsMembersAssetClassOption PortfolioMarginGroupsMembersAssetClass = "option"
	PortfolioMarginGroupsMembersAssetClassDebt   PortfolioMarginGroupsMembersAssetClass = "debt"
)

func (r PortfolioMarginGroupsMembersAssetClass) IsKnown() bool {
	switch r {
	case PortfolioMarginGroupsMembersAssetClassOther, PortfolioMarginGroupsMembersAssetClassEquity, PortfolioMarginGroupsMembersAssetClassOption, PortfolioMarginGroupsMembersAssetClassDebt:
		return true
	}
	return false
}

type RegtMargin struct {
	// The remaining amount of start_of_day_buying_power that captures any day-trading
	// activity.
	DayTradeBuyingPower float64 `json:"day_trade_buying_power,required"`
	// The enforced margin requirement in effect.
	EffectiveRequirement float64 `json:"effective_requirement,required"`
	// Margin requirements based on regulatory rules.
	ExchangeRequirement float64 `json:"exchange_requirement,required"`
	// Reg-T margin groups
	Groups []RegtMarginGroup `json:"groups,required"`
	// The margin amount by taking the difference between total equity and the house
	// requirement. A negative number reflects a house margin deficit.
	HouseExcess float64 `json:"house_excess,required"`
	// Margin requirements based on Clear Street's house margin methodology.
	HouseRequirement float64 `json:"house_requirement,required"`
	// Market value net of long and short market values.
	NetMarketValue float64 `json:"net_market_value,required"`
	// The limit, or "up-to" amount, of securities value that can be purchased and held
	// overnight.
	OvernightBuyingPower float64 `json:"overnight_buying_power,required"`
	// Special Memorandum Account (SMA). The regulatory line of credit amount for
	// margin trading based on market value, trading activity, and available cash.
	Sma float64 `json:"sma,required"`
	// The limit, or "up-to" amount, of securities value that can be day-traded for a
	// given trading day.
	SodBuyingPower float64 `json:"sod_buying_power,required"`
	// Timestamp of when this margin was calculated.
	Timestamp int64 `json:"timestamp,required"`
	// Unique identifier for this margin calculation.
	Version string `json:"version,required"`
	// The maring amount by taking the difference between total equity and the
	// effective requirement. A negative number reflects an effective margin deficit.
	EffectiveExcess float64 `json:"effective_excess"`
	// The margin amount by taking the difference between total equity and the exchange
	// requirement. A negative number reflects an regulatory margin deficit.
	ExchangeExcess float64        `json:"exchange_excess"`
	JSON           regtMarginJSON `json:"-"`
}

// regtMarginJSON contains the JSON metadata for the struct [RegtMargin]
type regtMarginJSON struct {
	DayTradeBuyingPower  apijson.Field
	EffectiveRequirement apijson.Field
	ExchangeRequirement  apijson.Field
	Groups               apijson.Field
	HouseExcess          apijson.Field
	HouseRequirement     apijson.Field
	NetMarketValue       apijson.Field
	OvernightBuyingPower apijson.Field
	Sma                  apijson.Field
	SodBuyingPower       apijson.Field
	Timestamp            apijson.Field
	Version              apijson.Field
	EffectiveExcess      apijson.Field
	ExchangeExcess       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RegtMargin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regtMarginJSON) RawJSON() string {
	return r.raw
}

type RegtMarginGroup struct {
	// The enforced margin requirement in effect for the symbol group.
	EffectiveRequirement float64 `json:"effective_requirement,required"`
	// Margin requirements based on regulatory rules for the symbol group.
	ExchangeRequirement float64 `json:"exchange_requirement,required"`
	// Margin requirements based on Clear Street's house margin methodology for the
	// symbol group.
	HouseRequirement float64 `json:"house_requirement,required"`
	// The percentage effective margin requirement in terms of the symbol group market
	// value. Formula: `(effective_requirement / net_market_value)`
	MarginPercent float64 `json:"margin_percent,required"`
	// The percentage effective margin requirement in terms of the total effective
	// requirement. Formula: `(effective_requirement / sum(effective_requirement))`
	MarginPercentContribution float64 `json:"margin_percent_contribution,required"`
	// The aggregated market value of all instruments for the symbol group.
	MarketValue float64 `json:"market_value,required"`
	// The percentage market value of the symbol group in terms of the total
	// net_market_value of all positions. Formula: `(market_value / net_market_value)`
	MarketValuePercent float64 `json:"market_value_percent,required"`
	// A list of securities that comprise this group.
	Members []RegtMarginGroupsMember `json:"members,required"`
	// Unique name of the group, typically the symbol of the underlier.
	Name string              `json:"name,required"`
	JSON regtMarginGroupJSON `json:"-"`
}

// regtMarginGroupJSON contains the JSON metadata for the struct [RegtMarginGroup]
type regtMarginGroupJSON struct {
	EffectiveRequirement      apijson.Field
	ExchangeRequirement       apijson.Field
	HouseRequirement          apijson.Field
	MarginPercent             apijson.Field
	MarginPercentContribution apijson.Field
	MarketValue               apijson.Field
	MarketValuePercent        apijson.Field
	Members                   apijson.Field
	Name                      apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *RegtMarginGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regtMarginGroupJSON) RawJSON() string {
	return r.raw
}

type RegtMarginGroupsMember struct {
	// The asset class of the symbol.
	AssetClass RegtMarginGroupsMembersAssetClass `json:"asset_class,required"`
	// Market value of the instrument.
	MarketValue float64 `json:"market_value,required"`
	// The percentage market value of the instrument in terms of the total
	// `net_market_value` of all positions held. Formula:
	// `market_value / net_market_value`
	MarketValuePercent float64 `json:"market_value_percent,required"`
	// The quantity held for this instrument.
	Quantity string `json:"quantity,required"`
	// The symbol for the instrument.
	Symbol string                     `json:"symbol,required"`
	JSON   regtMarginGroupsMemberJSON `json:"-"`
}

// regtMarginGroupsMemberJSON contains the JSON metadata for the struct
// [RegtMarginGroupsMember]
type regtMarginGroupsMemberJSON struct {
	AssetClass         apijson.Field
	MarketValue        apijson.Field
	MarketValuePercent apijson.Field
	Quantity           apijson.Field
	Symbol             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RegtMarginGroupsMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regtMarginGroupsMemberJSON) RawJSON() string {
	return r.raw
}

// The asset class of the symbol.
type RegtMarginGroupsMembersAssetClass string

const (
	RegtMarginGroupsMembersAssetClassOther  RegtMarginGroupsMembersAssetClass = "other"
	RegtMarginGroupsMembersAssetClassEquity RegtMarginGroupsMembersAssetClass = "equity"
	RegtMarginGroupsMembersAssetClassOption RegtMarginGroupsMembersAssetClass = "option"
	RegtMarginGroupsMembersAssetClassDebt   RegtMarginGroupsMembersAssetClass = "debt"
)

func (r RegtMarginGroupsMembersAssetClass) IsKnown() bool {
	switch r {
	case RegtMarginGroupsMembersAssetClassOther, RegtMarginGroupsMembersAssetClassEquity, RegtMarginGroupsMembersAssetClassOption, RegtMarginGroupsMembersAssetClassDebt:
		return true
	}
	return false
}

type EntityListResponse struct {
	Data []Entity               `json:"data"`
	JSON entityListResponseJSON `json:"-"`
}

// entityListResponseJSON contains the JSON metadata for the struct
// [EntityListResponse]
type entityListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityListResponseJSON) RawJSON() string {
	return r.raw
}

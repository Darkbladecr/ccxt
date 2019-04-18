
// Instrument Tradeable Contracts, Indices, and History
export default interface Instrument {
    askPrice: number;
	bankruptLimitDownPrice: number;
	bankruptLimitUpPrice: number;
	bidPrice: number;
	buyLeg: string;
	calcInterval: Date;
	capped: boolean;
	closingTimestamp: Date;
	deleverage: boolean;
	expiry: Date;
	fairBasis: number;
	fairBasisRate: number;
	fairMethod: string;
	fairPrice: number;
	foreignNotional24h: number;
	front: Date;
	fundingBaseSymbol: string;
	fundingInterval: Date;
	fundingPremiumSymbol: string;
	fundingRate: number;
	fundingTimestamp: Date;
	hasLiquidity: boolean;
	highPrice: number;
	homeNotional24h: number;
	impactAskPrice: number;
	impactBidPrice: number;
	impactMidPrice: number;
	indicativeFundingRate: number;
	indicativeSettlePrice: number;
	indicativeTaxRate: number;
	initMargin: number;
	insuranceFee: number;
	inverseLeg: string;
	isInverse: boolean;
	isQuanto: boolean;
	lastChangePcnt: number;
	lastPrice: number;
	lastPriceProtected: number;
	lastTickDirection: string;
	limit: number;
	limitDownPrice: number;
	limitUpPrice: number;
	listing: Date;
	lotSize: number;
	lowPrice: number;
	maintMargin: number;
	makerFee: number;
	markMethod: string;
	markPrice: number;
	maxOrderQty: number;
	maxPrice: number;
	midPrice: number;
	multiplier: number;
	openInterest: number;
	openValue: number;
	openingTimestamp: Date;
	optionMultiplier: number;
	optionStrikePcnt: number;
	optionStrikePrice: number;
	optionStrikeRound: number;
	optionUnderlyingPrice: number;
	positionCurrency: string;
	prevClosePrice: number;
	prevPrice24h: number;
	prevTotalTurnover: number;
	prevTotalVolume: number;
	publishInterval: Date;
	publishTime: Date;
	quoteCurrency: string;
	quoteToSettleMultiplier: number;
	rebalanceInterval: Date;
	rebalanceTimestamp: Date;
	reference: string;
	referenceSymbol: string;
	relistInterval: Date;
	riskLimit: number;
	riskStep: number;
	rootSymbol: string;
	sellLeg: string;
	sessionInterval: Date;
	settlCurrency: string;
	settle: Date;
	settledPrice: number;
	settlementFee: number;
	state: string;
	symbol: string;
	takerFee: number;
	taxed: boolean;
	tickSize: number;
	timestamp: Date;
	totalTurnover: number;
	totalVolume: number;
	turnover: number;
	turnover24h: number;
	typ: string;
	underlying: string;
	underlyingSymbol: string;
	underlyingToPositionMultiplier: number;
	underlyingToSettleMultiplier: in64;
	volume: number;
	volume24h: number;
	vwap: number;
}

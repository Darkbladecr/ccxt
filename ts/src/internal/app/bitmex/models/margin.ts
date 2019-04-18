
// Margin margin
export default interface Margin {
	account: number;
	action: string;
	amount: number;
	availableMargin: number;
	comission: number;
	confirmedDebit: number;
	currency: string;
	excessMargin: number;
	excessMarginPcnt: number;
	grossComm: number;
	grossExecCost: number;
	grossLastValue: number;
	grossMarkValue: number;
	grossOpenCost: number;
	grossOpenPremium: number;
	indicativeTax: number;
	initMargin: number;
	maintMargin: number;
	marginBalance: number;
	marginBalancePcnt: number;
	marginLeverage: number;
	marginUsedPcnt: number;
	pendingCredit: number;
	pendingDebit: number;
	prevRealisedPnl: number;
	prevState: string;
	prevUnrealisedPnl: number;
	realisedPnl: number;
	riskLimit: number;
	riskValue: number;
	sessionMargin: number;
	state: string;
	syntheticMargin: number;
	targetExcessMargin: number;
	taxableMargin: number;
	timestamp: Date;
	unrealisedPnl: number;
	unrealisedProfit: number;
	varMargin: number;
	walletBalance: number;
	withdrawableMargin: number;
}
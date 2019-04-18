// Execution Raw Order and Balance Data
export default interface Execution {
  account: number;
  avgPx: number;
  clOrdID: string;
  clOrdLinkID: string;
  Commission: number;
  contingencyType: string;
  cumQty: number;
  currency: string;
  displayQty: number;
  exDestination: string;
  execComm: number;
  execCost: number;
  execID: string;
  execInst: string;
  execType: string;
  foreignNotional: number;
  homeNotional: number;
  lastLiquidityInd: string;
  lastMkt: string;
  lastPx: number;
  lastQty: number;
  leavesQty: number;
  multiLegReportingType: string;
  ordRejReason: string;
  ordStatus: string;
  ordType: string;
  orderQty: number;
  pegOffsetValue: number;
  pegPriceType: string;
  price: number;
  settlCurrency: string;
  side: string;
  simpleCumQty: number;
  simpleLeavesQty: number;
  simpleOrderQty: number;
  stopPx: number;
  symbol: string;
  text: string;
  timeInForce: string;
  timestamp: Date;
  tradePublishIndicator: string;
  transactTime: Date;
  trdMatchID: string;
  triggered: string;
  underlyingLastPx: number;
  wokringIndicator: boolean;
}

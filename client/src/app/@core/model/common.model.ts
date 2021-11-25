export interface BaseResponse<T> {
  data: T;
}

interface NonceData {
  nonce: number;
}

interface LogonData {
  token: string;
  created: string;
  expiry: number;
}

export type NonceRespose = BaseResponse<NonceData>;
export type LogonRespose = BaseResponse<LogonData>;

export interface EthAccounts {
  jsonrpc: string;
  result: string[];
}

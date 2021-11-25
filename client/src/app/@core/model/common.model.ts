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

export interface WalletData {
  id: number;
  FakeId: string;
  created_at: string;
  updated_at: string;
  balance: number;
  status: number;
  userId: number;
}

export type NonceRespose = BaseResponse<NonceData>;
export type LogonRespose = BaseResponse<LogonData>;
export type WalletRespose = BaseResponse<WalletData>;

export interface EthAccounts {
  jsonrpc: string;
  result: string[];
}

export interface IUser {
  userId: string;
  email: string;
  walletAddress: string;
}


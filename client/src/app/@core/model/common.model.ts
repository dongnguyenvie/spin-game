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

export interface MyPackageData {
  id: number;
  FakeId: string;
  created_at: string;
  updated_at: string;
  userId: number;
  package: number;
  status: number;
}

export interface PlayGameData {
  index: number;
  award: number;
}

interface AwardOptions {
  awards: number[];
}

export type NonceRespose = BaseResponse<NonceData>;
export type LogonRespose = BaseResponse<LogonData>;
export type MyWalletRespose = BaseResponse<WalletData>;
export type MyPackageRespose = BaseResponse<MyPackageData>;
export type PlayGameResponse = BaseResponse<PlayGameData>;
export type AwardOptionResponse = BaseResponse<AwardOptions>;

export interface EthAccounts {
  jsonrpc: string;
  result: string[];
}

export interface IUser {
  userId: string;
  email: string;
  walletAddress: string;
}

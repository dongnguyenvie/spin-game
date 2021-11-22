export interface BaseResponse<T> {
  data: T;
}

interface NonceData {
  nonce: number;
}

export type NonceRespose = BaseResponse<NonceData>;

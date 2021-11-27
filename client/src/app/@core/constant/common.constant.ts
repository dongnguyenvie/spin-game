export const ENDPOINT = 'http://localhost:8081/';
export const API_VERSION = 'api/v1/';
export const API = {
  nonce: ENDPOINT + API_VERSION + 'users/nonce/:address',
  login: ENDPOINT + API_VERSION + 'signin',
  register: ENDPOINT + API_VERSION + 'signup',
  myWallet: ENDPOINT + API_VERSION + 'wallets/me',
  buyPackage: ENDPOINT + API_VERSION + 'spin/buy-package',
  myPackage: ENDPOINT + API_VERSION + 'spin/me',
  playGame: ENDPOINT + API_VERSION + 'spin/play',
  awards: ENDPOINT + API_VERSION + 'spin/awards',
  withdraw: ENDPOINT + API_VERSION + 'wallets/withdraw',
};

export const packagePrice = 1000000000000;
export const COIN_DECIMAL = 1e18;
export const PACKAGES = [1, 2, 3, 4, 5];
export const SPIN_PROCESSER = '0x733facC2E45a8f6078E6e9649761c2b13899d5c8';

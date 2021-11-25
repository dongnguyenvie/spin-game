export const ENDPOINT = 'http://localhost:8081/';
export const API_VERSION = 'api/v1/';
export const API = {
  nonce: ENDPOINT + API_VERSION + 'users/nonce/:address',
  login: ENDPOINT + API_VERSION + 'signin',
  register: ENDPOINT + API_VERSION + 'signup',
};

export const REGISTER_SIGN = '25251325';

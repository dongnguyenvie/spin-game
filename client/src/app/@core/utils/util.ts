import Web3 from 'web3';
import { TransactionReceipt } from 'web3-core';
import { COIN_DECIMAL } from '../constant/common.constant';

export function parseJwt(token: string) {
  var base64Url = token.split('.')[1];
  var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
  var jsonPayload = decodeURIComponent(
    atob(base64)
      .split('')
      .map(function (c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
      })
      .join('')
  );

  return JSON.parse(jsonPayload);
}

export function toPlainString(val: number) {
  const num = val / COIN_DECIMAL;
  return ('' + +num).replace(
    /(-?)(\d*)\.?(\d*)e([+-]\d+)/,
    function (a, b, c, d, e) {
      return e < 0
        ? b + '0.' + Array(1 - e - c.length).join('0') + c + d
        : b + c + d + Array(e - d.length + 1).join('0');
    }
  );
}

export function toGasprice(price: number) {
  return price * COIN_DECIMAL;
}

export function tsxChecker(
  web3: Web3,
  tsxHash: string
): Promise<TransactionReceipt> {
  const fetchTsx = () => web3.eth.getTransactionReceipt(tsxHash);

  return new Promise((resolve, reject) => {
    let count = 0;
    const interval = setInterval(async () => {
      if (count > 50) {
        clearInterval(interval);
      }
      fetchTsx()
        .then((result) => {
          console.log('pool', tsxHash);
          if (result) {
            clearInterval(interval);
            resolve(result);
          }
        })
        .catch((err) => {
          clearInterval(interval);
          reject(err);
        });
      count++;
    }, 10000);
  });
}

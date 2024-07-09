import CryptoJS from 'crypto-js';

export function encryptPassword(password: string): string {
  return CryptoJS.MD5(password).toString().toLowerCase();
}

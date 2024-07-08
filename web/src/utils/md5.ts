import { MD5 } from 'crypto-js';

export function encryptPassword(password: string): string {
  return MD5(password).toString().toLowerCase();
}


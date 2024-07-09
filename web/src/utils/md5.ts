import { Md5 } from 'ts-md5';

export function encryptPassword(input: string): string {
  return Md5.hashStr(input);
}



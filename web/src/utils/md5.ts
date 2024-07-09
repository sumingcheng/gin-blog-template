import md5 from 'js-md5';

export function encryptPassword(input: string): string {
  return md5(input);
}



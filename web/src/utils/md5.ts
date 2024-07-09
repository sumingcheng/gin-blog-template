import { createHash } from 'crypto';

export function encryptPassword(input: string): string {
  return createHash('md5').update(input).digest('hex');
}



import forge from 'node-forge';

export function encryptPassword(input: string): string {
  let md = forge.md.md5.create();
  md.update(input);
  return md.digest().toHex();
}



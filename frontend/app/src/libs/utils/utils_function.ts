export function isEmptyOrWhitespace(str: string): boolean {
  return str.replace(/[\s\u3000]/g, "").length === 0;
}

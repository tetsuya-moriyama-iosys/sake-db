export function isEmpty(val: string | number | null | undefined): boolean {
  return val === null || val === undefined || val === '';
}

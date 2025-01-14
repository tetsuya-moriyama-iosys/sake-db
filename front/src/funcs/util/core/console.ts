export function debug(...args: unknown[]): void {
  if (process.env.NODE_ENV !== 'production') {
    console.log(...args);
  }
}

export function errorDebug(...args: unknown[]): void {
  if (process.env.NODE_ENV !== 'production') {
    console.error(...args);
  }
}

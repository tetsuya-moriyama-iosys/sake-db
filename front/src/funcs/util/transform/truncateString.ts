/**
 * 指定された文字数を超過したら切り詰めるヘルパメソッド
 */
interface truncateStringArgs {
  str: string; // 切り詰めたいstring
  maxLength?: number; // 最大文字数
  truncate?: string; // 省略時、文の最後に表示させる省略記号
}

function truncateString({
  str,
  maxLength = 20,
  truncate = '…',
}: truncateStringArgs): string {
  if (str.length <= maxLength) return str;
  return str.slice(0, maxLength) + truncate;
}

export default truncateString;

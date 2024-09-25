import { toZonedTime } from 'date-fns-tz';

// timeZoneをブラウザから取得する関数
const getBrowserTimeZone = () => {
  return Intl.DateTimeFormat().resolvedOptions().timeZone;
};

// timeZoneのデフォルト値を getBrowserTimeZone() に設定
const date = (date: Date, timeZone: string = getBrowserTimeZone()) => {
  return toZonedTime(date, timeZone);
};

export default date;

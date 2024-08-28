type ValueType = string | number | Blob | Date | boolean | null;

function convertFormData<T extends Record<string, ValueType>>(
  obj: T,
  form?: FormData,
  namespace?: string,
): FormData {
  const formData = form || new FormData();

  for (const property of Object.keys(obj)) {
    const key = namespace ? `${namespace}[${property}]` : property;
    const value: ValueType = obj[property];

    if (value instanceof Date) {
      // DateオブジェクトはISO文字列に変換
      formData.append(key, value.toISOString());
    } else if (Array.isArray(value)) {
      // 配列の場合は、各要素を再帰的に処理
      value.forEach((element, index) => {
        const arrayKey = `${key}[${index}]`;
        convertFormData({ [arrayKey]: element }, formData);
      });
    } else if (
      typeof value === 'object' &&
      value !== null &&
      !(value instanceof File) &&
      !(value instanceof Blob)
    ) {
      // オブジェクトの場合は、再帰的に処理
      convertFormData(value, formData, key);
    } else {
      // その他の基本的な型の場合は、そのまま追加
      const processedValue =
        value === null || value === false
          ? '' // null や false の場合は空文字
          : typeof value === 'number' || value === true
            ? String(value) // numberもしくはtrue は文字列に変換
            : value; // string または Blob の場合はそのまま

      formData.append(key, processedValue);
    }
  }

  return formData;
}

export default convertFormData;

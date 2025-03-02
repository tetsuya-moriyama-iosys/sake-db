// カテゴリの経路を取得する関数
import type { Category } from '@/graphQL/Category/categories';

export function getDisplayCategoryIds(
  categories: Category[],
  targetId: number,
): number[] {
  let result: number[] = [];
  // 再帰的に経路を探すヘルパー関数
  const findPath = (category: Category, path: number[]): number[] => {
    // 経路に現在のカテゴリのIDを追加
    path.push(category.id);

    // ターゲットIDが見つかった場合、その、childrenを加えて返す
    if (category.id === targetId) {
      const children: number[] =
        category.children?.map((child) => child.id) ?? [];
      return [...path, ...children];
    }

    // 子カテゴリを探索する
    if (category.children) {
      for (const child of category.children) {
        result = findPath(child, [...path]); // 新しい経路を作成して再帰
        if (result.length > 0) {
          return result;
        }
      }
    }

    // ターゲットが見つからない場合は空配列を返す
    return [];
  };

  // ルートカテゴリから再帰的に探索を開始
  for (const category of categories) {
    //大カテゴリだけは無条件で追加
    result = [category.id, ...result, ...findPath(category, result)];
  }

  // ターゲットIDが見つからなければnullを返す
  return result;
}

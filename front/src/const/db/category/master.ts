/* eslint-disable no-unused-vars*/
//TODO:将来的に不要になる可能性はある(API経由することのコストを考慮すると、頻繁にアクセスされるここはハードコーディングでいい気もする)
export enum MasterCategory {
  SPIRITS = 1, //蒸留酒
  BREW = 2, //醸造酒
  LIQUEUR = 3, //リキュール
  //OTHER = 4, //その他 何がこれに属するか不明なので、必要になるまでコメントアウト
}

export const MasterCategoryLabel = {
  [MasterCategory.SPIRITS]: '蒸留酒',
  [MasterCategory.BREW]: '醸造酒',
  [MasterCategory.LIQUEUR]: 'リキュール',
  //[MasterCategory.OTHER]: 'その他',
};

export enum SPIRITS {
  WHISKEY = 11, //ウィスキー
  GIN = 12, //ジン
  RUM = 13, //ラム
  BRANDY = 14, //ブランデー
  VODKA = 15, //ウォッカ
  TEQUILA = 16, //テキーラ
  OTHER = 17, //その他
}

export enum BREW {
  SAKE = 51, //日本酒
  BEER = 52, //ビール
  WINE = 53, //ワイン
  OTHER = 54, //その他
}

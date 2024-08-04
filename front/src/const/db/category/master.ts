/* eslint-disable no-unused-vars*/

export enum MasterCategory {
  SPIRITS = 1, //蒸留酒
  BREW = 2, //醸造酒
  LIQUEUR = 3, //リキュール
  OTHER = 4, //その他
}

export const MasterCategoryLabel = {
  [MasterCategory.SPIRITS]: '蒸留酒',
  [MasterCategory.BREW]: '醸造酒',
  [MasterCategory.LIQUEUR]: 'リキュール',
  [MasterCategory.OTHER]: 'その他',
};

export enum SPIRITS {
  WHISKEY = 11, //ウィスキー
  BRANDY = 12, //ブランデー
  RUM = 13, //ラム
  VODKA = 14, //ウォッカ
  GIN = 15, //ジン
  TEQUILA = 16, //テキーラ
  OTHER = 17, //その他
}

export enum BREW {
  SAKE = 51, //日本酒
  BEER = 52, //ビール
  WINE = 53, //ワイン
  OTHER = 54, //その他
}

import type { DocumentNode } from '@apollo/client/core';
import { gql } from '@apollo/client/core';

import type { Liquor } from '@/graphQL/Index/random';

export interface RecommendLiquorList {
  readonly getRecommendLiquorList: Recommend[];
}

export interface Recommend {
  readonly rate: number;
  readonly comment: string;
  readonly liquor: Omit<Liquor, 'updatedAt'>;
  readonly user: User;
  readonly updatedAt: Date;
}

export interface User {
  readonly id: string;
  readonly name: string;
  readonly imageBase64: string | undefined;
}

//ブックマークを加味したおすすめリスト
export const RECOMMEND_LIST_FROM_BOOKMARK: DocumentNode = gql`
  query getRecommendLiquorList {
    getRecommendLiquorList {
      rate
      comment
      liquor {
        id
        name
        categoryId
        categoryName
        description
        imageBase64
      }
      user {
        id
        name
        imageBase64
      }
      updatedAt
    }
  }
`;

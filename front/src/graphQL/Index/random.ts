import { gql, type DocumentNode } from '@apollo/client/core';

export interface RecommendLiquorResponse {
  readonly randomRecommendList: Liquor[];
}

export interface Liquor {
  readonly id: string;
  readonly name: string;
  readonly categoryId: number;
  readonly categoryName: string;
  readonly imageBase64: string;
  readonly description: string;
  readonly updatedAt: Date;
}

export const RANDOM_RECOMMEND_LIST: DocumentNode = gql`
  query RandomRecommendList($limit: Int!) {
    randomRecommendList(limit: $limit) {
      id
      name
      categoryId
      categoryName
      description
      imageBase64
      updatedAt
    }
  }
`;

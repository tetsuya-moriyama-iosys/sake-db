import { gql, type DocumentNode } from '@apollo/client/core';

export interface RecommendLiquorResponse {
  randomRecommendList: Liquor[];
}

export interface Liquor {
  id: string;
  name: string;
  categoryId: number;
  categoryName: string;
  imageBase64: string;
  description: string;
  updatedAt: Date;
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

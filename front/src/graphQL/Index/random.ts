import { gql, type DocumentNode } from '@apollo/client/core';

export interface RecommendLiquorResponse {
  id: number;
  name: string;
  image_base64: Blob;
  description: string;
  updated_at: Date;
}

export const RANDOM_RECOMMEND_LIST: DocumentNode = gql`
  query RandomRecommendList($limit: Int!) {
    randomRecommendList(limit: $limit) {
      id
      name
      image_base64
      description
      updated_at
    }
  }
`;

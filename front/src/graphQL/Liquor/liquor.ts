import { gql, type DocumentNode } from '@apollo/client/core';
import { type Liquor as RandomLiquor } from '../Index/random';

export interface LiquorResponse {
  liquor: Liquor;
}

export interface Liquor extends RandomLiquor {
  imageUrl: string;
  updatedAt: Date;
}

export const LIQUOR_DETAIL_GET: DocumentNode = gql`
  query Liquor($id: String!) {
    liquor(id: $id) {
      id
      name
      categoryId
      categoryName
      description
      imageBase64
      imageUrl
      createdAt
      updatedAt
    }
  }
`;

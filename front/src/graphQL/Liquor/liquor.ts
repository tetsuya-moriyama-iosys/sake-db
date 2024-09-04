import { gql, type DocumentNode } from '@apollo/client/core';
import { type Liquor as RandomLiquor } from '../Index/random';

export interface LiquorResponse<T> {
  liquor: T;
}

export interface Liquor extends RandomLiquor {
  imageUrl: string;
  updatedAt: Date;
  versionNo: number;
}

export type LiquorForEdit = Omit<
  Liquor,
  'updatedAt' | 'imageUrl' | 'categoryName'
>;

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
      versionNo
    }
  }
`;

export const LIQUOR_DETAIL_FOR_EDIT: DocumentNode = gql`
  query Liquor($id: String!) {
    liquor(id: $id) {
      id
      name
      categoryId
      description
      imageBase64
      versionNo
    }
  }
`;

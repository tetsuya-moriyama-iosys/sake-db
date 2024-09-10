import { gql, type DocumentNode } from '@apollo/client/core';
import { type Liquor as CardLiquor } from '../Index/random';

export interface LiquorResponse<T> {
  liquor: T;
}

export interface ListFromCategoryResponse {
  listFromCategory: {
    categoryName: string;
    categoryDescription: string;
    liquors: CardLiquor[];
  };
}

export interface Liquor extends CardLiquor {
  imageUrl: string;
  updatedAt: Date;
  versionNo: number;
  categoryTrail: {
    id: number;
    name: string;
  }[];
}

export type LiquorForEdit = Omit<
  Liquor,
  'updatedAt' | 'imageUrl' | 'categoryName' | 'categoryTrail'
>;

export const LIQUOR_DETAIL_GET: DocumentNode = gql`
  query ($id: String!) {
    liquor(id: $id) {
      id
      name
      categoryId
      categoryName
      categoryTrail {
        id
        name
      }
      description
      imageBase64
      imageUrl
      createdAt
      updatedAt
      versionNo
    }
  }
`;

export const LIQUOR_LIST_FROM_CATEGORY: DocumentNode = gql`
  query ($id: Int!) {
    listFromCategory(categoryId: $id) {
      categoryName
      categoryDescription
      liquors {
        id
        name
        categoryId
        categoryName
        description
        imageBase64
        updatedAt
      }
    }
  }
`;

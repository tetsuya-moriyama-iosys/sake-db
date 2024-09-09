import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql';

export interface Category {
  id: number;
  name: string;
  parent: number;
  imageBase64: string;
  imageUrl: string;
  description: string;
  versionNo: number;
  createdAt: Date | null;
  updatedAt: Date | null;
  children: Category[] | null;
}

//export type CategoryForEdit = Omit<Category, 'imageUrl' | 'children'>;

export interface Categories {
  categories: Category[];
}

export interface CategoryResponse<T> {
  category: T;
}

export const GET_DETAIL: DocumentNode = gql`
  query ($id: Int!) {
    category(id: $id) {
      id
      name
      description
      imageBase64
      imageUrl
      description
      children {
        id
        name
      }
    }
  }
`;

export const GET_DETAIL_FOR_EDIT: DocumentNode = gql`
  query ($id: Int!) {
    category(id: $id) {
      id
      name
      parent
      description
      imageBase64
      description
      versionNo
    }
  }
`;

//TODO:もうちょっといい方法がないか考えたいが、一旦保留
export const GET_QUERY: DocumentNode = gql`
  query {
    categories {
      id
      name
      children {
        id
        name
        children {
          id
          name
          children {
            id
            name
            children {
              id
              name
              children {
                id
                name
                children {
                  id
                  name
                }
              }
            }
          }
        }
      }
    }
  }
`;

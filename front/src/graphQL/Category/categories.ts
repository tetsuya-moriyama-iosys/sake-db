import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql';

export interface Category {
  readonly id: number;
  readonly name: string;
  readonly parent: number;
  readonly imageBase64: string;
  readonly imageUrl: string;
  readonly description: string;
  readonly versionNo: number;
  readonly readonly: boolean;
  readonly createdAt: Date | null;
  readonly updatedAt: Date | null;
  readonly children: Category[] | null;
}

export interface Categories {
  readonly categories: Category[];
}

export interface CategoryResponse<T> {
  readonly category: T;
}

export const GET_DETAIL: DocumentNode = gql`
  query category($id: Int!) {
    category(id: $id) {
      id
      name
      description
      imageBase64
      imageUrl
      description
      readonly
      children {
        id
        name
      }
    }
  }
`;

//TODO:もうちょっといい方法がないか考えたいが、一旦保留(フラグメントとか使えばもうちょっとマシになるかも)
export const GET_QUERY: DocumentNode = gql`
  query categories {
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
              }
            }
          }
        }
      }
    }
  }
`;

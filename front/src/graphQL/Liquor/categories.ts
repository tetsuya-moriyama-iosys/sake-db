import type { Category } from '@/type/common/liquor/Category';
import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql';

export interface Categories {
  categories: Category[];
}

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

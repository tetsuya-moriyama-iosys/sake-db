import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';
import type { Category } from '@/graphQL/Category/categories';

export interface HistoryResponse {
  histories: CategoryHistoryData;
}

export interface CategoryHistoryData {
  now: Category;
  histories: Category[] | null;
}

export const GET_LOGS_FOR_ROLLBACK: DocumentNode = gql`
  query ($id: Int!) {
    histories(id: $id) {
      now {
        id
        name
        parent
        description
        imageBase64
        versionNo
        updatedAt
      }
      histories {
        id
        name
        parent
        description
        imageBase64
        versionNo
        updatedAt
      }
    }
  }
`;
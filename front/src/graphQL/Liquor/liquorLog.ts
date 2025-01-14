import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql/index';

import type { Liquor } from '@/graphQL/Liquor/liquor';

export interface HistoryResponse {
  readonly liquorHistories: LiquorHistoryData;
}

export interface LiquorHistoryData {
  readonly now: Liquor;
  readonly histories: Liquor[] | null;
}

export const GET_LOGS_FOR_ROLLBACK: DocumentNode = gql`
  query liquorHistories($id: String!) {
    liquorHistories(id: $id) {
      now {
        id
        name
        categoryId
        description
        imageBase64
        versionNo
        updatedAt
      }
      histories {
        id
        name
        categoryId
        description
        imageBase64
        versionNo
        updatedAt
      }
    }
  }
`;

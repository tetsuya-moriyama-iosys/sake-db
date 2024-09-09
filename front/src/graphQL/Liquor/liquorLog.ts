import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';
import type { Liquor } from '@/graphQL/Liquor/liquor';

export interface HistoryResponse {
  liquorHistories: LiquorHistoryData;
}

export interface LiquorHistoryData {
  now: Liquor;
  histories: Liquor[] | null;
}

export const GET_LOGS_FOR_ROLLBACK: DocumentNode = gql`
  query ($id: Int!) {
    liquorHistories(id: $id) {
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
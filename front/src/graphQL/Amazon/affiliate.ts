import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql';

export interface AffiliateData {
  readonly items: AffiliateItem[];
  readonly lowestPrice?: number;
}

export interface AffiliateItem {
  readonly name: string;
  readonly URL: string;
  readonly imageURL?: string;
  readonly price?: number;
}
export interface AffiliateResponse {
  readonly data: AffiliateData;
}

export const GET_AFFILIATE_LIST: DocumentNode = gql`
  query ($keyword: String!) {
    data(name: $keyword) {
      items {
        name
        URL
        imageURL
        price
      }
      lowestPrice
    }
  }
`;

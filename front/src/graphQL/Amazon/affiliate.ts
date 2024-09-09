import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql';

export interface AffiliateData {
  items: AffiliateItem[];
  lowestPrice?: number;
}

export interface AffiliateItem {
  name: string;
  URL: string;
  imageURL?: string;
  price?: number;
}
export interface AffiliateResponse {
  data: AffiliateData;
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

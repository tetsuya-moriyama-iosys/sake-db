import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';

export interface CheckResponse {
  readonly getIsBookMarked: boolean;
}
export interface AddResponse {
  readonly addBookMark: boolean;
}
export interface RemoveResponse {
  readonly removeBookMark: boolean;
}

export const CHECK: DocumentNode = gql`
  query ($id: String!) {
    getIsBookMarked(id: $id)
  }
`;
export const ADD: DocumentNode = gql`
  mutation ($id: String!) {
    addBookMark(id: $id)
  }
`;
export const REMOVE: DocumentNode = gql`
  mutation ($id: String!) {
    removeBookMark(id: $id)
  }
`;

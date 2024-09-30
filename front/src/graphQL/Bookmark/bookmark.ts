import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';

export interface GetBookmarkListResponse {
  readonly getBookMarkList: Bookmark[] | null;
}
export interface GetBookmarkedListResponse {
  readonly getBookMarkedList: Bookmark[] | null;
}

export interface Bookmark {
  readonly userId: string;
  readonly name: string;
  readonly imageBase64: string | null;
  readonly createdAt: Date;
}
export interface CheckResponse {
  readonly getIsBookMarked: boolean;
}
export interface AddResponse {
  readonly addBookMark: boolean;
}
export interface RemoveResponse {
  readonly removeBookMark: boolean;
}

export const LIST: DocumentNode = gql`
  query {
    getBookMarkList {
      userId
      name
      imageBase64
      createdAt
    }
  }
`;

//被ブックマークのリスト
export const BOOKMARKED_LIST: DocumentNode = gql`
  query ($id: ID!) {
    getBookMarkedList(id: $id) {
      userId
      name
      imageBase64
      createdAt
    }
  }
`;

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

import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';

export interface BoardResponse {
  readonly board: Post[];
}

export interface Post {
  readonly id: string;
  readonly name: string | null;
  readonly userId: string | null;
  readonly text: string;
  readonly rate: number | null;
  readonly updatedAt: Date;
}

export const Post: DocumentNode = gql`
  mutation ($input: BoardInput!) {
    postBoard(input: $input)
  }
`;

export const GET_BOARD: DocumentNode = gql`
  query ($liquorId: String!) {
    board(liquorId: $liquorId) {
      id
      name
      userId
      text
      rate
      updatedAt
    }
  }
`;

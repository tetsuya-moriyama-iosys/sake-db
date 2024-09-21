import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';

export interface BoardResponse {
  readonly board: Post[] | null;
}
export interface MyBoardResponse {
  readonly getMyBoard: PostCore | null;
}

export interface PostCore {
  readonly text: string;
  readonly rate: number | null;
}

export interface Post extends PostCore {
  readonly userName: string | null;
  readonly userId: string | null;
  readonly updatedAt: Date;
}
export const Post: DocumentNode = gql`
  mutation ($input: BoardInput!) {
    postBoard(input: $input)
  }
`;
export const GetMyPostByLiquorId: DocumentNode = gql`
  query ($id: String!) {
    getMyBoard(liquorId: $id) {
      text
      rate
    }
  }
`;

export const GET_BOARD: DocumentNode = gql`
  query ($liquorId: String!) {
    board(liquorId: $liquorId) {
      userName
      userId
      text
      rate
      updatedAt
    }
  }
`;

import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql/index';

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
  mutation postBoard($input: BoardInput!) {
    postBoard(input: $input)
  }
`;
export const GetMyPostByLiquorId: DocumentNode = gql`
  query getMyBoard($id: String!) {
    getMyBoard(liquorId: $id) {
      text
      rate
    }
  }
`;

export const GET_BOARD: DocumentNode = gql`
  query board($liquorId: String!) {
    board(liquorId: $liquorId) {
      userId
      userName
      text
      rate
      updatedAt
    }
  }
`;

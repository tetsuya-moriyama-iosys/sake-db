import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';
import { wrapVariables } from '@/graphQL/core';

export interface BoardResponse<T extends PostCore = Post> {
  readonly board: T[];
}

export interface PostCore {
  readonly text: string;
  readonly rate: number | null;
}

export interface Post extends PostCore {
  readonly id: string;
  readonly name: string | null;
  readonly userId: string | null;
  readonly updatedAt: Date;
}

export function myBoardRequest(liquorId: string) {
  return wrapVariables({
    id: liquorId,
  });
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
      id
      name
      userId
      text
      rate
      updatedAt
    }
  }
`;

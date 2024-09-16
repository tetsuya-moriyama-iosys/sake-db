import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';

export const Post: DocumentNode = gql`
  mutation ($input: BoardInput!) {
    postBoard(input: $input)
  }
`;

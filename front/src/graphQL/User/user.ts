import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';
import type { User } from '@/graphQL/Auth/auth';

export interface GetUserdataResponse {
  getUser: User;
}

export const GET_USERDATA: DocumentNode = gql`
  query {
    getUser {
      id
      name
      email
    }
  }
`;

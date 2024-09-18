import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';
import type { AuthUser } from '@/graphQL/User/user';

export interface LoginResponse {
  readonly login: LoginResult;
}
export interface LoginResult {
  readonly token: string;
  readonly user: AuthUser;
}

export const Register: DocumentNode = gql`
  mutation ($input: RegisterInput!) {
    register(input: $input) {
      id
      name
      email
    }
  }
`;

export const LOGIN: DocumentNode = gql`
  mutation ($input: LoginInput!) {
    login(input: $input) {
      token
      user {
        id
        name
        email
      }
    }
  }
`;

import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';

export interface User {
  id: string;
  name: string;
  email: string;
}

export interface LoginResponse {
  login: LoginResult;
}
export interface LoginResult {
  token: string;
  user: User;
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

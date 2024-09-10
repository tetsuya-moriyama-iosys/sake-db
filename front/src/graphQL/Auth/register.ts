import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';

export interface User {
  id: string;
  name: string;
  email: string;
}

export const REGISTER: DocumentNode = gql`
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
      id
      name
      email
    }
  }
`;

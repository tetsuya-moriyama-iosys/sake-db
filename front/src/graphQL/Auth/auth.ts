import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql';

import { AuthFragment } from '@/graphQL/Auth/fragment';
import type { AuthPayload } from '@/graphQL/auto-generated';

export interface ResetEmailExeResponse {
  readonly resetExe: Omit<AuthPayload, '__typename'>;
}

export const Register: DocumentNode = gql`
  mutation registerUser($input: RegisterInput!) {
    registerUser(input: $input) {
      ...AuthPayload
    }
  }
  ${AuthFragment}
`;

export const LOGIN: DocumentNode = gql`
  mutation login($input: LoginInput!) {
    login(input: $input) {
      ...AuthPayload
    }
  }
  ${AuthFragment}
`;

//自身のフルデータ
export const GET_MY_USERDATA_FULL: DocumentNode = gql`
  query getMyDataFull {
    getMyData {
      id
      name
      email
      profile
      imageBase64
      roles
    }
  }
`;

export const PASSWORD_RESET: DocumentNode = gql`
  mutation resetEmail($email: String!) {
    resetEmail(email: $email)
  }
`;
export const PASSWORD_RESET_EXE: DocumentNode = gql`
  mutation resetExe($token: String!, $password: String!) {
    resetExe(token: $token, password: $password) {
      ...AuthPayload
    }
  }
  ${AuthFragment}
`;

export const REFRESH_TOKEN = gql`
  mutation RefreshToken {
    refreshToken
  }
`;

export const LOGIN_WITH_REFRESH_TOKEN = gql`
  mutation loginWithRefreshToken {
    loginWithRefreshToken {
      ...AuthPayload
    }
  }
  ${AuthFragment}
`;

export const LOGOUT = gql`
  mutation logout {
    logout
  }
`;

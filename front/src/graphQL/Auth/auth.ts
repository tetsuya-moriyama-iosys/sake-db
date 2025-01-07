import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql';

import type { AuthPayload } from '@/graphQL/auto-generated';

//ログイン時に返ってくるデータ
export interface RegisterResponse {
  readonly registerUser: LoginResult;
}
export interface GetUserdataResponse {
  readonly getUser: AuthUserFull;
}

export const Roles = {
  Admin: 'admin',
} as const;
export type Role = (typeof Roles)[keyof typeof Roles];

//認証済みユーザー情報(要はパスワード以外のデータ)
export interface AuthUser {
  id: string;
  name: string;
  imageBase64: string | null; //アイコン表示に必要
  roles: Role[];
}
export interface AuthUserFull extends AuthUser {
  email: string;
  profile: string;
}

export interface ResetEmailExeResponse {
  readonly resetExe: Omit<AuthPayload, '__typename'>;
}

export interface LoginResult {
  readonly accessToken: string;
  readonly user: AuthUser;
}

export interface GetUserResponse {
  readonly getUser: AuthUser;
}

export const AuthUserDataFragment = gql`
  fragment AuthUserDataFragment on User {
    id
    name
    imageBase64
    roles
  }
`;

export const AuthFragment = gql`
  fragment AuthFragment on AuthPayload {
    accessToken
    user {
      ...AuthUserDataFragment
    }
  }
  ${AuthUserDataFragment}
`;

export const Register: DocumentNode = gql`
  mutation registerUser($input: RegisterInput!) {
    registerUser(input: $input) {
      ...AuthFragment
    }
  }
  ${AuthFragment}
`;

export const LOGIN: DocumentNode = gql`
  mutation login($input: LoginInput!) {
    login(input: $input) {
      ...AuthFragment
    }
  }
  ${AuthFragment}
`;

//memo:idはトークンから取るので、inputはRegisterと同値でかまわないが、ログイン判定を必要とするため呼び出すリゾルバが異なる
export const Update: DocumentNode = gql`
  mutation updateUser($input: RegisterInput!) {
    updateUser(input: $input)
  }
`;

//最低限のデータ(再ログイン)
export const GET_USER: DocumentNode = gql`
  query getUser {
    getUser {
      ...AuthUserDataFragment
    }
  }
  ${AuthUserDataFragment}
`;

//自身のフルデータ
export const GET_MY_USERDATA_FULL: DocumentNode = gql`
  query getUserFull {
    getUser {
      id
      name
      email
      profile
      imageBase64
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
      ...AuthFragment
    }
  }
  ${AuthFragment}
`;

export const REFRESH_TOKEN = gql`
  mutation RefreshToken {
    refreshToken {
      accessToken
    }
  }
`;

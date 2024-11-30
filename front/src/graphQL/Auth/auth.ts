import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql';

//ログイン時に返ってくるデータ
export interface RegisterResponse {
  readonly registerUser: LoginResult;
}
export interface GetUserdataResponse {
  readonly getUser: AuthUserFull;
}
export interface LoginResponse {
  readonly login: LoginResult;
}

export const Roles = {
  Admin: 'admin',
} as const;
export type Role = (typeof Roles)[keyof typeof Roles];

//認証済みユーザー情報(要はパスワード以外のデータ)
export interface AuthUser {
  id: string;
  name: string;
  imageBase64: string | undefined; //アイコン表示に必要
  roles: Role[];
}
export interface AuthUserFull extends AuthUser {
  email: string;
  profile: string;
}

export interface ResetEmailExeResponse {
  readonly resetExe: LoginResult;
}
export interface LoginResult {
  readonly token: string;
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
    token
    user {
      ...AuthUserDataFragment
    }
  }
  ${AuthUserDataFragment}
`;

export const Register: DocumentNode = gql`
  mutation ($input: RegisterInput!) {
    registerUser(input: $input) {
      ...AuthFragment
    }
  }
  ${AuthFragment}
`;

export const LOGIN: DocumentNode = gql`
  mutation ($input: LoginInput!) {
    login(input: $input) {
      ...AuthFragment
    }
  }
  ${AuthFragment}
`;

//memo:idはトークンから取るので、inputはRegisterと同値でかまわないが、ログイン判定を必要とするため呼び出すリゾルバが異なる
export const Update: DocumentNode = gql`
  mutation ($input: RegisterInput!) {
    updateUser(input: $input)
  }
`;

//最低限のデータ(再ログイン)
export const GET_USER: DocumentNode = gql`
  query {
    getUser {
      ...AuthUserDataFragment
    }
  }
  ${AuthUserDataFragment}
`;

//自身のフルデータ
export const GET_MY_USERDATA_FULL: DocumentNode = gql`
  query {
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
  mutation ($email: String!) {
    resetEmail(email: $email)
  }
`;
export const PASSWORD_RESET_EXE: DocumentNode = gql`
  mutation ($token: String!, $password: String!) {
    resetExe(token: $token, password: $password) {
      ...AuthFragment
    }
  }
  ${AuthFragment}
`;

import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';

//ログイン時に返ってくるデータ
export interface GetUserdataResponse {
  readonly getUser: AuthUser;
}

//認証済みユーザー情報(要はパスワード以外のデータ)
export interface AuthUser {
  id: string;
  name: string;
  email: string;
  imageBase64: string | undefined; //アイコン表示に必要
}

export interface LoginResponse {
  readonly login: LoginResult;
}
export interface LoginResult {
  readonly token: string;
  readonly user: AuthUser;
}

export interface GetUserResponse {
  readonly getUser: AuthUser;
}

//TODO:トークンを取得してログインまで終わらせる
export const Register: DocumentNode = gql`
  mutation ($input: RegisterInput!) {
    registerUser(input: $input) {
      id
      name
      email
    }
  }
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
      id
      name
      imageBase64
    }
  }
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

export const LOGIN: DocumentNode = gql`
  mutation ($input: LoginInput!) {
    login(input: $input) {
      token
      user {
        id
        name
        imageBase64
      }
    }
  }
`;

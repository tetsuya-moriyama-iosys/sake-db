import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';

export interface GetUserdataResponse<T extends AuthUser> {
  readonly getUser: T;
}

//認証に最低限必要なユーザー情報
export interface AuthUser {
  readonly id: string;
  readonly name: string;
  readonly email: string;
}

//外部に公開可能なユーザー情報
export interface User extends Omit<AuthUser, 'email'> {
  readonly profile: string;
  readonly imageBase64: string;
}

//ユーザー情報(パスワード以外の情報)
export type UserFullData = AuthUser & User;

export interface GetUserByIdResponse {
  readonly getUserById: User;
}

//自身のユーザーデータ
export const GET_MY_USERDATA: DocumentNode = gql`
  query {
    getUser {
      id
      name
      email
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

//指定したIDのユーザーデータ
export const GET_USERDATA: DocumentNode = gql`
  query ($id: String!) {
    getUserById(id: $id) {
      id
      name
      profile
      imageBase64
    }
  }
`;

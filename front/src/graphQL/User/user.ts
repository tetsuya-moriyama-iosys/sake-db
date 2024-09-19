import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';

export interface GetUserdataResponse<T extends AuthUser> {
  readonly getUser: T;
}

//ログイン時に返ってくるデータ
export interface AuthUser {
  readonly id: string;
  readonly name: string;
  readonly imageBase64: string | undefined; //アイコン表示に必要
}

//外部に公開可能なユーザー情報
export interface User extends AuthUser {
  readonly profile: string;
}

//ユーザー情報(パスワード以外の情報)
export type UserFullData = AuthUser & User;

export interface GetUserByIdResponse {
  readonly getUserById: User;
}

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

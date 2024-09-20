import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';

export interface GetUserdataResponse<T extends AuthUser> {
  readonly getUser: T;
}
export interface GetUserDataFullResponse {
  readonly getUserByIdDetail: UserFullData;
}

// export interface DetailType{
//   readonly getUserByIdDetail:
// }

// export interface Detail{
//   readonly rate5Liquors: Liquor[];
//   readonly rate4Liquors: Liquor[];
//   readonly rate3Liquors: Liquor[];
//   readonly rate2Liquors: Liquor[];
//   readonly rate1Liquors: Liquor[];
// }

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
  readonly getUserByIdDetail: User;
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

export const GET_USERDATA_FULL: DocumentNode = gql`
  query ($id: String!) {
    getUserByIdDetail(id: $id) {
      detail {
        rate5Liquors {
          id
          name
          categoryId
          categoryName
          imageBase64
        }
        rate4Liquors {
          id
          name
          categoryId
          categoryName
          imageBase64
        }
        rate3Liquors {
          id
          name
          categoryId
          categoryName
          imageBase64
        }
        rate2Liquors {
          id
          name
          categoryId
          categoryName
          imageBase64
        }
        rate1Liquors {
          id
          name
          categoryId
          categoryName
          imageBase64
        }
      }
      user {
        id
        name
        profile
        imageBase64
      }
    }
  }
`;

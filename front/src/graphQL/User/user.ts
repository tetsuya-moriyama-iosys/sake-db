import type { DocumentNode } from 'graphql/index';
import { gql } from '@apollo/client/core';
import type { AuthUser } from '@/graphQL/Auth/auth';

//ユーザーのフルデータ
export interface GetUserDetailResponse {
  readonly getUserByIdDetail: UserDetail;
}

//ユーザー情報と評価情報が入っているインターフェース
export interface UserDetail {
  readonly evaluateList: EvaluateList;
  readonly user: User;
}

//外部に公開可能なユーザー情報
export interface User extends Omit<AuthUser, 'email'> {
  readonly profile: string;
}

//評価履歴のデータベース
export interface EvaluateList {
  readonly recentComments: UserLiquor[] | null;
  readonly rate5Liquors: UserLiquor[] | null;
  readonly rate4Liquors: UserLiquor[] | null;
  readonly rate3Liquors: UserLiquor[] | null;
  readonly rate2Liquors: UserLiquor[] | null;
  readonly rate1Liquors: UserLiquor[] | null;
  readonly noRateLiquors: UserLiquor[] | null;
}

//評価レコード
export interface UserLiquor {
  readonly id: string;
  readonly categoryId: string;
  readonly categoryName: string;
  readonly liquorId: string;
  readonly name: string;
  readonly imageBase64: string;
  readonly comment: string;
  readonly rate: number;
  readonly updatedAt: Date;
}

//ユーザーページ用のフルデータ
export const GET_USERDATA_FULL: DocumentNode = gql`
  query ($id: String!) {
    getUserByIdDetail(id: $id) {
      evaluateList {
        recentComments {
          id
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
        }
        rate5Liquors {
          id
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
        }
        rate4Liquors {
          id
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
        }
        rate3Liquors {
          id
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
        }
        rate2Liquors {
          id
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
        }
        rate1Liquors {
          id
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
        }
        noRateLiquors {
          id
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
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

import { gql } from '@apollo/client';
import type { DocumentNode } from 'graphql/index';

import type { User } from '@/graphQL/auto-generated';

//ユーザーのフルデータ
export interface GetUserDetailResponse {
  readonly getUserByIdDetail: UserDetail;
}

//ユーザー情報と評価情報が入っているインターフェース
export interface UserDetail {
  readonly evaluateList: EvaluateList;
  readonly user: Omit<User, '__typename' | 'email'>;
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
  query getUserByIdDetail($id: String!) {
    getUserByIdDetail(id: $id) {
      evaluateList {
        recentComments {
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

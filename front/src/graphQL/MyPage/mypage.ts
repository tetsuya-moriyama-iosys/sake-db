//memo:idはトークンから取るので、inputはRegisterと同値でかまわないが、ログイン判定を必要とするため呼び出すリゾルバが異なる
import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql/index';

export const Update: DocumentNode = gql`
  mutation updateUser($input: RegisterInput!) {
    updateUser(input: $input)
  }
`;

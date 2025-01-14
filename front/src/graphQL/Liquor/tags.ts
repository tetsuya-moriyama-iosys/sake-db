import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql/index';

import type { Tag } from '@/graphQL/Liquor/liquor';

//Formでも使い回すため、相互参照を防止するためにこちらで定義
export const PostTagKeys = {
  LiquorId: 'liquorId',
  Tag: 'text',
} as const;

export interface PostTag {
  readonly [PostTagKeys.LiquorId]: string;
  readonly [PostTagKeys.Tag]: string;
}

export interface GetTagsResponse {
  readonly getTags: Tag[];
}
export interface PostTagResponse {
  readonly postTag: Tag[];
}

export const FetchTags: DocumentNode = gql`
  query getTags($liquorId: ID!) {
    getTags(liquorId: $liquorId) {
      id
      text
    }
  }
`;
export const PostTag: DocumentNode = gql`
  mutation postTag($input: TagInput!) {
    postTag(input: $input) {
      id
      text
    }
  }
`;

export const DeleteTag: DocumentNode = gql`
  mutation deleteTag($id: ID!) {
    deleteTag(id: $id)
  }
`;

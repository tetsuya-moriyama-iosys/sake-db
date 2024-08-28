import { gql, type DocumentNode } from '@apollo/client/core';

export const CREATE_POST_MUTATION: DocumentNode = gql`
  mutation CreateLiquor(
    $name: String!
    $category_id: Int!
    $description: String
    $image: Upload
  ) {
    createLiquor(
      inputs: {
        name: $name
        category_id: $category_id
        description: $description
        image: $image
      }
    ) {
      id
      name
      category_id
      description
      created_at
      updated_at
    }
  }
`;

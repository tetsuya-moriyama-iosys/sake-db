import { gql } from '@apollo/client/core';

export const GET_MESSAGES = gql`
  query {
    messages {
      id
      message
    }
  }
`;

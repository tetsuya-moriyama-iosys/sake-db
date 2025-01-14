import { gql } from '@apollo/client/core';

export const AuthUserDataFragment = gql`
  fragment AuthUserDataFragment on User {
    id
    name
    imageBase64
    roles
  }
`;

export const AuthFragment = gql`
  fragment AuthPayload on AuthPayload {
    accessToken
    user {
      ...AuthUserDataFragment
    }
  }
  ${AuthUserDataFragment}
`;

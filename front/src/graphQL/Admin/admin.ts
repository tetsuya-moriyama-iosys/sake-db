import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql/index';

export const AdminCheck: DocumentNode = gql`
  query adminCheck {
    checkAdmin
  }
`;

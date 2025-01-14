import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql/index';

export type FlavorMapResponse = {
  getFlavorMap: FlavorMap;
};
export type VotedResponse = {
  getVoted: Coordinates;
};

export type FlavorMap = {
  categoryId: number;
  xNames: [string, string];
  yNames: [string, string];
  userFullAmount: number;
  guestFullAmount: number;
  mapData: FlavorCell[];
};
export interface FlavorCell extends Coordinates {
  rate: number;
  userAmount: number;
  guestAmount: number;
}

export interface PostFlavorMap extends Coordinates {
  liquorId: string;
}

export type Coordinates = {
  x: number;
  y: number;
};

export const GetFlavorMap: DocumentNode = gql`
  query getFlavorMap($liquorId: ID!) {
    getFlavorMap(liquorId: $liquorId) {
      categoryId
      xNames
      yNames
      userFullAmount
      guestFullAmount
      mapData {
        x
        y
        rate
        userAmount
        guestAmount
      }
    }
  }
`;

export const PostFlavorMap: DocumentNode = gql`
  mutation postFlavor($input: PostFlavorMap!) {
    postFlavor(input: $input)
  }
`;

export const GetVoted: DocumentNode = gql`
  query getVoted($liquorId: ID!) {
    getVoted(liquorId: $liquorId) {
      x
      y
    }
  }
`;

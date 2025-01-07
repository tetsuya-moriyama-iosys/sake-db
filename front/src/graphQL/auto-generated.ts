import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Coordinate: { input: any; output: any; }
  DateTime: { input: any; output: any; }
};

export type AffiliateData = {
  __typename: 'AffiliateData';
  items?: Maybe<Array<AffiliateItem>>;
  lowestPrice?: Maybe<Scalars['Int']['output']>;
};

export type AffiliateItem = {
  __typename: 'AffiliateItem';
  URL: Scalars['String']['output'];
  imageURL?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
  price?: Maybe<Scalars['Int']['output']>;
};

export type AuthPayload = {
  __typename: 'AuthPayload';
  accessToken: Scalars['String']['output'];
  user: User;
};

export type BoardInput = {
  liquorID: Scalars['String']['input'];
  rate?: InputMaybe<Scalars['Int']['input']>;
  text: Scalars['String']['input'];
};

export type BoardPost = {
  __typename: 'BoardPost';
  categoryId: Scalars['Int']['output'];
  categoryName: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  liquorId: Scalars['ID']['output'];
  liquorName: Scalars['String']['output'];
  rate?: Maybe<Scalars['Int']['output']>;
  text: Scalars['String']['output'];
  updatedAt: Scalars['DateTime']['output'];
  userId?: Maybe<Scalars['ID']['output']>;
  userName?: Maybe<Scalars['String']['output']>;
  youtube?: Maybe<Scalars['String']['output']>;
};

export type BookMarkListUser = {
  __typename: 'BookMarkListUser';
  createdAt: Scalars['DateTime']['output'];
  imageBase64?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
  userId: Scalars['ID']['output'];
};

export type Category = {
  __typename: 'Category';
  children?: Maybe<Array<Category>>;
  description?: Maybe<Scalars['String']['output']>;
  id: Scalars['Int']['output'];
  imageBase64?: Maybe<Scalars['String']['output']>;
  imageUrl?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
  parent?: Maybe<Scalars['Int']['output']>;
  readonly: Scalars['Boolean']['output'];
  updatedAt?: Maybe<Scalars['DateTime']['output']>;
  versionNo?: Maybe<Scalars['Int']['output']>;
};

export type CategoryHistory = {
  __typename: 'CategoryHistory';
  histories?: Maybe<Array<Maybe<Category>>>;
  now: Category;
};

export type CategoryTrail = {
  __typename: 'CategoryTrail';
  id: Scalars['Int']['output'];
  name: Scalars['String']['output'];
};

export type FlavorCellData = {
  __typename: 'FlavorCellData';
  guestAmount: Scalars['Int']['output'];
  rate: Scalars['Float']['output'];
  userAmount: Scalars['Int']['output'];
  x: Scalars['Coordinate']['output'];
  y: Scalars['Coordinate']['output'];
};

export type FlavorMapData = {
  __typename: 'FlavorMapData';
  categoryId: Scalars['Int']['output'];
  guestFullAmount: Scalars['Int']['output'];
  mapData: Array<FlavorCellData>;
  userFullAmount: Scalars['Int']['output'];
  xNames: Array<Scalars['String']['output']>;
  yNames: Array<Scalars['String']['output']>;
};

export type Liquor = {
  __typename: 'Liquor';
  categoryId: Scalars['Int']['output'];
  categoryName: Scalars['String']['output'];
  categoryTrail?: Maybe<Array<CategoryTrail>>;
  description?: Maybe<Scalars['String']['output']>;
  id: Scalars['String']['output'];
  imageBase64?: Maybe<Scalars['String']['output']>;
  imageUrl?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
  rate1Users: Array<Scalars['ID']['output']>;
  rate2Users: Array<Scalars['ID']['output']>;
  rate3Users: Array<Scalars['ID']['output']>;
  rate4Users: Array<Scalars['ID']['output']>;
  rate5Users: Array<Scalars['ID']['output']>;
  updatedAt: Scalars['DateTime']['output'];
  versionNo: Scalars['Int']['output'];
  youtube?: Maybe<Scalars['String']['output']>;
};

export type LiquorHistory = {
  __typename: 'LiquorHistory';
  histories?: Maybe<Array<Maybe<Liquor>>>;
  now: Liquor;
};

export type ListFromCategory = {
  __typename: 'ListFromCategory';
  categoryDescription?: Maybe<Scalars['String']['output']>;
  categoryName: Scalars['String']['output'];
  liquors: Array<Maybe<Liquor>>;
};

export type LoginInput = {
  email: Scalars['String']['input'];
  password: Scalars['String']['input'];
};

export type Mutation = {
  __typename: 'Mutation';
  addBookMark: Scalars['Boolean']['output'];
  deleteTag: Scalars['Boolean']['output'];
  login: AuthPayload;
  postBoard: Scalars['Boolean']['output'];
  postFlavor: Scalars['Boolean']['output'];
  postTag: Tag;
  refreshToken: Scalars['String']['output'];
  registerUser: AuthPayload;
  removeBookMark: Scalars['Boolean']['output'];
  resetEmail: Scalars['Boolean']['output'];
  resetExe: AuthPayload;
  updateUser: Scalars['Boolean']['output'];
};


export type MutationAddBookMarkArgs = {
  id: Scalars['String']['input'];
};


export type MutationDeleteTagArgs = {
  id: Scalars['ID']['input'];
};


export type MutationLoginArgs = {
  input: LoginInput;
};


export type MutationPostBoardArgs = {
  input: BoardInput;
};


export type MutationPostFlavorArgs = {
  input: PostFlavorMap;
};


export type MutationPostTagArgs = {
  input: TagInput;
};


export type MutationRegisterUserArgs = {
  input: RegisterInput;
};


export type MutationRemoveBookMarkArgs = {
  id: Scalars['String']['input'];
};


export type MutationResetEmailArgs = {
  email: Scalars['String']['input'];
};


export type MutationResetExeArgs = {
  password: Scalars['String']['input'];
  token: Scalars['String']['input'];
};


export type MutationUpdateUserArgs = {
  input: RegisterInput;
};

export type PostFlavorMap = {
  liquorId: Scalars['ID']['input'];
  x: Scalars['Coordinate']['input'];
  y: Scalars['Coordinate']['input'];
};

export type Query = {
  __typename: 'Query';
  board?: Maybe<Array<BoardPost>>;
  categories: Array<Category>;
  category: Category;
  checkAdmin: Scalars['Boolean']['output'];
  data: AffiliateData;
  getBookMarkList?: Maybe<Array<BookMarkListUser>>;
  getBookMarkedList?: Maybe<Array<BookMarkListUser>>;
  getFlavorMap?: Maybe<FlavorMapData>;
  getIsBookMarked: Scalars['Boolean']['output'];
  getMyBoard?: Maybe<BoardPost>;
  getRecommendLiquorList: Array<Recommend>;
  getTags: Array<Tag>;
  getUser: User;
  getUserById: User;
  getUserByIdDetail: UserPageData;
  getVoted?: Maybe<VotedData>;
  histories?: Maybe<CategoryHistory>;
  liquor: Liquor;
  liquorHistories?: Maybe<LiquorHistory>;
  listFromCategory: ListFromCategory;
  randomRecommendList: Array<Liquor>;
};


export type QueryBoardArgs = {
  liquorId: Scalars['String']['input'];
  page?: InputMaybe<Scalars['Int']['input']>;
};


export type QueryCategoryArgs = {
  id: Scalars['Int']['input'];
};


export type QueryDataArgs = {
  limit?: InputMaybe<Scalars['Int']['input']>;
  name: Scalars['String']['input'];
};


export type QueryGetBookMarkedListArgs = {
  id: Scalars['ID']['input'];
};


export type QueryGetFlavorMapArgs = {
  liquorId: Scalars['ID']['input'];
};


export type QueryGetIsBookMarkedArgs = {
  id: Scalars['String']['input'];
};


export type QueryGetMyBoardArgs = {
  liquorId: Scalars['String']['input'];
};


export type QueryGetTagsArgs = {
  liquorId: Scalars['ID']['input'];
};


export type QueryGetUserByIdArgs = {
  id: Scalars['String']['input'];
};


export type QueryGetUserByIdDetailArgs = {
  id: Scalars['String']['input'];
};


export type QueryGetVotedArgs = {
  liquorId: Scalars['ID']['input'];
};


export type QueryHistoriesArgs = {
  id: Scalars['Int']['input'];
};


export type QueryLiquorArgs = {
  id: Scalars['String']['input'];
};


export type QueryLiquorHistoriesArgs = {
  id: Scalars['String']['input'];
};


export type QueryListFromCategoryArgs = {
  categoryId: Scalars['Int']['input'];
};


export type QueryRandomRecommendListArgs = {
  limit: Scalars['Int']['input'];
};

export type Recommend = {
  __typename: 'Recommend';
  comment: Scalars['String']['output'];
  liquor: RecommendLiquor;
  rate: Scalars['Int']['output'];
  updatedAt: Scalars['DateTime']['output'];
  user: RecommendUser;
};

export type RecommendLiquor = {
  __typename: 'RecommendLiquor';
  categoryId: Scalars['Int']['output'];
  categoryName: Scalars['String']['output'];
  description: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  imageBase64?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
};

export type RecommendUser = {
  __typename: 'RecommendUser';
  id: Scalars['ID']['output'];
  imageBase64?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
};

export type RegisterInput = {
  email: Scalars['String']['input'];
  imageBase64?: InputMaybe<Scalars['String']['input']>;
  name: Scalars['String']['input'];
  password?: InputMaybe<Scalars['String']['input']>;
  profile?: InputMaybe<Scalars['String']['input']>;
};

export type Tag = {
  __typename: 'Tag';
  id: Scalars['ID']['output'];
  text: Scalars['String']['output'];
};

export type TagInput = {
  liquorId: Scalars['ID']['input'];
  text: Scalars['String']['input'];
};

export type User = {
  __typename: 'User';
  email: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  imageBase64?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
  profile?: Maybe<Scalars['String']['output']>;
  roles?: Maybe<Array<Scalars['String']['output']>>;
};

export type UserEvaluateList = {
  __typename: 'UserEvaluateList';
  noRateLiquors?: Maybe<Array<UserLiquor>>;
  rate1Liquors?: Maybe<Array<UserLiquor>>;
  rate2Liquors?: Maybe<Array<UserLiquor>>;
  rate3Liquors?: Maybe<Array<UserLiquor>>;
  rate4Liquors?: Maybe<Array<UserLiquor>>;
  rate5Liquors?: Maybe<Array<UserLiquor>>;
  recentComments?: Maybe<Array<UserLiquor>>;
};

export type UserLiquor = {
  __typename: 'UserLiquor';
  categoryId: Scalars['Int']['output'];
  categoryName: Scalars['String']['output'];
  comment?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  imageBase64?: Maybe<Scalars['String']['output']>;
  liquorId: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  rate?: Maybe<Scalars['Int']['output']>;
  updatedAt: Scalars['DateTime']['output'];
};

export type UserPageData = {
  __typename: 'UserPageData';
  evaluateList: UserEvaluateList;
  user: User;
};

export type VotedData = {
  __typename: 'VotedData';
  categoryId: Scalars['Int']['output'];
  liquorId: Scalars['ID']['output'];
  updatedAt: Scalars['DateTime']['output'];
  userId: Scalars['ID']['output'];
  x: Scalars['Coordinate']['output'];
  y: Scalars['Coordinate']['output'];
};

export type AdminCheckQueryVariables = Exact<{ [key: string]: never; }>;


export type AdminCheckQuery = { __typename: 'Query', checkAdmin: boolean };

export type DataQueryVariables = Exact<{
  keyword: Scalars['String']['input'];
  limit?: InputMaybe<Scalars['Int']['input']>;
}>;


export type DataQuery = { __typename: 'Query', data: { __typename: 'AffiliateData', lowestPrice?: number | null, items?: Array<{ __typename: 'AffiliateItem', name: string, URL: string, imageURL?: string | null, price?: number | null }> | null } };

export type AuthUserDataFragmentFragment = { __typename: 'User', id: string, name: string, imageBase64?: string | null, roles?: Array<string> | null };

export type AuthFragmentFragment = { __typename: 'AuthPayload', accessToken: string, user: { __typename: 'User', id: string, name: string, imageBase64?: string | null, roles?: Array<string> | null } };

export type RegisterUserMutationVariables = Exact<{
  input: RegisterInput;
}>;


export type RegisterUserMutation = { __typename: 'Mutation', registerUser: { __typename: 'AuthPayload', accessToken: string, user: { __typename: 'User', id: string, name: string, imageBase64?: string | null, roles?: Array<string> | null } } };

export type LoginMutationVariables = Exact<{
  input: LoginInput;
}>;


export type LoginMutation = { __typename: 'Mutation', login: { __typename: 'AuthPayload', accessToken: string, user: { __typename: 'User', id: string, name: string, imageBase64?: string | null, roles?: Array<string> | null } } };

export type UpdateUserMutationVariables = Exact<{
  input: RegisterInput;
}>;


export type UpdateUserMutation = { __typename: 'Mutation', updateUser: boolean };

export type GetUserQueryVariables = Exact<{ [key: string]: never; }>;


export type GetUserQuery = { __typename: 'Query', getUser: { __typename: 'User', id: string, name: string, imageBase64?: string | null, roles?: Array<string> | null } };

export type GetUserFullQueryVariables = Exact<{ [key: string]: never; }>;


export type GetUserFullQuery = { __typename: 'Query', getUser: { __typename: 'User', id: string, name: string, email: string, profile?: string | null, imageBase64?: string | null } };

export type ResetEmailMutationVariables = Exact<{
  email: Scalars['String']['input'];
}>;


export type ResetEmailMutation = { __typename: 'Mutation', resetEmail: boolean };

export type ResetExeMutationVariables = Exact<{
  token: Scalars['String']['input'];
  password: Scalars['String']['input'];
}>;


export type ResetExeMutation = { __typename: 'Mutation', resetExe: { __typename: 'AuthPayload', accessToken: string, user: { __typename: 'User', id: string, name: string, imageBase64?: string | null, roles?: Array<string> | null } } };

export type RefreshTokenMutationVariables = Exact<{ [key: string]: never; }>;


export type RefreshTokenMutation = { __typename: 'Mutation', refreshToken: string };

export type GetBookMarkListQueryVariables = Exact<{ [key: string]: never; }>;


export type GetBookMarkListQuery = { __typename: 'Query', getBookMarkList?: Array<{ __typename: 'BookMarkListUser', userId: string, name: string, imageBase64?: string | null, createdAt: any }> | null };

export type GetBookMarkedListQueryVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type GetBookMarkedListQuery = { __typename: 'Query', getBookMarkedList?: Array<{ __typename: 'BookMarkListUser', userId: string, name: string, imageBase64?: string | null, createdAt: any }> | null };

export type GetIsBookMarkedQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type GetIsBookMarkedQuery = { __typename: 'Query', getIsBookMarked: boolean };

export type AddBookMarkMutationVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type AddBookMarkMutation = { __typename: 'Mutation', addBookMark: boolean };

export type RemoveBookMarkMutationVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type RemoveBookMarkMutation = { __typename: 'Mutation', removeBookMark: boolean };

export type CategoryQueryVariables = Exact<{
  id: Scalars['Int']['input'];
}>;


export type CategoryQuery = { __typename: 'Query', category: { __typename: 'Category', id: number, name: string, description?: string | null, imageBase64?: string | null, imageUrl?: string | null, readonly: boolean, children?: Array<{ __typename: 'Category', id: number, name: string }> | null } };

export type CategoriesQueryVariables = Exact<{ [key: string]: never; }>;


export type CategoriesQuery = { __typename: 'Query', categories: Array<{ __typename: 'Category', id: number, name: string, children?: Array<{ __typename: 'Category', id: number, name: string, children?: Array<{ __typename: 'Category', id: number, name: string, children?: Array<{ __typename: 'Category', id: number, name: string, children?: Array<{ __typename: 'Category', id: number, name: string, children?: Array<{ __typename: 'Category', id: number, name: string, children?: Array<{ __typename: 'Category', id: number, name: string, children?: Array<{ __typename: 'Category', id: number, name: string, children?: Array<{ __typename: 'Category', id: number, name: string, children?: Array<{ __typename: 'Category', id: number, name: string, children?: Array<{ __typename: 'Category', id: number, name: string, children?: Array<{ __typename: 'Category', id: number, name: string, children?: Array<{ __typename: 'Category', id: number, name: string, children?: Array<{ __typename: 'Category', id: number, name: string }> | null }> | null }> | null }> | null }> | null }> | null }> | null }> | null }> | null }> | null }> | null }> | null }> | null }> };

export type HistoriesQueryVariables = Exact<{
  id: Scalars['Int']['input'];
}>;


export type HistoriesQuery = { __typename: 'Query', histories?: { __typename: 'CategoryHistory', now: { __typename: 'Category', id: number, name: string, parent?: number | null, description?: string | null, imageBase64?: string | null, versionNo?: number | null, readonly: boolean, updatedAt?: any | null }, histories?: Array<{ __typename: 'Category', id: number, name: string, parent?: number | null, description?: string | null, imageBase64?: string | null, versionNo?: number | null, updatedAt?: any | null } | null> | null } | null };

export type RandomRecommendListQueryVariables = Exact<{
  limit: Scalars['Int']['input'];
}>;


export type RandomRecommendListQuery = { __typename: 'Query', randomRecommendList: Array<{ __typename: 'Liquor', id: string, name: string, categoryId: number, categoryName: string, description?: string | null, imageBase64?: string | null, updatedAt: any }> };

export type GetRecommendLiquorListQueryVariables = Exact<{ [key: string]: never; }>;


export type GetRecommendLiquorListQuery = { __typename: 'Query', getRecommendLiquorList: Array<{ __typename: 'Recommend', rate: number, comment: string, updatedAt: any, liquor: { __typename: 'RecommendLiquor', id: string, name: string, categoryId: number, categoryName: string, description: string, imageBase64?: string | null }, user: { __typename: 'RecommendUser', id: string, name: string, imageBase64?: string | null } }> };

export type PostBoardMutationVariables = Exact<{
  input: BoardInput;
}>;


export type PostBoardMutation = { __typename: 'Mutation', postBoard: boolean };

export type GetMyBoardQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type GetMyBoardQuery = { __typename: 'Query', getMyBoard?: { __typename: 'BoardPost', text: string, rate?: number | null } | null };

export type BoardQueryVariables = Exact<{
  liquorId: Scalars['String']['input'];
}>;


export type BoardQuery = { __typename: 'Query', board?: Array<{ __typename: 'BoardPost', userId?: string | null, userName?: string | null, text: string, rate?: number | null, updatedAt: any }> | null };

export type GetFlavorMapQueryVariables = Exact<{
  liquorId: Scalars['ID']['input'];
}>;


export type GetFlavorMapQuery = { __typename: 'Query', getFlavorMap?: { __typename: 'FlavorMapData', categoryId: number, xNames: Array<string>, yNames: Array<string>, userFullAmount: number, guestFullAmount: number, mapData: Array<{ __typename: 'FlavorCellData', x: any, y: any, rate: number, userAmount: number, guestAmount: number }> } | null };

export type PostFlavorMutationVariables = Exact<{
  input: PostFlavorMap;
}>;


export type PostFlavorMutation = { __typename: 'Mutation', postFlavor: boolean };

export type GetVotedQueryVariables = Exact<{
  liquorId: Scalars['ID']['input'];
}>;


export type GetVotedQuery = { __typename: 'Query', getVoted?: { __typename: 'VotedData', x: any, y: any } | null };

export type LiquorQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type LiquorQuery = { __typename: 'Query', liquor: { __typename: 'Liquor', id: string, name: string, categoryId: number, categoryName: string, description?: string | null, youtube?: string | null, imageBase64?: string | null, imageUrl?: string | null, updatedAt: any, versionNo: number, categoryTrail?: Array<{ __typename: 'CategoryTrail', id: number, name: string }> | null } };

export type ListFromCategoryQueryVariables = Exact<{
  id: Scalars['Int']['input'];
}>;


export type ListFromCategoryQuery = { __typename: 'Query', listFromCategory: { __typename: 'ListFromCategory', categoryName: string, categoryDescription?: string | null, liquors: Array<{ __typename: 'Liquor', id: string, name: string, categoryId: number, categoryName: string, description?: string | null, imageBase64?: string | null, updatedAt: any } | null> } };

export type LiquorHistoriesQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type LiquorHistoriesQuery = { __typename: 'Query', liquorHistories?: { __typename: 'LiquorHistory', now: { __typename: 'Liquor', id: string, name: string, categoryId: number, description?: string | null, imageBase64?: string | null, versionNo: number, updatedAt: any }, histories?: Array<{ __typename: 'Liquor', id: string, name: string, categoryId: number, description?: string | null, imageBase64?: string | null, versionNo: number, updatedAt: any } | null> | null } | null };

export type GetTagsQueryVariables = Exact<{
  liquorId: Scalars['ID']['input'];
}>;


export type GetTagsQuery = { __typename: 'Query', getTags: Array<{ __typename: 'Tag', id: string, text: string }> };

export type PostTagMutationVariables = Exact<{
  input: TagInput;
}>;


export type PostTagMutation = { __typename: 'Mutation', postTag: { __typename: 'Tag', id: string, text: string } };

export type DeleteTagMutationVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type DeleteTagMutation = { __typename: 'Mutation', deleteTag: boolean };

export type GetUserByIdDetailQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type GetUserByIdDetailQuery = { __typename: 'Query', getUserByIdDetail: { __typename: 'UserPageData', evaluateList: { __typename: 'UserEvaluateList', recentComments?: Array<{ __typename: 'UserLiquor', categoryId: number, categoryName: string, liquorId: string, name: string, imageBase64?: string | null, comment?: string | null, rate?: number | null, updatedAt: any }> | null, rate5Liquors?: Array<{ __typename: 'UserLiquor', categoryId: number, categoryName: string, liquorId: string, name: string, imageBase64?: string | null, comment?: string | null, rate?: number | null, updatedAt: any }> | null, rate4Liquors?: Array<{ __typename: 'UserLiquor', categoryId: number, categoryName: string, liquorId: string, name: string, imageBase64?: string | null, comment?: string | null, rate?: number | null, updatedAt: any }> | null, rate3Liquors?: Array<{ __typename: 'UserLiquor', categoryId: number, categoryName: string, liquorId: string, name: string, imageBase64?: string | null, comment?: string | null, rate?: number | null, updatedAt: any }> | null, rate2Liquors?: Array<{ __typename: 'UserLiquor', categoryId: number, categoryName: string, liquorId: string, name: string, imageBase64?: string | null, comment?: string | null, rate?: number | null, updatedAt: any }> | null, rate1Liquors?: Array<{ __typename: 'UserLiquor', categoryId: number, categoryName: string, liquorId: string, name: string, imageBase64?: string | null, comment?: string | null, rate?: number | null, updatedAt: any }> | null, noRateLiquors?: Array<{ __typename: 'UserLiquor', categoryId: number, categoryName: string, liquorId: string, name: string, imageBase64?: string | null, comment?: string | null, rate?: number | null, updatedAt: any }> | null }, user: { __typename: 'User', id: string, name: string, profile?: string | null, imageBase64?: string | null } } };

export const AuthUserDataFragmentFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"AuthUserDataFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"User"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"roles"}}]}}]} as unknown as DocumentNode<AuthUserDataFragmentFragment, unknown>;
export const AuthFragmentFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"AuthFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"AuthPayload"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"accessToken"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"AuthUserDataFragment"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"AuthUserDataFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"User"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"roles"}}]}}]} as unknown as DocumentNode<AuthFragmentFragment, unknown>;
export const AdminCheckDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"adminCheck"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"checkAdmin"}}]}}]} as unknown as DocumentNode<AdminCheckQuery, AdminCheckQueryVariables>;
export const DataDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"data"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"keyword"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"limit"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"data"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"keyword"}}},{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"Variable","name":{"kind":"Name","value":"limit"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"URL"}},{"kind":"Field","name":{"kind":"Name","value":"imageURL"}},{"kind":"Field","name":{"kind":"Name","value":"price"}}]}},{"kind":"Field","name":{"kind":"Name","value":"lowestPrice"}}]}}]}}]} as unknown as DocumentNode<DataQuery, DataQueryVariables>;
export const RegisterUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"registerUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"RegisterInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"registerUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"AuthFragment"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"AuthUserDataFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"User"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"roles"}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"AuthFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"AuthPayload"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"accessToken"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"AuthUserDataFragment"}}]}}]}}]} as unknown as DocumentNode<RegisterUserMutation, RegisterUserMutationVariables>;
export const LoginDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"login"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"LoginInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"login"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"AuthFragment"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"AuthUserDataFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"User"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"roles"}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"AuthFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"AuthPayload"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"accessToken"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"AuthUserDataFragment"}}]}}]}}]} as unknown as DocumentNode<LoginMutation, LoginMutationVariables>;
export const UpdateUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"updateUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"RegisterInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}]}]}}]} as unknown as DocumentNode<UpdateUserMutation, UpdateUserMutationVariables>;
export const GetUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"AuthUserDataFragment"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"AuthUserDataFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"User"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"roles"}}]}}]} as unknown as DocumentNode<GetUserQuery, GetUserQueryVariables>;
export const GetUserFullDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getUserFull"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}}]}}]}}]} as unknown as DocumentNode<GetUserFullQuery, GetUserFullQueryVariables>;
export const ResetEmailDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"resetEmail"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"email"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"resetEmail"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"email"},"value":{"kind":"Variable","name":{"kind":"Name","value":"email"}}}]}]}}]} as unknown as DocumentNode<ResetEmailMutation, ResetEmailMutationVariables>;
export const ResetExeDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"resetExe"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"token"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"password"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"resetExe"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"token"},"value":{"kind":"Variable","name":{"kind":"Name","value":"token"}}},{"kind":"Argument","name":{"kind":"Name","value":"password"},"value":{"kind":"Variable","name":{"kind":"Name","value":"password"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"AuthFragment"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"AuthUserDataFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"User"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"roles"}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"AuthFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"AuthPayload"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"accessToken"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"AuthUserDataFragment"}}]}}]}}]} as unknown as DocumentNode<ResetExeMutation, ResetExeMutationVariables>;
export const RefreshTokenDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"RefreshToken"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"refreshToken"}}]}}]} as unknown as DocumentNode<RefreshTokenMutation, RefreshTokenMutationVariables>;
export const GetBookMarkListDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getBookMarkList"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getBookMarkList"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]}}]} as unknown as DocumentNode<GetBookMarkListQuery, GetBookMarkListQueryVariables>;
export const GetBookMarkedListDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getBookMarkedList"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getBookMarkedList"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]}}]} as unknown as DocumentNode<GetBookMarkedListQuery, GetBookMarkedListQueryVariables>;
export const GetIsBookMarkedDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getIsBookMarked"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getIsBookMarked"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}]}]}}]} as unknown as DocumentNode<GetIsBookMarkedQuery, GetIsBookMarkedQueryVariables>;
export const AddBookMarkDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"addBookMark"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"addBookMark"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}]}]}}]} as unknown as DocumentNode<AddBookMarkMutation, AddBookMarkMutationVariables>;
export const RemoveBookMarkDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"removeBookMark"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"removeBookMark"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}]}]}}]} as unknown as DocumentNode<RemoveBookMarkMutation, RemoveBookMarkMutationVariables>;
export const CategoryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"category"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"category"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"imageUrl"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"readonly"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]} as unknown as DocumentNode<CategoryQuery, CategoryQueryVariables>;
export const CategoriesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"categories"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"categories"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]}}]}}]}}]}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<CategoriesQuery, CategoriesQueryVariables>;
export const HistoriesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"histories"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"histories"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"now"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"parent"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"versionNo"}},{"kind":"Field","name":{"kind":"Name","value":"readonly"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"histories"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"parent"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"versionNo"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]}}]} as unknown as DocumentNode<HistoriesQuery, HistoriesQueryVariables>;
export const RandomRecommendListDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"RandomRecommendList"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"limit"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"randomRecommendList"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"Variable","name":{"kind":"Name","value":"limit"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"categoryName"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]} as unknown as DocumentNode<RandomRecommendListQuery, RandomRecommendListQueryVariables>;
export const GetRecommendLiquorListDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getRecommendLiquorList"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getRecommendLiquorList"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}},{"kind":"Field","name":{"kind":"Name","value":"liquor"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"categoryName"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}}]}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}}]}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]} as unknown as DocumentNode<GetRecommendLiquorListQuery, GetRecommendLiquorListQueryVariables>;
export const PostBoardDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"postBoard"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"BoardInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"postBoard"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}]}]}}]} as unknown as DocumentNode<PostBoardMutation, PostBoardMutationVariables>;
export const GetMyBoardDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getMyBoard"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getMyBoard"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"liquorId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"text"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}}]}}]}}]} as unknown as DocumentNode<GetMyBoardQuery, GetMyBoardQueryVariables>;
export const BoardDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"board"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"liquorId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"board"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"liquorId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"liquorId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userId"}},{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"text"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]} as unknown as DocumentNode<BoardQuery, BoardQueryVariables>;
export const GetFlavorMapDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getFlavorMap"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"liquorId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getFlavorMap"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"liquorId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"liquorId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"xNames"}},{"kind":"Field","name":{"kind":"Name","value":"yNames"}},{"kind":"Field","name":{"kind":"Name","value":"userFullAmount"}},{"kind":"Field","name":{"kind":"Name","value":"guestFullAmount"}},{"kind":"Field","name":{"kind":"Name","value":"mapData"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"x"}},{"kind":"Field","name":{"kind":"Name","value":"y"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"userAmount"}},{"kind":"Field","name":{"kind":"Name","value":"guestAmount"}}]}}]}}]}}]} as unknown as DocumentNode<GetFlavorMapQuery, GetFlavorMapQueryVariables>;
export const PostFlavorDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"postFlavor"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"PostFlavorMap"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"postFlavor"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}]}]}}]} as unknown as DocumentNode<PostFlavorMutation, PostFlavorMutationVariables>;
export const GetVotedDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getVoted"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"liquorId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getVoted"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"liquorId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"liquorId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"x"}},{"kind":"Field","name":{"kind":"Name","value":"y"}}]}}]}}]} as unknown as DocumentNode<GetVotedQuery, GetVotedQueryVariables>;
export const LiquorDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"liquor"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"liquor"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"categoryName"}},{"kind":"Field","name":{"kind":"Name","value":"categoryTrail"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"youtube"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"imageUrl"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"versionNo"}}]}}]}}]} as unknown as DocumentNode<LiquorQuery, LiquorQueryVariables>;
export const ListFromCategoryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"listFromCategory"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"listFromCategory"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"categoryId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"categoryName"}},{"kind":"Field","name":{"kind":"Name","value":"categoryDescription"}},{"kind":"Field","name":{"kind":"Name","value":"liquors"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"categoryName"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]}}]} as unknown as DocumentNode<ListFromCategoryQuery, ListFromCategoryQueryVariables>;
export const LiquorHistoriesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"liquorHistories"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"liquorHistories"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"now"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"versionNo"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"histories"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"versionNo"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]}}]} as unknown as DocumentNode<LiquorHistoriesQuery, LiquorHistoriesQueryVariables>;
export const GetTagsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getTags"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"liquorId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getTags"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"liquorId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"liquorId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"text"}}]}}]}}]} as unknown as DocumentNode<GetTagsQuery, GetTagsQueryVariables>;
export const PostTagDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"postTag"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"TagInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"postTag"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"text"}}]}}]}}]} as unknown as DocumentNode<PostTagMutation, PostTagMutationVariables>;
export const DeleteTagDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"deleteTag"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"deleteTag"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}]}]}}]} as unknown as DocumentNode<DeleteTagMutation, DeleteTagMutationVariables>;
export const GetUserByIdDetailDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getUserByIdDetail"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getUserByIdDetail"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"evaluateList"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"recentComments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"categoryName"}},{"kind":"Field","name":{"kind":"Name","value":"liquorId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"rate5Liquors"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"categoryName"}},{"kind":"Field","name":{"kind":"Name","value":"liquorId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"rate4Liquors"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"categoryName"}},{"kind":"Field","name":{"kind":"Name","value":"liquorId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"rate3Liquors"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"categoryName"}},{"kind":"Field","name":{"kind":"Name","value":"liquorId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"rate2Liquors"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"categoryName"}},{"kind":"Field","name":{"kind":"Name","value":"liquorId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"rate1Liquors"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"categoryName"}},{"kind":"Field","name":{"kind":"Name","value":"liquorId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"noRateLiquors"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"categoryId"}},{"kind":"Field","name":{"kind":"Name","value":"categoryName"}},{"kind":"Field","name":{"kind":"Name","value":"liquorId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}},{"kind":"Field","name":{"kind":"Name","value":"imageBase64"}}]}}]}}]}}]} as unknown as DocumentNode<GetUserByIdDetailQuery, GetUserByIdDetailQueryVariables>;
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
  __typename?: 'AffiliateData';
  items?: Maybe<Array<AffiliateItem>>;
  lowestPrice?: Maybe<Scalars['Int']['output']>;
};

export type AffiliateItem = {
  __typename?: 'AffiliateItem';
  URL: Scalars['String']['output'];
  imageURL?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
  price?: Maybe<Scalars['Int']['output']>;
};

export type AuthPayload = {
  __typename?: 'AuthPayload';
  token: Scalars['String']['output'];
  user: User;
};

export type BoardInput = {
  liquorID: Scalars['String']['input'];
  rate?: InputMaybe<Scalars['Int']['input']>;
  text: Scalars['String']['input'];
};

export type BoardPost = {
  __typename?: 'BoardPost';
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
  __typename?: 'BookMarkListUser';
  createdAt: Scalars['DateTime']['output'];
  imageBase64?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
  userId: Scalars['ID']['output'];
};

export type Category = {
  __typename?: 'Category';
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
  __typename?: 'CategoryHistory';
  histories?: Maybe<Array<Maybe<Category>>>;
  now: Category;
};

export type CategoryTrail = {
  __typename?: 'CategoryTrail';
  id: Scalars['Int']['output'];
  name: Scalars['String']['output'];
};

export type FlavorCellData = {
  __typename?: 'FlavorCellData';
  guestAmount: Scalars['Int']['output'];
  rate: Scalars['Float']['output'];
  userAmount: Scalars['Int']['output'];
  x: Scalars['Coordinate']['output'];
  y: Scalars['Coordinate']['output'];
};

export type FlavorMapData = {
  __typename?: 'FlavorMapData';
  categoryId: Scalars['Int']['output'];
  guestFullAmount: Scalars['Int']['output'];
  mapData: Array<FlavorCellData>;
  userFullAmount: Scalars['Int']['output'];
  xNames: Array<Scalars['String']['output']>;
  yNames: Array<Scalars['String']['output']>;
};

export type Liquor = {
  __typename?: 'Liquor';
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
  __typename?: 'LiquorHistory';
  histories?: Maybe<Array<Maybe<Liquor>>>;
  now: Liquor;
};

export type ListFromCategory = {
  __typename?: 'ListFromCategory';
  categoryDescription?: Maybe<Scalars['String']['output']>;
  categoryName: Scalars['String']['output'];
  liquors: Array<Maybe<Liquor>>;
};

export type LoginInput = {
  email: Scalars['String']['input'];
  password: Scalars['String']['input'];
};

export type Mutation = {
  __typename?: 'Mutation';
  addBookMark: Scalars['Boolean']['output'];
  deleteTag: Scalars['Boolean']['output'];
  login: AuthPayload;
  postBoard: Scalars['Boolean']['output'];
  postFlavor: Scalars['Boolean']['output'];
  postTag: Tag;
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
  __typename?: 'Query';
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
  __typename?: 'Recommend';
  comment: Scalars['String']['output'];
  liquor: RecommendLiquor;
  rate: Scalars['Int']['output'];
  updatedAt: Scalars['DateTime']['output'];
  user: RecommendUser;
};

export type RecommendLiquor = {
  __typename?: 'RecommendLiquor';
  categoryId: Scalars['Int']['output'];
  categoryName: Scalars['String']['output'];
  description: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  imageBase64?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
};

export type RecommendUser = {
  __typename?: 'RecommendUser';
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
  __typename?: 'Tag';
  id: Scalars['ID']['output'];
  text: Scalars['String']['output'];
};

export type TagInput = {
  liquorId: Scalars['ID']['input'];
  text: Scalars['String']['input'];
};

export type User = {
  __typename?: 'User';
  email: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  imageBase64?: Maybe<Scalars['String']['output']>;
  name: Scalars['String']['output'];
  profile?: Maybe<Scalars['String']['output']>;
  roles: Array<Maybe<Scalars['String']['output']>>;
};

export type UserEvaluateList = {
  __typename?: 'UserEvaluateList';
  noRateLiquors?: Maybe<Array<UserLiquor>>;
  rate1Liquors?: Maybe<Array<UserLiquor>>;
  rate2Liquors?: Maybe<Array<UserLiquor>>;
  rate3Liquors?: Maybe<Array<UserLiquor>>;
  rate4Liquors?: Maybe<Array<UserLiquor>>;
  rate5Liquors?: Maybe<Array<UserLiquor>>;
  recentComments?: Maybe<Array<UserLiquor>>;
};

export type UserLiquor = {
  __typename?: 'UserLiquor';
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
  __typename?: 'UserPageData';
  evaluateList: UserEvaluateList;
  user: User;
};

export type VotedData = {
  __typename?: 'VotedData';
  categoryId: Scalars['Int']['output'];
  liquorId: Scalars['ID']['output'];
  updatedAt: Scalars['DateTime']['output'];
  userId: Scalars['ID']['output'];
  x: Scalars['Coordinate']['output'];
  y: Scalars['Coordinate']['output'];
};

export type AdminCheckQueryVariables = Exact<{ [key: string]: never; }>;


export type AdminCheckQuery = { __typename?: 'Query', checkAdmin: boolean };


export const AdminCheckDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"adminCheck"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"checkAdmin"}}]}}]} as unknown as DocumentNode<AdminCheckQuery, AdminCheckQueryVariables>;
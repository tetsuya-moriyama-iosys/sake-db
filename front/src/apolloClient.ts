import { ApolloClient, InMemoryCache } from '@apollo/client/core';

const client = new ApolloClient({
  uri: import.meta.env.VITE_API_URL + '/query', // GraphQLサーバーのURL
  cache: new InMemoryCache(),
});

export default client;

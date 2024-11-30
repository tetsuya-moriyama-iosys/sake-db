import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  schema: '../backend/graph/schema/*.graphqls',
  documents: './src/graphQL/Admin/*.ts', //Admin配下でお試し適用している
  generates: {
    './src/graphQL/auto-generated.ts': {
      plugins: ['typescript', 'typescript-operations', 'typed-document-node'],
    },
  },
};
export default config;

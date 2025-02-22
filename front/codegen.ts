import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  schema: '../backend/graph/schema/*.graphqls',
  documents: './src/graphQL/**/*.ts',
  generates: {
    './src/graphQL/auto-generated.ts': {
      plugins: ['typescript', 'typescript-operations', 'typed-document-node'],
      config: {
        nonOptionalTypename: true,
      },
    },
  },
};
export default config;

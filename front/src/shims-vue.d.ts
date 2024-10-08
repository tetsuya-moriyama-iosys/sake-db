// src/shims-vue.d.ts
declare module '*.vue' {
  import { DefineComponent } from 'vue';
  const component: DefineComponent<
    NonNullable<unknown>,
    NonNullable<unknown>,
    never
  >;
  export default component;
}

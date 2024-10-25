/**
 * ユーザーページ共通関数(composableってほど汎用的でないもの)
 */

import type { DocumentNode } from 'graphql/index';
import { onMounted, type Ref, ref } from 'vue';

import useQuery from '@/funcs/composable/useQuery';
import {
  type AuthUserFull,
  GET_MY_USERDATA_FULL,
  type GetUserdataResponse,
} from '@/graphQL/Auth/auth';

function core(query: DocumentNode) {
  const user: Ref<AuthUserFull | null | undefined> = ref<AuthUserFull | null>();
  const { fetch } = useQuery<GetUserdataResponse>(query, {
    isAuth: true,
  });

  const fetchUser = async (): Promise<void> => {
    const response: GetUserdataResponse = await fetch();
    user.value = response.getUser;
  };

  onMounted(fetchUser);

  return { user };
}

export function getUserDetail() {
  const { user } = core(GET_MY_USERDATA_FULL);

  return { user };
}

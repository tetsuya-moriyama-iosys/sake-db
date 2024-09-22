/**
 * ログインユーザー情報取得
 */

import { ref, onMounted, type Ref } from 'vue';
import useQuery from '@/funcs/composable/useQuery';
import type { DocumentNode } from 'graphql/index';
import {
  type AuthUser,
  GET_MY_USERDATA_FULL,
  type GetUserdataResponse,
} from '@/graphQL/Auth/auth';

function core(query: DocumentNode) {
  const user: Ref<AuthUser | null | undefined> = ref<AuthUser | null>();
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

export function useUser() {
  const { user } = core(GET_MY_USERDATA_FULL);

  return { user };
}

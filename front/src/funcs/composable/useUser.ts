/**
 * ログインユーザー情報取得
 */

import { ref, onMounted, type Ref } from 'vue';
import useQuery from '@/funcs/composable/useQuery';
import {
  type AuthUser,
  GET_MY_USERDATA,
  GET_MY_USERDATA_FULL,
  type GetUserdataResponse,
  type UserFullData,
} from '@/graphQL/User/user';
import type { DocumentNode } from 'graphql/index';

function core<T extends AuthUser>(query: DocumentNode) {
  const user: Ref<T | null | undefined> = ref<T | null>();
  const { fetch } = useQuery<GetUserdataResponse<T>>(query, {
    isAuth: true,
  });

  const fetchUser = async (): Promise<void> => {
    const response: GetUserdataResponse<T> = await fetch();
    user.value = response.getUser;
  };

  onMounted(fetchUser);

  return { user };
}

export function useUser() {
  const { user } = core<AuthUser>(GET_MY_USERDATA);

  return { user };
}

export function useUserFullData() {
  const { user } = core<UserFullData>(GET_MY_USERDATA_FULL);

  return { user };
}

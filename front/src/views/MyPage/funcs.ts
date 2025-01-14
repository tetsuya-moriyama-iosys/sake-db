/**
 * ユーザーページ共通関数(composableってほど汎用的でないもの)
 */
import { onMounted, type Ref, ref } from 'vue';

import useQuery from '@/funcs/composable/useQuery/useQuery';
import { stripTypeName } from '@/funcs/util/stripTypeName';
import { GET_MY_USERDATA_FULL } from '@/graphQL/Auth/auth';
import { type AuthUserFull, getAuthUserFull } from '@/graphQL/Auth/types';
import type { GetMyDataFullQuery } from '@/graphQL/auto-generated';

export function useUserDetail() {
  const user: Ref<AuthUserFull | null | undefined> = ref<AuthUserFull | null>();
  const { fetch } = useQuery<GetMyDataFullQuery>(GET_MY_USERDATA_FULL, {
    isAuth: true,
  });

  const fetchUser = async (): Promise<void> => {
    const response = await fetch();
    user.value = getAuthUserFull(stripTypeName(response.getMyData));
  };

  onMounted(fetchUser);

  return { user };
}

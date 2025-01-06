import type { Role } from '@/graphQL/Auth/auth';
import type {
  LoginMutation,
  RegisterUserMutation,
} from '@/graphQL/auto-generated';

// ユニオン型が自動生成できないため、手動で定義
export type AuthPayloadForUI = {
  accessToken: string;
  user: {
    id: string;
    name: string;
    imageBase64: string | null;
    roles: Role[];
  };
};

export function getAuthPayloadForUI(
  payload: RegisterUserMutation['registerUser'] | LoginMutation['login'],
): AuthPayloadForUI {
  const { accessToken, user } = payload;
  const { __typename, ...rest } = user;
  return {
    accessToken,
    user: {
      id: rest.id,
      name: rest.name,
      imageBase64: rest.imageBase64 ?? null,
      roles: (rest.roles ?? []) as Role[], //一旦横着したタイプガード
    },
  };
}

import type { GetMyDataFullQuery } from '@/graphQL/auto-generated';

export const Roles = {
  Admin: 'admin',
} as const;
export type Role = (typeof Roles)[keyof typeof Roles];

//認証済みユーザー情報(要はパスワード以外のデータ)
export interface AuthUser {
  id: string;
  name: string;
  imageBase64: string | null; //アイコン表示に必要
  roles: Role[];
}
export interface AuthUserFull extends AuthUser {
  email: string;
  profile: string;
}

export function getAuthUserFull(
  user: Omit<GetMyDataFullQuery['getMyData'], '__typename'>,
): AuthUserFull {
  return {
    ...user,
    imageBase64: user.imageBase64 ?? null,
    profile: user.profile ?? '',
    roles: (user.roles ?? []) as Role[],
  };
}

<template>
  <table>
    <tr>
      <th>ユーザー名</th>
      <th v-if="props.isShowCreatedAt">追加日</th>
      <th v-if="$slots.default">
        <slot name="slotHeading"></slot>
      </th>
    </tr>
    <tr v-for="user in props.userList" :key="user.userId">
      <td>
        <router-link :to="{ name: 'UserPage', params: { id: user.userId } }">{{
          user.name
        }}</router-link>
      </td>
      <td v-if="props.isShowCreatedAt">
        {{ format(date(user.createdAt), 'yyyy/MM/dd') }}
      </td>
      <td v-if="$slots.default">
        <!-- slotを使用して親コンポーネントがカスタマイズできるようにする -->
        <slot :user="user" />
      </td>
    </tr>
  </table>
</template>

<script setup lang="ts">
import { type Bookmark } from '@/graphQL/Bookmark/bookmark';
import date from '@/funcs/util/date';
import { format } from 'date-fns';

const props = defineProps<{
  userList: Bookmark[];
  isShowCreatedAt?: boolean;
}>();
</script>

<style scoped></style>

<script setup lang="ts">
import { User } from 'lucide-vue-next'

import NavMain from '@/components/layout/NavMain.vue'

import {
  Sidebar,
  SidebarHeader,
  type SidebarProps,
  SidebarRail,
} from '@/components/ui/sidebar'

const props = defineProps<SidebarProps>()

// This is sample data.
const data = {
  navMain: [
    { title: 'Dashboard', url: '/dashboard', icon: 'i-heroicons-home' },
    { title: 'Calories', url: '/dashboard/calories', icon: 'i-heroicons-fire' },
    { title: 'Workouts', url: '/dashboard/workouts', icon: 'i-heroicons-muscle' },
    { title: 'Profile', url: '/dashboard/profile', icon: User },
  ],
}
</script>

<template>
  <Sidebar class="border-r-0" v-bind="props">
    <SidebarHeader>
      <Logo />
      <NavMain :items="data.navMain" />
    </SidebarHeader>
    <SidebarRail />
  </Sidebar>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const isCollapsed = ref(false)

const navigation = [
  { name: 'Dashboard', href: '/dashboard', icon: 'i-heroicons-home' },
  { name: 'Calories', href: '/dashboard/calories', icon: 'i-heroicons-fire' },
  { name: 'Workouts', href: '/dashboard/workouts', icon: 'i-heroicons-muscle' },
  { name: 'Profile', href: '/dashboard/profile', icon: 'i-heroicons-user' },
]
</script>

<template>
  <div class="min-h-screen bg-background">
    <!-- Sidebar -->
    <aside
      class="fixed inset-y-0 left-0 z-50 w-64 bg-card border-r border-border transition-transform duration-200"
      :class="{ '-translate-x-full': isCollapsed }"
    >
      <div class="flex h-16 items-center justify-between px-4 border-b border-border">
        <h1 class="text-xl font-bold">
          Rebound
        </h1>
        <button
          class="p-2 rounded-md hover:bg-accent"
          @click="isCollapsed = !isCollapsed"
        >
          <i class="i-heroicons-bars-3" />
        </button>
      </div>

      <nav class="p-4 space-y-1">
        <RouterLink
          v-for="item in navigation"
          :key="item.href"
          :to="item.href"
          class="flex items-center px-4 py-2 text-sm font-medium rounded-md hover:bg-accent"
          :class="{ 'bg-accent': $route.path === item.href }"
        >
          <i :class="item.icon" class="mr-3 h-5 w-5" />
          {{ item.name }}
        </RouterLink>
      </nav>
    </aside>

    <!-- Main content -->
    <main
      class="transition-all duration-200"
      :class="{ 'ml-64': !isCollapsed, 'ml-0': isCollapsed }"
    >
      <div class="p-8">
        <RouterView />
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
const stats = [
  { name: 'Today\'s Calories', value: '1,850', change: '+12%', icon: 'i-heroicons-fire' },
  { name: 'Workouts This Week', value: '3', change: '+1', icon: 'i-heroicons-muscle' },
  { name: 'Water Intake', value: '2.5L', change: '-0.5L', icon: 'i-heroicons-water' },
  { name: 'Steps', value: '8,432', change: '+1,234', icon: 'i-heroicons-shoe' },
]

const recentActivities = [
  { type: 'workout', name: 'Morning Run', time: '2 hours ago', calories: '320' },
  { type: 'meal', name: 'Lunch', time: '1 hour ago', calories: '450' },
  { type: 'water', name: 'Water Intake', time: '30 mins ago', amount: '500ml' },
]
</script>

<template>
  <div class="space-y-8">
    <div>
      <h1 class="text-3xl font-bold">
        Dashboard
      </h1>
      <p class="text-muted-foreground">
        Welcome back! Here's your fitness summary.
      </p>
    </div>

    <!-- Stats Grid -->
    <div class="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
      <div
        v-for="stat in stats"
        :key="stat.name"
        class="p-6 bg-card rounded-lg border border-border"
      >
        <div class="flex items-center justify-between">
          <div>
            <p class="text-sm font-medium text-muted-foreground">
              {{ stat.name }}
            </p>
            <p class="text-2xl font-bold">
              {{ stat.value }}
            </p>
          </div>
          <div class="p-2 bg-accent rounded-full">
            <i :class="stat.icon" class="h-6 w-6" />
          </div>
        </div>
        <p
          class="mt-2 text-sm"
          :class="stat.change.startsWith('+') ? 'text-green-500' : 'text-red-500'"
        >
          {{ stat.change }}
        </p>
      </div>
    </div>

    <!-- Recent Activity -->
    <div class="bg-card rounded-lg border border-border p-6">
      <h2 class="text-xl font-semibold mb-4">
        Recent Activity
      </h2>
      <div class="space-y-4">
        <div
          v-for="activity in recentActivities"
          :key="activity.name"
          class="flex items-center justify-between p-4 bg-background rounded-lg"
        >
          <div class="flex items-center space-x-4">
            <div
              class="p-2 rounded-full"
              :class="{
                'bg-blue-100': activity.type === 'workout',
                'bg-green-100': activity.type === 'meal',
                'bg-cyan-100': activity.type === 'water',
              }"
            >
              <i
                :class="{
                  'i-heroicons-muscle': activity.type === 'workout',
                  'i-heroicons-fire': activity.type === 'meal',
                  'i-heroicons-water': activity.type === 'water',
                }"
                class="h-5 w-5"
              />
            </div>
            <div>
              <p class="font-medium">
                {{ activity.name }}
              </p>
              <p class="text-sm text-muted-foreground">
                {{ activity.time }}
              </p>
            </div>
          </div>
          <div class="text-sm font-medium">
            {{ activity.calories || activity.amount }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<route lang="yaml">
meta:
  layout: dashboard
</route>

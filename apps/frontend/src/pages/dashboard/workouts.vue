<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import { computed, ref } from 'vue'

const workouts = ref([
  {
    name: 'Morning Run',
    type: 'Cardio',
    duration: 45,
    calories: 320,
    date: '2024-03-20',
    exercises: ['Running', 'Stretching'],
  },
  {
    name: 'Upper Body',
    type: 'Strength',
    duration: 60,
    calories: 450,
    date: '2024-03-19',
    exercises: ['Bench Press', 'Pull-ups', 'Shoulder Press'],
  },
])

const newWorkout = ref({
  name: '',
  type: '',
  duration: '',
  calories: '',
  date: '',
  exercises: '',
})

const workoutTypes = ['Cardio', 'Strength', 'Flexibility', 'HIIT', 'Yoga']

function addWorkout() {
  if (!newWorkout.value.name || !newWorkout.value.type)
    return

  workouts.value.push({
    name: newWorkout.value.name,
    type: newWorkout.value.type,
    duration: Number.parseInt(newWorkout.value.duration),
    calories: Number.parseInt(newWorkout.value.calories),
    date: newWorkout.value.date,
    exercises: newWorkout.value.exercises.split(',').map(ex => ex.trim()),
  })

  // Reset form
  newWorkout.value = {
    name: '',
    type: '',
    duration: '',
    calories: '',
    date: '',
    exercises: '',
  }
}

const totalWorkoutsThisWeek = computed(() => {
  const oneWeekAgo = new Date()
  oneWeekAgo.setDate(oneWeekAgo.getDate() - 7)
  return workouts.value.filter(w => new Date(w.date) >= oneWeekAgo).length
})

const totalCaloriesBurned = computed(() => {
  return workouts.value.reduce((sum, w) => sum + w.calories, 0)
})
</script>

<template>
  <div class="space-y-8">
    <div>
      <h1 class="text-3xl font-bold">
        Workout Tracking
      </h1>
      <p class="text-muted-foreground">
        Track your workouts and fitness progress.
      </p>
    </div>

    <!-- Workout Summary -->
    <div class="grid gap-4 md:grid-cols-2">
      <Card>
        <CardContent class="p-6">
          <p class="text-sm font-medium text-muted-foreground">
            Workouts This Week
          </p>
          <p class="text-2xl font-bold">
            {{ totalWorkoutsThisWeek }}
          </p>
        </CardContent>
      </Card>
      <Card>
        <CardContent class="p-6">
          <p class="text-sm font-medium text-muted-foreground">
            Total Calories Burned
          </p>
          <p class="text-2xl font-bold">
            {{ totalCaloriesBurned }} cal
          </p>
        </CardContent>
      </Card>
    </div>

    <!-- Add Workout Form -->
    <Card>
      <CardHeader>
        <h2 class="text-xl font-semibold">
          Add Workout
        </h2>
      </CardHeader>
      <CardContent>
        <form class="space-y-4" @submit.prevent="addWorkout">
          <div class="grid gap-4 md:grid-cols-2">
            <div class="space-y-2">
              <Label>Workout Name</Label>
              <Input
                v-model="newWorkout.name"
                placeholder="e.g., Morning Run"
              />
            </div>
            <div class="space-y-2">
              <Label>Type</Label>
              <Select v-model="newWorkout.type">
                <SelectTrigger>
                  <SelectValue placeholder="Select type" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem
                    v-for="type in workoutTypes"
                    :key="type"
                    :value="type"
                  >
                    {{ type }}
                  </SelectItem>
                </SelectContent>
              </Select>
            </div>
            <div class="space-y-2">
              <Label>Duration (minutes)</Label>
              <Input
                v-model="newWorkout.duration"
                type="number"
                placeholder="e.g., 45"
              />
            </div>
            <div class="space-y-2">
              <Label>Calories Burned</Label>
              <Input
                v-model="newWorkout.calories"
                type="number"
                placeholder="e.g., 320"
              />
            </div>
            <div class="space-y-2">
              <Label>Date</Label>
              <Input
                v-model="newWorkout.date"
                type="date"
              />
            </div>
            <div class="space-y-2">
              <Label>Exercises (comma-separated)</Label>
              <Input
                v-model="newWorkout.exercises"
                placeholder="e.g., Running, Stretching"
              />
            </div>
          </div>
          <Button type="submit">
            Add Workout
          </Button>
        </form>
      </CardContent>
    </Card>

    <!-- Workouts List -->
    <Card>
      <CardHeader>
        <h2 class="text-xl font-semibold">
          Recent Workouts
        </h2>
      </CardHeader>
      <CardContent>
        <div class="space-y-4">
          <div
            v-for="workout in workouts"
            :key="workout.name"
            class="p-4 bg-muted rounded-lg"
          >
            <div class="flex items-center justify-between">
              <div>
                <h3 class="font-medium">
                  {{ workout.name }}
                </h3>
                <p class="text-sm text-muted-foreground">
                  {{ workout.date }}
                </p>
                <div class="mt-2 flex flex-wrap gap-2">
                  <span
                    v-for="exercise in workout.exercises"
                    :key="exercise"
                    class="px-2 py-1 text-xs bg-accent rounded-full"
                  >
                    {{ exercise }}
                  </span>
                </div>
              </div>
              <div class="text-right">
                <div class="text-lg font-semibold">
                  {{ workout.calories }} cal
                </div>
                <div class="text-sm text-muted-foreground">
                  {{ workout.duration }} min
                </div>
                <div class="text-sm text-muted-foreground">
                  {{ workout.type }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  </div>
</template>

<route lang="yaml">
meta:
  layout: dashboard
</route>

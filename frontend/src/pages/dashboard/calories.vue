<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { computed, ref } from 'vue'
import { Card, CardContent, CardHeader } from '../../components/ui/card'

const meals = ref([
  { name: 'Breakfast', calories: 450, time: '08:00', foods: ['Oatmeal', 'Banana', 'Coffee'] },
  { name: 'Lunch', calories: 650, time: '12:30', foods: ['Chicken Salad', 'Whole Grain Bread'] },
  { name: 'Snack', calories: 200, time: '15:00', foods: ['Apple', 'Almonds'] },
])

const newMeal = ref({
  name: '',
  calories: '',
  time: '',
  foods: '',
})

function addMeal() {
  if (!newMeal.value.name || !newMeal.value.calories)
    return

  meals.value.push({
    name: newMeal.value.name,
    calories: Number.parseInt(newMeal.value.calories),
    time: newMeal.value.time,
    foods: newMeal.value.foods.split(',').map(food => food.trim()),
  })

  // Reset form
  newMeal.value = {
    name: '',
    calories: '',
    time: '',
    foods: '',
  }
}

const totalCalories = computed(() => {
  return meals.value.reduce((sum, meal) => sum + meal.calories, 0)
})

const calorieGoal = 2000
const remainingCalories = computed(() => calorieGoal - totalCalories.value)
</script>

<template>
  <div class="space-y-8">
    <div>
      <h1 class="text-3xl font-bold">
        Calorie Tracking
      </h1>
      <p class="text-muted-foreground">
        Track your daily food intake and calories.
      </p>
    </div>

    <!-- Calorie Summary -->
    <div class="grid gap-4 md:grid-cols-3">
      <Card>
        <CardContent class="p-6">
          <p class="text-sm font-medium text-muted-foreground">
            Daily Goal
          </p>
          <p class="text-2xl font-bold">
            {{ calorieGoal }} cal
          </p>
        </CardContent>
      </Card>
      <Card>
        <CardContent class="p-6">
          <p class="text-sm font-medium text-muted-foreground">
            Consumed
          </p>
          <p class="text-2xl font-bold">
            {{ totalCalories }} cal
          </p>
        </CardContent>
      </Card>
      <Card>
        <CardContent class="p-6">
          <p class="text-sm font-medium text-muted-foreground">
            Remaining
          </p>
          <p class="text-2xl font-bold" :class="remainingCalories < 0 ? 'text-destructive' : 'text-green-500'">
            {{ remainingCalories }} cal
          </p>
        </CardContent>
      </Card>
    </div>

    <!-- Add Meal Form -->
    <Card>
      <CardHeader>
        <h2 class="text-xl font-semibold">
          Add Meal
        </h2>
      </CardHeader>
      <CardContent>
        <form class="space-y-4" @submit.prevent="addMeal">
          <div class="grid gap-4 md:grid-cols-2">
            <div class="space-y-2">
              <Label>Meal Name</Label>
              <Input
                v-model="newMeal.name"
                placeholder="e.g., Breakfast, Lunch"
              />
            </div>
            <div class="space-y-2">
              <Label>Calories</Label>
              <Input
                v-model="newMeal.calories"
                type="number"
                placeholder="e.g., 450"
              />
            </div>
            <div class="space-y-2">
              <Label>Time</Label>
              <Input
                v-model="newMeal.time"
                type="time"
              />
            </div>
            <div class="space-y-2">
              <Label>Foods (comma-separated)</Label>
              <Input
                v-model="newMeal.foods"
                placeholder="e.g., Oatmeal, Banana, Coffee"
              />
            </div>
          </div>
          <Button type="submit">
            Add Meal
          </Button>
        </form>
      </CardContent>
    </Card>

    <!-- Meals List -->
    <Card>
      <CardHeader>
        <h2 class="text-xl font-semibold">
          Today's Meals
        </h2>
      </CardHeader>
      <CardContent>
        <div class="space-y-4">
          <div
            v-for="meal in meals"
            :key="meal.name"
            class="p-4 bg-muted rounded-lg"
          >
            <div class="flex items-center justify-between">
              <div>
                <h3 class="font-medium">
                  {{ meal.name }}
                </h3>
                <p class="text-sm text-muted-foreground">
                  {{ meal.time }}
                </p>
                <div class="mt-2 flex flex-wrap gap-2">
                  <span
                    v-for="food in meal.foods"
                    :key="food"
                    class="px-2 py-1 text-xs bg-accent rounded-full"
                  >
                    {{ food }}
                  </span>
                </div>
              </div>
              <div class="text-lg font-semibold">
                {{ meal.calories }} cal
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

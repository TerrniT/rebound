<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { Calendar } from '@/components/ui/calendar'
import { Card, CardContent, CardHeader } from '@/components/ui/card'
import { Progress } from '@/components/ui/progress'
import { Plus } from 'lucide-vue-next'
import { computed, ref } from 'vue'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

// Calendar
const date = ref(new Date())

// Summary Stats
const dailyStats = ref({
  caloriesGoal: 2000,
  caloriesEaten: 1450,
  caloriesBurned: 320,
  macros: {
    carbs: { current: 180, goal: 250 },
    protein: { current: 95, goal: 120 },
    fat: { current: 45, goal: 60 },
  },
})

const remainingCalories = computed(() =>
  dailyStats.value.caloriesGoal - dailyStats.value.caloriesEaten + dailyStats.value.caloriesBurned,
)

const caloriesProgress = computed(() =>
  (dailyStats.value.caloriesEaten / dailyStats.value.caloriesGoal) * 100,
)

// Meals
const meals = ref([
  {
    name: 'Breakfast',
    calories: 450,
    time: '08:00',
    foods: ['Oatmeal', 'Banana', 'Coffee'],
  },
  {
    name: 'Lunch',
    calories: 650,
    time: '12:30',
    foods: ['Chicken Salad', 'Whole Grain Bread'],
  },
  {
    name: 'Dinner',
    calories: 350,
    time: '18:00',
    foods: ['Grilled Salmon', 'Quinoa', 'Vegetables'],
  },
  {
    name: 'Snacks',
    calories: 200,
    time: '15:00',
    foods: ['Apple', 'Almonds'],
  },
])

// Add Food Dialog
const showAddFoodDialog = ref(false)
const selectedMeal = ref('')
const newFood = ref({
  name: '',
  calories: '',
  time: '',
})

const addFood = () => {
  if (!newFood.value.name || !newFood.value.calories) return

  const meal = meals.value.find(m => m.name === selectedMeal.value)
  if (meal) {
    meal.foods.push(newFood.value.name)
    meal.calories += parseInt(newFood.value.calories)
    dailyStats.value.caloriesEaten += parseInt(newFood.value.calories)
  }

  // Reset form
  newFood.value = {
    name: '',
    calories: '',
    time: '',
  }
  showAddFoodDialog.value = false
}

// Water Tracker
const waterStats = ref({
  current: 1.5, // liters
  goal: 2.5, // liters
})

const waterProgress = computed(() =>
  (waterStats.value.current / waterStats.value.goal) * 100,
)

// Add Water Dialog
const showAddWaterDialog = ref(false)
const waterAmount = ref('0.25')

const addWater = () => {
  const amount = parseFloat(waterAmount.value)
  waterStats.value.current = Math.min(waterStats.value.current + amount, waterStats.value.goal)
  showAddWaterDialog.value = false
}
</script>

<template>
  <div class="space-y-8">
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold">
        Dashboard
      </h1>
      <p class="text-muted-foreground">
        {{ date.toLocaleDateString('en-US', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' }) }}
      </p>
    </div>

    <!-- Calendar Section -->
    <Card>
      <CardHeader>
        <h2 class="text-xl font-semibold">
          Calendar
        </h2>
      </CardHeader>
      <CardContent>
        <Calendar
          v-model="date"
          class="rounded-md border"
        />
      </CardContent>
    </Card>

    <!-- Summary Section -->
    <div class="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
      <Card>
        <CardContent class="p-6">
          <div class="space-y-2">
            <p class="text-sm font-medium text-muted-foreground">
              Calories Eaten
            </p>
            <p class="text-2xl font-bold">
              {{ dailyStats.caloriesEaten }} / {{ dailyStats.caloriesGoal }}
            </p>
            <Progress :value="caloriesProgress" class="h-2" />
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardContent class="p-6">
          <div class="space-y-2">
            <p class="text-sm font-medium text-muted-foreground">
              Calories Burned
            </p>
            <p class="text-2xl font-bold">
              {{ dailyStats.caloriesBurned }}
            </p>
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardContent class="p-6">
          <div class="space-y-2">
            <p class="text-sm font-medium text-muted-foreground">
              Remaining
            </p>
            <p class="text-2xl font-bold" :class="remainingCalories < 0 ? 'text-destructive' : 'text-green-500'">
              {{ remainingCalories }}
            </p>
          </div>
        </CardContent>
      </Card>

      <Card>
        <CardContent class="p-6">
          <div class="space-y-2">
            <p class="text-sm font-medium text-muted-foreground">
              Macros
            </p>
            <div class="space-y-1">
              <div class="flex justify-between text-sm">
                <span>Carbs</span>
                <span>{{ dailyStats.macros.carbs.current }}g / {{ dailyStats.macros.carbs.goal }}g</span>
              </div>
              <Progress :value="(dailyStats.macros.carbs.current / dailyStats.macros.carbs.goal) * 100" class="h-1" />
            </div>
            <div class="space-y-1">
              <div class="flex justify-between text-sm">
                <span>Protein</span>
                <span>{{ dailyStats.macros.protein.current }}g / {{ dailyStats.macros.protein.goal }}g</span>
              </div>
              <Progress :value="(dailyStats.macros.protein.current / dailyStats.macros.protein.goal) * 100" class="h-1" />
            </div>
            <div class="space-y-1">
              <div class="flex justify-between text-sm">
                <span>Fat</span>
                <span>{{ dailyStats.macros.fat.current }}g / {{ dailyStats.macros.fat.goal }}g</span>
              </div>
              <Progress :value="(dailyStats.macros.fat.current / dailyStats.macros.fat.goal) * 100" class="h-1" />
            </div>
          </div>
        </CardContent>
      </Card>
    </div>

    <!-- Nutrition Section -->
    <Card>
      <CardHeader>
        <h2 class="text-xl font-semibold">
          Today's Meals
        </h2>
      </CardHeader>
      <CardContent>
        <div class="space-y-6">
          <div
            v-for="meal in meals"
            :key="meal.name"
            class="flex items-center justify-between p-4 bg-muted rounded-lg"
          >
            <div class="space-y-1">
              <div class="flex items-center gap-2">
                <h3 class="font-medium">
                  {{ meal.name }}
                </h3>
                <span class="text-sm text-muted-foreground">{{ meal.time }}</span>
              </div>
              <div class="flex flex-wrap gap-2">
                <span
                  v-for="food in meal.foods"
                  :key="food"
                  class="px-2 py-1 text-xs bg-accent rounded-full"
                >
                  {{ food }}
                </span>
              </div>
            </div>
            <div class="flex items-center gap-4">
              <span class="text-lg font-semibold">{{ meal.calories }} cal</span>
              <Dialog v-model:open="showAddFoodDialog">
                <DialogTrigger as-child>
                  <Button
                    variant="outline"
                    size="icon"
                    @click="selectedMeal = meal.name"
                  >
                    <Plus class="h-4 w-4" />
                  </Button>
                </DialogTrigger>
                <DialogContent>
                  <DialogHeader>
                    <DialogTitle>Add Food to {{ meal.name }}</DialogTitle>
                  </DialogHeader>
                  <div class="space-y-4 py-4">
                    <div class="space-y-2">
                      <Label>Food Name</Label>
                      <Input
                        v-model="newFood.name"
                        placeholder="e.g., Oatmeal"
                      />
                    </div>
                    <div class="space-y-2">
                      <Label>Calories</Label>
                      <Input
                        v-model="newFood.calories"
                        type="number"
                        placeholder="e.g., 150"
                      />
                    </div>
                    <div class="space-y-2">
                      <Label>Time</Label>
                      <Input
                        v-model="newFood.time"
                        type="time"
                      />
                    </div>
                  </div>
                  <div class="flex justify-end">
                    <Button @click="addFood">
                      Add Food
                    </Button>
                  </div>
                </DialogContent>
              </Dialog>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>

    <!-- Water Tracker Section -->
    <Card>
      <CardHeader>
        <h2 class="text-xl font-semibold">
          Water Intake
        </h2>
      </CardHeader>
      <CardContent>
        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <div class="space-y-1">
              <p class="text-sm font-medium text-muted-foreground">
                Daily Goal
              </p>
              <p class="text-2xl font-bold">
                {{ waterStats.current }}L / {{ waterStats.goal }}L
              </p>
            </div>
            <Dialog v-model:open="showAddWaterDialog">
              <DialogTrigger as-child>
                <Button variant="outline" size="icon">
                  <Plus class="h-4 w-4" />
                </Button>
              </DialogTrigger>
              <DialogContent>
                <DialogHeader>
                  <DialogTitle>Add Water</DialogTitle>
                </DialogHeader>
                <div class="space-y-4 py-4">
                  <div class="space-y-2">
                    <Label>Amount (Liters)</Label>
                    <Input
                      v-model="waterAmount"
                      type="number"
                      step="0.25"
                      min="0.25"
                      max="1"
                    />
                  </div>
                </div>
                <div class="flex justify-end">
                  <Button @click="addWater">
                    Add Water
                  </Button>
                </div>
              </DialogContent>
            </Dialog>
          </div>
          <Progress :value="waterProgress" class="h-2" />
        </div>
      </CardContent>
    </Card>
  </div>
</template>

<route lang="yaml">
meta:
  layout: dashboard
</route> 
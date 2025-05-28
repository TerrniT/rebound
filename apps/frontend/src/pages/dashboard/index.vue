<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader } from '@/components/ui/card'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Progress } from '@/components/ui/progress'
import { Tab, TabContent, TabList } from '@/components/ui/tabs'
import { type DateValue, getLocalTimeZone, today } from '@internationalized/date'
import { Clock, Plus, Search, Star } from 'lucide-vue-next'

import { computed, ref } from 'vue'

// Types
interface Food {
  name: string
  calories: number
  protein: number
  carbs: number
  fat: number
}

interface Meal {
  name: string
  calories: number
  time: string
  foods: string[]
}

// Calendar
const date = ref(today(getLocalTimeZone())) as Ref<DateValue>

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
const carbsProgress = computed(() => (dailyStats.value.macros.carbs.current / dailyStats.value.macros.carbs.goal) * 100)

const proteinProgress = computed(() =>
  (dailyStats.value.macros.protein.current / dailyStats.value.macros.protein.goal) * 100,
)

const fatProgress = computed(() =>
  (dailyStats.value.macros.fat.current / dailyStats.value.macros.fat.goal) * 100,
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

// Meal Details Modal
const showMealDetailsDialog = ref(false)
const selectedMealDetails = ref<Meal | null>(null)

function openMealDetails(meal: Meal) {
  selectedMealDetails.value = meal
  showMealDetailsDialog.value = true
}

// Add Food Modal
const showAddFoodModal = ref(false)
const searchQuery = ref('')
const activeTab = ref('frequent')
const selectedMeal = ref('')

// Sample data for tabs
const frequentFoods = ref<Food[]>([
  { name: 'Oatmeal', calories: 150, protein: 5, carbs: 27, fat: 3 },
  { name: 'Banana', calories: 105, protein: 1, carbs: 27, fat: 0 },
  { name: 'Chicken Breast', calories: 165, protein: 31, carbs: 0, fat: 3.6 },
])

const recentFoods = ref<Food[]>([
  { name: 'Salmon', calories: 208, protein: 22, carbs: 0, fat: 13 },
  { name: 'Quinoa', calories: 120, protein: 4, carbs: 21, fat: 2 },
])

const favoriteFoods = ref<Food[]>([
  { name: 'Greek Yogurt', calories: 130, protein: 12, carbs: 9, fat: 4 },
  { name: 'Almonds', calories: 160, protein: 6, carbs: 6, fat: 14 },
])

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

function addWater() {
  const amount = Number.parseFloat(waterAmount.value)
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
        {{ date.toString() }}
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
        <CommonCalendar v-model="date" />
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
            <Progress v-model="caloriesProgress" class="h-2" />
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
            <p class="text-2xl font-bold" :class="remainingCalories < 0 ? 'text-destructive' : 'text-brand'">
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
              <Progress v-model="carbsProgress" class="h-1" />
            </div>
            <div class="space-y-1">
              <div class="flex justify-between text-sm">
                <span>Protein</span>
                <span>{{ dailyStats.macros.protein.current }}g / {{ dailyStats.macros.protein.goal }}g</span>
              </div>
              <Progress v-model="proteinProgress" class="h-1" />
            </div>
            <div class="space-y-1">
              <div class="flex justify-between text-sm">
                <span>Fat</span>
                <span>{{ dailyStats.macros.fat.current }}g / {{ dailyStats.macros.fat.goal }}g</span>
              </div>
              <Progress v-model="fatProgress" class="h-1" />
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
        <div class="space-y-4">
          <div
            v-for="meal in meals"
            :key="meal.name"
            class="flex items-center justify-between bg-muted rounded-lg cursor-pointer hover:bg-muted/80 transition-colors"
            @click="openMealDetails(meal)"
          >
            <div class="space-y-1 p-4">
              <div class="flex items-center gap-2 justify-center">
                <h3 class="font-medium text-md">
                  {{ meal.name }}
                </h3>
                <span class="block text-sm text-muted-foreground">{{ meal.time }}</span>
              </div>
            </div>
            <div class="flex items-center gap-4 p-2">
              <span class="text-lg font-semibold">{{ meal.calories }} cal</span>
              <Button
                variant="default"
                size="icon"
                @click.stop="showAddFoodModal = true; selectedMeal = meal.name"
              >
                <Plus class="h-4 w-4" />
              </Button>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>

    <!-- Meal Details Modal -->
    <Dialog v-model:open="showMealDetailsDialog">
      <DialogContent class="max-w-2xl">
        <DialogHeader>
          <DialogTitle>{{ selectedMealDetails?.name }} Details</DialogTitle>
        </DialogHeader>
        <div class="space-y-6 py-4">
          <div class="grid grid-cols-2 gap-4">
            <div class="space-y-2">
              <p class="text-sm font-medium text-muted-foreground">
                Total Calories
              </p>
              <p class="text-2xl font-bold">
                {{ selectedMealDetails?.calories }} cal
              </p>
            </div>
            <div class="space-y-2">
              <p class="text-sm font-medium text-muted-foreground">
                Time
              </p>
              <p class="text-2xl font-bold">
                {{ selectedMealDetails?.time }}
              </p>
            </div>
          </div>

          <div class="space-y-4">
            <h3 class="font-semibold">
              Foods
            </h3>
            <div class="space-y-2">
              <div
                v-for="food in selectedMealDetails?.foods"
                :key="food"
                class="flex items-center justify-between p-2 bg-muted rounded-lg"
              >
                <span>{{ food }}</span>
                <Button variant="ghost" size="sm">
                  Remove
                </Button>
              </div>
            </div>
          </div>

          <div class="flex justify-end gap-2">
            <Button variant="outline" @click="showMealDetailsDialog = false">
              Close
            </Button>
            <Button @click="showAddFoodModal = true; showMealDetailsDialog = false">
              Add Food
            </Button>
          </div>
        </div>
      </DialogContent>
    </Dialog>

    <!-- Add Food Modal -->
    <Dialog v-model:open="showAddFoodModal">
      <DialogContent class="max-w-2xl">
        <DialogHeader>
          <DialogTitle>Add Food to {{ selectedMeal }}</DialogTitle>
        </DialogHeader>
        <div class="space-y-6 py-4">
          <div class="flex gap-2">
            <div class="relative flex-1">
              <Search class="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input
                v-model="searchQuery"
                placeholder="Search foods or scan barcode..."
                class="pl-8"
              />
            </div>
            <Button variant="outline">
              Scan Barcode
            </Button>
          </div>

          <TabList :default-value="activeTab" class="w-full">
            <template #tabs="{ active }">
              <div class="grid w-full grid-cols-3">
                <Tab
                  value="frequent"
                  :class="active === 'frequent' ? 'border-b-2 border-primary' : ''"
                  class="flex items-center justify-center gap-2 py-2"
                >
                  <Clock class="h-4 w-4" />
                  Frequent
                </Tab>
                <Tab
                  value="recent"
                  :class="active === 'recent' ? 'border-b-2 border-primary' : ''"
                  class="flex items-center justify-center gap-2 py-2"
                >
                  <Clock class="h-4 w-4" />
                  Recent
                </Tab>
                <Tab
                  value="favorites"
                  :class="active === 'favorites' ? 'border-b-2 border-primary' : ''"
                  class="flex items-center justify-center gap-2 py-2"
                >
                  <Star class="h-4 w-4" />
                  Favorites
                </Tab>
              </div>
            </template>

            <template #content>
              <TabContent value="frequent" class="space-y-4">
                <div
                  v-for="food in frequentFoods"
                  :key="food.name"
                  class="flex items-center justify-between p-4 bg-muted rounded-lg cursor-pointer hover:bg-muted/80"
                >
                  <div class="space-y-1">
                    <p class="font-medium">
                      {{ food.name }}
                    </p>
                    <p class="text-sm text-muted-foreground">
                      {{ food.calories }} cal | P: {{ food.protein }}g | C: {{ food.carbs }}g | F: {{ food.fat }}g
                    </p>
                  </div>
                  <Button variant="ghost" size="sm">
                    Add
                  </Button>
                </div>
              </TabContent>

              <TabContent value="recent" class="space-y-4">
                <div
                  v-for="food in recentFoods"
                  :key="food.name"
                  class="flex items-center justify-between p-4 bg-muted rounded-lg cursor-pointer hover:bg-muted/80"
                >
                  <div class="space-y-1">
                    <p class="font-medium">
                      {{ food.name }}
                    </p>
                    <p class="text-sm text-muted-foreground">
                      {{ food.calories }} cal | P: {{ food.protein }}g | C: {{ food.carbs }}g | F: {{ food.fat }}g
                    </p>
                  </div>
                  <Button variant="ghost" size="sm">
                    Add
                  </Button>
                </div>
              </TabContent>

              <TabContent value="favorites" class="space-y-4">
                <div
                  v-for="food in favoriteFoods"
                  :key="food.name"
                  class="flex items-center justify-between p-4 bg-muted rounded-lg cursor-pointer hover:bg-muted/80"
                >
                  <div class="space-y-1">
                    <p class="font-medium">
                      {{ food.name }}
                    </p>
                    <p class="text-sm text-muted-foreground">
                      {{ food.calories }} cal | P: {{ food.protein }}g | C: {{ food.carbs }}g | F: {{ food.fat }}g
                    </p>
                  </div>
                  <Button variant="ghost" size="sm">
                    Add
                  </Button>
                </div>
              </TabContent>
            </template>
          </TabList>

          <div class="flex justify-end gap-2">
            <Button variant="outline" @click="showAddFoodModal = false">
              Cancel
            </Button>
            <Button @click="showAddFoodModal = false">
              Done
            </Button>
          </div>
        </div>
      </DialogContent>
    </Dialog>

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
          <Progress v-model="waterProgress" class="h-2" />
        </div>
      </CardContent>
    </Card>
  </div>
</template>

<route lang="yaml">
meta:
  layout: dashboard
</route>

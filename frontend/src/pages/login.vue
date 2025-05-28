<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardFooter, CardHeader } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { useAuthStore } from '@/stores/auth'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const auth = useAuthStore()

const email = ref('')
const password = ref('')
const isLoading = ref(false)
const error = ref('')

async function handleLogin() {
  if (!email.value || !password.value) {
    error.value = 'Please fill in all fields'
    return
  }

  try {
    isLoading.value = true
    error.value = ''
    await auth.login(email.value, password.value)
    router.push('/dashboard')
  }
  catch (e) {
    error.value = 'Invalid credentials'
  }
  finally {
    isLoading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-background">
    <Card class="w-full max-w-md">
      <CardHeader class="text-center">
        <h1 class="text-3xl font-bold">
          Welcome Back
        </h1>
        <p class="text-muted-foreground mt-2">
          Sign in to your account
        </p>
      </CardHeader>

      <CardContent>
        <form class="space-y-6" @submit.prevent="handleLogin">
          <div class="space-y-4">
            <div class="space-y-2">
              <label class="text-sm font-medium">Email</label>
              <Input
                v-model="email"
                type="email"
                placeholder="Enter your email"
              />
            </div>

            <div class="space-y-2">
              <label class="text-sm font-medium">Password</label>
              <Input
                v-model="password"
                type="password"
                placeholder="Enter your password"
              />
            </div>

            <div v-if="error" class="text-sm text-destructive">
              {{ error }}
            </div>
          </div>

          <Button
            type="submit"
            class="w-full"
            :disabled="isLoading"
          >
            <span v-if="isLoading">Signing in...</span>
            <span v-else>Sign in</span>
          </Button>
        </form>
      </CardContent>

      <CardFooter>
        <p class="text-center text-sm text-muted-foreground w-full">
          For demo purposes, use any email and password
        </p>
      </CardFooter>
    </Card>
  </div>
</template>

<route lang="yaml">
meta:
  layout: default
</route>

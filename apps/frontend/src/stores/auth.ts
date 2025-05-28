import { defineStore } from 'pinia'
import { ref } from 'vue'

interface User {
  id: string
  email: string
  name: string
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const isAuthenticated = ref(false)

  // Mock login function
  const login = async (email: string, password: string) => {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000))

    // Mock successful login
    user.value = {
      id: '1',
      email,
      name: 'John Doe',
    }
    isAuthenticated.value = true
  }

  // Mock logout function
  const logout = async () => {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 500))

    user.value = null
    isAuthenticated.value = false
  }

  return {
    user,
    isAuthenticated,
    login,
    logout,
  }
})

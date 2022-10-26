import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import api from '@/api'

export const usePlayStore = defineStore('play', () => {
  const languageId = ref('')
  const availableLanguages = ref([])

  return {
    availableLanguages,

    language: computed(() => {
      return availableLanguages.value.find(
        (language) => language.id === languageId.value
      )
    }),

    fetch: () => {
      availableLanguages.value = api.languages.list()
    },

    run: () => {
      api.runs.create()
    },

    changeLanguage: (newLanguageId) => {
      languageId.value = newLanguageId
    }
  }
})

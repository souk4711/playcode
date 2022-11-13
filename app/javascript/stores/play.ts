import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import api from '@/api'

const enum RunStatus {
  Norun,
  Running,
  Completed
}

export const usePlayStore = defineStore('play', () => {
  const languageId = ref('')
  const availableLanguages = ref([])

  const file = ref({
    content: ''
  })

  const runStatus = ref(RunStatus.Norun)
  const runResult = ref(null)

  return {
    availableLanguages,
    file,
    runStatus,
    runResult,

    language: computed(() => {
      return availableLanguages.value.find(
        (language) => language.id === languageId.value
      )
    }),

    isNorun: computed(() => runStatus.value === RunStatus.Norun),
    isRunning: computed(() => runStatus.value === RunStatus.Running),
    isCompleted: computed(() => runStatus.value === RunStatus.Completed),

    fetch: async () => {
      availableLanguages.value = await api.languages.list()
    },

    run: async () => {
      try {
        runStatus.value = RunStatus.Running
        runResult.value = await api.runs.create({ lang: languageId, file })
      } finally {
        runStatus.value = RunStatus.Completed
      }
    },

    changeLanguage: (newLanguageId) => {
      languageId.value = newLanguageId
    }
  }
})

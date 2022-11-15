import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { useSessionStorage } from '@vueuse/core'
import api from '@/api'

const enum RunStatus {
  Norun,
  Running,
  Completed,
  Unexpected
}

export const usePlayStore = defineStore('play', () => {
  const languageId = useSessionStorage('playStore#languageId', '')
  const availableLanguages = useSessionStorage('playStore#availableLanguages', [])

  const fileInitialValue = { content: '' }
  const file = useSessionStorage('playStore#file', fileInitialValue, { mergeDefaults: true })

  const runStatus = ref(RunStatus.Norun)
  const runResult = ref(null)

  return {
    availableLanguages,
    file,
    runStatus,
    runResult,

    language: computed(() => {
      return availableLanguages.value.find((language) => language.id === languageId.value)
    }),

    isNorun: computed(() => runStatus.value === RunStatus.Norun),
    isRunning: computed(() => runStatus.value === RunStatus.Running),
    isCompleted: computed(() => runStatus.value === RunStatus.Completed),
    isUnexpected: computed(() => runStatus.value === RunStatus.Unexpected),

    fetch: async () => {
      availableLanguages.value = await api.languages.list()
    },

    run: async () => {
      try {
        runStatus.value = RunStatus.Running
        runResult.value = await api.runs.create({ lang: languageId, file })
        runStatus.value = RunStatus.Completed
      } catch {
        runStatus.value = RunStatus.Unexpected
      }
    },

    changeLanguage: (newLanguageId) => {
      languageId.value = newLanguageId
    }
  }
})

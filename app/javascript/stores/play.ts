import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { useSessionStorage } from '@vueuse/core'
import api from '@/api'
import type { Language, RunFile, RunResult } from '@/types/models'

const enum PlayStatus {
  Norun,
  Running,
  Completed,
  Unexpected
}

export const usePlayStore = defineStore('play', () => {
  const availableLanguages = useSessionStorage<Language[]>('playStore#availableLanguages', [])
  const languageId = useSessionStorage<string>('playStore#languageId', '')
  const language = computed<Language | null>(
    () => availableLanguages.value.find((language) => language.id === languageId.value) ?? null
  )
  const changeLanguage = (newLanguageId: string) => {
    languageId.value = newLanguageId
  }

  const playStatus = ref<PlayStatus>(PlayStatus.Norun)
  const isNorun = computed<boolean>(() => playStatus.value === PlayStatus.Norun)
  const isRunning = computed<boolean>(() => playStatus.value === PlayStatus.Running)
  const isCompleted = computed<boolean>(() => playStatus.value === PlayStatus.Completed)
  const isUnexpected = computed<boolean>(() => playStatus.value === PlayStatus.Unexpected)

  const runFileInitialValue: RunFile = { content: '' }
  const runFile = useSessionStorage<RunFile>('playStore#runFile', runFileInitialValue, { mergeDefaults: true })
  const runResult = ref<RunResult | null>(null)

  const fetch = async () => {
    const r = await api.languages.list()
    availableLanguages.value = r.items
  }

  const run = async () => {
    playStatus.value = PlayStatus.Running
    try {
      const r = await api.runs.create({
        language_id: language.value?.id ?? '',
        file: runFile.value
      })
      runResult.value = r
      playStatus.value = PlayStatus.Completed
    } catch {
      playStatus.value = PlayStatus.Unexpected
    }
  }

  return {
    availableLanguages,
    language,
    changeLanguage,

    isNorun,
    isRunning,
    isCompleted,
    isUnexpected,

    runFile,
    runResult,

    fetch,
    run
  }
})

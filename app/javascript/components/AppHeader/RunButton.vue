<script setup lang="ts">
import { computed } from 'vue'
import { SpinnerIcon } from '@/components'
import { usePlayStore } from '@/stores'

const store = usePlayStore()

const disabled = computed(() => {
  if (!store.language) return true
  if (store.isRunning) return true
  return false
})
</script>

<template>
  <button class="btn btn-primary w-32" :disabled="disabled" @click="store.run">
    <template v-if="store.isRunning">
      <SpinnerIcon class="animate-spin h-4 w-4 mr-2" />
      <span class="uppercase">{{ $t('running') }}</span>
    </template>
    <template v-else>
      <span class="uppercase">{{ $t('run') }}</span>
    </template>
  </button>
</template>

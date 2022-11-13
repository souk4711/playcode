<script setup lang="ts">
import { usePlayStore } from '@/stores'

const store = usePlayStore()
</script>

<template>
  <button class="btn btn-primary w-24" :disabled="!store.language">
    <span class="uppercase">{{ $t('run') }}</span>
  </button>
</template>

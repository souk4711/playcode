<script setup lang="ts">
import { computed } from 'vue'
import { Codemirror } from 'vue-codemirror'
import { usePlayStore } from '@/stores'
import langs from '@/config/langs'

const store = usePlayStore()

const lang = computed(() => langs.get(store.language?.langcode ?? ''))
</script>

<template>
  <div class="w-full">
    <Codemirror v-model="store.runFile.content" :extensions="lang.extensions" />
  </div>
</template>

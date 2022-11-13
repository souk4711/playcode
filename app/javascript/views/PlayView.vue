<script setup lang="ts">
import { onMounted } from 'vue'
import { usePlayStore } from '@/stores'
import CodeEditor from './PlayView/CodeEditor.vue'
import ConsoleOutput from './PlayView/ConsoleOutput.vue'

const store = usePlayStore()

onMounted(async () => {
  await store.fetch()
})
</script>

<template>
  <div class="flex h-[calc(100vh-4rem)] p-4 gap-4">
    <CodeEditor class="border-4 border-base-300" />
    <ConsoleOutput class="border-2 border-base-300" />
  </div>
</template>

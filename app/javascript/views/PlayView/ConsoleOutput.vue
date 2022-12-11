<script setup lang="ts">
import { SpinnerIcon } from '@/components'
import { usePlayStore } from '@/stores'

const store = usePlayStore()
</script>

<template>
  <div class="w-full flex flex-col items-center px-4 py-2 scrollbar">
    <h1 class="font-bold">Execution</h1>

    <div v-show="store.isRunning" class="w-full">
      <div class="divider text-base-content/40">Progress</div>
      <div>
        <SpinnerIcon class="animate-spin h-4 w-4 ml-2" />
      </div>
    </div>

    <div v-show="store.isUnexpected" class="w-full">
      <div class="divider text-base-content/40">Errors</div>
      <div>
        <pre class="whitespace-pre-wrap">{{ store.runException }}</pre>
      </div>
    </div>

    <div v-show="store.isCompleted" class="w-full">
      <div class="divider text-base-content/40">Standard Error</div>
      <div>
        <pre class="whitespace-pre-wrap">{{ store.runResult?.stderr }}</pre>
      </div>
    </div>

    <div v-show="store.isCompleted" class="w-full">
      <div class="divider text-base-content/40">Standard Output</div>
      <div>
        <pre class="whitespace-pre-wrap">{{ store.runResult?.stdout }}</pre>
      </div>
    </div>
  </div>
</template>

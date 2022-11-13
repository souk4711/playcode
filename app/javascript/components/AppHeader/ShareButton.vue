<script setup lang="ts">
import { ShareIcon } from '@heroicons/vue/24/solid'
</script>

<template>
  <div class="tooltip tooltip-left" :data-tip="$t('shareLink')">
    <a class="btn btn-ghost">
      <ShareIcon class="h-6 w-6" />
    </a>
  </div>
</template>

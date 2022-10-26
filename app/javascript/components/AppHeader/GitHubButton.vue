<script setup lang="ts">
import { GitHubIcon } from '@/components'
</script>

<template>
  <div class="tooltip tooltip-left" :data-tip="$t('viewOnGitHub')">
    <a
      class="btn btn-ghost"
      href="https://github.com/souk4711/playcode"
      target="_blank"
      rel="noopener noreferrer"
    >
      <GitHubIcon class="h-6 w-6" />
    </a>
  </div>
</template>

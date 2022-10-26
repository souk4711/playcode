<script setup lang="ts">
import CodeLanguageDropdown from './AppHeader/CodeLanguageDropdown.vue'
import RunButton from './AppHeader/RunButton.vue'
import ShareButton from './AppHeader/ShareButton.vue'
import ThemeDropdown from './AppHeader/ThemeDropdown.vue'
import LanguageDropdown from './AppHeader/LanguageDropdown.vue'
import GitHubButton from './AppHeader/GitHubButton.vue'
</script>

<template>
  <div class="navbar bg-base-100 shadow">
    <div class="flex-none px-2 gap-2">
      <CodeLanguageDropdown />
      <RunButton />
    </div>
    <div class="flex-1 flex justify-end px-2">
      <ShareButton />
      <ThemeDropdown />
      <LanguageDropdown />
      <GitHubButton />
    </div>
  </div>
</template>

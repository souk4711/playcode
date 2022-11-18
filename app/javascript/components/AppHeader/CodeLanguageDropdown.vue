<script setup lang="ts">
import { ChevronDownIcon } from '@heroicons/vue/24/solid'
import { usePlayStore } from '@/stores'

const store = usePlayStore()

const changeLanguage = (languageId: string) => {
  store.changeLanguage(languageId)

  if (document.activeElement instanceof HTMLElement) {
    document.activeElement.blur()
  }
}
</script>

<template>
  <div class="dropdown dropdown-end">
    <label class="btn btn-secondary flex w-48" tabindex="0">
      <span class="flex-1 text-left truncate" :class="{ uppercase: !store.language }">{{
        store.language ? store.language.name : $t('changeCodeLanguage')
      }}</span>
      <ChevronDownIcon class="h-4 w-4" />
    </label>
    <ul
      class="dropdown-content max-h-96 h-[70vh] w-48 overflow-y-auto bg-base-100 text-base-content shadow"
      tabindex="0"
    >
      <li v-for="language in store.availableLanguages" :key="language.id">
        <div class="btn btn-ghost w-full justify-start" tabindex="0" @click="changeLanguage(language.id)">
          <span class="truncate">{{ language.name }}</span>
        </div>
      </li>
    </ul>
  </div>
</template>

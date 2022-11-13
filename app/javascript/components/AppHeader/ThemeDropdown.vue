<script setup lang="ts">
import { onMounted } from 'vue'
import { themeChange } from 'theme-change'
import { SwatchIcon } from '@heroicons/vue/24/solid'
import themes from '@/config/themes'

onMounted(() => {
  themeChange(false)
})
</script>

<template>
  <div class="dropdown dropdown-end">
    <div class="tooltip tooltip-left" :data-tip="$t('changeTheme')">
      <label class="btn btn-ghost" tabindex="0">
        <SwatchIcon class="h-6 w-6" />
      </label>
    </div>
    <ul
      class="dropdown-content max-h-96 h-[70vh] w-52 overflow-y-auto p-2 mt-4 bg-base-200 text-base-content shadow grid gap-2"
    >
      <li v-for="theme in themes" :key="theme">
        <button
          class="w-full rounded-md outline outline-2"
          data-act-class="outline"
          :data-set-theme="theme"
        >
          <div
            class="w-full rounded-md bg-base-100 text-base-content font-sans cursor-pointer"
            :data-theme="theme"
          >
            <div class="grid grid-cols-5 grid-rows-3">
              <div
                class="col-span-5 row-span-3 row-start-1 flex gap-1 py-3 px-4"
              >
                <div class="flex-grow text-left text-sm font-bold">
                  {{ theme }}
                </div>
                <div class="flex flex-shrink-0 flex-wrap gap-1">
                  <div class="bg-primary w-2 rounded" />
                  <div class="bg-secondary w-2 rounded" />
                  <div class="bg-accent w-2 rounded" />
                  <div class="bg-neutral w-2 rounded" />
                </div>
              </div>
            </div>
          </div>
        </button>
      </li>
    </ul>
  </div>
</template>

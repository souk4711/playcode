<script setup lang="ts">
import { onMounted } from 'vue'
import { LanguageIcon, SwatchIcon } from '@heroicons/vue/24/solid'
import { themeChange } from 'theme-change'
import { GitHubIcon } from '@/components/icons'
import { updateI18nLocale } from '@/locales'

const themes = [
  'light',
  'dark',
  'cupcake',
  'bumblebee',
  'emerald',
  'corporate',
  'synthwave',
  'retro',
  'cyberpunk',
  'valentine',
  'halloween',
  'garden',
  'forest',
  'aqua',
  'lofi',
  'pastel',
  'fantasy',
  'wireframe',
  'black',
  'luxury',
  'dracula',
  'cmyk',
  'autumn',
  'business',
  'acid',
  'lemonade',
  'night',
  'coffee',
  'winter'
]

onMounted(() => {
  themeChange(false)
})
</script>

<template>
  <div class="navbar bg-base-100 shadow">
    <div class="flex-none px-2">
      <a class="text-lg font-bold uppercase">Playcode</a>
    </div>
    <div class="flex-1 flex justify-end px-2">
      <div class="dropdown dropdown-end">
        <div class="tooltip tooltip-left" :data-tip="$t('Change Theme')">
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
      <div class="dropdown dropdown-end">
        <div class="tooltip tooltip-left" :data-tip="$t('Change Language')">
          <label class="btn btn-ghost" tabindex="0">
            <LanguageIcon class="h-6 w-6" />
          </label>
        </div>
        <ul
          class="dropdown-content menu w-52 p-2 mt-4 bg-base-200 text-base-content shadow gap-1"
        >
          <li v-for="locale in $i18n.availableLocales" :key="locale">
            <div
              class="btn btn-ghost"
              :class="{ active: locale === $i18n.locale }"
              tabindex="0"
              @click="updateI18nLocale(locale)"
              @keydown.enter="updateI18nLocale(locale)"
            >
              <img
                class="w-6 h-6"
                :src="$t('locale.icon', '', { locale: locale })"
              />
              <span>{{ $t('locale.name', '', { locale: locale }) }}</span>
            </div>
          </li>
        </ul>
      </div>
      <div class="tooltip tooltip-left" :data-tip="$t('View on GitHub')">
        <a
          class="btn btn-ghost"
          href="https://github.com/souk4711/playcode"
          target="_blank"
          rel="noopener noreferrer"
        >
          <GitHubIcon class="h-6 w-6" />
        </a>
      </div>
    </div>
  </div>
</template>

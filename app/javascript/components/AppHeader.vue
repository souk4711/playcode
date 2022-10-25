<script setup lang="ts">
import { LanguageIcon } from '@heroicons/vue/24/solid'
import { GitHubIcon } from '@/components/icons'
import { updateI18nLocale } from '@/locales'
</script>

<template>
  <div class="navbar bg-base-100 shadow">
    <div class="flex-none px-2">
      <a class="text-lg font-bold uppercase">Playcode</a>
    </div>
    <div class="flex-1 flex justify-end px-2">
      <div class="dropdown dropdown-end">
        <div class="tooltip tooltip-left" :data-tip="$t('Change Language')">
          <label class="btn btn-ghost" tabindex="0">
            <LanguageIcon class="h-6 w-6" />
          </label>
        </div>
        <ul class="dropdown-content menu w-52 p-2 mt-3 bg-base-100 shadow">
          <li
            v-for="locale in $i18n.availableLocales"
            :key="`locale-${locale}`"
            :value="locale"
          >
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

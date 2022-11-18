<script setup lang="ts">
import { LanguageIcon } from '@heroicons/vue/24/solid'
import { changeI18nLocale } from '@/locales'

const changeLocale = (locale: string) => {
  changeI18nLocale(locale)

  if (document.activeElement instanceof HTMLElement) {
    document.activeElement.blur()
  }
}
</script>

<template>
  <div class="dropdown dropdown-end">
    <div class="tooltip tooltip-left" :data-tip="$t('changeLanguage')">
      <label class="btn btn-ghost" tabindex="0">
        <LanguageIcon class="h-6 w-6" />
      </label>
    </div>
    <ul class="dropdown-content menu w-52 p-2 mt-4 bg-base-100 text-base-content shadow gap-1" tabindex="0">
      <li v-for="locale in $i18n.availableLocales" :key="locale">
        <div
          class="btn btn-ghost"
          :class="{ active: locale === $i18n.locale }"
          tabindex="0"
          @click="changeLocale(locale)"
        >
          <img class="w-6 h-6" :src="$t('locale.icon', '', { locale: locale })" />
          <span>{{ $t('locale.name', '', { locale: locale }) }}</span>
        </div>
      </li>
    </ul>
  </div>
</template>

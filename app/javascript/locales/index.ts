import { createI18n as _createI18n } from 'vue-i18n'
import en from './en.json'
import zhHans from './zh-Hans.json'
import zhHant from './zh-Hant.json'

const i18n = (() => {
  let locale = localStorage.getItem('i18n.locale')
  if (locale === null) locale = 'en'

  return _createI18n({
    legacy: false,
    locale,
    fallbackLocale: 'en',
    messages: {
      en,
      zhHans,
      zhHant
    }
  })
})()

const updateI18nLocale = (locale): void => {
  i18n.global.locale.value = locale
  localStorage.setItem('i18n.locale', locale)
}

export default i18n
export { updateI18nLocale }

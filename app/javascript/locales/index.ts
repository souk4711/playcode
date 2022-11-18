import { createI18n as _createI18n } from 'vue-i18n'
import en from './en.json'
import zhHans from './zh-Hans.json'
import zhHant from './zh-Hant.json'

type Locale = 'en' | 'zhHans' | 'zhHant'

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

const changeI18nLocale = (locale: string) => {
  i18n.global.locale.value = locale as Locale
  localStorage.setItem('i18n.locale', locale)
}

export default i18n
export { changeI18nLocale }

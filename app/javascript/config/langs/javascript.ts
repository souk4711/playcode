import { StreamLanguage } from '@codemirror/language'
import { javascript } from '@codemirror/legacy-modes/mode/javascript'

export default {
  extensions: [StreamLanguage.define(javascript)]
}

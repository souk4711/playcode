import { StreamLanguage } from '@codemirror/language'
import { go } from '@codemirror/legacy-modes/mode/go'

export default {
  extensions: [StreamLanguage.define(go)]
}

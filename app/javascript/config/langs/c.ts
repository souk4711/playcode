import { StreamLanguage } from '@codemirror/language'
import { c } from '@codemirror/legacy-modes/mode/clike'

export default {
  extensions: [StreamLanguage.define(c)]
}

import { StreamLanguage } from '@codemirror/language'
import { java } from '@codemirror/legacy-modes/mode/clike'

export default {
  extensions: [StreamLanguage.define(java)]
}

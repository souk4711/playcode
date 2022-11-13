import { StreamLanguage } from '@codemirror/language'
import { scala } from '@codemirror/legacy-modes/mode/clike'

export default {
  extensions: [StreamLanguage.define(scala)]
}

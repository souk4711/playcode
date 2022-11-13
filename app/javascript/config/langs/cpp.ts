import { StreamLanguage } from '@codemirror/language'
import { cpp } from '@codemirror/legacy-modes/mode/clike'

export default {
  extensions: [StreamLanguage.define(cpp)]
}

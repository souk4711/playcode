import { StreamLanguage } from '@codemirror/language'
import { typescript } from '@codemirror/legacy-modes/mode/javascript'

export default {
  extensions: [StreamLanguage.define(typescript)]
}

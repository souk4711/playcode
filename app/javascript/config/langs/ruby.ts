import { StreamLanguage } from '@codemirror/language'
import { ruby } from '@codemirror/legacy-modes/mode/ruby'

export default {
  extensions: [StreamLanguage.define(ruby)]
}

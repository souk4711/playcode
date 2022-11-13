import { StreamLanguage } from '@codemirror/language'
import { d } from '@codemirror/legacy-modes/mode/d'

export default {
  extensions: [StreamLanguage.define(d)]
}

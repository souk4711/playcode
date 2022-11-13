import { StreamLanguage } from '@codemirror/language'
import { python } from '@codemirror/legacy-modes/mode/python'

export default {
  extensions: [StreamLanguage.define(python)]
}

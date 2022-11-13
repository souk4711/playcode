import { StreamLanguage } from '@codemirror/language'
import { rust } from '@codemirror/legacy-modes/mode/rust'

export default {
  extensions: [StreamLanguage.define(rust)]
}

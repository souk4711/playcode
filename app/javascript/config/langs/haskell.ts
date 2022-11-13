import { StreamLanguage } from '@codemirror/language'
import { haskell } from '@codemirror/legacy-modes/mode/haskell'

export default {
  extensions: [StreamLanguage.define(haskell)]
}

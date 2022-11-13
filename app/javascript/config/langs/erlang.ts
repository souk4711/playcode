import { StreamLanguage } from '@codemirror/language'
import { erlang } from '@codemirror/legacy-modes/mode/erlang'

export default {
  extensions: [StreamLanguage.define(erlang)]
}

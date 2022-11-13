import c from './c'
import cpp from './cpp'
import d from './d'
import erlang from './erlang'
import go from './go'
import haskell from './haskell'
import java from './java'
import javascript from './javascript'
import plaintext from './plaintext'
import python from './python'
import ruby from './ruby'
import rust from './rust'
import scala from './scala'
import typescript from './typescript'

const langs = {
  c,
  cpp,
  d,
  erlang,
  go,
  haskell,
  java,
  javascript,
  plaintext,
  python,
  ruby,
  rust,
  scala,
  typescript
}

const get = (langcode) => {
  return langs[langcode] ?? plaintext
}

export default {
  get
}

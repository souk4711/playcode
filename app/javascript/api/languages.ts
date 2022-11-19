import http from '@/api/http'
import type { IndexResponse } from '@/types/api/languages'

const list = async (): Promise<IndexResponse> => await http.get('/languages')

export default {
  list
}

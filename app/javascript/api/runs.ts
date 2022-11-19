import http from '@/api/http'
import type { CreateRequest, CreateResponse } from '@/types/api/runs'

const create = async (request: CreateRequest): Promise<CreateResponse> => await http.post('/runs')

export default {
  create
}

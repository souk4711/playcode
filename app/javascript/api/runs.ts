import type { RunResult } from '@/types/models'
import type { CreateRequest, CreateResponse } from '@/types/api/runs'

const create = async (request: CreateRequest): Promise<CreateResponse> => {
  const result: RunResult = {
    status: 'ok',
    reason: '',
    stdout: 'Hello, World',
    stderr: ''
  }

  return await new Promise((resolve) =>
    setTimeout(() => {
      resolve(result)
    }, 2000)
  )
}

export default {
  create
}

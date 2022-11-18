import type { RunFile, RunResult } from '@/types/models'

interface CreateRequest {
  language_id: string
  file: RunFile
}

type CreateResponse = RunResult

export type { CreateRequest, CreateResponse }

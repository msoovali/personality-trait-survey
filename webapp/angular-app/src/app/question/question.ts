export interface Option {
  id: string
  description: string
  score: number
}

export interface Question {
  id: string
  description: string
  options: Option[]
}

import { Trait } from "../trait/trait";

export interface Survey {
  id: string
  score: number
  trait: Trait
  answers: SurveyAnswer[]
}

export interface SurveyAnswer {
  questionId: string
  optionId: string
}

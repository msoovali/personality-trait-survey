import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { QuestionAddEditComponent } from './question/question-add-edit/question-add-edit.component';
import { QuestionListComponent } from './question/question-list/question-list.component';
import { ReviewSurveyComponent } from './survey/review-survey/review-survey.component';
import { SurveyComponent } from './survey/survey/survey.component';
import { TraitAddEditComponent } from './trait/trait-add-edit/trait-add-edit.component';
import { TraitListComponent } from './trait/trait-list/trait-list.component';

const routes: Routes = [
  { path: 'questions', component: QuestionListComponent },
  { path: 'question/add', component: QuestionAddEditComponent },
  { path: 'question/edit/:id', component: QuestionAddEditComponent },
  { path: 'traits', component: TraitListComponent },
  { path: 'trait/add', component: TraitAddEditComponent },
  { path: 'trait/edit/:id', component: TraitAddEditComponent },
  { path: 'survey/:id', component: ReviewSurveyComponent },
  { path: '', component: SurveyComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

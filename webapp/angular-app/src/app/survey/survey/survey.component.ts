import { Component, OnInit } from '@angular/core';
import { FormArray, FormBuilder, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AlertService } from 'src/app/alert/service/alert.service';
import { Question } from 'src/app/question/question';
import { QuestionService } from 'src/app/question/service/question.service';
import { SurveyService } from '../service/survey.service';
import { Survey } from '../survey';

@Component({
  selector: 'app-survey',
  templateUrl: './survey.component.html',
  styleUrls: ['./survey.component.css']
})
export class SurveyComponent implements OnInit {
  searchForm = this.formBuilder.group({
    surveyId: ['', Validators.required]
  });
  form = this.formBuilder.group({
    answers: this.formBuilder.array([])
  });
  questions: Question[] | undefined;

  constructor(
    private service: SurveyService,
    private questionService: QuestionService,
    private alertService: AlertService,
    private formBuilder: FormBuilder,
    private router: Router,
  ) { }

  ngOnInit(): void {
    this.questionService.getQuestions().subscribe(
      {
        next: (data: Question[]) => {
          data.forEach(q => {
            const answer = this.formBuilder.group({
              questionId: [q.id, Validators.required],
              optionId: ['', Validators.required]
            });
            this.answers.push(answer);
          });
          this.questions = data;
        }
      }
    )
  }

  get answers() {
    return this.form.controls["answers"] as FormArray;
  }

  reviewSurvey() {
    const {surveyId} = this.searchForm.value;
    this.router.navigate(["survey", surveyId])
  }

  saveForm() {
    this.alertService.clear();
    this.service.submitSurvey(this.form.value).subscribe(
      {
        next: (data: Survey) => {
          this.router.navigate(["survey", data.id]);
        },
        error: (err) => {
          this.alertService.error(err.error)
        }
      }
    )
  }
}

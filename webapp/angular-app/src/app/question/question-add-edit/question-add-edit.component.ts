import { Component, OnInit } from '@angular/core';
import { FormArray, FormBuilder, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { AlertService } from 'src/app/alert/service/alert.service';
import { Question } from '../question';
import { QuestionService } from '../service/question.service';

@Component({
  selector: 'app-question-edit',
  templateUrl: './question-add-edit.component.html',
  styleUrls: ['./question-add-edit.component.css']
})
export class QuestionAddEditComponent implements OnInit {
  form = this.formBuilder.group({
    description: ['', Validators.required],
    options: this.formBuilder.array([]),
  });
  question: Question | undefined;
  id!: string | null;

  constructor(
    private service: QuestionService,
    private route: ActivatedRoute,
    private alertService: AlertService,
    private formBuilder: FormBuilder,
    private router: Router,
  ) { }

  ngOnInit(): void {
    this.id = this.route.snapshot.paramMap.get('id');
    if (this.id) {
      this.getQuestion(this.id);
    } else {
      this.addOption();
      this.addOption();
    }
  }

  getQuestion(id: string | null): void {
    if (id == null) {
      this.alertService.error("Please specify question id You wish to edit!");
      return
    }
    this.service.getQuestion(id).subscribe(
      {
        next: (data) => {
          this.question = data;
          // fill form
          this.form.controls["description"].setValue(this.question.description);
          this.question.options.forEach(option => {
            const optionForm = this.formBuilder.group({
              id: [option.id],
              description: [option.description, Validators.required],
              score: [option.score, Validators.min(0)]
            });
            this.options.push(optionForm);
          });
        },
        error: () => {
          this.alertService.error("Something went wrong!");
        },
        complete: () => {
        }
      }
    )
  }

  get options() {
    return this.form.controls["options"] as FormArray;
  }

  addOption() {
    const optionForm = this.formBuilder.group({
      description: ['', Validators.required],
      score: [0, Validators.min(0)]
    });
    this.options.push(optionForm);
  }

  deleteOption(optionIndex: number) {
    this.options.removeAt(optionIndex);
  }

  saveForm() {
    this.alertService.clear();
    if (this.question) {
      this.service.editQuestion(this.form.value, this.question.id).subscribe(
        {
          next: (data) => {
            this.router.navigate(["questions"]);
            this.alertService.success("Question saved successfully")
          },
          error: (err) => {
            this.alertService.error(err.error);
          }
        }
      );
    } else {
      this.service.addNewQuestion(this.form.value).subscribe(
        {
          next: (data) => {
            this.router.navigate(["questions"]);
            this.alertService.success("Question added successfully")
          },
          error: (err) => {
            this.alertService.error(err.error);
          }
        }
      );
    }
  }
}

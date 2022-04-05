import { Component, OnInit } from '@angular/core';
import { AlertService } from 'src/app/alert/service/alert.service';
import { QuestionService } from '../service/question.service';

@Component({
  selector: 'app-question-list',
  templateUrl: './question-list.component.html',
  styleUrls: ['./question-list.component.css']
})
export class QuestionListComponent implements OnInit {

  questions = this.service.getQuestions();

  constructor(private service: QuestionService, private alertService: AlertService) { }

  ngOnInit(): void {
  }

  deleteQuestion(questionId: string) {
    this.alertService.clear();
    this.service.deleteQuestion(questionId).subscribe(
      {
        error: () => {
          this.alertService.error("Something went wrong!");
        },
        complete: () => {
          this.questions = this.service.getQuestions();
          this.alertService.success("Successfully deleted question!");
        }
      }
    );
  }
}

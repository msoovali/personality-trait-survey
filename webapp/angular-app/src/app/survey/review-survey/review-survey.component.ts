import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { AlertService } from 'src/app/alert/service/alert.service';
import { SurveyService } from '../service/survey.service';
import { Survey } from '../survey';

@Component({
  selector: 'app-review-survey',
  templateUrl: './review-survey.component.html',
  styleUrls: ['./review-survey.component.css']
})
export class ReviewSurveyComponent implements OnInit {
  id!: string | null;
  survey: Survey | undefined;
  loading = true;

  constructor(
    private service: SurveyService,
    private route: ActivatedRoute,
    private alertService: AlertService,
  ) { }

  ngOnInit(): void {
    this.id = this.route.snapshot.paramMap.get('id');
    if (this.id) {
      this.getSurvey(this.id);
    } else {
      this.alertService.error("Survey ID not specified!");
      this.loading = false;
    }
  }

  getSurvey(id: string | null): void {
    if (id == null) {
      this.alertService.error("Please specify trait id You wish to edit!");
      return
    }
    this.service.getSurvey(id).subscribe(
      {
        next: (data) => {
          this.survey = data;
        },
        error: () => {
          this.alertService.error("Something went wrong!");
          this.loading = false;
        },
        complete: () => {
          this.loading = false;
        }
      }
    )
  }
}

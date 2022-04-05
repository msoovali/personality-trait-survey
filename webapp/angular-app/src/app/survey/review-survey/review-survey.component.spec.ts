import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReviewSurveyComponent } from './review-survey.component';

describe('ReviewSurveyComponent', () => {
  let component: ReviewSurveyComponent;
  let fixture: ComponentFixture<ReviewSurveyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReviewSurveyComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ReviewSurveyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

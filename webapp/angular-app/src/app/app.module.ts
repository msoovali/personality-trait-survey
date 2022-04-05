import { HttpClientModule } from '@angular/common/http';
import { APP_INITIALIZER, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { QuestionListComponent } from './question/question-list/question-list.component';
import { SettingsService } from './services/settings/settings.service';
import { QuestionAddEditComponent } from './question/question-add-edit/question-add-edit.component';
import { AlertComponent } from './alert/component/alert.component';
import { ReactiveFormsModule } from '@angular/forms';
import { TraitListComponent } from './trait/trait-list/trait-list.component';
import { TraitAddEditComponent } from './trait/trait-add-edit/trait-add-edit.component';
import { ReviewSurveyComponent } from './survey/review-survey/review-survey.component';
import { SurveyComponent } from './survey/survey/survey.component';

export function initApp(settingsService: SettingsService) {
  return () => settingsService.loadAppSettings();
}

@NgModule({
  declarations: [
    AppComponent,
    QuestionListComponent,
    QuestionAddEditComponent,
    AlertComponent,
    TraitListComponent,
    TraitAddEditComponent,
    ReviewSurveyComponent,
    SurveyComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule
  ],
  providers: [
    {
      provide: APP_INITIALIZER,
      deps: [SettingsService],
      multi: true,
      useFactory: initApp
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }

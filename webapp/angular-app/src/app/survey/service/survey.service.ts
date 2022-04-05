import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { SettingsService } from 'src/app/services/settings/settings.service';
import { Survey } from '../survey';

@Injectable({
  providedIn: 'root'
})
export class SurveyService {

  constructor(private http: HttpClient, private settings: SettingsService) { }

  getSurvey(id: string) {
    return this.http.get<Survey>(`${this.settings.SETTINGS.API_URL}/v1/survey/result/${id}`)
  }

  submitSurvey(survey: Survey) {
    const headers = new HttpHeaders({ "Content-Type": "application/json" });
    return this.http.post<Survey>(`${this.settings.SETTINGS.API_URL}/v1/survey/finish`, JSON.stringify(survey), { headers: headers })
  }
}

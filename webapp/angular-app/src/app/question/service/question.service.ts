import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { SettingsService } from 'src/app/services/settings/settings.service';
import { Question } from '../question';

@Injectable({
  providedIn: 'root'
})
export class QuestionService {

  constructor(private http: HttpClient, private settings: SettingsService) { }

  getQuestions() {
    return this.http.get<Question[]>(`${this.settings.SETTINGS.API_URL}/v1/questions`);
  }

  getQuestion(id: string) {
    return this.http.get<Question>(`${this.settings.SETTINGS.API_URL}/v1/question/${id}`)
  }

  deleteQuestion(id: string) {
    return this.http.delete(`${this.settings.SETTINGS.API_URL}/v1/question/${id}`);
  }

  addNewQuestion(question: Question) {
    const headers = new HttpHeaders({ "Content-Type": "application/json" });
    return this.http.post(`${this.settings.SETTINGS.API_URL}/v1/question`, JSON.stringify(question), { headers: headers });
  }

  editQuestion(question: Question, id: string) {
    question.id = id;
    const headers = new HttpHeaders({ "Content-Type": "application/json" });
    return this.http.put(`${this.settings.SETTINGS.API_URL}/v1/question`, JSON.stringify(question), { headers: headers });
  }
}

import { lastValueFrom } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class SettingsService {
  private settings: any;

  constructor(private http: HttpClient) {}

  loadAppSettings(): Promise<any> {
    return lastValueFrom(this.http.get('/assets/settings/settings.json')).then(data => this.settings = data);
  }

  get SETTINGS () {
    return this.settings;
  }
}

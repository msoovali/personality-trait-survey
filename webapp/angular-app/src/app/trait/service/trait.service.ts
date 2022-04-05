import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { SettingsService } from 'src/app/services/settings/settings.service';
import { Trait } from '../trait';

@Injectable({
  providedIn: 'root'
})
export class TraitService {

  constructor(private http: HttpClient, private settings: SettingsService) { }

  getTraits() {
    return this.http.get<Trait[]>(`${this.settings.SETTINGS.API_URL}/v1/traits`);
  }

  getTrait(id: string) {
    return this.http.get<Trait>(`${this.settings.SETTINGS.API_URL}/v1/trait/${id}`)
  }

  deleteTrait(id: string) {
    return this.http.delete(`${this.settings.SETTINGS.API_URL}/v1/trait/${id}`);
  }

  addNewTrait(trait: Trait) {
    const headers = new HttpHeaders({ "Content-Type": "application/json" });
    return this.http.post(`${this.settings.SETTINGS.API_URL}/v1/trait`, JSON.stringify(trait), { headers: headers });
  }

  editTrait(trait: Trait, id: string) {
    trait.id = id;
    const headers = new HttpHeaders({ "Content-Type": "application/json" });
    return this.http.put(`${this.settings.SETTINGS.API_URL}/v1/trait`, JSON.stringify(trait), { headers: headers });
  }
}

import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';
import { Alert, AlertType } from '../alert';

@Injectable({
  providedIn: 'root'
})
export class AlertService {
  private subject = new Subject<Alert>();

  constructor() { }

  onAlert(): Observable<Alert> {
    return this.subject.asObservable();
  }

  alert(alert: Alert) {
    this.subject.next(alert);
  }

  success(message: string) {
    this.alert({ type: AlertType.Success, message });
  }

  error(message: string) {
    this.alert({ type: AlertType.Error, message });
  }

  clear() {
    this.subject.next({ type: AlertType.Clear, message: "" });
  }
}

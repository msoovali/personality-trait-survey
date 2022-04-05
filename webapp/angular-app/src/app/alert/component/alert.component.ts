import { Component, OnDestroy, OnInit } from '@angular/core';
import { NavigationStart, Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { Alert, AlertType } from '../alert';
import { AlertService } from '../service/alert.service';

@Component({
  selector: 'app-alert',
  templateUrl: './alert.component.html',
  styleUrls: ['./alert.component.css']
})
export class AlertComponent implements OnDestroy {

  alerts: Alert[] = [];
  alertSubscription: Subscription;
  routeSubscription: Subscription;

  constructor(private router: Router, private service: AlertService) {
    this.alertSubscription = this.service.onAlert().subscribe(
      alert => {
        if (alert) {
          if (alert.type == AlertType.Clear) {
            this.alerts = [];
          } else {
            this.alerts.push(alert);
          }
        } else {
          this.alerts = [];
        }
      }
    );

    this.routeSubscription = this.router.events.subscribe(event => {
      if (event instanceof NavigationStart) {
        this.service.clear()
      }
    });
  }

  ngOnDestroy(): void {
    this.alertSubscription.unsubscribe();
    this.routeSubscription.unsubscribe();
  }

  cssClass(alert: Alert) {
    if (!alert || alert.type == AlertType.Clear) {
      return;
    }

    const classes = ['alert'];

    const alertTypeClass = {
      [AlertType.Success]: 'alert-success',
      [AlertType.Error]: 'alert-danger',
    }

    classes.push(alertTypeClass[alert.type])

    return classes.join(' ');
  }

}

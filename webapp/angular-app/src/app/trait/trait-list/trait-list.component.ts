import { Component, OnInit } from '@angular/core';
import { AlertService } from 'src/app/alert/service/alert.service';
import { TraitService } from '../service/trait.service';

@Component({
  selector: 'app-trait-list',
  templateUrl: './trait-list.component.html',
  styleUrls: ['./trait-list.component.css']
})
export class TraitListComponent implements OnInit {

  traits = this.service.getTraits();

  constructor(private service: TraitService, private alertService: AlertService) { }

  ngOnInit(): void {
  }

  deleteTrait(traitId: string) {
    this.alertService.clear();
    this.service.deleteTrait(traitId).subscribe(
      {
        error: () => {
          this.alertService.error("Something went wrong!");
        },
        complete: () => {
          this.traits = this.service.getTraits();
          this.alertService.success("Successfully deleted trait!");
        }
      }
    );
  }
}

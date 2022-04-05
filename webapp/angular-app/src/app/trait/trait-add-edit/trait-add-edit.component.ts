import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { AlertService } from 'src/app/alert/service/alert.service';
import { TraitService } from '../service/trait.service';
import { Trait } from '../trait';

@Component({
  selector: 'app-trait-add-edit',
  templateUrl: './trait-add-edit.component.html',
  styleUrls: ['./trait-add-edit.component.css']
})
export class TraitAddEditComponent implements OnInit {
  form = this.formBuilder.group({
    type: ['', Validators.required],
    minScoreRequirement: [0, Validators.min(0)],
  });
  trait: Trait | undefined;
  id!: string | null;

  constructor(
    private service: TraitService,
    private route: ActivatedRoute,
    private alertService: AlertService,
    private formBuilder: FormBuilder,
    private router: Router,
  ) { }

  ngOnInit(): void {
    this.id = this.route.snapshot.paramMap.get('id');
    if (this.id) {
      this.getTrait(this.id);
    }
  }

  getTrait(id: string | null): void {
    if (id == null) {
      this.alertService.error("Please specify trait id You wish to edit!");
      return
    }
    this.service.getTrait(id).subscribe(
      {
        next: (data) => {
          this.trait = data;
          // fill form
          this.form.controls["type"].setValue(this.trait.type);
          this.form.controls["minScoreRequirement"].setValue(this.trait.minScoreRequirement);
        },
        error: () => {
          this.alertService.error("Something went wrong!");
        },
        complete: () => {
        }
      }
    )
  }

  saveForm() {
    this.alertService.clear();
    if (this.trait) {
      this.service.editTrait(this.form.value, this.trait.id).subscribe(
        {
          next: (data) => {
            this.router.navigate(["traits"]);
            this.alertService.success("Trait saved successfully")
          },
          error: (err) => {
            this.alertService.error(err.error);
          }
        }
      );
    } else {
      this.service.addNewTrait(this.form.value).subscribe(
        {
          next: (data) => {
            this.router.navigate(["traits"]);
            this.alertService.success("Trait added successfully")
          },
          error: (err) => {
            this.alertService.error(err.error);
          }
        }
      );
    }
  }
}

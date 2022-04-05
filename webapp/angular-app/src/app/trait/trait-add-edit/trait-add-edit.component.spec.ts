import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TraitAddEditComponent } from './trait-add-edit.component';

describe('TraitAddEditComponent', () => {
  let component: TraitAddEditComponent;
  let fixture: ComponentFixture<TraitAddEditComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TraitAddEditComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TraitAddEditComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

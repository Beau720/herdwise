import { Component } from '@angular/core';
import { ReactiveFormsModule, FormsModule, } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { Location } from '@angular/common';
@Component({
  selector: 'app-help-desk',
  standalone: true,
  imports: [CommonModule,FormsModule],
  templateUrl: './help-desk.component.html',
  styleUrl: './help-desk.component.css'
})
export class HelpDeskComponent {

  constructor(private router: Router,private location: Location) {}
  
  goBack(){
    this.location.back();
  }

}

import { Component } from '@angular/core';
import { ReactiveFormsModule, FormsModule, } from '@angular/forms';
import { CommonModule ,NgOptimizedImage} from '@angular/common';
import { Router } from '@angular/router';
import { Farmer } from '../models/farmer';
import { Location } from '@angular/common';
//import { FarmersService } from '../services/farmers.service';
@Component({
  selector: 'app-update-profile',
  standalone: true,
  imports: [ReactiveFormsModule, FormsModule, CommonModule],
  templateUrl: './update-profile.component.html',
  styleUrl: './update-profile.component.css'
})
export class UpdateProfileComponent {
  isEditMode = false;
  user = {
    name: 'John Doe',
    email: 'john.doe@example.com',
    contact: '+123456789'
  };
  name: string = '';
  email: string = '';
  password: string = '';
  errorMessage: string = '';
  success: string = '';
  farmer: Farmer ={
    farmerID: 0,
    first_name: '',
    last_name: '',
    email: '',
    contact: '',
    password: '',
    latFarm: '',
    longFarm: '',
    size: 0
  }
  constructor(private router: Router ,private location: Location) { }

  ngOnInit(): void { 
   // user$ = this.authService.getFarmer(this.authService.getUserIdFromToken())
   }
  // Toggle between edit and view modes
  toggleEditMode() {
    this.isEditMode = !this.isEditMode;
  }

  // Function to update the profile
  updateProfile() {
    // this.farmer.farmerID = this.authService.getUserIdFromToken();
    //   this.fService.putFarmer(1, this.farmer).subscribe({
    //     next:(updateFarmer: Farmer) => {
    //       this.success = 'Farmer updated successfully';
    //       this.farmer = updateFarmer;
    //       this.router.navigate(['/ViewFarmer']);
    //     },
    //     error:(err: Farmer) => {
    //       this.error = 'Error updating  user data';
    //       console.error(err);
    //     }
    //     })
    this.isEditMode = false;
  }

  goBack(){
    this.location.back();
  }

}

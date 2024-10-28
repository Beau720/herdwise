import { Component } from '@angular/core';
import { ReactiveFormsModule, FormsModule, } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { Farmer } from '../models/farmer';
import { Location } from '@angular/common';
import { FarmersService } from '../services/farmers.service';
import { AuthService } from '../services/auth.service';
@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [ReactiveFormsModule, FormsModule,CommonModule],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css'
})
export class ProfileComponent {
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
  
  editableFields: Record<EditableField, boolean> = {
    first_name: false,
    last_name: false,
    email: false,
    contact: false,
    latFarm: false,
    longFarm: false,
    size: false
  };

  anyFieldEditable = false;


  constructor( private fService: FarmersService,
    private authService: AuthService,
    private router: Router ,
    private location: Location) { }
    ngOnInit(): void {
      this.getFarmerDetails();
    }

    getFarmerDetails(){
      var farmerID = this.authService.getUserIdFromToken();
      console.log("Farmer Id: " ,farmerID)
    this.fService.getFarmer(farmerID).subscribe({
      next: (data: Farmer) => {
        console.log("Data received: ", data); 
        this.farmer = data;
        this.success = 'User is here successfully';
      },
      error: (err: any) => {
        this.errorMessage = 'Error fetching employee data';
        console.error("Error occurred: ", err);
      }
    })
    }
    
      // Toggle the edit state for a specific field
  toggleEdit(field: EditableField) {
    if (this.editableFields[field]) {
      // Save the changes if the field was in editable mode
      this.saveField(field);
    }
    this.editableFields[field] = !this.editableFields[field];
    this.anyFieldEditable = Object.values(this.editableFields).some(value => value); // Check if any field is editable
  }
  
    saveField(field: string) {
      this.farmer.farmerID = this.authService.getUserIdFromToken();
            this.fService.putFarmer(this.farmer.farmerID, this.farmer).subscribe({
              next:(updateFarmer: Farmer) => {
                this.success = 'Farmer updated successfully';
                this.farmer = updateFarmer;
                this.router.navigate(['/ViewFarmer']);
              },
              error:(err: Farmer) => {
                this.errorMessage = 'Error updating  user data';
                console.error(err);
              }
              })
    }
  
    cancelEdit() {
      this.getFarmerDetails();  // Reset form to original data
      this.editableFields = {
        first_name: false,
        last_name: false,
        email: false,
        contact: false,
        latFarm: false,
        longFarm: false,
        size: false
      };
      this.anyFieldEditable = false;
    }

    goBack(){
      this.location.back();
    }

}
type EditableField = 'first_name' | 'last_name' | 'email' | 'contact' | 'latFarm' | 'longFarm' | 'size';

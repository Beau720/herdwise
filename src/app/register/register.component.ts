import { Component } from '@angular/core';
import { ReactiveFormsModule, FormsModule, } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { Farmer } from '../models/farmer';
import { Location } from '@angular/common';
import { FarmersService } from '../services/farmers.service';
@Component({
  selector: 'app-register',
  standalone: true,
  imports: [ReactiveFormsModule, FormsModule,CommonModule],
  templateUrl: './register.component.html',
  styleUrl: './register.component.css'
})
export class RegisterComponent {

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
  confPassword: any;
  showPassword = false;
  showConfirmPassword = false;
  passwordMismatchMessage = '';
  constructor( private fService: FarmersService,
    private router: Router ,
    private location: Location) { }

  ngOnInit(): void {  }
  register(){

    this.fService.postFarmer(this.farmer).subscribe({
      next: (data: Farmer) => {
        console.log("my data passsed:",data)
        this.farmer = data;
        this.success = 'User created successfully';
        this.router.navigate(['/login']);
      },
      error: (err) => {
        this.errorMessage = 'Error creating user';
        console.error(err);
        
      }
    });
  }

  // Toggle password visibility
  togglePasswordVisibility() {
    this.showPassword = !this.showPassword;
  }

  // Toggle confirm password visibility
  toggleConfirmPasswordVisibility() {
    this.showConfirmPassword = !this.showConfirmPassword;
  }

  // Check if the password and confirmPassword fields match
  passwordsMatch(): boolean {
    return this.farmer.password === this.confPassword;
  }

  goBack(){
    this.location.back();
  }
  

}

import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, NgForm } from '@angular/forms';
import { ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { Farmer } from '../models/farmer';
import { AuthService } from '../services/auth.service';
import { FarmersService } from '../services/farmers.service';
@Component({
  selector: 'app-login',
  standalone: true,
  imports: [CommonModule, FormsModule,ReactiveFormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {
  success: string| undefined;
  error: string | undefined;
  userId : number | undefined;
  showPassword = false;
  farmer : Farmer ={
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

  constructor(
    private authservice: AuthService,
    private router: Router
  ){}

  onSubmit(loginForm: NgForm) {
    if (loginForm.valid) {
      //You can add your login logic here, e.g., send data to a backend server.
      this.authservice.login(this.farmer.email ,this.farmer.password ).subscribe({
          next:( result) => {
          this.authservice.setToken(result.token);
          this.success = "Successfully logged in"
          // Navigate to dashboard or home page after successful login
          this.router.navigate(['/profile'])
          }
        },
        
      );
    } else {
      console.log('Form is not valid');
    }
  }

  togglePasswordVisibility() {
    this.showPassword = !this.showPassword;
  }

  goRegister(){
    this.router.navigate(['/register']);
  }

}

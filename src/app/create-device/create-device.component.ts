import { Component } from '@angular/core';
import { ReactiveFormsModule, FormsModule, } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { Farmer } from '../models/farmer';
import { Device } from '../models/device';
import { Location } from '@angular/common';
import { routes } from '../app.routes';
import { DeviceService } from '../services/device.service';
import { AuthService } from '../services/auth.service';
@Component({
  selector: 'app-create-device',
  standalone: true,
  imports: [ReactiveFormsModule, FormsModule,CommonModule],
  templateUrl: './create-device.component.html',
  styleUrl: './create-device.component.css'
})
export class CreateDeviceComponent {
  errorMessage: string = '';
  success: string = '';
  
  device : Device ={
    deviceId: 0,
    ref: '',
    long: '',
    lati: '',
    temp: '',
    type: '',
    farmerId: 0,
    highTemp: '',
    lowTemp: '',
    model: ''
  }

  constructor(private authService: AuthService,
    private dService: DeviceService,
    private router: Router ,
    private location: Location) {
      this.device.farmerId = 1
     }


addDevice() {
  this.dService.postDevice(this.device).subscribe({
      next: (data: Device) => {
        this.device = data;
        this.success = 'Device Added Successfully';
        this.router.navigate(['/all-devices']);
      },
      error: (err) => {
        this.errorMessage = 'Error Creating Device';
        console.error(err);
        
      }
    });
  
}

goBack(){
  this.location.back();
}


}

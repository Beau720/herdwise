import { Component } from '@angular/core';
import { ReactiveFormsModule, FormsModule, } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { Device } from '../models/device';
import { Location } from '@angular/common';
import { DeviceService } from '../services/device.service';
import { AuthService } from '../services/auth.service';
import { Observable } from 'rxjs';


@Component({
  selector: 'app-devices',
  standalone: true,
  imports: [ReactiveFormsModule,FormsModule,CommonModule],
  templateUrl: './devices.component.html',
  styleUrl: './devices.component.css'
})
export class DevicesComponent {
  selectedComponent: string = 'devices'; 
  success?: string;
  error?: string;
  errorMessage?: string;
  i= 0;
  selectComponent(component: string) {
    this.selectedComponent = component;
  }

  isEditable: boolean = false;
  editableStates: boolean[] = [];
  originalDevice: Partial<Device>[] = []; 
  devices: Device[] = [];
  device: Device ={
    deviceId: 0,
    ref: '',
    long: '',
    lati: '',
    temp: '',
    type: '',
    highTemp: '',
    lowTemp: '',
    model: '',
    farmerId: 0
    
  }

  devices$: Observable<Device[]> | undefined;
  constructor( private dService: DeviceService,
    //private fService: AuthService,
    private router: Router,
    private location: Location) {}

  ngOnInit(): void { 
     this.fetchDevice();
    // this.devices$ = this.dService.getDevices(1)
    // this.devices$.subscribe(
    //   (data: Device[]) => console.log('Users Data:', data),
    //   error => console.error('Error fetching users data:', error)
    // );
     }

 
     editDevice(index: number) {
      const device = this.devices[index];
      if (!this.editableStates[index]) {
        // Store the original data if starting edit mode
        this.originalDevice[index] = { ...device };
        this.editableStates[index] = true;
      } else {
        // Save changes when toggling off edit mode
        this.UpdateDevice(device.deviceId, device);
        this.editableStates[index] = false;
      }
    }
  cancelEdit(index: number) {
    // Restore the original data and exit edit mode
    this.devices[index] = { ...this.originalDevice[index] } as Device;
    this.editableStates[index] = false;
  }
  fetchDevice(){
    //const userId = this.authService.getUserId();
    this.dService.getDevices(1).subscribe({
      next: (data: Device[]) => {
        this.devices = data;
        this.editableStates = new Array(this.devices.length).fill(false); 
        this.originalDevice = new Array(this.devices.length).fill(null); 
        this.success = 'User is here successfully';
      },
      error: (err: any) => {
        this.error = 'Error fetching employee data';
        console.error(err);
      }
    })
  }
  UpdateDevice(deviceId: number,device: Device) {
      this.dService.putDevice(deviceId,device).subscribe({
        next:(data: Device) => {
          this.success = 'Device  updated successfully';
          this.device = data;
          this.router.navigate(['/all-devices']);
        },
        error:(err: Device) => {
          this.error = 'Error updating  user data';
          console.error(err);
        }
        })
    }

    addDevice() {
        this.router.navigate(['/add-device']);
      }

  goBack(){
    this.location.back();
  }

}

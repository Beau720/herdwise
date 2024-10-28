import { Component } from '@angular/core';
import { MapComponent } from "../map/map.component";
import { DevicesComponent } from "../devices/devices.component";
import { HelpDeskComponent } from '../help-desk/help-desk.component';
import { ProfileComponent } from '../profile/profile.component';
import { ReactiveFormsModule, FormsModule, } from '@angular/forms';
import { CommonModule ,NgOptimizedImage} from '@angular/common';
import { RouterModule, Router } from '@angular/router';
import { filter } from 'rxjs';
import { AuthService } from '../services/auth.service';


@Component({
  selector: 'app-home',
  standalone: true,
  imports: [MapComponent,HelpDeskComponent, DevicesComponent,ProfileComponent,
    NgOptimizedImage,ReactiveFormsModule, FormsModule,CommonModule,RouterModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {
  showHeader: boolean = false;
  notificationMessage: string | null = null;
  selectedComponent: string = 'home';  // Default selection
 constructor(private router: Router ,
  private authService: AuthService,
 ){}

 ngOnInit(): void {

}
  selectComponent(component: string) {
    this.selectedComponent = component;
  }
  isNavOpen = true; // Initially the navigation is closed

  toggleNav() {
    this.isNavOpen = !this.isNavOpen; // Toggle the value of isNavOpen
  }
  goHelp(){
    this.router.navigate(['/help']);
  }
  goDevice(){
    this.router.navigate(['/all-devices']);
  }
  goMap(){
    this.router.navigate(['/map']);
  }
  goProfile(){
    this.router.navigate(['/profile']);
  }

  signOut(){
    this.authService.logout();
    this.router.navigate(['/login']);
  }

  showNotification(message: string): void {
    this.notificationMessage = message;  // Set the notification message
    setTimeout(() => {
      this.notificationMessage = null;   // Clear the notification after 3 seconds
    }, 3000);
  }


}

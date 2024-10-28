import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
@Component({
  selector: 'app-editdevice',
  standalone: true,
  imports: [],
  templateUrl: './editdevice.component.html',
  styleUrl: './editdevice.component.css'
})
export class EditdeviceComponent {
  selectedComponent: string = 'Edit-Device'; 
  selectComponent(component: string) {
    this.selectedComponent = component;
  }

  userId!: number;

  constructor(private route: ActivatedRoute) {}

  ngOnInit(): void {
    // Get the user ID from the route
    this.userId = +this.route.snapshot.paramMap.get('id')!;
    // Fetch the user details based on this ID (e.g., via a service)
    console.log('Editing user with ID:', this.userId);
  }
}

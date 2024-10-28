import { Component, OnInit,Inject, PLATFORM_ID  } from '@angular/core';
import { GoogleMapsModule } from '@angular/google-maps';
import { isPlatformBrowser } from '@angular/common';

@Component({
  selector: 'app-map',
  standalone: true,
  imports: [GoogleMapsModule],
  templateUrl: './map.component.html',
  styleUrl: './map.component.css'
})
export class MapComponent implements OnInit {


  center: google.maps.LatLngLiteral = { lat: -25.7479, lng: 28.2293 };
  zoom = 15;
  polygon!: google.maps.Polygon ;
  marker!: google.maps.Marker;
  private circle!: google.maps.Circle;
  constructor(@Inject(PLATFORM_ID) private platformId: Object) {}
   ngOnInit(): void { 
    if (isPlatformBrowser(this.platformId)) {
      this.loadMap();
      this.startTrackingPetLocation();
    }
  }
  loadMap(): void {
    const map = new google.maps.Map(document.getElementById('map') as HTMLElement, {
      zoom: this.zoom,
      center: this.center,
      mapTypeId: google.maps.MapTypeId.TERRAIN,
    });
  
    this.marker = new google.maps.Marker({
      position: this.center,
      map: map,
      title: 'Pet Location',
      draggable: true, // Make marker not draggable
      
    });
  
    // Define a circle with a given center and radius
    this.circle = new google.maps.Circle({
      strokeColor: '#FF0000',
      strokeOpacity: 0.9,
      strokeWeight: 2,
      fillColor: '#000000',
      fillOpacity: 0.35,
      map: map,
      center: this.center,
      radius: 50, // Radius in meters
    });
  
    this.circle.setMap(map);
  }
  startTrackingPetLocation(): void {
    // Example: Update the marker position every 5 seconds with new random coordinates
    setInterval(() => {
      const newPosition = this.getNewPetLocation();
      this.updateMarkerPosition(newPosition);
    }, 5000); // Update every 5 seconds
  }

  getNewPetLocation(): google.maps.LatLngLiteral {
    // Simulating a new random position (this would be replaced with actual GPS data)
    const randomLat = this.center.lat + (Math.random() - 0.5) * 0.001; // Randomly adjust latitude
    const randomLng = this.center.lng + (Math.random() - 0.5) * 0.001; // Randomly adjust longitude
    return { lat: randomLat, lng: randomLng };
  }
  
  updateMarkerPosition(newPosition: google.maps.LatLngLiteral): void {
    this.marker.setPosition(newPosition); // Update marker position
  
    const position = new google.maps.LatLng(newPosition.lat, newPosition.lng);
    const circleCenter = this.circle.getCenter();  // Get the center of the circle
  
    if (circleCenter) {  // Check if the center is not null
      const isInside = google.maps.geometry.spherical.computeDistanceBetween(position, circleCenter) <= this.circle.getRadius();
  
      console.log('Marker Position:', position.toString()); // Log marker position for debugging
      console.log('Is Inside Circle:', isInside); // Log check result for debugging
  
      if (!isInside) {
        window.alert("Cattle are outside of the circle");
      }
    } else {
      console.error('Circle center is not defined');
    }
  }
  
}


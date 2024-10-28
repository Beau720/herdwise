import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable} from 'rxjs';
import { Device } from '../models/device';

@Injectable({
  providedIn: 'root'
})
export class DeviceService {

  readonly apiUrl = 'http://localhost:8085';

  constructor(private http: HttpClient) { }

  postDevice(device: Device): Observable<Device> {
    return this.http.post<Device>(`${this.apiUrl}/device/create`, device, {headers: this.getHeaders()});
  }

  getDevices(farmerID: number): Observable<Device[]> {
    return this.http.get<Device[]>(`${this.apiUrl}/device/list/${farmerID}`, { headers: this.getHeaders() });
  }

  getDeviceRef(ref: string): Observable<Device> {
    return this.http.get<Device>(`${this.apiUrl}/device/ref/${ref}`, { headers: this.getHeaders() });
  }

  getLocation(deviceID: number): Observable<{ lat: string; lng: string }> {
    return this.http.get<{ lat: string; lng: string }>(`${this.apiUrl}/device/deviceID/${deviceID}`, { headers: this.getHeaders() });
  }
  getTemp(deviceID: number): Observable<{ temp: string}> {
    return this.http.get<{ temp: string }>(`${this.apiUrl}/device/deviceID/${deviceID}`, { headers: this.getHeaders() });
  }

  putDevice(deviceId: number,device: Device): Observable<Device> {
    return this.http.put<Device>(`${this.apiUrl}/device/update/${deviceId}`,device, { headers: this.getHeaders() });
  }

  private getHeaders(): HttpHeaders {
    return new HttpHeaders({'Content-Type': 'application/json'});
  }
}

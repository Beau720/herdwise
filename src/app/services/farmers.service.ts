import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable} from 'rxjs';
import { Farmer } from '../models/farmer';
@Injectable({
  providedIn: 'root'
})
export class FarmersService {
  readonly apiUrl = 'http://localhost:8085';

  constructor(private http: HttpClient) { }

  postFarmer(farmer: Farmer): Observable<Farmer> {
    return this.http.post<Farmer>(`${this.apiUrl}/farmer/create`, farmer, {headers: this.getHeaders()});
  }

  getFarmer(farmerID: number): Observable<Farmer> {
    return this.http.get<Farmer>(`${this.apiUrl}/farmer/farmerId/${farmerID}`, { headers: this.getHeaders() });
  }

  putFarmer(farmerID:number, farmer: Farmer): Observable<Farmer> {
    return this.http.put<Farmer>(`${this.apiUrl}/farmer/update/${farmerID}`,farmer, { headers: this.getHeaders() });
  }

  private getHeaders(): HttpHeaders {
    return new HttpHeaders({'Content-Type': 'application/json'});
  }
}

import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, BehaviorSubject } from 'rxjs';
import {jwtDecode} from 'jwt-decode';
@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private apiUrl = 'http://localhost:8085';
  private tokenSubject = new BehaviorSubject<string | null>(null);
  token$ = this.tokenSubject.asObservable();
  constructor(private http: HttpClient) { }

  

  //loging api 
  login(email: string, password: string): Observable<{ token: string; userID: number }> {
    return this.http.post<{ token: string; userID: number }>(`${this.apiUrl}/farmer/login`,{email,password},{headers: this.getAuthHeaders()});
  }

  // Save JWT token to local storage
  setToken(token: string): void {
    if (typeof window !== 'undefined'&& window.localStorage) {
     localStorage.setItem('authToken', token);
    }
     this.tokenSubject.next(token);
    
   }

  // Get token from local storage
  getToken(): string | null {
    if (typeof window !== 'undefined'&& window.localStorage) {
    return localStorage.getItem('authToken');
    }
    return null;
 }
 // Set headers with JWT for authenticated requests
 getAuthHeaders(): HttpHeaders {
  const token = this.getToken();
  return new HttpHeaders({
    'Authorization': `Bearer ${token}`
  });
}

// Remove token (logout)
logout(): void {
  localStorage.removeItem('authToken');
  this.tokenSubject.next(null);
}


getUserIdFromToken(): number {
  const token = localStorage.getItem('authToken');
  if (token) {
    const decoded: any = jwtDecode(token);
    console.log("the decode is",decoded);
    return decoded.id;  // Extract userId from the token payload
  }
  return 0;
}



}

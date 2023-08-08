import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  getMessages(room: string): Observable<Message[]> {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
      'Accept': 'application/json',
    });
    const body = { Room: room };
    return this.http.post<Message[]>(`${this.apiUrl}/message`, body, { headers });
  }

  postMessage(userId: number, room: string, message: string): Observable<void> {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
    });    

    const body = { UserID: userId, Room: room };
    return this.http.post<void>(`${this.apiUrl}/message/stock=${message}`, body, { headers });
  }
}

export interface Message {
  ID: number;
  UserName: string;
  Message: string;
  Room: string;
  Timestamp: string; 
}
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private apiUrl = 'http://localhost:8080/message';

  constructor(private http: HttpClient) {}

  getMessages(userId: number, room: string): Observable<string[]> {
    const url = `${this.apiUrl}/message?stock=${userId}&Room=${room}`;
    return this.http.get<string[]>(url);
  }

  postMessage(userId: number, room: string, message: string): Observable<void> {
    const url = `${this.apiUrl}/message`;
    const body = { UserID: userId, Room: room, Message: message };
    return this.http.post<void>(url, body);
  }
 
}
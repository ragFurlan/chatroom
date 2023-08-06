import { Component, OnInit } from '@angular/core';
import { ApiService } from '../services/api.service';

@Component({
  selector: 'app-chat',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
  
})
export class ChatComponent implements OnInit {
  users = [
    { id: 1, name: 'User 1' },
    { id: 2, name: 'User 2' }
  ];
  
  rooms = ['ROOM1', 'ROOM2'];
  selectedUser: number = this.users[0].id; 
  selectedRoom: string = this.rooms[0];    
  messages: string[] = [];

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {
    this.updateMessages();
  }

  updateMessages(): void {
    this.apiService.getMessages(this.selectedUser, this.selectedRoom)
      .subscribe((messages: string[]) => {
        this.messages = messages;
      });
  }

  sendMessage(message: string): void {
    this.apiService.postMessage(this.selectedUser, this.selectedRoom, message)
      .subscribe(() => {
        this.updateMessages();
      });
  }
}
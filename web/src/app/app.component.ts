import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { ApiService, Message } from '../services/api.service';

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
  formattedMessages: string[] = [];

  constructor(private apiService: ApiService) {}
  
  @ViewChild('messageInput') messageInputRef!: ElementRef<HTMLInputElement>;

  ngOnInit(): void {
    //this.updateMessages();
  } 

  updateMessages(): void {
    this.apiService.getMessages(this.selectedRoom)    
      .subscribe((messages: Message[]) => {
        this.formattedMessages = messages.map(message => {
          const timestamp = new Date(message.Timestamp).toLocaleString();
          return `${timestamp} - ${message.UserName} - ${message.Message}`;
        });
      });
  }

  sendMessage(message: string): void {
    this.apiService.postMessage(this.selectedUser, this.selectedRoom, message)
      .subscribe(() => {
        debugger;
        this.updateMessages();
        this.messageInputRef.nativeElement.value = '';
      });
  }
}

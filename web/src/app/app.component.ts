import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { ApiService, Message } from '../services/api.service';

@Component({
  selector: 'app-chat',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
  
})
export class ChatComponent implements OnInit {
  users = [
    { id: "1", name: 'User 1' },
    { id: "2", name: 'User 2' }
  ];

  rooms = [
    { id: "ROOM1", name: 'ROOM 1' },
    { id: "ROOM2", name: 'ROOM 2' }
  ];
  
 
  selectedUser: string = this.users[0].id; 
  selectedRoom: string = this.rooms[0].id; 
  formattedMessages: string[] = [];

  constructor(private apiService: ApiService) {}
  
  @ViewChild('messageInput') messageInputRef!: ElementRef<HTMLInputElement>;

  ngOnInit(): void {
    this.updateMessages();
  } 

  updateMessages(): void {
    this.apiService.getMessages(this.selectedRoom)    
      .subscribe((messages: Message[]) => {
        this.formattedMessages = [];
        if (messages != null) {
          this.formattedMessages = messages.map(message => {
            const timestamp = new Date(message.Timestamp).toLocaleString();
            return `${timestamp} - ${message.UserName} - ${message.Message}`;
          });
      }
      },
      () => {
        this.formattedMessages = [];
      });
  }

  sendMessage(message: string): void {
    console.log("message: "+ message +" user: " + this.selectedUser +" room: ", this.selectedRoom)
    this.apiService.postMessage(this.selectedUser, this.selectedRoom, message)
      .subscribe(() => {      
        this.updateMessages();
        this.messageInputRef.nativeElement.value = '';
        
      },
      (err) => {
        console.error("error: "+ JSON.stringify(err))
        this.formattedMessages = [];
      });
  }
}

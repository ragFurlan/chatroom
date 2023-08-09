import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms'; 

import { AppRoutingModule } from './app-routing.module';
import { ChatComponent } from './app.component'; 
import { ApiService } from '../services/api.service';

@NgModule({
  declarations: [
    ChatComponent 
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule 
  ],
  providers: [ApiService],
  bootstrap: [ChatComponent]
})
export class AppModule { }

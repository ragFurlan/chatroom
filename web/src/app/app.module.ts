import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms'; // Importe este módulo

import { AppRoutingModule } from './app-routing.module';
import { ChatComponent } from './app.component'; 

@NgModule({
  declarations: [
    ChatComponent // Adicione o componente de chat aqui
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule // Adicione este módulo
  ],
  providers: [],
  bootstrap: [ChatComponent]
})
export class AppModule { }

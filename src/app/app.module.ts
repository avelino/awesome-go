import { NgModule } from '@angular/core';
import { BrowserModule, HammerModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { IgxCardModule, IgxButtonModule, IgxRippleModule, IgxIconModule, IgxGridModule, IgxListModule, IgxAvatarModule, IgxToggleModule, IgxDialogModule, IgxCheckboxModule, IgxSwitchModule } from 'igniteui-angular';
import { IgxCategoryChartModule } from 'igniteui-angular-charts';
import { FormsModule } from '@angular/forms';
import { TeamComponent } from './team/team.component';
import { EventsComponent } from './events/events.component';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent,
    DashboardComponent,
    TeamComponent,
    EventsComponent
  ],
  imports: [
    BrowserModule,
    HammerModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    IgxCardModule,
    IgxButtonModule,
    IgxRippleModule,
    IgxIconModule,
    IgxGridModule,
    IgxCategoryChartModule,
    FormsModule,
    IgxListModule,
    IgxAvatarModule,
    IgxToggleModule,
    IgxDialogModule,
    IgxCheckboxModule,
    IgxSwitchModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}

import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subject, takeUntil } from 'rxjs';
import { NewTeamMembersListType } from './models/dashboard-data/new-team-members-list-type';
import { DashboardDataService } from './services/dashboard-data.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit, OnDestroy {
  private destroy$: Subject<void> = new Subject<void>();
  public listItemVisible: boolean = false;
  public dashboardDataNewTeamMembersList: NewTeamMembersListType[] = [];
  public checked: boolean = true;

  constructor(
    private dashboardDataService: DashboardDataService,
  ) {}

  ngOnInit() {
    this.dashboardDataService.getNewTeamMembersListList().pipe(takeUntil(this.destroy$)).subscribe({
      next: (data) => this.dashboardDataNewTeamMembersList = data,
      error: (_err: any) => this.dashboardDataNewTeamMembersList = []
    });
  }

  ngOnDestroy() {
    this.destroy$.next();
    this.destroy$.complete();
  }
}

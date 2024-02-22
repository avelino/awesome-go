import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subject, takeUntil } from 'rxjs';
import { AllTeamMembersType } from '../models/dashboard-data/all-team-members-type';
import { DashboardDataService } from '../services/dashboard-data.service';

@Component({
  selector: 'app-team',
  templateUrl: './team.component.html',
  styleUrls: ['./team.component.scss']
})
export class TeamComponent implements OnInit, OnDestroy {
  private destroy$: Subject<void> = new Subject<void>();
  public dashboardDataAllTeamMembers: AllTeamMembersType[] = [];

  constructor(
    private dashboardDataService: DashboardDataService,
  ) {}

  ngOnInit() {
    this.dashboardDataService.getAllTeamMembersList().pipe(takeUntil(this.destroy$)).subscribe({
      next: (data) => this.dashboardDataAllTeamMembers = data,
      error: (_err: any) => this.dashboardDataAllTeamMembers = []
    });
  }

  ngOnDestroy() {
    this.destroy$.next();
    this.destroy$.complete();
  }
}

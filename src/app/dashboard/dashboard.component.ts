import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subject, takeUntil } from 'rxjs';
import { TeamGrowthType } from '../models/dashboard-data/team-growth-type';
import { TeamMembersGridType } from '../models/dashboard-data/team-members-grid-type';
import { DashboardDataService } from '../services/dashboard-data.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit, OnDestroy {
  private destroy$: Subject<void> = new Subject<void>();
  public dashboardDataTeamMembersGrid: TeamMembersGridType[] = [];
  public dashboardDataTeamGrowth: TeamGrowthType[] = [];

  constructor(
    private dashboardDataService: DashboardDataService,
  ) {}

  ngOnInit() {
    this.dashboardDataService.getTeamMembersGridList().pipe(takeUntil(this.destroy$)).subscribe({
      next: (data) => this.dashboardDataTeamMembersGrid = data,
      error: (_err: any) => this.dashboardDataTeamMembersGrid = []
    });
    this.dashboardDataService.getTeamGrowthList().pipe(takeUntil(this.destroy$)).subscribe({
      next: (data) => this.dashboardDataTeamGrowth = data,
      error: (_err: any) => this.dashboardDataTeamGrowth = []
    });
  }

  ngOnDestroy() {
    this.destroy$.next();
    this.destroy$.complete();
  }
}

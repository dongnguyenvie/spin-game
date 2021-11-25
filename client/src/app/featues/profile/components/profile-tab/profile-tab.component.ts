import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-profile-tab',
  templateUrl: './profile-tab.component.html',
  styleUrls: ['./profile-tab.component.scss'],
})
export class ProfileTabComponent implements OnInit {
  @Input() tabs: string[] = [];
  @Input() active = 0;
  @Output() changeTab = new EventEmitter<number>();

  constructor() {}

  onChangeTab(tab: number) {
    this.changeTab.emit(tab);
  }

  ngOnInit() {}
}

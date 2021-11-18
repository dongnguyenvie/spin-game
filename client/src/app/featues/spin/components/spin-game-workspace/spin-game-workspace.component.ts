import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-spin-game-workspace',
  templateUrl: './spin-game-workspace.component.html',
  styleUrls: ['./spin-game-workspace.component.scss'],
})
export class SpinGameWorkspaceComponent implements OnInit {
  wheelRotation = 0;
  constructor() {}

  ngOnInit() {}

  rotate() {
    const x = 1024;
    const y = 9999;

    const deg = Math.floor(Math.random() * (x - y)) + y;
    this.wheelRotation = deg;
  }
}

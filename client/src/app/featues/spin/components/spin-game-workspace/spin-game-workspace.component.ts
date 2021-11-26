import { Component, OnInit } from '@angular/core';
import { GameService } from 'src/app/@core/services/game.service';

@Component({
  selector: 'app-spin-game-workspace',
  templateUrl: './spin-game-workspace.component.html',
  styleUrls: ['./spin-game-workspace.component.scss'],
})
export class SpinGameWorkspaceComponent implements OnInit {
  wheelRotation = 500;
  isRunning = false;
  isStopping = false;
  options: number[] = [];
  constructor(private gameSvc: GameService) {}

  ngOnInit() {
    this.gameSvc.fetchAwardOptions$().subscribe((resp) => {
      if (!resp) return;
      this.options = resp.awards;
    });
  }

  private _reset() {}

  rotate() {
    this.isRunning = true;
    this.isStopping = false;
    const x = 1024;
    const y = 9999;

    const deg = Math.floor(Math.random() * (x - y)) + y;
    this.wheelRotation = deg;
    this.gameSvc.fetchPlayGame$().subscribe((val) => {
      console.log(val);
      setTimeout(() => {
        this.isStopping = true;
      }, 3000);
    });
  }
}

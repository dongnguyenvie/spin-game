import { Component, OnInit } from '@angular/core';
import { of, switchMap, take } from 'rxjs';
import { AuthService } from 'src/app/@core/services/auth.service';
import { GameService } from 'src/app/@core/services/game.service';
import { WalletService } from 'src/app/@core/services/wallet.service';

const AWARDS_POSITION = [0, 9, 277, 330, 229, 330, 138, 47];
const AWARDS_ICON = ['😊', '😝', '😍', '😎', '🍩', '🍭', '🍰', '🍬'];
@Component({
  selector: 'app-spin-game-workspace',
  templateUrl: './spin-game-workspace.component.html',
  styleUrls: ['./spin-game-workspace.component.scss'],
})
export class SpinGameWorkspaceComponent implements OnInit {
  wheelRotation = 500;
  isRunning = false;
  isStopping = false;
  options: string[] = [];
  result: string | null = null;
  constructor(
    private gameSvc: GameService,
    private authSvc: AuthService,
    private wallSvc: WalletService
  ) {}

  ngOnInit() {
    this.gameSvc.fetchAwardOptions$().subscribe((resp) => {
      if (!resp) return;
      this.options = resp.awards.map((d) => d + 'ETH');
      console.log(resp);
    });
  }

  get awards() {
    return this.options.map((opt, i) =>
      i % 2 === 0 ? `${opt} ${AWARDS_ICON[i]}` : `${AWARDS_ICON[i]} ${opt}`
    );
  }

  rotate() {
    const token$ = this.authSvc.token$.pipe(take(1));
    this.gameSvc.startProcessing();

    token$
      .pipe(
        switchMap((token) => {
          if (!token) {
            alert('Vui lòng login tài khoản bằng ví MetaMask');
            return of(null);
          }
          this.isRunning = true;
          this.isStopping = false;
          this.result = null;
          return this.gameSvc.fetchPlayGame$();
        })
      )
      .subscribe((resp) => {
        if (!resp) return;
        setTimeout(() => {
          this.isStopping = true;
          this.wheelRotation = AWARDS_POSITION[resp.index];
          this.result = resp.award + 'ETH';
          setTimeout(() => {
            this.gameSvc.stopProcessing();
            this.gameSvc.fetchMyPackage$().subscribe();
          }, 6000);
        }, 3000);
      });
  }
}

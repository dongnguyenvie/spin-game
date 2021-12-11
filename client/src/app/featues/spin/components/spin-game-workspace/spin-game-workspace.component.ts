import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
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
    console.log('xxxxxx', this);
  }

  get awards() {
    return this.options.map((opt, i) =>
      i % 2 === 0 ? `${opt} ${AWARDS_ICON[i]}` : `${AWARDS_ICON[i]} ${opt}`
    );
  }

  rotate() {
    const package$ = this.gameSvc.package$.pipe(take(1));
    this.gameSvc.startProcessing();

    package$
      .pipe(
        switchMap((packageData) => {
          if (!packageData || packageData.package < 1) {
            alert(
              'Tài khoản của bạn chưa đăng nhập hoặc không đủ số lượt quay'
            );
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
            this.wallSvc.fetchMyWallet$().subscribe();
          }, 6000);
        }, 3000);
      });
  }

  ngOnChanges() {
    console.log('ngOnChanges');
  }

  ngDoCheck() {
    console.log('ngDoCheck');
  }

  ngAfterContentInit() {
    console.log('ngAfterContentInit');
  }

  ngAfterContentChecked() {
    console.log('ngAfterContentChecked');
  }

  ngAfterViewInit() {
    console.log('ngAfterViewInit');
  }

  ngAfterViewChecked() {
    console.log('ngAfterViewChecked');
  }

  ngOnDestroy() {
    console.log('ngOnDestroy');
  }
}

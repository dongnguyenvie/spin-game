import {
  ChangeDetectionStrategy,
  Component,
  OnDestroy,
  OnInit,
} from '@angular/core';
import { map, Subscriber, Subscription, take } from 'rxjs';
import { packagePrice, PACKAGES } from 'src/app/@core/constant/common.constant';
import { VarDirective } from 'src/app/@core/directives/ng-var.directive';
import { GameService } from 'src/app/@core/services/game.service';
import { SignerService } from 'src/app/@core/services/signer.service';
import { WalletService } from 'src/app/@core/services/wallet.service';
import { ProfileService } from '../../profile.service';

// 'Package',
const TABS = ['Wallet'];
@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss'],
  viewProviders: [VarDirective],
})
export class ProfileComponent implements OnInit, OnDestroy {
  isShow$ = this.profileSvc.show$;
  tabs = TABS;
  tabActive = 0;
  packages = PACKAGES;
  depositOptions = [0.00001, 0.00002, 0.00003];

  isShowSub: Subscription;

  constructor(
    private profileSvc: ProfileService,
    private walletSvc: WalletService,
    private gameSvc: GameService,
    private signer: SignerService
  ) {}

  ngOnDestroy(): void {}

  ngOnInit(): void {}

  onOpen(): void {
    this.profileSvc.onOpen();
  }

  onClose(): void {
    this.profileSvc.onClose();
  }
}

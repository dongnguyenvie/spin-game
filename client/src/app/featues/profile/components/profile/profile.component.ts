import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscriber, Subscription } from 'rxjs';
import { WalletService } from 'src/app/@core/services/wallet.service';
import { ProfileService } from '../../profile.service';

const TABS = ['Package', 'Wallet'];
@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss'],
})
export class ProfileComponent implements OnInit, OnDestroy {
  isShow$ = this.profileSvc.show$;
  tabs = TABS;
  tabActive = 0;
  wallet$ = this.walletSvc.wallet$;

  isShowSub: Subscription;

  constructor(
    private profileSvc: ProfileService,
    private walletSvc: WalletService
  ) {}

  ngOnInit(): void {
    this.isShowSub = this.isShow$.subscribe((isShow) => {
      if (!isShow) return;
      this.walletSvc.fetchMyWallet().subscribe();
    });
  }

  ngOnDestroy(): void {
    this.isShowSub.unsubscribe();
  }

  onOpen(): void {
    this.profileSvc.onOpen();
  }

  onClose(): void {
    this.profileSvc.onClose();
  }

  onChangeTab(index: number): void {
    this.tabActive = index;
  }
}

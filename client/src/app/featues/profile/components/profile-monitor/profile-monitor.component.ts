import { Component, OnInit } from '@angular/core';
import { map, Subscription, take } from 'rxjs';
import { packagePrice, PACKAGES } from 'src/app/@core/constant/common.constant';
import { AuthService } from 'src/app/@core/services/auth.service';
import { GameService } from 'src/app/@core/services/game.service';
import { SignerService } from 'src/app/@core/services/signer.service';
import { WalletService } from 'src/app/@core/services/wallet.service';
import { toGasprice, toPlainString } from 'src/app/@core/utils/util';
import { ProfileService } from '../../profile.service';

const TABS = ['Wallet'];
@Component({
  selector: 'app-profile-monitor',
  templateUrl: './profile-monitor.component.html',
  styleUrls: ['./profile-monitor.component.scss'],
})
export class ProfileMonitorComponent implements OnInit {
  isShow$ = this.profileSvc.show$;
  tabs = TABS;
  tabActive = 0;
  packages = PACKAGES;
  depositOptions = [0.00001, 0.00002, 0.00003];

  tokenSub: Subscription;

  constructor(
    private profileSvc: ProfileService,
    private walletSvc: WalletService,
    private gameSvc: GameService,
    private signer: SignerService,
    private authSvc: AuthService
  ) {}

  ngOnInit(): void {
    this.tokenSub = this.token$.subscribe((token) => {
      if (!token) return;
      this.walletSvc.fetchMyWallet$().subscribe();
      this.gameSvc.fetchMyPackage$().subscribe();
    });
  }

  ngOnDestroy(): void {
    this.tokenSub.unsubscribe();
  }

  get wallet$() {
    return this.walletSvc.wallet$.pipe(
      map((data) => {
        if (!data) return null;
        return {
          ...data,
          balanceToStr: toPlainString(data.balance),
        };
      })
    );
  }

  get package$() {
    return this.gameSvc.package$;
  }

  get token$() {
    return this.authSvc.token$;
  }

  isInsufficientFunds(balance: number, num: number) {
    const price = num * packagePrice;
    return price > balance;
  }

  onDeposit(amount: number) {
    const addr$ = this.signer.address$.pipe(take(1));
    addr$.subscribe((addr) => {
      this.signer.contract.methods
        .deposit()
        .send({
          value: toGasprice(amount),
          from: addr,
          gas: 42234,
        })
        .then(() => {
          alert(
            'Nạp thành công, hệ thống đang cập nhập số tiền, Vui lòng đợi trong vài phút'
          );
        })
        .catch(() => {
          alert('Nạp thất bại!!');
        });
    });
  }

  onBuyPackage(num: number) {
    this.gameSvc.fetchBuyPackage$(num).subscribe((resp) => {
      if (!resp) return;
      alert('Mua thành công');
    });
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

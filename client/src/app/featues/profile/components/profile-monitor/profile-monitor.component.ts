import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { map, Subscription, take } from 'rxjs';
import {
  COIN_DECIMAL,
  packagePrice,
  PACKAGES,
} from 'src/app/@core/constant/common.constant';
import { AuthService } from 'src/app/@core/services/auth.service';
import { GameService } from 'src/app/@core/services/game.service';
import { SignerService } from 'src/app/@core/services/signer.service';
import { WalletService } from 'src/app/@core/services/wallet.service';
import {
  toGasprice,
  toPlainString,
  tsxChecker,
} from 'src/app/@core/utils/util';
import { ProfileService } from '../../profile.service';

const GAS_DEFAULT = 30000000000;
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

  withdrawProcessing: Record<
    string,
    { status: boolean; amount: string; info: any }
  > = {};

  tokenSub: Subscription;
  withdrawAmout = 0;
  @ViewChild('withdrawInput') withdrawInput: ElementRef;

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

  get note() {
    return `Giá 1 lượt = ${toPlainString(packagePrice)}ETH`;
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
          gasPrice: GAS_DEFAULT,
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

  onChangeWithdrawAmount(e: any) {
    const amount = e.target.value.replace(/[^0-9\.]+/g, '') || 0;
    this.withdrawAmout = amount * COIN_DECIMAL;
    this.withdrawInput.nativeElement.value = amount;
  }

  onSetMaxWithdraw() {
    const sub = this.walletSvc.wallet$.subscribe((wallet) => {
      const amount = wallet.balance - GAS_DEFAULT;
      this.withdrawAmout = amount;
      this.withdrawInput.nativeElement.value = toPlainString(amount);
    });
    sub.unsubscribe();
  }

  onWithdraw() {
    const result = confirm(`Bạn muốn rút ${toPlainString(this.withdrawAmout)}`);
    if (!result) return;
    this.walletSvc.fetchWithdraw$(this.withdrawAmout).subscribe((result) => {
      if (!result) return;
      this.withdrawProcessing[result.tx] = {
        amount: toPlainString(result.amount),
        status: false,
        info: null,
      };

      tsxChecker(this.signer.signer, result.tx)
        .then((tsResult) => {
          this.withdrawProcessing[result.tx].status = true;
          this.withdrawProcessing[result.tx].info = tsResult;
          console.log('transaction result', tsResult);
        })
        .catch((err) => {
          console.log('transaction error', err);
        });

      console.log('onWithdraw', result);
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

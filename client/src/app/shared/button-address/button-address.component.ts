import { ChangeDetectorRef, Component, OnDestroy, OnInit } from '@angular/core';
import { AuthService } from 'src/app/@core/services/auth.service';
import { SignerService } from 'src/app/@core/services/signer.service';

@Component({
  selector: 'app-button-address',
  templateUrl: './button-address.component.html',
  styleUrls: ['./button-address.component.scss'],
  providers: [],
})
export class ButtonAddressComponent implements OnInit, OnDestroy {
  address$ = this.signer.address$;
  token$ = this.authSvc.token$;

  constructor(
    private signer: SignerService,
    private authSvc: AuthService,
    private cdn: ChangeDetectorRef
  ) {}

  ngOnInit() {
    this.authSvc.user$.subscribe({
      next: (resp) => {
        console.log('resp', resp);
      },
    });
  }

  ngOnDestroy() {}

  connectWallet() {
    this.signer.connectWallet().subscribe({
      complete: () => {
        console.log('connectWallet complete');
      },
      next: (resp) => {
        console.log('connectWallet', resp);
      },
    });
  }

  login() {
    this.authSvc.login$().subscribe((resp) => {
      console.log({ resp });
    });
  }
}

import { Component, Input, OnDestroy, OnInit } from '@angular/core';
import { AuthService } from 'src/app/@core/services/auth.service';
import { SignerService } from 'src/app/@core/services/signer.service';
import { ProfileService } from 'src/app/featues/profile/profile.service';

@Component({
  selector: 'app-button-address',
  templateUrl: './button-address.component.html',
  styleUrls: ['./button-address.component.scss'],
})
export class ButtonAddressComponent implements OnInit, OnDestroy {
  address$ = this.signer.address$;
  token$ = this.authSvc.token$;
  user$ = this.authSvc.user$;

  constructor(
    private signer: SignerService,
    private authSvc: AuthService,
    private profileSvc: ProfileService
  ) {}

  ngOnInit(): void {
    this.authSvc.user$.subscribe({
      next: (resp) => {
        console.log('resp', resp);
      },
    });
  }

  ngOnDestroy(): void {}

  connectWallet(): void {
    this.signer.connectWallet().subscribe({
      complete: () => {
        console.log('connectWallet complete');
      },
      next: (resp) => {
        console.log('connectWallet', resp);
      },
    });
  }

  login(): void {
    this.authSvc.login$().subscribe((resp) => {
      console.log({ resp });
    });
  }

  onOpenProfile(): void {
    this.profileSvc.onOpen();
  }
}

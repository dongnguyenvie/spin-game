import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subject } from 'rxjs';
import { AuthService } from 'src/app/@core/services/auth.service';
import { SignerService } from 'src/app/@core/services/signer.service';

@Component({
  selector: 'app-button-address',
  templateUrl: './button-address.component.html',
  styleUrls: ['./button-address.component.scss'],
  providers: [],
})
export class ButtonAddressComponent implements OnInit, OnDestroy {
  balance$ = new Subject<void>();
  account$ = new Subject<void>();

  address$ = this.signer.address$;
  isLogon$ = this.signer.isLogon$;

  constructor(private signer: SignerService, private authSvc: AuthService) {}

  ngOnInit() {}

  ngOnDestroy() {}

  connectWallet() {
    this.signer.connectWallet();
  }

  login() {
    this.authSvc.login$().subscribe((resp) => {
      console.log({ resp });
    });
  }
}

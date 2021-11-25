import { Component } from '@angular/core';
import { defer, map, switchMap, take, timer } from 'rxjs';
import { EthAccounts } from './@core/model/common.model';
import { SignerService } from './@core/services/signer.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  title = 'client';
  constructor(private singer: SignerService) {
  }
}

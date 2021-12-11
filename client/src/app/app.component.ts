import { ChangeDetectorRef, Component, NgZone } from '@angular/core';
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
  constructor(
    private singer: SignerService,
    private zone: NgZone,
    private cdr: ChangeDetectorRef
  ) {
    this.zone.runOutsideAngular(() => {
      this.singer.setup(cdr);
    });
  }

  ngOnInit() {
    // console.log('xxxxxx', this);
  }
}

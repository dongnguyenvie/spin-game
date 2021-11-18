import { Component, OnInit } from '@angular/core';
import { Subject } from 'rxjs';

@Component({
  selector: 'app-button-address',
  templateUrl: './button-address.component.html',
  styleUrls: ['./button-address.component.scss'],
})
export class ButtonAddressComponent implements OnInit {
  $balance = new Subject<void>();
  $account = new Subject<void>();

  constructor() {}

  ngOnInit() {}

  connectWallet() {}
}

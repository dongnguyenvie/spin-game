import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-profile-modal',
  templateUrl: './profile-modal.component.html',
  styleUrls: ['./profile-modal.component.scss'],
})
export class ProfileModalComponent implements OnInit {
  @Input() isOpen: Observable<boolean>;
  @Output() close = new EventEmitter();

  constructor() {}

  ngOnInit() {}

  onClose() {
    this.close.emit();
  }
}

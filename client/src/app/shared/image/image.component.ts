import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Md5 } from 'src/app/@core/libs/md5';

const prefixMail = 'http://www.gravatar.com/avatar/';
@Component({
  selector: 'app-image',
  templateUrl: './image.component.html',
  styleUrls: ['./image.component.scss'],
})
export class ImageComponent implements OnInit {
  @Input() src: string;
  @Input() class: string = '';

  @Output() click = new EventEmitter();

  constructor() {}

  get link() {
    const emailMd5 = Md5.hashStr(this.src);
    return prefixMail + emailMd5;
  }

  onClick(): void {
    this.click.emit();
  }

  ngOnInit(): void {}
}

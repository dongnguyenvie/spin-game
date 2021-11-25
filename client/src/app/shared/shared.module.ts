import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AddressPipe } from './pipes/address.pipe';
import { ButtonAddressComponent } from './button-address/button-address.component';
import { ImageComponent } from './image/image.component';

@NgModule({
  imports: [CommonModule],
  declarations: [AddressPipe, ButtonAddressComponent, ImageComponent],
  exports: [AddressPipe, ButtonAddressComponent, ImageComponent],
})
export class SharedModule {}

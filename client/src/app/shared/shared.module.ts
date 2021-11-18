import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AddressPipe } from './pipes/address.pipe';
import { ButtonAddressComponent } from './button-address/button-address.component';

@NgModule({
  imports: [CommonModule],
  declarations: [AddressPipe, ButtonAddressComponent],
  exports: [AddressPipe, ButtonAddressComponent],
})
export class SharedModule {}

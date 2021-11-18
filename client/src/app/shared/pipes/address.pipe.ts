import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'address',
})
export class AddressPipe implements PipeTransform {
  transform(value: any, args?: any): any {
    const first4Digits = value.slice(0, 4);
    const last4Digits = value.slice(-4);
    return first4Digits.concat('...', last4Digits);
  }
}

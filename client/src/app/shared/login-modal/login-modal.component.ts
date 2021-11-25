import {
  Component,
  ElementRef,
  HostListener,
  OnDestroy,
  OnInit,
  ViewChild,
} from '@angular/core';
import {
  AbstractControl,
  FormBuilder,
  FormsModule,
  ValidationErrors,
  ValidatorFn,
  Validators,
} from '@angular/forms';
import { Subscriber, Subscription } from 'rxjs';
import { AuthService } from 'src/app/@core/services/auth.service';
import { SignerService } from 'src/app/@core/services/signer.service';
import { LoginModalService } from './login-modal.service';

const checkPasswords: ValidatorFn = (
  group: AbstractControl
): ValidationErrors | null => {
  let pass = group.get('password')?.value;
  let confirmPass = group.get('confirmPassword')?.value;
  return pass === confirmPass ? null : { notSame: true };
};

@Component({
  selector: 'app-login-modal',
  templateUrl: './login-modal.component.html',
  styleUrls: ['./login-modal.component.scss'],
  providers: [],
})
export class LoginModalComponent implements OnInit, OnDestroy {
  constructor(
    private modalSvc: LoginModalService,
    private fb: FormBuilder,
    private signer: SignerService,
    private authSvc: AuthService
  ) {}

  @ViewChild('signupFormRef') signupFormRef: ElementRef;

  isShow$ = this.modalSvc.isShow$;
  addrSub!: Subscription;

  signupForm = this.fb.group(
    {
      email: ['test@mail.com', [Validators.required, Validators.minLength(4)]],
      walletAddress: ['', Validators.required],
      password: ['', [Validators.required, Validators.minLength(6)]],
      // confirmPassword: ['', Validators.required],
    },
    {
      // validator: checkPasswords,
    }
  );

  ngOnInit(): void {
    this.addrSub = this.signer.address$.subscribe((addr) => {
      if (!addr) return;
      this.signupForm.get('');
      this.signupForm.patchValue({
        walletAddress: addr,
      });
    });
  }

  ngOnDestroy(): void {
    this.addrSub.unsubscribe();
  }

  // @HostListener('document:click', ['$event.target'])
  // public onClick(target: any) {
  //   const clickedInside = this.signupFormRef.nativeElement.contains(target);
  //   if (!clickedInside) {
  //     this.modalSvc.close();
  //   }
  // }

  onClose(): void {
    this.modalSvc.close();
  }

  onRegister(): void {
    const isValid = this.signupForm.valid;
    if (!isValid) return;
    const { walletAddress, email, password } = this.signupForm.value;
    this.authSvc
      .register$({
        addr: walletAddress,
        email: email,
        password: password,
        messageSign: String(Date.now()),
      })
      .subscribe((resp) => {
        if (!resp) return;
        this.modalSvc.close();
        this.authSvc.login$().subscribe((resp2) => {
          console.log({ resp2 });
        });
      });
  }

  onOpen(): void {
    this.modalSvc.open();
  }
}

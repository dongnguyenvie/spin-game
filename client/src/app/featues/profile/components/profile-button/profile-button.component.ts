import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/@core/services/auth.service';

@Component({
  selector: 'app-profile-button',
  templateUrl: './profile-button.component.html',
  styleUrls: ['./profile-button.component.scss'],
})
export class ProfileButtonComponent implements OnInit {
  constructor(private authSvc: AuthService) {}
  user$ = this.authSvc.user$;

  ngOnInit() {
    this.user$.subscribe((user) => {
      console.log(user);
    });
  }
}

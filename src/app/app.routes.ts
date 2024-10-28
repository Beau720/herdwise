import { RouterModule, Routes } from '@angular/router';
import { MapComponent } from './map/map.component';
import { HomeComponent } from './home/home.component';
import { DevicesComponent } from './devices/devices.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { EditdeviceComponent } from './editdevice/editdevice.component';
import { ProfileComponent } from './profile/profile.component';
import { UpdateProfileComponent } from './update-profile/update-profile.component';
import { HelpDeskComponent } from './help-desk/help-desk.component';
import { CreateDeviceComponent } from './create-device/create-device.component';
import { NgModule } from '@angular/core';

export const routes: Routes = [
    { path: '', redirectTo: 'login', pathMatch: 'full'},
    { path: 'map', component: MapComponent },
    { path: 'home', component: HomeComponent },
    { path: 'login', component: LoginComponent },
    { path: 'all-devices', component: DevicesComponent },
    { path: 'add-device', component: CreateDeviceComponent},
    //{ path: 'edit-Device', component: EditdeviceComponent},
    { path: 'register', component: RegisterComponent },
    { path: 'profile', component: ProfileComponent},
    //{ path: 'editProfile', component: UpdateProfileComponent},
    { path: 'help', component: HelpDeskComponent},
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {}
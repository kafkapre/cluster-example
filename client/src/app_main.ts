import {bootstrap} from "angular2/platform/browser"     //importing bootstrap function
import {HTTP_PROVIDERS} from 'angular2/http';

import {AppComponent} from "./app.component"  //importing component function

bootstrap(AppComponent, [HTTP_PROVIDERS]);

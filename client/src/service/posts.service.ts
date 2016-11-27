import {Injectable} from "angular2/core"
import {Http, Response, Headers, RequestOptions} from 'angular2/http';
import {Observable} from 'rxjs/Rx';

import {Person} from '../person';





@Injectable()
export class PostsService{
    ping_endpoint_url:string = "http://localhost/ping";
    pings_endpoint_url:string = "http://localhost/pings";
    http:Http = null

    constructor(http: Http){
        this.http = http;
    }

    getPing (){
        return this.http.get(this.ping_endpoint_url)
              .map((res:Response) => res.json())
              .catch((error:any) => Observable.throw(error.json().error || 'Server error'));
    }

    postPing(body: Object): Observable<Person[]> {
        let bodyString = JSON.stringify(body); // Stringify payload
        let headers      = new Headers({ 'Content-Type': 'application/json' }); // ... Set content type to JSON
        let options       = new RequestOptions({ headers: headers }); // Create a request option

        return this.http.post(this.pings_endpoint_url, bodyString, options) // ...using post request
                         .map((res:Response) => res.json()) // ...and calling .json() on the response to return data
                         .catch((error:any) => Observable.throw(error.json().error || 'Server error')); //...errors if
    }


}

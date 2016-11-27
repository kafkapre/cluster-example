import {Component} from 'angular2/core';

import {PostsService} from './service/posts.service'

import {Person} from './person';



@Component({
  selector: 'my-app',
  templateUrl: './src/app.component.tpl.html',
  providers: [PostsService]
})

export class AppComponent {

    error:String = "some_error";
    person:Person = new Person("1", "Natasa");
    postsService: PostsService

    constructor(private _postsService: PostsService){
        this.postsService = _postsService;
    }

    getCountriesByRegion(){
        this.error = "";

        // this.postsService.getPing();

        this.postsService.getPing()
         .subscribe(
            data => this.person = data,
            error => this.error = "Region is invalid."
         );
    }


    postPing(){
        this.error = "";
        let p = new Person("Jiri", "Fabian");

        // this.postsService.getPing();

        this.postsService.postPing(p)
         .subscribe(
            error => this.error = "post is invalid."
         );
    }

}

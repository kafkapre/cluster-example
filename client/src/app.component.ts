import {Component} from 'angular2/core';

import {PostsService} from './service/posts.service'

import {Person} from './person';
import {Comment} from './model/comment';



@Component({
  selector: 'my-app',
  templateUrl: './src/app.component.tpl.html',
  providers: [PostsService]
})

export class AppComponent {

    error:String = "some_error";
    person:Person = new Person("1", "Natasa");
    postsService: PostsService
    comments: Comment[] = new Array();

    constructor(private _postsService: PostsService){
        this.postsService = _postsService;
        this.comments = [
                Comment.create("1", "author 1", "aaaaa"),
                Comment.create("2", "author 2", "bbbbb"),
                Comment.create("3", "author 3", "ccccc"),
                Comment.create("4", "author 4", "ddddd"),
                Comment.create("5", "author 5", "eeeee"),
                Comment.create("6", "author 6", "fffff")

        ]
        for (var c of this.comments) {
            #console.log(c) //TODO
        }

    }

    getCountriesByRegion(){
        this.error = "";

        // this.postsService.getPing();

        //this.postsService.getPing()
        // .subscribe(
        //    data => this.person = data,
        //    error => this.error = "Region is invalid."
        // );
    }


    postPing(){
        this.error = "";
        let p = new Person("Jiri", "Fabian");

        // this.postsService.getPing();

        //this.postsService.postPing(p)
        // .subscribe(
        //    error => this.error = "post is invalid."
        // );
    }

}

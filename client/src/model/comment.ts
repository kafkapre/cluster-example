export class Comment {
    constructor(
        public id: string,
        public author: string,
        public text: string,
        public likes: number,
        public unlikes: number,
        public parentId: string,
        public timestamp: string
        ){}


        static create(id: string, author: string, text: string) {
            return new Comment(id, author, text, 0, 0, "none", "1111")
        }

}

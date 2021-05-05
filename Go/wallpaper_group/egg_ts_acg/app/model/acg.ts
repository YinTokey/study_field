import {Application} from 'egg';

export default (app: Application) => {
    const mongoose = app.mongoose;
    const Schema = mongoose.Schema;
    const PostSchema = new Schema({
        pictureId: {
            type: String,
        },
        imageUrl: {
            type: String,
        },
        largeImageUrl: {
            type: String,
        },
        name: {
            type: String,
        },
        description: {
            type: String,
        },
        author: {
            type: String,
        },
        width: {
            type: Number,
        },
        height: {
            type: Number,
        },
        likes: {
            type: Number,
        },
        category: {
            type: String,
        },
        tags: {
            type: [Object],
        },
        columnId: {
            type: Schema.Types.ObjectId,
        },
    });

    return mongoose.model('Acg', PostSchema);

};

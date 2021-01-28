'use strict';
module.exports = app => {
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
        categories: {
            type: String,
        },
        columnId: {
            type: Schema.Types.ObjectId,
        },
    });
    return mongoose.model('Acg', PostSchema);
};

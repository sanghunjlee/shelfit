# ShelfIt
A simple way to manage lists (or 'shelves') of items. 

## About 

Shelfit is a project I have wanted to implement for myself to organize various lists that started to grow exponentially ever since the COVID year. I've begun to form a concrete idea about building this CLI app after encountering the [ultralist](https://github.com/ultralist/ultralist) by Grant Ammons, whose open source code helped me learn a lot about how to implement this concept. 

Here are concepts that I wanted to focus the implemtation around:

1. Using "tags" to group & filter items
2. Ditching rating system to group & filter (I found rating distracting in terms of organizing)
3. Organizing items the way a library would organize its collection


## Documentation



## How to use
#### Initializing the shelf
First, you have to initialize the *shelf*.
```bash
> shelfit init
```

This will create the `shelf.json` file where all the data will be stored.

#### Adding books
Add a *book* with a command `shelfit add`, and details of the book.

You can describe the book with the tags: `@category`, `.genre`, and/or `!status`. These tags should only be one-word-long. The `.genre` tag can be repeated to add more genres to the detail.

You can also expand the book into *volumes* by using the tag `+subitem`. The plus symbole (`+`) will trigger the parser to recognize everything there after (until another `+` or the end of the line is reached) as a sub-item. Each subitem can be described with the `!status` tag.

###### Examples

```bash
> shelfit add Chainsaw Man @manga .good +v1 !finished +v2 !unread
> shelfit add @book The Lord of the Ring !finished +The Fellowship of the Ring +The Two Towers +The Return of the King
> shelfit add @movie Drive My Car .favorite .japanese .murakami !finished
```

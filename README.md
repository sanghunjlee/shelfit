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
### Initializing the *Shelf*
First, you have to initialize the *shelf*.
```bash
> shelfit init
```

This will create the `shelf.json` file where all the data will be stored. If `shelf.json` file already exists, the command will do nothing.

### Adding Items
Add an item with a command `shelfit add <item>`. `<item>` is consisted of the *title* of the item (simple description of the item) and *qualifiers* of the item.

There are two qualifiers that can be used:

1. `!` - Category (**required**)
 
    - This quantifier is required
    - If you try to input multiple categories, only the first category will be registered.

2. `.` - Tag(s)

> [!NOTE]
> The qualifier can only be used a continuous string (no space). Please use dash(s) (`-`) or underscore(s) (`_`) to input multi-words qualifier(s)

**Examples**

```go

> shelfit add The Lord of the Ring !book .fantasy .fiction

// you can use `-` or `_` to make a multi-word qualifier
> shelfit add Apples (3) !grocery-list .vons .fruits_veggies

// you can use `--note` or `-n` flag to add a one-line note
// encapsulated in quotes
> shelfit add Chainsaw Man !manga .good_stuff --note "Written by the same author as Fire Punch"
```

### Listing Items

### Editting an Item

### Deleting Item(s)

### Clearing Items

## How to Contribute

Please refer to TODO.md for list of features to add
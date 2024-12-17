package class

type Post struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float32 `json:"price"`
    ImageURL    string  `json:"image_url"`
}

// Getter methods for the struct fields

func(post *Post) GetID() int {
    return post.ID
}

func(post *Post) GetName() string {
    return post.Name
}

func(post *Post) GetDescription() string {
    return post.Description
}

func(post *Post) GetPrice() float32 {
    return post.Price
}

func(post *Post) GetImageUrl() string {
    return post.ImageURL
}



func(post *Post) GetIDAdress() *int {
    return &post.ID
}

func(post *Post) GetNameAdress() *string {
    return &post.Name
}

func(post *Post) GetDescriptionAdress() *string {
    return &post.Description
}

func(post *Post) GetPriceAdress() *float32 {
    return &post.Price
}

func(post *Post) GetImageUrlAdress() *string {
    return &post.ImageURL
}

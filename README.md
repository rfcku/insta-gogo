```
$ go get github.com/rfcku/insta-gogo
```


```go
    ig := ig.New("IG_TOKEN", "IG_ID")
	
    // Create & Publish Story
	story, err := ig.CreateAndPublish(imageURL, "Sample Caption", "STORIES")
	if err != nil {
		fmt.Println("Error creating container:", err)
		return
	}
	fmt.Printf("Created container with ID: %s\n", feed)

	feed, err = ig.CreateAndPublish(imageURL, "Sample Description", "IMAGE")
	if err != nil {
		fmt.Println("Error publishing container:", err)
		return
	}
	fmt.Printf("Published container with response: %s\n", story)
```

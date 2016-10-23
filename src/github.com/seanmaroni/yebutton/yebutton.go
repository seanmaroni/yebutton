package slash

import "fmt"
import "net/http"
import "time"
import "math/rand"


func init() {
     http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

     //Read the Request Parameter "command"
     command := r.FormValue("command")

     //Ideally do other checks for tokens/username/etc
    
     if command == "/ye" {
     	//generates a random int
     	rand.Seed(time.Now().Unix())

     	quotes := []string{
    		"I am a god.",
    		"My greatest award is what I’m about to do.",
    		"You can’t look at a glass half full or empty if it’s overflowing",
    		"Nobody can tell me where I can and can’t go",
    		"Most people are slowed down by the perception of themselves. If your taught you can't do anything, you won't do anything.",
    		"If your a Kanye West fan you’re not a fan of me, you’re a fan of yourself.",
    		"The only luxury is time.",
    		"People in this world shun people for being great, for being a bright color, for standing out. But the time is now, to be OK with being the greatest you.",
    		"You’ve got to be really dialed into exactly who you are to the one hundredth power or you’re just everyone else.",
    		"It’s ok to create. It’s ok to not be boxed in. I want people to feel like, awesome is possible.",
    		"The time is now. The time is now to express yourself. The time is now to believe in yourself.",
    		"People always tell you to be humble, be humble, be humble. When was the last time someone told you to be great. To be amazing. To be awesome.",
	}

		n := rand.Int() % len(quotes)
       	fmt.Fprint(w,"'", quotes[n],"'")

     } else {
         fmt.Fprint(w,"I don't understand your command.")
     }
}
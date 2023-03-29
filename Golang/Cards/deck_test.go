package card_api

import (
    "encoding/json"
    "net/http"
    "fmt"
    "testing"
    "reflect"
    "flag"
    "math/rand"
    "time"
    "strings"
)

var meta_test = flag.Bool("meta_test", false, "meta_test flag")

var verbosity = flag.Int("v", 0, "verbosity level")


func fill_ss_deck () []string  {
    values  := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "0", "J", 
                "Q", "K"}
    suits   := []string{"S", "D", "C", "H"}
    var ss_deck []string

    for _, suit := range suits {
    	for _, value := range values {
            ss_deck = append(ss_deck, value+suit)
    	}
    }
    return ss_deck
}

var synthetic_sorted_deck = fill_ss_deck()

var new_unshuffled_deck_id string
var new_shuffled_deck_id   string

type NewDeckResponse struct {
    Success    bool   `json:"success"`
    DeckID     string `json:"deck_id"`
    Remaining  int    `json:"remaining"`
    Shuffled   bool   `json:"shuffled"`
}

type DrawNCardsResponse struct {
    Success    bool   `json:"success"`
    DeckID     string `json:"deck_id"`
    Remaining  int    `json:"remaining"`
    Cards      []Card `json:"cards"`
}

type myCardObject struct {
    ArrayValue []int `json:"arrayValue"`
}

type Card struct {
    Code       string `json:"code"`
    Value      string `json:"value"`
    Suit       string `json:"suit"`
    Image      string `json:"image"`
}

type AddToPilesResponse struct {
    Success    bool   `json:"success"`
    DeckId     string `json:"deck_id"`
    Remaining  int `json:"remaining"`
    Piles      map[string]Pile `json:"piles"`
}

type Pile struct {
    Remaining  int `json:"remaining"`
    Cards      []Card `json:"cards"`
}

func NewDeckUnshuffled() (*NewDeckResponse, error) {
    /* Make a GET request to the "https://deckofcardsapi.com/api/deck/new/" 
      endpoint. */
    resp, err := http.Get("https://deckofcardsapi.com/api/deck/new/")
    if err != nil {
    	return nil, err
    }

    defer resp.Body.Close()
    
    // Check the response status code
    if resp.StatusCode != http.StatusOK {
    	return nil, fmt.Errorf("Expected status code %d, but got %d", 
            http.StatusOK, resp.StatusCode)
    }
    
    // Parse the response body into a NewDeckResponse struct
    var newDeckResponse NewDeckResponse
    err = json.NewDecoder(resp.Body).Decode(&newDeckResponse)
    if err != nil {
    	return nil, err
    }

    
    return &newDeckResponse, nil
}

func TestNewDeckUnshuffled(t *testing.T) {
    /* Call the NewDeckUnshuffled function to get a new deck. Since the request 
      does not include the "shuffled" parameter, we expect the card order to be
      in the ordered state described in the Deck of Cards API page. For suits,
      this is Spades, Diamonds, Clubs, and Hearts. */ 
    newDeckResponse, err := NewDeckUnshuffled()
    if err != nil {
    	t.Fatal(err)
    }

    if *verbosity > 0 {
        println(fmt.Sprintf("The deck_id value is \"%s\".", 
          newDeckResponse.DeckID ) ) 
    }

    // assign the Deck ID to be used in a later test.
    new_unshuffled_deck_id = newDeckResponse.DeckID
    if *verbosity > 0 {
        println(fmt.Sprintf("The new_unshuffled_deck_id value is \"%s\".", 
          new_unshuffled_deck_id)) 
    }

    // Check the "success" field in the response.
    if !newDeckResponse.Success {
    	t.Errorf("Expected \"success\" to be true, but got false.")
    }
    
    // Check that "shuffled" is not true (in an UNshuffled deck).
    if newDeckResponse.Shuffled {
        t.Errorf("Expected \"shuffled\" to be false, but got true.")
    }

    // Check that the "deck_id" field is not empty.
    if newDeckResponse.DeckID == "" {
    	t.Errorf("Expected \"deck_id\" to be non-empty, but got empty.")
    }

    // Check that the "remaining" field is a full deck of 52 cards.
    if newDeckResponse.Remaining != 52 {
    	t.Errorf("Expected \"remaining\" to be 52, but got %d.", 
        newDeckResponse.Remaining)
    }

}
    
func NewDeckShuffled() (*NewDeckResponse, error) {
    // Make a GET request to the "https://deckofcardsapi.com/api/deck/new/"
    // endpointwith "shuffle" in the path.
    resp, err := http.Get("https://deckofcardsapi.com/api/deck/new/shuffle")
    // println(resp)
    if err != nil {
    	return nil, err
    }

    defer resp.Body.Close()
    
    // Check the response status code
    if resp.StatusCode != http.StatusOK {
    	return nil, fmt.Errorf("Expected status code %d, but got %d", 
        http.StatusOK, resp.StatusCode)
    }
    
    // Parse the response body into a NewDeckResponse struct
    var newDeckResponse NewDeckResponse
    err = json.NewDecoder(resp.Body).Decode(&newDeckResponse)
    if err != nil {
    	return nil, err
    }
    
    return &newDeckResponse, nil
}

func TestNewDeckShuffled(t *testing.T) {
    // Call the NewDeckShuffled function to get a new deck, shuffled.
    newDeckResponse, err := NewDeckShuffled()
    if err != nil {
    	t.Fatal(err)
    }

    if *verbosity > 0 {
        println(fmt.Sprintf("The deck_id value is \"%s\".", 
          newDeckResponse.DeckID)) 
    }
    if *verbosity > 0 {
        println(fmt.Sprintf("The deck_id value is \"%s\".", 
          newDeckResponse.DeckID)) 
    }

    // assign the Deck ID to be used in a later test.
    new_shuffled_deck_id = newDeckResponse.DeckID
    if *verbosity > 0 {
        println(fmt.Sprintf("The new_shuffled_deck_id value is \"%s\".",
          new_shuffled_deck_id ) ) 
    }

    // Check the "success" field in the response.
    if !newDeckResponse.Success {
    	t.Errorf("Expected success to be true, but got false")
    }
    
    // Check that "shuffled" is true (in this, a shuffled deck).
    if !newDeckResponse.Shuffled {
        t.Errorf("Expected shuffled to be true, but got false")
    }

    // Check that the "deck_id" field is not empty.
    if newDeckResponse.DeckID == "" {
    	t.Errorf("Expected deck_id to be non-empty, but got empty")
    }

    // Check that the "remaining" field is equal to the deck_multiplier times
    // 52.
    if newDeckResponse.Remaining != 52 {
    	t.Errorf("Expected deck_id to be %d, but got %d.", 
        52, newDeckResponse.Remaining )
    }

}

func DrawNCardsFromUnshuffledDeck(cards_count int) (*DrawNCardsResponse, 
    error) {
    
    // Call the draw function.
    resp2, err := http.Get(
        fmt.Sprintf("https://deckofcardsapi.com/api/deck/%s/draw/?count=%d", 
          new_unshuffled_deck_id, cards_count))

    // Parse the response body into a drawNCardsResponse struct
    var drawNCardsResponse DrawNCardsResponse
    err = json.NewDecoder(resp2.Body).Decode(&drawNCardsResponse)
    if err != nil {
    	return nil, err
    }

    return &drawNCardsResponse, nil

}

func TestDrawNCardsFromUnshuffledDeck(t *testing.T) {
    if *verbosity > 0 {
        println(fmt.Sprintf(
          "Using the new_unshuffled_deck_id value of \"%s\".", 
          new_unshuffled_deck_id)) 
    }

    DrawNCardsResponse, err := DrawNCardsFromUnshuffledDeck(52)
    if err != nil {
        t.Fatal(err)
    }

    if !DrawNCardsResponse.Success {
    	t.Errorf("Expected success to be true, but got false")
    }

    // Check that the "deck_id" field is not empty
    if DrawNCardsResponse.DeckID == "" {
    	t.Errorf("Expected deck_id to be non-empty, but got empty")
    }
    if len(DrawNCardsResponse.Cards) != 52 {
        t.Errorf("Did not get the expected number of cards!")
    }
    indiv_card_counts := make(map[string]int)
    var actual_new_card_sort []string
    for i := 0; i < 52; i++ {
        indiv_card_counts[DrawNCardsResponse.Cards[i].Code]++
        actual_new_card_sort = append(actual_new_card_sort, 
            DrawNCardsResponse.Cards[i].Code)
    }
    if *verbosity >= 3 {
        for _, card := range actual_new_card_sort {
            println(fmt.Sprintf("%s", card))
        }
    }
    
    if *meta_test {
        rand.Seed(time.Now().UnixNano())
        
        rand.Shuffle(len(actual_new_card_sort), func(i, j int) {
            actual_new_card_sort[i], 
            actual_new_card_sort[j] = actual_new_card_sort[j], 
                                      actual_new_card_sort[i]
        })
    }
        
    if !reflect.DeepEqual(actual_new_card_sort, synthetic_sorted_deck) {
        t.Errorf("The new unshuffled card deck is not in sorted order.")
        if *verbosity >= 1 {
            println(strings.Join(actual_new_card_sort, " "))
        }
    }

    if *meta_test {
        indiv_card_counts["5D"]++
        indiv_card_counts["QH"] = 0
    }

    for key, value := range indiv_card_counts {
        if value != 1 {
            t.Errorf(
                "The card: \"%s\" was returned other than once! (%d times)",
                key, value,
            )
        }
    }

}

func NewMultipleDecksShuffled(deck_multiplier int) (*NewDeckResponse, error) {
    // Make a GET request to the "https://deckofcardsapi.com/api/deck/new/"
    //  endpoint with "shuffle" in the path.
    resp, err := http.Get(fmt.Sprintf(
      "https://deckofcardsapi.com/api/deck/new/shuffle?deck_count=%d", 
      deck_multiplier))
    if err != nil {
    	return nil, err
    }
    defer resp.Body.Close()
    
    // Check the response status code
    if resp.StatusCode != http.StatusOK {
    	return nil, fmt.Errorf("Expected status code %d, but got %d", 
          http.StatusOK, resp.StatusCode)
    }
    
    // Parse the response body into a NewDeckResponse struct
    var newDeckResponse NewDeckResponse
    err = json.NewDecoder(resp.Body).Decode(&newDeckResponse)
    if err != nil {
    	return nil, err
    }
    
    return &newDeckResponse, nil
}

func TestNewMultipleDecksShuffled(t *testing.T) {
    // Call the NewDeckShuffled function to get a new deck, shuffled.
    deck_multiplier := 20 // Testing to date shows that 20 is the maximum.
    newDeckResponse, err := NewMultipleDecksShuffled(deck_multiplier)
    if err != nil {
    	t.Fatal(err)
    }

    if *verbosity > 0 {
        println(fmt.Sprintf("The deck_id value is \"%s\".", 
          newDeckResponse.DeckID ) ) 
    }
   
    // Check the "success" field in the response.
    if !newDeckResponse.Success {
    	t.Errorf("Expected success to be true, but got false")
    }
    
    // Check that "shuffled" is true (in an shuffled deck).
    if !newDeckResponse.Shuffled {
        t.Errorf("Expected shuffled to be true, but got false")
    }
    
    // Check that the "deck_id" field is not empty.
    if newDeckResponse.DeckID == "" {
    	t.Errorf("Expected deck_id to be non-empty, but got empty")
    }
    // Check that the "remaining" field is equal to the deck_multiplier times
    // 52.
    if newDeckResponse.Remaining != (52 * deck_multiplier) {
    	t.Errorf("Expected deck_id to be %d, but got %d.", 
          (52 * deck_multiplier), newDeckResponse.Remaining )
    }

}

func DrawNCardsFromShuffledDeck(cards_count int) (*DrawNCardsResponse, error) {

    // call the draw function.
    // print(fmt.Sprintf("%s", newDeckResponse.DeckID))
    resp2, err := http.Get(fmt.Sprintf(
      "https://deckofcardsapi.com/api/deck/%s/draw/?count=%d",
      new_shuffled_deck_id, cards_count))
    // Parse the response body into a drawNCardsResponse struct
    var drawNCardsResponse DrawNCardsResponse
    err = json.NewDecoder(resp2.Body).Decode(&drawNCardsResponse)
    if err != nil {
    	return nil, err
    }
       
    return &drawNCardsResponse, nil
}

func TestDrawNCardsFromShuffledDeck(t *testing.T) {
    DrawNCardsResponse, err := DrawNCardsFromShuffledDeck(52)
    if err != nil {
        t.Fatal(err)
    }

    if !DrawNCardsResponse.Success {
    	t.Errorf("Expected success to be true, but got false")
    }
    // Check that the "deck_id" field is not empty
    if DrawNCardsResponse.DeckID == "" {
    	t.Errorf("Expected deck_id to be non-empty, but got empty")
    }
    if len(DrawNCardsResponse.Cards) != 52 {
        t.Errorf("Did not get the expected number of cards!")
    }
    indiv_card_counts := make(map[string]int)
    var actual_new_card_sort []string
    for i := 0; i < 52; i++ {
        indiv_card_counts[DrawNCardsResponse.Cards[i].Code]++
        actual_new_card_sort = append(actual_new_card_sort, 
            DrawNCardsResponse.Cards[i].Code)
    }

    if *meta_test {
        actual_new_card_sort = synthetic_sorted_deck
    }

    
    if reflect.DeepEqual(actual_new_card_sort, synthetic_sorted_deck) {
        t.Errorf("The new card deck is (improperly) in an unshuffled order.")
         if *verbosity >= 1 {
            println(strings.Join(actual_new_card_sort, " "))
        }
    }

    if *meta_test {
        indiv_card_counts["5D"]++
        indiv_card_counts["QH"] = 0
    }

    for key, value := range indiv_card_counts {
        if value != 1 {
            t.Errorf(
                "The card: \"%s\" was returned other than once! (%d times)",
                key, value,
            )
        }
    }
}
    

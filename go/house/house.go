// Package house procedurally generates a song.
package house

import(
	"strings"
	"bytes"
)


const testVersion = 2

// Embed embeds a noun phrase as the object of relative clause with a
// transitive verb.
func Embed(clause, noun string) string {
	if noun == "" { 
		return clause
	}
	return clause + " " + noun
}

// Verse generates a verse of a song with relative clauses that have
// a recursive structure.
func Verse(subject string, phrases []string, noun string) string {
	phrase := []string{}
	for _, p := range phrases {
		phrase = append(phrase, Embed(p, ""))
	}
	out := []string{ subject }
	out = append(out, phrase...)
	out = append(out, noun)
	return strings.Join(out, " ")
}

// constructVerse constructs a verse for the nth position. Verse should
// probably do this, but its structure is prescribed by the challenge
// and I don't understand it.
func constructVerse(n int) string {
	things := []string{
		"house that Jack built.", "malt", "rat", "cat", "dog",
		"cow with the crumpled horn",
		"maiden all forlorn",
		"man all tattered and torn",
		"priest all shaven and shorn",
		"rooster that crowed in the morn",
		"farmer sowing his corn",
		"horse and the hound and the horn",
	}
	verbs := []string {
		"", "lay in", "ate", "killed", "worried", "tossed", "milked",
		"kissed", "married", "woke", "kept", "belonged to",
	}
	var buf bytes.Buffer
	buf.WriteString("This is the "+things[n])
	for i := n; i > 0; i-- {
        buf.WriteString("\nthat "+verbs[i]+" the "+things[i-1])
    }
    return buf.String() + "\n\n"
}

// Song generates the full text of "The House That Jack Built".
func Song() string {
	var buf bytes.Buffer
	for i := 0; i <= 11; i++ {
		buf.WriteString(constructVerse(i))
	}
	return strings.TrimSpace(buf.String())
}

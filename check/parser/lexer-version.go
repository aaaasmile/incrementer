package parser

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

const (
	EOFRune = -1 // serve al parser per riconoscere la fine del source
)

/// RUN Stack used in Rewind
type runeNode struct {
	r    rune
	next *runeNode
}

type runeStack struct {
	start *runeNode
}

func newRuneStack() runeStack {
	return runeStack{}
}

func (s *runeStack) push(r rune) {
	node := &runeNode{r: r}
	if s.start == nil {
		s.start = node
	} else {
		node.next = s.start
		s.start = node
	}
}

func (s *runeStack) pop() rune {
	if s.start == nil {
		return EOFRune
	} else {
		n := s.start
		s.start = n.next
		return n.r
	}
}

func (s *runeStack) clear() {
	s.start = nil
}

// ********** TokenType ******************

type TokenType int

type Token struct {
	Type  TokenType
	Value string
}

func (tk *Token) String() string {
	switch tk.Type {
	case itemEOF:
		return "EOF"
	case itemError:
		return tk.Value
	}
	if len(tk.Value) > 30 {
		return fmt.Sprintf("%.10q...", tk.Value)
	}
	return fmt.Sprintf("%q", tk.Value)
}

// ********** Lexer ******************
type StateFunc func(*L) StateFunc

type PropInfo struct {
	Keyword   string
	TokenType TokenType
}

type L struct {
	source          string
	start, position int
	state           StateFunc
	tokens          chan Token
	runstack        runeStack
}

func NewL(src string, start StateFunc) *L {
	l := L{
		source:   src,
		state:    start,
		start:    0,
		position: 0,
		runstack: newRuneStack(),
	}
	buffSize := len(l.source) / 2
	if buffSize <= 0 {
		buffSize = 1
	}
	l.tokens = make(chan Token, buffSize)

	return &l
}

func (l *L) current() string {
	return l.source[l.start:l.position]
}

func (l *L) emit(t TokenType) {
	tok := Token{
		Type:  t,
		Value: l.current(),
	}
	l.tokens <- tok
	l.start = l.position
	l.runstack.clear()
}

func (l *L) ignore() {
	l.runstack.clear()
	l.start = l.position
}

func (l *L) peek() rune {
	r := l.next()
	l.rewind()
	return r
}

func (l *L) rewind() {
	r := l.runstack.pop()
	if r > EOFRune {
		size := utf8.RuneLen(r)
		l.position -= size
		if l.position < l.start {
			l.position = l.start
		}
	}
}

func (l *L) next() rune {
	var (
		r rune
		s int
	)
	str := l.source[l.position:]
	if len(str) == 0 {
		r, s = EOFRune, 0
	} else {
		r, s = utf8.DecodeRuneInString(str)
		if r == utf8.RuneError && s == 1 {
			r, s = EOFRune, 0
		}
	}
	l.position += s
	l.runstack.push(r)

	return r
}

func (l *L) nextItem() Token {
	for {
		select {
		case item := <-l.tokens:
			return item
		default:
			if l.state != nil {
				l.state = l.state(l)
			} else {
				return Token{Type: itemEOF, Value: ""}
			}
		}
	}
}

func (l *L) errorf(format string, args ...interface{}) StateFunc {
	l.tokens <- Token{
		itemError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

////////////////////////////////////////////////////////////////////////
/// Lexer Custom Part

const (
	importKey     = "import"
	fromKey       = "from"
	metaPathOpen  = "'"
	metaPathClose = "'"
)

//rebuild token type with stringer -type=TokenType
const (
	itemImportKey TokenType = iota
	itemfromKey
	itemError
	itemEOF
)

/////////////////// State functions ///////////////////////////////////

/////////////////// Parser /////////////////
type TrackedSource struct {
	Name           string
	Dir            string
	ImportChildren []string
	DebugParser    bool
}

func (ts *TrackedSource) ParseSource(source, name, dir string) error {
	ts.Name = name
	ts.Dir = dir
	ts.ImportChildren = make([]string, 0)
	ll := NewL(source, lexStateInit)
	defer close(ll.tokens)
	for {
		item := ll.nextItem()
		if ts.DebugParser {
			fmt.Println("*** type: ", item.Type.String(), item.String())
		}
		switch item.Type {
		case itemError:
			return errors.New(item.Value)
		case itemEOF:
			return nil
		}
	}
	return nil
}

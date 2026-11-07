package books

import (
	
	A	"git.duniter.org/gerard94/util/avl"
	E	"git.duniter.org/gerard94/util/graphQL/exec"
	F	"path/filepath"
	G	"git.duniter.org/gerard94/util/graphQL"
	M	"git.duniter.org/gerard94/util/misc"
	R	"git.duniter.org/gerard94/util/resources"
	S	"git.duniter.org/gerard94/util/sets"
		"math"
		"strings"
		
		/*
		"fmt"
		*/
	
)

const (
	
	typeSystemName = "Books"
	
	tsDir = "util/graphQL/examples/books"
	tsName = "TypeSystem.txt"

)

var (
	
	tsPath = F.Join(R.FindDir(tsDir), tsName)

)

type (
	
	scalarer struct {
	}
	
	typeSystem struct {
		G.TypeSystem
	}

)

type (
	
	iDn int
	
	human interface {
		humanF () *humanC
	}
	
	humanC struct {
		name string
		books *A.Tree // *titleId
	}
	
	author struct {
		humanC
	}
	
	publisher struct {
		humanC
	}
	
	book struct {
		id iDn
		title string
		publishDate string
		authors *A.Tree // *nameSt
		publisher *publisher
	}
	
	nameSt struct {
		name string
		ref human
	}
	
	titleId struct {
		id iDn
		title string
		ref *book
	}
	
	idSt struct {
		id iDn
	}
	
	idBook struct {
		idSt
		ref *book
	}
	
	idStream struct {
		idSt
		stream *G.EventStream
	}
	
	dataBase struct {
		authors, // *nameSt
		publishers, // *nameSt
		books, // *titleId
		booksById *A.Tree // *idBook
		authorsCreatedStreams,
		publishersCreatedStreams,
		booksCreatedStreams *A.Tree // *idStream
	}
	
	eventStreamer interface {
		G.EventStreamer
		eventStreamF () *eventStreamC
	}
	
	eventStreamC struct {
		id iDn
		rootValue *G.OutputObjectValue
		argumentValues *A.Tree // *GValMapItem
		sen G.SourceEventNotification
	}
	
	bookCreatedStream struct {
		eventStreamC
	}
	
	authorCreatedStream struct {
		eventStreamC
	}
	
	publisherCreatedStream struct {
		eventStreamC
	}

)

var (
	
	sl scalarer
	
	idS = S.Full()
	iter = idS.Attach()
	db = new(dataBase)

)

func (a *author) humanF () *humanC {
	return &a.humanC
}

func (p *publisher) humanF () *humanC {
	return &p.humanC
}

func (b *bookCreatedStream) eventStreamF () *eventStreamC {
	return &b.eventStreamC
}

func (a *authorCreatedStream) eventStreamF () *eventStreamC {
	return &a.eventStreamC
}

func (p *publisherCreatedStream) eventStreamF () *eventStreamC {
	return &p.eventStreamC
}

func (n1 *nameSt) Compare (n2 A.Comparer) A.Comp {
	nn2 := n2.(*nameSt)
	name1 := strings.ToUpper(n1.name)
	name2 := strings.ToUpper(nn2.name)
	if name1 < name2 {
		return A.Lt
	}
	if name1 > name2 {
		return A.Gt
	}
	if n1.name < nn2.name {
		return A.Lt
	}
	if n1.name > nn2.name {
		return A.Gt
	}
	return A.Eq
} //Compare

func (n1 *titleId) Compare (n2 A.Comparer) A.Comp {
	nn2 := n2.(*titleId)
	title1 := strings.ToUpper(n1.title)
	title2 := strings.ToUpper(nn2.title)
	if title1 < title2 {
		return A.Lt
	}
	if title1 > title2 {
		return A.Gt
	}
	if n1.title < nn2.title {
		return A.Lt
	}
	if n1.title > nn2.title {
		return A.Gt
	}
	if n1.id < nn2.id {
		return A.Lt
	}
	if n1.id > nn2.id {
		return A.Gt
	}
	return A.Eq
} //Compare

func (i1 *idSt) Compare (i2 A.Comparer) A.Comp {
	ii2 := i2.(*idSt)
	if i1.id < ii2.id {
		return A.Lt
	}
	if i1.id > ii2.id {
		return A.Gt
	}
	return A.Eq
} //Compare

func (i1 *idBook) Compare (i2 A.Comparer) A.Comp {
	ii2 := i2.(*idBook)
	return (&i1.idSt).Compare(&ii2.idSt)
} //Compare

func (i1 *idStream) Compare (i2 A.Comparer) A.Comp {
	ii2 := i2.(*idStream)
	return (&i1.idSt).Compare(&ii2.idSt)
} //Compare

func prefix (s1, s2 string) bool {
	ss1 := []rune(strings.ToUpper(s1))
	ss2 := []rune(strings.ToUpper(s2))
	if len(ss1) > len(ss2) {
		return false
	}
	for i, c := range ss1 {
		if ss2[i] != c {
			return false
		}
	}
	return true
} //prefix

func newId () iDn {
	e, b := iter.FirstE(); M.Assert(b)
	idS.Excl(e)
	return iDn(e)
} //newId

func disposeId (id iDn) {
	idS.Incl(int(id))
} //disposeId

func addStream (t *A.Tree, stream *G.EventStream) {
	id := newId()
	stream.EventStreamer.(eventStreamer).eventStreamF().id = id
	_, ok, _ := t.SearchIns(&idStream{idSt{id}, stream}); M.Assert(!ok)
} //addStream

func disposeStream (t *A.Tree, id iDn) {
	ok := t.Delete(&idStream{idSt: idSt{id}}); M.Assert(ok)
	disposeId(id)
} //disposeStream

func (es *bookCreatedStream) StreamName () string {
	return "bookCreated"
} //StreamName

func (es *bookCreatedStream) RecordNotificationProc (notification G.SourceEventNotification) {
	es.sen = notification
} //RecordNotificationProc

func (es *bookCreatedStream) CloseEvent () {
	disposeStream(db.booksCreatedStreams, es.id)
} //CloseEvent

func (es *authorCreatedStream) StreamName () string {
	return "authorCreated"
} //StreamName

func (es *authorCreatedStream) RecordNotificationProc (notification G.SourceEventNotification) {
	es.sen = notification
} //RecordNotificationProc

func (es *authorCreatedStream) CloseEvent () {
	disposeStream(db.authorsCreatedStreams, es.id)
} //CloseEvent

func (es *publisherCreatedStream) StreamName () string {
	return "publisherCreated"
} //StreamName

func (es *publisherCreatedStream) RecordNotificationProc (notification G.SourceEventNotification) {
	es.sen = notification
} //RecordNotificationProc

func (es *publisherCreatedStream) CloseEvent () {
	disposeStream(db.publishersCreatedStreams, es.id)
} //CloseEvent

func insertName (t *A.Tree, h *human) bool {
	e, b, _ := t.SearchIns(&nameSt{(*h).humanF().name, *h})
	*h = e.Val().(*nameSt).ref
	return !b
} //insertName

func newAuthor (name string) *author {
	return &author{humanC{name, A.New()}}
} //newAuthor

func newPublisher (name string) *publisher {
	return &publisher{humanC{name, A.New()}}
} //newPublisher

func insertTitleId (t *A.Tree, b *book) {
	_, ok, _ := t.SearchIns(&titleId{b.id, b.title, b}); M.Assert(!ok, 100)
} //insertTitleId

func insertIdBook (t *A.Tree, b *book) {
	_, ok, _ := t.SearchIns(&idBook{idSt{b.id}, b}); M.Assert(!ok)
} //insertIdBook

func newBook (title string, id iDn) *book {
	return &book{id: id, title: title, authors: A.New()}
} //newBook

func insertBook (books, publishers, ids, publishersCreatedStreams *A.Tree, title, publishDate, publisherName string) *book {
	id := newId()
	b := newBook(title, id)
	insertIdBook(ids, b)
	insertTitleId(books, b)
	b.publishDate = publishDate
	var h human = newPublisher(publisherName)
	if insertName(publishers, &h) {
		e := publishersCreatedStreams.Next(nil)
		for e != nil {
			el := e.Val().(*idStream)
			sen := el.stream.EventStreamer.(eventStreamer).eventStreamF().sen; M.Assert(sen != nil)
			ov := G.NewOutputObjectValue()
			ov.InsertOutputField("publisherCreated", G.MakeAnyValue(h))
			sen(el.stream, ov)
			e = publishersCreatedStreams.Next(e)
		}
	}
	p := h.(*publisher)
	b.publisher = p
	insertTitleId(p.books, b)
	return b
} //insertBook

func addBookAuthor (authors, authorsCreatedStreams *A.Tree, b *book, authorName string) {
	var h human = newAuthor(authorName)
	if insertName(authors, &h) {
		e := authorsCreatedStreams.Next(nil)
		for e != nil {
			el := e.Val().(*idStream)
			sen := el.stream.EventStreamer.(eventStreamer).eventStreamF().sen; M.Assert(sen != nil)
			ov := G.NewOutputObjectValue()
			ov.InsertOutputField("authorCreated", G.MakeAnyValue(h))
			sen(el.stream, ov)
			e = authorsCreatedStreams.Next(e)
		}
	}
	a := h.(*author)
	insertName(b.authors, &h)
	insertTitleId(a.books, b)
} //addBookAuthor

func createBook (books, authors, publishers, ids, booksCreatedStreams, authorsCreatedStreams, publishersCreatedStreams *A.Tree, ov *G.InputObjectValue) G.Value {
	var (vA *G.ListValue; publishDate, title, publisherName string)
	for f := ov.First(); f != nil; f = ov.Next(f) {
		v := f.Value
		switch f.Name.S {
		case "authors":
			vA = v.(*G.ListValue); M.Assert(vA.First() != nil)
		case "publishDate":
			if v == nil {
				publishDate = ""
			} else {
				publishDate = v.(*G.StringValue).String.S
			}
		case "title":
			title = v.(*G.StringValue).String.S
		case "publisher":
			publisherName = v.(*G.InputObjectValue).First().Value.(*G.StringValue).String.S
		}
	}
	b := insertBook(books, publishers, ids, publishersCreatedStreams, title, publishDate, publisherName)
	for l := vA.First(); l != nil; l = vA.Next(l) {
		name := l.Value.(*G.InputObjectValue).First().Value.(*G.StringValue).String.S
		addBookAuthor(authors, authorsCreatedStreams, b, name)
	}
	e := booksCreatedStreams.Next(nil)
	for e != nil {
		el := e.Val().(*idStream)
		sen := el.stream.EventStreamer.(eventStreamer).eventStreamF().sen; M.Assert(sen != nil)
		ov := G.NewOutputObjectValue()
		ov.InsertOutputField("bookCreated", G.MakeAnyValue(b))
		sen(el.stream, ov)
		e = booksCreatedStreams.Next(e)
	}
	return G.MakeAnyValue(b)
} //createBook

func createBookR (_ G.TypeSystem, rootValue *G.OutputObjectValue, as *A.Tree) G.Value {
	var (v G.Value; ov *G.InputObjectValue)
	ok := G.GetValue(as, "book", &v)
	if ok {
		ov, ok = v.(*G.InputObjectValue)
	}
	if !ok {
		return G.MakeNullValue()
	}
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	return createBook(db.books, db.authors, db.publishers, db.booksById, db.booksCreatedStreams, db.authorsCreatedStreams, db.publishersCreatedStreams, ov)
} //createBookR

func createBooksR (_ G.TypeSystem, rootValue *G.OutputObjectValue, as *A.Tree) G.Value {
	var (v G.Value; lv *G.ListValue)
	ok := G.GetValue(as, "books", &v)
	if ok {
		lv, ok = v.(*G.ListValue)
	}
	if !ok {
		return G.MakeNullValue()
	}
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	l := G.NewListValue()
	for li := lv.First(); li != nil; li = lv.Next(li) {
		l.Append(createBook(db.books, db.authors, db.publishers, db.booksById, db.booksCreatedStreams, db.authorsCreatedStreams, db.publishersCreatedStreams, li.Value.(*G.InputObjectValue)))
	}
	return l
} //createBooksR

func findBook (t *A.Tree, title string) *G.ListValue {
	e, _, _ := t.SearchNext(&titleId{id: iDn(0), title: title})
	l := G.NewListValue()
	for e != nil && prefix(title, e.Val().(*titleId).title) {
		l.Append(G.MakeAnyValue(e.Val().(*titleId).ref))
		e = t.Next(e)
	}
	return l
} //findBook

func findHuman (t *A.Tree, name string) *G.ListValue {
	e, _, _ := t.SearchNext(&nameSt{name: name})
	l := G.NewListValue()
	for e != nil && prefix(name, e.Val().(*nameSt).name) {
		l.Append(G.MakeAnyValue(e.Val().(*nameSt).ref))
		e = t.Next(e)
	}
	return l
} //findHuman

func findAllBooks (t *A.Tree) *G.ListValue {
	l := G.NewListValue()
	e := t.Next(nil)
	for e != nil {
		l.Append(G.MakeAnyValue(e.Val().(*titleId).ref))
		e = t.Next(e)
	}
	return l
} //findAllBooks

func findAllHumans (t *A.Tree) *G.ListValue {
	l := G.NewListValue()
	e := t.Next(nil)
	for e != nil {
		l.Append(G.MakeAnyValue(e.Val().(*nameSt).ref))
		e = t.Next(e)
	}
	return l
} //findAllHumans

func humansR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	l := G.NewListValue()
	lA := findAllHumans(db.authors)
	lP := findAllHumans(db.publishers)
	la := lA.First(); lp := lP.First()
	for la != nil || lp != nil {
		if la != nil && lp != nil {
			if la.Value.(*G.AnyValue).Any.(human).humanF().name <= lp.Value.(*G.AnyValue).Any.(human).humanF().name {
				l.Append(la.Value)
				la = lA.Next(la)
			} else {
				l.Append(lp.Value)
				lp = lP.Next(lp)
			}
		} else if la != nil {
			l.Append(la.Value)
			la = lA.Next(la)
		} else {
			l.Append(lp.Value)
			lp = lP.Next(lp)
		}
	}
	return l
} //humansR

func authorR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	var v G.Value
	if !G.GetValue(argumentValues, "name", &v) {
		return G.MakeNullValue()
	}
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	return findHuman(db.authors, v.(*G.StringValue).String.S)
} //authorR

func authorsR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	return findAllHumans(db.authors)
} //authorsR

func publisherR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	var v G.Value
	if !G.GetValue(argumentValues, "name", &v) {
		return G.MakeNullValue()
	}
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	return findHuman(db.publishers, v.(*G.StringValue).String.S)
} //publisherR

func publishersR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	return findAllHumans(db.publishers)
} //publishersR

func bookR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	var v G.Value
	if !G.GetValue(argumentValues, "title", &v) {
		return G.MakeNullValue()
	}
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	return findBook(db.books, v.(*G.StringValue).String.S)
} //bookR

func booksR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	return findAllBooks(db.books)
} //booksR

func humanNameR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	v := rootValue.First().Value
	switch v := v.(type) {
	case *G.AnyValue:
		return G.MakeStringValue(v.Any.(human).humanF().name)
	case *G.NullValue:
		return v
	default:
		M.Halt(100)
		return nil
	}
} //humanNameR

func humanBookR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	var vv G.Value
	if !G.GetValue(argumentValues, "title", &vv) {
		return G.MakeNullValue()
	}
	v := rootValue.First().Value
	switch v := v.(type) {
	case *G.AnyValue:
		books := v.Any.(human).humanF().books
		return findBook(books, vv.(*G.StringValue).String.S)
	case *G.NullValue:
		return v
	default:
		M.Halt(100)
		return nil
	}
} //humanBookR

func humanBooksR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	v := rootValue.First().Value
	switch v := v.(type) {
	case *G.AnyValue:
		l := G.NewListValue()
		books := v.Any.(human).humanF().books
		e := books.Next(nil)
		for e != nil {
			l.Append(G.MakeAnyValue(e.Val().(*titleId).ref))
			e = books.Next(e)
		}
		return l
	case *G.NullValue:
		return v
	default:
		M.Halt(100)
		return nil
	}
} //humanBooksR

func bookIdR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	v := rootValue.First().Value
	switch v := v.(type) {
	case *G.AnyValue:
		return G.MakeIntValue(int(v.Any.(*book).id))
	case *G.NullValue:
		return v
	default:
		M.Halt(100)
		return nil
	}
} //bookIdR

func bookTitleR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {

	v := rootValue.First().Value
	switch v := v.(type) {
	case *G.AnyValue:
		return G.MakeStringValue(v.Any.(*book).title)
	case *G.NullValue:
		return v
	default:
		M.Halt(100)
		return nil
	}
} //bookTitleR

func bookPublishDateR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	v := rootValue.First().Value
	switch v := v.(type) {
	case *G.AnyValue:
		date := v.Any.(*book).publishDate
		if date == "" {
			return G.MakeNullValue()
		} else {
			return G.MakeStringValue(date)
		}
	case *G.NullValue:
		return v
	default:
		M.Halt(100)
		return nil
	}
} //bookPublishDateR

func bookAuthorR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	var vv G.Value
	if !G.GetValue(argumentValues, "name", &vv) {
		return G.MakeNullValue()
	}
	v := rootValue.First().Value
	switch v := v.(type) {
	case *G.AnyValue:
		authors := v.Any.(*book).authors
		return findHuman(authors, vv.(*G.StringValue).String.S)
	case *G.NullValue:
		return v
	default:
		M.Halt(100)
		return nil
	}
} //bookAuthorR

func bookAuthorsR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	v := rootValue.First().Value
	switch v := v.(type) {
	case *G.AnyValue:
		l := G.NewListValue()
		authors := v.Any.(*book).authors
		e := authors.Next(nil)
		for e != nil {
			l.Append(G.MakeAnyValue(e.Val().(*nameSt).ref))
			e = authors.Next(e)
		}
		return l
	case *G.NullValue:
		return v
	default:
		M.Halt(100)
		return nil
	}
} //bookAuthorsR

func bookPublisherR (_ G.TypeSystem, rootValue *G.OutputObjectValue, argumentValues *A.Tree) G.Value {
	v := rootValue.First().Value
	switch v := v.(type) {
	case *G.AnyValue:
		return G.MakeAnyValue(v.Any.(*book).publisher)
	case *G.NullValue:
		return v
	default:
		M.Halt(100)
		return nil
	}
} //bookPublisherR

func deleteBook (books, authors, publishers, ids *A.Tree, id iDn) *book {
	
	DelBook := func (t *A.Tree, b *book) {
		ok := t.Delete(&titleId{title: b.title, id: b.id}); M.Assert(ok)
	} //DelBook
	
	DelBookById := func (t *A.Tree, b *book) {
		ok := t.Delete(&idBook{idSt: idSt{b.id}}); M.Assert(ok)
	} //DelBookById

	DelHuman := func (t *A.Tree, h human) {
		ok := t.Delete(&nameSt{name: h.humanF().name}); M.Assert(ok)
	} //DelHuman

	//deleteBook
	var (e *A.Elem; ok bool)
	if e, ok, _ = ids.Search(&idBook{idSt: idSt{id}}); !ok {
		return nil
	}
	b := e.Val().(*idBook).ref
	DelBook(b.publisher.books, b)
	if b.publisher.books.IsEmpty() {
		DelHuman(publishers, b.publisher)
	}
	e = b.authors.Next(nil)
	for e != nil {
		el := e.Val().(*nameSt)
		a := el.ref.(*author)
		DelBook(a.books, b)
		if a.books.IsEmpty() {
			DelHuman(authors, a)
		}
		e = b.authors.Next(e)
	}
	DelBook(books, b)
	DelBookById(ids, b)
	disposeId(b.id)
	return b
} //deleteBook

func deleteBookR (_ G.TypeSystem, rootValue *G.OutputObjectValue, as *A.Tree) G.Value {
	var (v G.Value; i *G.IntValue)
	ok := G.GetValue(as, "id", &v)
	if ok {
		i, ok = v.(*G.IntValue)
	}
	if !ok {
		return G.MakeNullValue()
	}
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	b := deleteBook(db.books, db.authors, db.publishers, db.booksById, iDn(i.Int))
	if b == nil {
		return G.MakeNullValue()
	}
	return G.MakeAnyValue(b)
} //deleteBookR

func deleteBooksR (_ G.TypeSystem, rootValue *G.OutputObjectValue, as *A.Tree) G.Value {
	var (v G.Value; lv *G.ListValue)
	ok := G.GetValue(as, "ids", &v)
	if ok {
		lv, ok = v.(*G.ListValue)
	}
	if !ok {
		return G.MakeNullValue()
	}
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	l := G.NewListValue()
	for li := lv.First(); li != nil; li = lv.Next(li) {
		v := li.Value
		ok := v != nil
		if ok {
			_, ok = v.(*G.InputObjectValue)
		}
		if !ok {
			l.Append(G.MakeNullValue())
		} else {
			b := deleteBook(db.books, db.authors, db.publishers, db.booksById, iDn(v.(*G.InputObjectValue).First().Value.(*G.IntValue).Int))
			if b == nil {
				l.Append(G.MakeNullValue())
			} else {
				l.Append(G.MakeAnyValue(b))
			}
		}
	}
	return l
} //deleteBooksR

func createdR (_ G.TypeSystem, rootValue *G.OutputObjectValue, as *A.Tree) G.Value {
	return rootValue.First().Value
} //createdR

func bookCreatedStreamResolver (rootValue *G.OutputObjectValue, argumentValues *A.Tree) *G.EventStream { // ValMapItem
	bs := &bookCreatedStream{eventStreamC{rootValue: rootValue, argumentValues: argumentValues, sen: nil}}
	es := G.MakeEventStream(bs)
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	addStream(db.booksCreatedStreams, es)
	return es
} //bookCreatedStreamResolver

func authorCreatedStreamResolver (rootValue *G.OutputObjectValue, argumentValues *A.Tree) *G.EventStream { // ValMapItem
	as := &authorCreatedStream{eventStreamC{rootValue: rootValue, argumentValues: argumentValues, sen: nil}}
	es := G.MakeEventStream(as)
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	addStream(db.authorsCreatedStreams, es)
	return es
} //authorCreatedStreamResolver

func publisherCreatedStreamResolver (rootValue *G.OutputObjectValue, argumentValues *A.Tree) *G.EventStream { // ValMapItem
	ps := &publisherCreatedStream{eventStreamC{rootValue: rootValue, argumentValues: argumentValues, sen: nil}}
	es := G.MakeEventStream(ps)
	db := rootValue.First().Value.(*G.AnyValue).Any.(*dataBase)
	addStream(db.publishersCreatedStreams, es)
	return es
} //publisherCreatedStreamResolver

func (ts *typeSystem) fixFieldResolvers () {
	ts.FixFieldResolver("Mutation", "createBook", createBookR)
	ts.FixFieldResolver("Mutation", "createBooks", createBooksR)
	ts.FixFieldResolver("Mutation", "deleteBook", deleteBookR)
	ts.FixFieldResolver("Mutation", "deleteBooks", deleteBooksR)
	ts.FixFieldResolver("Subscription", "bookCreated", createdR)
	ts.FixFieldResolver("Subscription", "authorCreated", createdR)
	ts.FixFieldResolver("Subscription", "publisherCreated", createdR)
	ts.FixFieldResolver("Query", "authors", authorsR)
	ts.FixFieldResolver("Query", "author", authorR)
	ts.FixFieldResolver("Query", "publishers", publishersR)
	ts.FixFieldResolver("Query", "publisher", publisherR)
	ts.FixFieldResolver("Query", "books", booksR)
	ts.FixFieldResolver("Query", "book", bookR)
	ts.FixFieldResolver("Query", "authorsOrPublishers", humansR)
	ts.FixFieldResolver("Query", "humans", humansR)
	ts.FixFieldResolver("Query", "hasBooks", humansR)
	ts.FixFieldResolver("Author", "name", humanNameR)
	ts.FixFieldResolver("Author", "book", humanBookR)
	ts.FixFieldResolver("Author", "books", humanBooksR)
	ts.FixFieldResolver("Publisher", "name", humanNameR)
	ts.FixFieldResolver("Publisher", "book", humanBookR)
	ts.FixFieldResolver("Publisher", "books", humanBooksR)
	ts.FixFieldResolver("Book", "id", bookIdR)
	ts.FixFieldResolver("Book", "title", bookTitleR)
	ts.FixFieldResolver("Book", "publishDate", bookPublishDateR)
	ts.FixFieldResolver("Book", "author", bookAuthorR)
	ts.FixFieldResolver("Book", "authors", bookAuthorsR)
	ts.FixFieldResolver("Book", "publisher", bookPublisherR)
} //fixFieldResolvers

func (ts *typeSystem) fixStreamResolvers () {
	ts.FixStreamResolver("bookCreated", bookCreatedStreamResolver)
	ts.FixStreamResolver("authorCreated", authorCreatedStreamResolver)
	ts.FixStreamResolver("publisherCreated", publisherCreatedStreamResolver)
} //fixStreamResolvers

func iDnCoercer (ts G.TypeSystem, value G.Value, path G.Path, cValue *G.Value) bool {
	switch value := value.(type) {
	case *G.IntValue:
		if math.MinInt32 <= value.Int && value.Int <= math.MaxInt32 {
			*cValue = value;
			return true
		}
		ts.MappedError("IntOutOfRange", "", "", nil, path)
	default:
		ts.MappedError("NonIntValueForIntType", "", "", nil, path)
	}
	return false
} //iDnCoercer

func (scalarer) FixScalarCoercer (scalarName string, sc *G.ScalarCoercer) {
	switch scalarName {
	case "IDn":
		*sc = iDnCoercer
	default:
	}
}

func unsubscribeAll (t *A.Tree) { // idStream
	e := t.Next(nil)
	for e != nil {
		ee := t.Next(e)
		el := e.Val().(*idStream)
		G.Unsubscribe(el.stream.ResponseStreams[0])
		e = ee
	}
} //unsubscribeAll

func UnsubscribeAllBookCreated () {
	unsubscribeAll(db.booksCreatedStreams)
} //UnsubscribeAllBookCreated

func UnsubscribeAllAuthorCreated () {
	unsubscribeAll(db.authorsCreatedStreams)
} //UnsubscribeAllAuthorCreated

func UnsubscribeAllPublisherCreated () {
	unsubscribeAll(db.publishersCreatedStreams)
} //UnsubscribeAllPublisherCreated

func abstractTypeResolver (ts G.TypeSystem, td G.TypeDefinition, ov *G.OutputObjectValue) *G.ObjectTypeDefinition {
	ok := ov != nil && ov.First() != nil && ov.Next(ov.First()) == nil
	var a *G.AnyValue
	if ok {
		a, ok = ov.First().Value.(*G.AnyValue)
	}
	if ok {
		ok = a.Any != nil
	}
	if !ok {
		return nil
	}
	var name string
	switch a.Any.(type) {
	case *author:
		name = "Author"
	case *publisher:
		name = "Publisher"
	default:
		M.Halt(100)
	}
	d := ts.GetTypeDefinition(name)
	ok = d != nil
	var od *G.ObjectTypeDefinition
	if ok {
		od, ok = d.(*G.ObjectTypeDefinition)
	}
	if !ok {
		return nil
	}
	return od
} //abstractTypeResolver

func fixResolvers (ts G.TypeSystem) {
	tss := ts.(*typeSystem)
	tss.fixFieldResolvers()
	tss.fixStreamResolvers()
	ts.FixAbstractTypeResolver(abstractTypeResolver)
} //fixResolvers

func NewTypeSystem () *E.Executor {

	const (
		
		rootName = "root"

	)
	
	ts := new(typeSystem)
	ts.TypeSystem = G.Dir.NewTypeSystem(sl)
	ov := G.NewOutputObjectValue()
	ov.InsertOutputField(rootName, G.MakeAnyValue(db))
	exe := E.NewExecutor()
	exe.SetTypeSystem(ts, typeSystemName, tsPath, fixResolvers, ov)
	return exe
} //NewTypeSystem

func Initialize () {
	db.authors = A.New()
	db.publishers = A.New()
	db.books = A.New()
	db.booksById = A.New()
	db.authorsCreatedStreams = A.New()
	db.publishersCreatedStreams = A.New()
	db.booksCreatedStreams = A.New()
} //Initialize

func init () {
	Initialize()
} //init

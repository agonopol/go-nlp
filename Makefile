include $(GOROOT)/src/Make.inc

TARG=nlp
GOFILES=\
	ngram.go \
	utils.go \
	document.go \
	stoplist.go \
	html.go \

include $(GOROOT)/src/Make.pkg

module bozdoz.com/network

go 1.14

replace bozdoz.com/spreadsheet => ./spreadsheet

replace bozdoz.com/logger => ./logger

require (
	bozdoz.com/logger v0.0.0-00010101000000-000000000000
	bozdoz.com/spreadsheet v0.0.0-00010101000000-000000000000
	cloud.google.com/go v0.72.0 // indirect
	golang.org/x/sys v0.0.0-20201113233024-12cec1faf1ba // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20201113130914-ce600e9a6f9e // indirect
)

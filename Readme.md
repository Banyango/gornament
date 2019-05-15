# Gornament

Web Component templator.

Use Gornament to inject your webcomponent `<tempalte>` as a bundle into your html pages.

`<link rel="import" href="/path/to/imports/stuff.html">` has been deprecated in Chrome and not implemented in Firefox. 

I wanted a simple way to bundle up all my WebComponents and place them into my html without having to resort to Webpack or otherwise. So I wrote this little command line tool. 

Hopefully someone else finds it handy.

### How to Build

Install Go

`go build gornament.go` 

### Usage

In your project dir 
`./gornament`

-> gornament will look in
 - ./components for html templates.
 - ./templates for the html to add bundle to. Add {{ .Templates }} in the `<head>` of your html page.
 
 -> gornament will output the html in ./
 
### Example

Check out the example dir for an example.


 

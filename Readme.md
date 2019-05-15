# Gornament

Web Component templator.

Use Gornament to inject your webcomponent `<tempalte>` as a bundle into your html pages.

`<link rel="import" href="/path/to/imports/stuff.html">` has been deprecated in Chrome and not implemented in Firefox. 

I wanted a simple way to bundle up all my WebComponents and place them into my html without having to resort to ES6 or otherwise (Doing really simple pure html pages for clients and I don't want to spin up all that npm stuff). So I wrote this little command line tool. 

Hopefully someone else finds it handy.

### Usage

In your project dir 
`./gornament`

-> gornament will look in
 - ./components for html templates.
 - ./templates for the html to add bundle to.
 
 -> gornament will output the html in ./
 
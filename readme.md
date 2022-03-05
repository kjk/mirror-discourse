# Static HTML mirror of Discourse forums

This is a Go program that creates a simple, static HTML mirror of Discourse forum.

Creating a static copy of Discourse forum with web mirroring tools like `wget` is pretty much impossible (I tried).

This tool uses Discourse's JSON APIs to overcome this.

Example mirror: https://forum.sumatrapdfreader.org/

# How to run it

Let's assume your Discourse forum is hosted on https://myforum.mydomain.com

Run:
```
mirror-discourse https://myforum.mydomain.com
```
This creates a static HTML mirror in `www` directory and starts a local server to preview it at http://localhost:8777

You can then deploy such website publicly using e.g. https://render.com or netlify or vercel.

## Options

Available cmd-line options:
* `-dir <directory>` : changes the output directory from `www`. Will also create `cache` directory as a sibiling directory
* `-limit <n>` : a large forum will take a long time to process. If you're just testing, use `-limit 1` to only process first page of results
* `-banner <banner.html>` : you can add a piece of HTML to be shown at the top of generated `index.html` e.g. a notice that this is read-only forum and maybe a link to new forum. See `sumatra_forum_banner.html` for example. Alternatively, you can edit `tmpl_main.html` or edit `index.html` after it was generated.

## If you have Go compiler installed

### Option 1

```
go install https://github.com/kjk/mirror-discourse
mirror-discourse https://myforum.mydomain.com
```

### Option 2

Download the code locally and run:
```
go run . https://myforum.mydomain.com
```

## If you don't have Go compiler installed

You can open the code in e.g. Gitpod: https://gitpod.io/#https://github.com/kjk/mirror-discourse

Then run:
```
go run . https://myforum.mydomain.com
```

You can then download `meta_discourse` directory locally.

# Customize the output

You can change:
* `main.css`
* `tmpl_main.html` : template for home page
* `tmpl_topic.html` : template for topic page

In `main.go` you can also change `banner_html` to html fragment displayed at the top of every page. Alternatively, change `tmpl_main.html` and `tmpl_topic.html`.

I used it to point people to new forum. 

# Inspiration

This program is a port of https://www.marksmath.org/ArchiveDiscourse/

ArchiveDiscourse is a Jupiter notebook and I just coulddn't figure out how to run it locally.

So I ported it instead.

# No support

This code worked for me.

If it doesn't work for you, tough.

I will not help you, so don't ask.

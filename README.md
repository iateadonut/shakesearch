## Working version

https://lit-anchorage-03612.herokuapp.com/ - Try searching "oh my sons horatio" for a set of results.

## Next steps (also, "If I had more time...")

- Click the line and it brings you to that place in the poem/play.
- Improve search engine algorithm.  Modern readers may expect to see 'O' when they type 'Oh', which is already built in.  What else might they want?

## apache solr

This is clearly a task for a search engine.

download here: https://www.apache.org/dyn/closer.lua/lucene/solr/8.7.0/solr-8.7.0.tgz

from webroot:
wget https://downloads.apache.org/lucene/solr/8.7.0/solr-8.7.0.tgz

tar xf solr-8.7.0.tgz

bin/solr delete -c shakespeare

bin/solr create -c shakespeare

git checkout solr-8.7.0/server/solr/shakespeare/conf/synonyms.txt solr-8.7.0/server/solr/shakespeare/conf/DIHconfigfile.shakespeare.xml solr-8.7.0/server/solr/shakespeare/conf/managed-schema solr-8.7.0/server/solr/shakespeare/conf/solrconfig.xml

to import:
curl -X POST -H 'Content-Type: application/json' http://localhost:8983/solr/shakespeare/dataimport?command=full-import

to delete before re-importing:
curl -X POST -H 'Content-Type: application/json' --data-binary '{"delete":{"query":"*:*" }}' http://localhost:8983/solr/shakespeare/update

modify the synonyms file:  Maybe the modern english-speaker would think that "o" and "oh" are the same words?  We want her to find "o, my son" if she searches "oh my son".

# ShakeSearch

Welcome to the Pulley Shakesearch Take-home Challenge! In this repository,
you'll find a simple web app that allows a user to search for a text string in
the complete works of Shakespeare.

You can see a live version of the app at
https://pulley-shakesearch.herokuapp.com/. Try searching for "Hamlet" to display
a set of results.

In it's current state, however, the app is just a rough prototype. The search is
case sensitive, the results are difficult to read, and the search is limited to
exact matches.

## Your Mission

Improve the search backend. Think about the problem from the user's perspective
and prioritize your changes according to what you think is most useful.

## Submission

1. Fork this repository and send us a link to your fork after pushing your changes. 
2. Heroku hosting - The project includes a Heroku Procfile and, in its
current state, can be deployed easily on Heroku's free tier.
3. In your submission, share with us what changes you made and how you would prioritize changes if you had more time.



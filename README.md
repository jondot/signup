# Signup

A Go based web service to collect signup emails. Great for when you need to
set up a quick beta launch/signup landing page.

Easily deployable to Heroku.

Demo <a href="http://jondot.github.com/signup">is here</a>, demo source <a href="https://github.com/jondot/signup/blob/master/example/signup.html">is here</a>.


## Getting Started

Signup can be easily deployed on Heroku:

    $ git clone git://github.com/jondot/signup.git
    $ heroku create -s cedar --buildpack git@github.com:kr/heroku-buildpack-go.git

Now go ahead and add a MongoDB plugin (which ever you prefer), then we
continue:

    $ heroku config:add HOST=http://<app-name>.herokuapp.com MONGO_HOST=mongodb://user:pw@<host> MONGO_DB=signup PROJECTS=<project-name>
    $ git push heroku master

Visit `http://app-name.herokuapp.com` and see if you get a `pong.`.

## Usage

You need to POST an `email` to a project that is in your `PROJECTS` list.

    $ curl -d "email=foo@example.com" http://app-name.herokuapp.com/project-name

### Ajax: Using Precooked Script

The signup server has a *special* endpoint that you can use to save
yourself from manually placing the javascript POST code.

It will inject a snippet of javascript which will hook onto your form:

    <script type="text/javascript" src="http://app-name.herokuapp.com/script/project-name/%23form-id"></script>

To handle flash errors/success messages, by default, on success it will show an element with the `signup-success`
ID, and on failure it will show an element with the `signup-failure` id.

To reset errors/success, it will hide elements which has the `signup-flash` class.





## Details

Here is a complete listing of the environment variables:

### Environment

    HOST        http://my-nifty-signup.heroku.com
    MONGO_HOST  http://example.com
    MONGO_DB    signup_prod
    PROJECTS    copycat bundleupon
    KIOSK_MODE* set/unset

Note that `HOST` is needed in the case where you're using the inject
feature. This is where you want the injected javascript code to `POST` at.

On heroku it is simply `$ heroku config:add VAR=VALUE` for each of
these that you want.

* You only need to set `KIOSK_MODE` to any value to avoid saving contacts (running in demo
mode).



### Build Pack and Hosting

This uses the Go build pack:

    $ heroku create -s cedar --buildpack git@github.com:kr/heroku-buildpack-go.git

Alternatively, You can deploy this service very easily on
your own server.   
Since this is Go, everything should be in a self contained binary. Just `go get`, `go install` and run `signup`.




# Contributing

Fork, implement, add tests, pull request, get my everlasting thanks and a respectable place here :).

* Thanks to zeebo, jessta and rogpeppe of #go-nuts for helpful feedback!.

# Copyright


Copyright (c) 2012 [Dotan Nahum](http://gplus.to/dotan) [@jondot](http://twitter.com/jondot). See MIT-LICENSE for further details.


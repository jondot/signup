# Signup

A Go based service to take in beta signups.


## Deploying

Signup can be deployed on Heroku, and using the free-tier
MongoDB plugins, it can provide for storing contacts, so that you get a 0-maintenance , 0-cost
service.


You'll need to add these environment variables. for example:

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



### Details
This uses the Go build pack. 

    heroku create -s cedar --buildpack git@github.com:kr/heroku-buildpack-go.git

Alternatively, You can deploy this service very easily on
your own server.   
Since this is Go, everything should be in a self contained binary. Just `go get`, `go install` and run `signup`.


## Usage

You need to POST an `email` to a project that is in your `PROJECTS` list.

    $ curl -vv -d "email=foo@example.com"
    http://signup.your-host.com:5000/your_project

## On a Page

If you have some kind of landing page, you can easily POST the same
details using ajax:

      $(function(){
          $('#signup-form').submit(function(){
            $.post('http://your-heroku-signup-app.herokuapp.com/your-project-id', $(this).serialize(), function(data){
                  // YAY!
              })
              .error(function(){
                  // BOOHOO :(
              });
            return false;
          });
      });

Given that you have a form element with `signup-form` as ID, and an
`input` with a `name` of `email` it should work seamlessly.


## Using Precooked Script

The signup server has a *special* endpoint that you can use to save
yourself from manually placing the javascript POST code.

It will inject a snippet of javascript which will hook onto your form:

    <script type="text/javascript" src="http://your-heroku-signup-app.herokuapp.com/script/your-project-id/%23form-id"></script>

By default, on success it will show an element with the `signup-success`
ID, and on failure it will show an element with the `signup-failure` id.

It will hide elements which has the `signup-flash` class.


You can tweak this behaviour:

    Signup.reset(function(){
      //CLEAR
    });

    Signup.success(function(){
      //YAY!
    });

    Signup.failure(function(){
      //BOOHOO!
    });

# Contributing

Fork, implement, add tests, pull request, get my everlasting thanks and a respectable place here :).


# Copyright


Copyright (c) 2012 [Dotan Nahum](http://gplus.to/dotan) [@jondot](http://twitter.com/jondot). See MIT-LICENSE for further details.


#!/bin/env python3

import flask

APP = flask.Flask(__name__)

@APP.route("/")
def index():
	return flask.render_template("index.htm")

@APP.route("/hello/<name>")
def hello(name):
	return flask.render_template("hello.htm", name=name)

if __name__ == "__main__" :
	APP.debug = True
	APP.run()


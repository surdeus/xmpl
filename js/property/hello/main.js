#!/usr/bin/env qjs

function Dog(name){
	this.bark = function() {
		console.log(`Bark! My name is ${this.name}`);
	};
	this.__defineSetter__("name", function(val){
		this._name = "Doggie " + val ;
	});
	this.__defineGetter__("name", () =>{
		return this._name ;
	});
	this.name = name ;
}

dog1 = new Dog("Dawg") ;
dog1.bark();

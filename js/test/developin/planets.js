#!/usr/bin/env qjs

"use strict"

function Planet(name, mass, sunDist, hasRings) {
	var fields = ["name", "mass", "sunDist", "hasRings"] ;
	fields.forEach(el => {
		this[el] = eval(el) ;
	});
	this.toString = () => {
		return JSON.stringify(this) ;
	}
}

var planets = [
	new Planet("Mercury", 2.285e23, 5.8e10, false),
	new Planet("Venus", 4.867e24, 1.082e11, false),
	new Planet("Earth", 5.972e24, 1.495e11, false),
	new Planet("Mars", 6.4e23, 2.279e11, false),
	new Planet("Jupiter", 1.9e27, 7.78e11, true),
	new Planet("Saturn", 5.68e26, 1.427e12, true),
	new Planet("Uran", 8.69e25, 2.871e12, true),
	new Planet("Neptune", 1.02e26, 4.497e12, true),
]

function sortArrayByField(arr, field, order=true) {
	var sortFunc = order ? (a, b) => {
			return a[field] > b[field] ? 1 : -1 ;
		} : (a, b) => {
			return a[field] < b[field] ? 1 : -1 ;
		}
	;
	arr.sort(sortFunc);
}

function randomInt(max) {
	return Math.floor(Math.random() * max)
}

function printArray(arr, field=null) {
	if(field)
		arr.forEach((el) => {
			console.log(el[field])
		});
	else
		arr.forEach((el) => {
			console.log(el);
		});
}

function filteredArray(arr, filterFunc) {
	var ret = [];
	arr.forEach( (el) => {
		if(filterFunc(el))
			ret.push(el);
	});
	return ret
}

// 3
function printPlanetsInOrderFromSun(arr){
	var cpy = arr.slice() ;
	sortArrayByField(cpy, "sunDist");
	printArray(cpy, "name");
}

// 4
function printPlanetsWithoutRings(arr) {
	var cpy = filteredArray(arr, (el) =>{
		return !el.hasRings ;
	}) ;
	printArray(cpy, "name");
}

function findMax(arr, cmpFn){
	var ret = arr[0] ;
	arr.forEach((el) => {
		if(cmpFn(el, ret) == 1)
			ret = el;
	});
	return ret ;
}

function findByVal(arr, val, field) {
	var ret = null ;
	arr.some( (el) => {
		if(el[field] == val){
			ret = el;
			return true;
		}
		return false ;
	});
	return ret;
}

// 5
function printPlanetsWithMinAndMaxMasses(arr) {
	var min = findMax(arr, (a, b) => {
		return a.mass < b.mass ? 1 : -1 ;
	}) ;
	var max = findMax(arr, (a, b) => {
		return a.mass > b.mass ? 1 : -1 ;
	}) ;
	console.log(min.name)
	console.log(max.name)
}

// 6
function minAndMaxDistancesBetweenPlanets(a, b) {
	var ret = [] ;
	ret.push(Math.abs(a.sunDist - b.sunDist));
	ret.push(a.sunDist + b.sunDist);
	return ret ;
}

// 7
function theMostDistancedPlanetAstronautCanReach(ps, fuelDist) {
	var earth = findByVal(ps, "Earth", "name")
	var cpy = planets.slice()
	sortArrayByField(cpy, "sunDist", false);
	var reachablePlanet = null ;
	cpy.some((el)=>{
		var [min, max] = minAndMaxDistancesBetweenPlanets(earth, el) ;
		if(min <= fuelDist ){
			reachablePlanet = el ;
			return true ;
		}
		return false ;
	});

	if(reachablePlanet != null && reachablePlanet != earth){
		console.log(`The most distanced planet we can reach with ${fuelDist} is ${reachablePlanet.name}`);
	} else {
		console.log(`We cannot reach any planet with ${fuelDist}`);
	}
}

console.log("Planets in order from the sun:")
printPlanetsInOrderFromSun(planets);
console.log();

console.log("Planets without rings:")
printPlanetsWithoutRings(planets);
console.log();

console.log("Planets with min and max masses:");
printPlanetsWithMinAndMaxMasses(planets);
console.log();

var [a, b] = [planets[randomInt(planets.length-1)], planets[randomInt(planets.length-1)]] ;
var [min, max] = minAndMaxDistancesBetweenPlanets(a, b);
console.log(`Min and max distances between ${a.name} and ${b.name} are ${min} and ${max}`);
console.log()

theMostDistancedPlanetAstronautCanReach(planets, 4e10)
theMostDistancedPlanetAstronautCanReach(planets, 1.5e11)
theMostDistancedPlanetAstronautCanReach(planets, 1.6e12)

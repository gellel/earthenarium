package climate

import (
	"github.com/gellel/earthenarium/hemisphere"
	"github.com/gellel/earthenarium/season"
)

var (
	Arctic = Climate{
		// polar
		hemisphere.Arctic[0]: {
			season.Winter: -15,
			season.Spring: -10,
			season.Summer: -5,
			season.Autumn: -10},
		// tundra
		hemisphere.Arctic[1]: {
			season.Winter: -9,
			season.Spring: -2,
			season.Summer: 1,
			season.Autumn: -1},
		// cooler-temperate
		hemisphere.Arctic[2]: {
			season.Winter: -2,
			season.Spring: 9,
			season.Summer: 15,
			season.Autumn: 6}}
)
var (
	Cancer = Climate{
		// cooler-temperate
		hemisphere.Cancer[0]: {
			season.Winter: 8,
			season.Spring: 13,
			season.Summer: 19,
			season.Autumn: 11},
		// mild-temperate
		hemisphere.Cancer[1]: {
			season.Winter: 12,
			season.Spring: 18,
			season.Summer: 25,
			season.Autumn: 14},
		// warm-temperate
		hemisphere.Cancer[2]: {
			season.Winter: 15,
			season.Spring: 24,
			season.Summer: 28,
			season.Autumn: 21}}
)
var (
	Equator = Climate{
		// northern-equator
		hemisphere.Equator[0]: {
			season.Winter: 25,
			season.Spring: 28,
			season.Summer: 31,
			season.Autumn: 26},
		// centre-equator
		hemisphere.Equator[1]: {
			season.Winter: 29,
			season.Spring: 32,
			season.Summer: 35,
			season.Autumn: 29},
		// lower-equator
		hemisphere.Equator[2]: {
			season.Winter: 23,
			season.Spring: 26,
			season.Summer: 29,
			season.Autumn: 25},
	}
)
var (
	Capricorn = Climate{
		// warm-temperate
		hemisphere.Capricorn[2]: {
			season.Winter: 15,
			season.Spring: 24,
			season.Summer: 28,
			season.Autumn: 21},
		// mild-temperate
		hemisphere.Capricorn[1]: {
			season.Winter: 12,
			season.Spring: 18,
			season.Summer: 25,
			season.Autumn: 14},
		// cool-temperate
		hemisphere.Capricorn[0]: {
			season.Winter: 7,
			season.Spring: 12,
			season.Summer: 16,
			season.Autumn: 12}}
)
var (
	Antarctic = Climate{
		// cooler-temperate
		hemisphere.Antarctic[2]: {
			season.Winter: -2,
			season.Spring: 6,
			season.Summer: 13,
			season.Autumn: 5},
		// tundra
		hemisphere.Antarctic[1]: {
			season.Winter: -10,
			season.Spring: -3,
			season.Summer: 1,
			season.Autumn: -2},
		// polar
		hemisphere.Antarctic[0]: {
			season.Winter: -20,
			season.Spring: -13,
			season.Summer: -10,
			season.Autumn: -7}}
)
var (
	Zones = Zone{
		hemisphere.Arctic.Key():    Arctic,
		hemisphere.Cancer.Key():    Cancer,
		hemisphere.Equator.Key():   Equator,
		hemisphere.Capricorn.Key(): Capricorn,
		hemisphere.Antarctic.Key(): Antarctic}
)

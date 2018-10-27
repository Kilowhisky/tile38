package tests

import (
	"testing"
)

func subTestSearch(t *testing.T, mc *mockServer) {
	runStep(t, mc, "KNN", keys_KNN_test)
	runStep(t, mc, "WITHIN_CIRCLE", keys_WITHIN_CIRCLE_test)
	runStep(t, mc, "INTERSECTS_CIRCLE", keys_INTERSECTS_CIRCLE_test)
	runStep(t, mc, "WITHIN", keys_WITHIN_test)
	runStep(t, mc, "INTERSECTS", keys_INTERSECTS_test)
}

func keys_KNN_test(mc *mockServer) error {
	return mc.DoBatch([][]interface{}{
		{"SET", "mykey", "1", "POINT", 5, 5}, {"OK"},
		{"SET", "mykey", "2", "POINT", 19, 19}, {"OK"},
		{"SET", "mykey", "3", "POINT", 12, 19}, {"OK"},
		{"SET", "mykey", "4", "POINT", -5, 5}, {"OK"},
		{"SET", "mykey", "5", "POINT", 33, 21}, {"OK"},
		{"NEARBY", "mykey", "LIMIT", 10, "POINTS", "POINT", 20, 20}, {
			"[0 [[2 [19 19]] [3 [12 19]] [5 [33 21]] [1 [5 5]] [4 [-5 5]]]]"},
		{"NEARBY", "mykey", "LIMIT", 10, "IDS", "POINT", 20, 20, 4000000}, {"[0 [2 3 5 1 4]]"},
		{"NEARBY", "mykey", "LIMIT", 10, "IDS", "POINT", 20, 20, 1500000}, {"[0 [2 3 5]]"},
	})
}

func keys_WITHIN_test(mc *mockServer) error {
	return mc.DoBatch([][]interface{}{
		{"SET", "mykey", "point1", "POINT", 37.7335, -122.4412}, {"OK"},
		{"SET", "mykey", "point2", "POINT", 37.7335, -122.44121}, {"OK"},
		{"SET", "mykey", "line3", "OBJECT", `{"type":"LineString","coordinates":[[-122.4408378,37.7341129],[-122.4408378,37.733]]}`}, {"OK"},
		{"SET", "mykey", "poly4", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.4408378,37.7341129],[-122.4408378,37.733],[-122.44,37.733],[-122.44,37.7341129],[-122.4408378,37.7341129]]]}`}, {"OK"},
		{"SET", "mykey", "multipoly5", "OBJECT", `{"type":"MultiPolygon","coordinates":[[[[-122.4408378,37.7341129],[-122.4408378,37.733],[-122.44,37.733],[-122.44,37.7341129],[-122.4408378,37.7341129]]],[[[-122.44091033935547,37.731981251280985],[-122.43994474411011,37.731981251280985],[-122.43994474411011,37.73254976045042],[-122.44091033935547,37.73254976045042],[-122.44091033935547,37.731981251280985]]]]}`}, {"OK"},
		{"SET", "mykey", "point6", "POINT", -5, 5}, {"OK"},
		{"SET", "mykey", "point7", "POINT", 33, 21}, {"OK"},
		{"SET", "mykey", "poly8", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.4408378,37.7341129],[-122.4408378,37.733],[-122.44,37.733],[-122.44,37.7341129],[-122.4408378,37.7341129]],[[-122.44060993194579,37.73345766902749],[-122.44044363498686,37.73345766902749],[-122.44044363498686,37.73355524732416],[-122.44060993194579,37.73355524732416],[-122.44060993194579,37.73345766902749]],[[-122.44060724973677,37.7336888869566],[-122.4402102828026,37.7336888869566],[-122.4402102828026,37.7339752567853],[-122.44060724973677,37.7339752567853],[-122.44060724973677,37.7336888869566]]]}`}, {"OK"},
		{"WITHIN", "mykey", "IDS", "OBJECT",
			`{
				"type": "Polygon",
				"coordinates": [
					[
						[-122.44126439094543,37.72906137107],
						[-122.43980526924135,37.72906137107],
						[-122.43980526924135,37.73421283683962],
						[-122.44126439094543,37.73421283683962],
						[-122.44126439094543,37.72906137107]
					]
				]
			}`}, {"[0 [point1 point2 line3 poly4 multipoly5 poly8]]"},

		{"SET", "key2", "poly9", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44037926197052,37.73313523548048],[-122.44017541408539,37.73313523548048],[-122.44017541408539,37.73336857568778],[-122.44037926197052,37.73336857568778],[-122.44037926197052,37.73313523548048]]]}`}, {"OK"},
		{"SET", "key2", "poly10", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44040071964262,37.73359343010089],[-122.4402666091919,37.73359343010089],[-122.4402666091919,37.73373767596864],[-122.44040071964262,37.73373767596864],[-122.44040071964262,37.73359343010089]]]}`}, {"OK"},
		{"WITHIN", "key2", "IDS", "GET", "mykey", "poly8"}, {"[0 [poly9]]"},

		{"SET", "key3", "poly11", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44059920310973,37.733279482240874],[-122.4402344226837,37.733279482240874],[-122.4402344226837,37.73375464605226],[-122.44059920310973,37.73375464605226],[-122.44059920310973,37.733279482240874]]]}`}, {"OK"},
		{"SET", "key3", "poly12", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44060993194579,37.73238005667773],[-122.44013786315918,37.73238005667773],[-122.44013786315918,37.73316917591997],[-122.44060993194579,37.73316917591997],[-122.44060993194579,37.73238005667773]]]}`}, {"OK"},
		{"WITHIN", "key3", "IDS", "GET", "mykey", "multipoly5"}, {"[0 [poly11]]"},

		{"SET", "key5", "poly13", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44073867797852,37.733211601447465],[-122.44011640548705,37.733211601447465],[-122.44011640548705,37.7340516218859],[-122.44073867797852,37.7340516218859],[-122.44073867797852,37.733211601447465]],[[-122.44060993194579,37.73345766902749],[-122.44060993194579,37.73355524732416],[-122.44044363498686,37.73355524732416],[-122.44044363498686,37.73345766902749],[-122.44060993194579,37.73345766902749]],[[-122.44060724973677,37.7336888869566],[-122.44060724973677,37.7339752567853],[-122.4402102828026,37.7339752567853],[-122.4402102828026,37.7336888869566],[-122.44060724973677,37.7336888869566]]]}`}, {"OK"},
		{"SET", "key5", "poly14", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44154334068298,37.73179457567642],[-122.43935465812682,37.73179457567642],[-122.43935465812682,37.7343740514423],[-122.44154334068298,37.7343740514423],[-122.44154334068298,37.73179457567642]],[[-122.44104981422423,37.73286371140448],[-122.44104981422423,37.73424677678513],[-122.43990182876587,37.73424677678513],[-122.43990182876587,37.73286371140448],[-122.44104981422423,37.73286371140448]],[[-122.44109272956847,37.731870943026074],[-122.43976235389708,37.731870943026074],[-122.43976235389708,37.7326855231885],[-122.44109272956847,37.7326855231885],[-122.44109272956847,37.731870943026074]]]}`}, {"OK"},
		{"WITHIN", "key5", "IDS", "GET", "mykey", "multipoly5"}, {"[0 [poly13]]"},

		{"SET", "key6", "multipoly5", "OBJECT", `{"type":"MultiPolygon","coordinates":[[[[-122.4408378,37.7341129],[-122.4408378,37.733],[-122.44,37.733],[-122.44,37.7341129],[-122.4408378,37.7341129]]],[[[-122.44091033935547,37.731981251280985],[-122.43994474411011,37.731981251280985],[-122.43994474411011,37.73254976045042],[-122.44091033935547,37.73254976045042],[-122.44091033935547,37.731981251280985]]]]}`}, {"OK"},
		{"SET", "key6", "poly13", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44073867797852,37.733211601447465],[-122.44011640548705,37.733211601447465],[-122.44011640548705,37.7340516218859],[-122.44073867797852,37.7340516218859],[-122.44073867797852,37.733211601447465]],[[-122.44060993194579,37.73345766902749],[-122.44060993194579,37.73355524732416],[-122.44044363498686,37.73355524732416],[-122.44044363498686,37.73345766902749],[-122.44060993194579,37.73345766902749]],[[-122.44060724973677,37.7336888869566],[-122.44060724973677,37.7339752567853],[-122.4402102828026,37.7339752567853],[-122.4402102828026,37.7336888869566],[-122.44060724973677,37.7336888869566]]]}`}, {"OK"},
		{"WITHIN", "key6", "IDS", "GET", "key5", "poly14"}, {"[0 []]"},

		{"SET", "key7", "multipoly15", "OBJECT", `{"type":"MultiPolygon","coordinates":[[[[-122.44056701660155,37.7332964524295],[-122.44029879570007,37.7332964524295],[-122.44029879570007,37.73375464605226],[-122.44056701660155,37.73375464605226],[-122.44056701660155,37.7332964524295]]],[[[-122.44067430496216,37.73217641163713],[-122.44034171104431,37.73217641163713],[-122.44034171104431,37.732430967850384],[-122.44067430496216,37.732430967850384],[-122.44067430496216,37.73217641163713]]]]}`}, {"OK"},
		{"SET", "key7", "multipoly16", "OBJECT", `{"type":"MultiPolygon","coordinates":[[[[-122.44056701660155,37.7332964524295],[-122.44029879570007,37.7332964524295],[-122.44029879570007,37.73375464605226],[-122.44056701660155,37.73375464605226],[-122.44056701660155,37.7332964524295]]],[[[-122.4402666091919,37.733109780140644],[-122.4401271343231,37.733109780140644],[-122.4401271343231,37.73323705675229],[-122.4402666091919,37.73323705675229],[-122.4402666091919,37.733109780140644]]]]}`}, {"OK"},
		{"SET", "key7", "multipoly17", "OBJECT", `{"type":"MultiPolygon","coordinates":[[[[-122.44056701660155,37.7332964524295],[-122.44029879570007,37.7332964524295],[-122.44029879570007,37.73375464605226],[-122.44056701660155,37.73375464605226],[-122.44056701660155,37.7332964524295]]],[[[-122.44032025337218,37.73267703802467],[-122.44013786315918,37.73267703802467],[-122.44013786315918,37.732838255971316],[-122.44032025337218,37.732838255971316],[-122.44032025337218,37.73267703802467]]]]}`}, {"OK"},
		{"WITHIN", "key7", "IDS", "GET", "mykey", "multipoly5"}, {"[0 [multipoly15 multipoly16]]"},
	})
}

func keys_INTERSECTS_test(mc *mockServer) error {
	return mc.DoBatch([][]interface{}{
		{"SET", "mykey", "point1", "POINT", 37.7335, -122.4412}, {"OK"},
		{"SET", "mykey", "point2", "POINT", 37.7335, -122.44121}, {"OK"},
		{"SET", "mykey", "line3", "OBJECT", `{"type":"LineString","coordinates":[[-122.4408378,37.7341129],[-122.4408378,37.733]]}`}, {"OK"},
		{"SET", "mykey", "poly4", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.4408378,37.7341129],[-122.4408378,37.733],[-122.44,37.733],[-122.44,37.7341129],[-122.4408378,37.7341129]]]}`}, {"OK"},
		{"SET", "mykey", "multipoly5", "OBJECT", `{"type":"MultiPolygon","coordinates":[[[[-122.4408378,37.7341129],[-122.4408378,37.733],[-122.44,37.733],[-122.44,37.7341129],[-122.4408378,37.7341129]]],[[[-122.44091033935547,37.731981251280985],[-122.43994474411011,37.731981251280985],[-122.43994474411011,37.73254976045042],[-122.44091033935547,37.73254976045042],[-122.44091033935547,37.731981251280985]]]]}`}, {"OK"},
		{"SET", "mykey", "point6", "POINT", -5, 5}, {"OK"},
		{"SET", "mykey", "point7", "POINT", 33, 21}, {"OK"},
		{"SET", "mykey", "poly8", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.4408378,37.7341129],[-122.4408378,37.733],[-122.44,37.733],[-122.44,37.7341129],[-122.4408378,37.7341129]],[[-122.44060993194579,37.73345766902749],[-122.44044363498686,37.73345766902749],[-122.44044363498686,37.73355524732416],[-122.44060993194579,37.73355524732416],[-122.44060993194579,37.73345766902749]],[[-122.44060724973677,37.7336888869566],[-122.4402102828026,37.7336888869566],[-122.4402102828026,37.7339752567853],[-122.44060724973677,37.7339752567853],[-122.44060724973677,37.7336888869566]]]}`}, {"OK"},
		{"INTERSECTS", "mykey", "IDS", "OBJECT",
			`{
				"type": "Polygon",
				"coordinates": [
					[
						[-122.44126439094543,37.732906137107],
						[-122.43980526924135,37.732906137107],
						[-122.43980526924135,37.73421283683962],
						[-122.44126439094543,37.73421283683962],
						[-122.44126439094543,37.732906137107]
					]
				]
			}`}, {"[0 [point1 point2 line3 poly4 multipoly5 poly8]]"},

		{"SET", "key2", "poly9", "OBJECT", `{"type": "Polygon","coordinates": [[[-122.44037926197052,37.73313523548048],[-122.44017541408539,37.73313523548048],[-122.44017541408539,37.73336857568778],[-122.44037926197052,37.73336857568778],[-122.44037926197052,37.73313523548048]]]}`}, {"OK"},
		{"SET", "key2", "poly10", "OBJECT", `{"type": "Polygon","coordinates": [[[-122.44040071964262,37.73359343010089],[-122.4402666091919,37.73359343010089],[-122.4402666091919,37.73373767596864],[-122.44040071964262,37.73373767596864],[-122.44040071964262,37.73359343010089]]]}`}, {"OK"},
		{"SET", "key2", "poly10.1", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44051605463028,37.73375464605226],[-122.44028002023695,37.73375464605226],[-122.44028002023695,37.733903134117966],[-122.44051605463028,37.733903134117966],[-122.44051605463028,37.73375464605226]]]}`}, {"OK"},
		{"INTERSECTS", "key2", "IDS", "GET", "mykey", "poly8"}, {"[0 [poly9 poly10]]"},

		{"SET", "key3", "poly11", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44059920310973,37.733279482240874],[-122.4402344226837,37.733279482240874],[-122.4402344226837,37.73375464605226],[-122.44059920310973,37.73375464605226],[-122.44059920310973,37.733279482240874]]]}`}, {"OK"},
		{"SET", "key3", "poly12", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44060993194579,37.73238005667773],[-122.44013786315918,37.73238005667773],[-122.44013786315918,37.73316917591997],[-122.44060993194579,37.73316917591997],[-122.44060993194579,37.73238005667773]]]}`}, {"OK"},
		{"SET", "key3", "poly12.1", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44074940681458,37.73270249351328],[-122.44008421897887,37.73270249351328],[-122.44008421897887,37.732872196546936],[-122.44074940681458,37.732872196546936],[-122.44074940681458,37.73270249351328]]]}`}, {"OK"},
		{"INTERSECTS", "key3", "IDS", "GET", "mykey", "multipoly5"}, {"[0 [poly11 poly12]]"},

		{"SET", "key5", "poly13", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44073867797852,37.733211601447465],[-122.44011640548705,37.733211601447465],[-122.44011640548705,37.7340516218859],[-122.44073867797852,37.7340516218859],[-122.44073867797852,37.733211601447465]],[[-122.44060993194579,37.73345766902749],[-122.44060993194579,37.73355524732416],[-122.44044363498686,37.73355524732416],[-122.44044363498686,37.73345766902749],[-122.44060993194579,37.73345766902749]],[[-122.44060724973677,37.7336888869566],[-122.44060724973677,37.7339752567853],[-122.4402102828026,37.7339752567853],[-122.4402102828026,37.7336888869566],[-122.44060724973677,37.7336888869566]]]}`}, {"OK"},
		{"SET", "key5", "poly14", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44154334068298,37.73179457567642],[-122.43935465812682,37.73179457567642],[-122.43935465812682,37.7343740514423],[-122.44154334068298,37.7343740514423],[-122.44154334068298,37.73179457567642]],[[-122.44104981422423,37.73286371140448],[-122.44104981422423,37.73424677678513],[-122.43990182876587,37.73424677678513],[-122.43990182876587,37.73286371140448],[-122.44104981422423,37.73286371140448]],[[-122.44109272956847,37.731870943026074],[-122.43976235389708,37.731870943026074],[-122.43976235389708,37.7326855231885],[-122.44109272956847,37.7326855231885],[-122.44109272956847,37.731870943026074]]]}`}, {"OK"},
		{"INTERSECTS", "key5", "IDS", "GET", "mykey", "multipoly5"}, {"[0 [poly13]]"},

		{"SET", "key6", "multipoly5", "OBJECT", `{"type":"MultiPolygon","coordinates":[[[[-122.4408378,37.7341129],[-122.4408378,37.733],[-122.44,37.733],[-122.44,37.7341129],[-122.4408378,37.7341129]]],[[[-122.44091033935547,37.731981251280985],[-122.43994474411011,37.731981251280985],[-122.43994474411011,37.73254976045042],[-122.44091033935547,37.73254976045042],[-122.44091033935547,37.731981251280985]]]]}`}, {"OK"},
		{"SET", "key6", "poly13", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.44073867797852,37.733211601447465],[-122.44011640548705,37.733211601447465],[-122.44011640548705,37.7340516218859],[-122.44073867797852,37.7340516218859],[-122.44073867797852,37.733211601447465]],[[-122.44060993194579,37.73345766902749],[-122.44060993194579,37.73355524732416],[-122.44044363498686,37.73355524732416],[-122.44044363498686,37.73345766902749],[-122.44060993194579,37.73345766902749]],[[-122.44060724973677,37.7336888869566],[-122.44060724973677,37.7339752567853],[-122.4402102828026,37.7339752567853],[-122.4402102828026,37.7336888869566],[-122.44060724973677,37.7336888869566]]]}`}, {"OK"},
		{"INTERSECTS", "key6", "IDS", "GET", "key5", "poly14"}, {"[0 []]"},

		{"SET", "key7", "multipoly15", "OBJECT", `{"type":"MultiPolygon","coordinates":[[[[-122.44056701660155,37.7332964524295],[-122.44029879570007,37.7332964524295],[-122.44029879570007,37.73375464605226],[-122.44056701660155,37.73375464605226],[-122.44056701660155,37.7332964524295]]],[[[-122.44067430496216,37.73217641163713],[-122.44034171104431,37.73217641163713],[-122.44034171104431,37.732430967850384],[-122.44067430496216,37.732430967850384],[-122.44067430496216,37.73217641163713]]]]}`}, {"OK"},
		{"SET", "key7", "multipoly16", "OBJECT", `{"type":"MultiPolygon","coordinates":[[[[-122.44056701660155,37.7332964524295],[-122.44029879570007,37.7332964524295],[-122.44029879570007,37.73375464605226],[-122.44056701660155,37.73375464605226],[-122.44056701660155,37.7332964524295]]],[[[-122.4402666091919,37.733109780140644],[-122.4401271343231,37.733109780140644],[-122.4401271343231,37.73323705675229],[-122.4402666091919,37.73323705675229],[-122.4402666091919,37.733109780140644]]]]}`}, {"OK"},
		{"SET", "key7", "multipoly17", "OBJECT", `{"type":"MultiPolygon","coordinates":[[[[-122.44056701660155,37.7332964524295],[-122.44029879570007,37.7332964524295],[-122.44029879570007,37.73375464605226],[-122.44056701660155,37.73375464605226],[-122.44056701660155,37.7332964524295]]],[[[-122.44032025337218,37.73267703802467],[-122.44013786315918,37.73267703802467],[-122.44013786315918,37.732838255971316],[-122.44032025337218,37.732838255971316],[-122.44032025337218,37.73267703802467]]]]}`}, {"OK"},
		{"SET", "key7", "multipoly17.1", "OBJECT", `{"type":"MultiPolygon","coordinates":[[[[-122.4407172203064,37.73270249351328],[-122.44049191474916,37.73270249351328],[-122.44049191474916,37.73286371140448],[-122.4407172203064,37.73286371140448],[-122.4407172203064,37.73270249351328]]],[[[-122.44032025337218,37.73267703802467],[-122.44013786315918,37.73267703802467],[-122.44013786315918,37.732838255971316],[-122.44032025337218,37.732838255971316],[-122.44032025337218,37.73267703802467]]]]}`}, {"OK"},
		{"INTERSECTS", "key7", "IDS", "GET", "mykey", "multipoly5"}, {"[0 [multipoly15 multipoly16 multipoly17]]"},
	})
}

func keys_WITHIN_CIRCLE_test(mc *mockServer) error {
	return mc.DoBatch([][]interface{}{
		{"SET", "mykey", "1", "POINT", 37.7335, -122.4412}, {"OK"},
		{"SET", "mykey", "2", "POINT", 37.7335, -122.44121}, {"OK"},
		{"SET", "mykey", "3", "OBJECT", `{"type":"LineString","coordinates":[[-122.4408378,37.7341129],[-122.4408378,37.733]]}`}, {"OK"},
		{"SET", "mykey", "4", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.4408378,37.7341129],[-122.4408378,37.733],[-122.44,37.733],[-122.44,37.7341129],[-122.4408378,37.7341129]]]}`}, {"OK"},
		{"SET", "mykey", "5", "OBJECT", `{"type":"MultiPolygon","coordinates":[[[[-122.4408378,37.7341129],[-122.4408378,37.733],[-122.44,37.733],[-122.44,37.7341129],[-122.4408378,37.7341129]]]]}`}, {"OK"},
		{"SET", "mykey", "6", "POINT", -5, 5}, {"OK"},
		{"SET", "mykey", "7", "POINT", 33, 21}, {"OK"},
		{"WITHIN", "mykey", "IDS", "CIRCLE", 37.7335, -122.4412, 1000}, {
			"[0 [1 2 3 4 5]]"},
		{"WITHIN", "mykey", "IDS", "CIRCLE", 37.7335, -122.4412, 10}, {
			"[0 [1 2]]"},
	})
}

func keys_INTERSECTS_CIRCLE_test(mc *mockServer) error {
	return mc.DoBatch([][]interface{}{
		{"SET", "mykey", "1", "POINT", 37.7335, -122.4412}, {"OK"},
		{"SET", "mykey", "2", "POINT", 37.7335, -122.44121}, {"OK"},
		{"SET", "mykey", "3", "OBJECT", `{"type":"LineString","coordinates":[[-122.4408378,37.7341129],[-122.4408378,37.733]]}`}, {"OK"},
		{"SET", "mykey", "4", "OBJECT", `{"type":"Polygon","coordinates":[[[-122.4408378,37.7341129],[-122.4408378,37.733],[-122.44,37.733],[-122.44,37.7341129],[-122.4408378,37.7341129]]]}`}, {"OK"},
		{"SET", "mykey", "5", "OBJECT", `{"type":"MultiPolygon","coordinates":[[[[-122.4408378,37.7341129],[-122.4408378,37.733],[-122.44,37.733],[-122.44,37.7341129],[-122.4408378,37.7341129]]]]}`}, {"OK"},
		{"SET", "mykey", "6", "POINT", -5, 5}, {"OK"},
		{"SET", "mykey", "7", "POINT", 33, 21}, {"OK"},
		{"INTERSECTS", "mykey", "IDS", "CIRCLE", 37.7335, -122.4412, 70}, {
			"[0 [1 2 3 4 5]]"},
		{"INTERSECTS", "mykey", "IDS", "CIRCLE", 37.7335, -122.4412, 10}, {
			"[0 [1 2]]"},
	})
}

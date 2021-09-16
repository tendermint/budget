package simulation_test

// TODO: write testcode after Randomize

//func TestParamChanges(t *testing.T) {
//	s := rand.NewSource(1)
//	r := rand.New(s)
//
//	expected := []struct {
//		composedKey string
//		key         string
//		simValue    string
//		subspace    string
//	}{
//		{"budget/", "", "\"\"", "budget"},
//	}
//
//	paramChanges := simulation.ParamChanges(r)
//
//	require.Len(t, paramChanges, 1)
//
//	for i, p := range paramChanges {
//		require.Equal(t, expected[i].composedKey, p.ComposedKey())
//		require.Equal(t, expected[i].key, p.Key())
//		require.Equal(t, expected[i].simValue, p.SimValue()(r))
//		require.Equal(t, expected[i].subspace, p.Subspace())
//	}
//}

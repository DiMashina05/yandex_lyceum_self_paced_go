package main

import(
"testing"
)



func TestMultiply(t *testing.T){
    
	test := []struct{
		input []byte
		output int
		error_wait error
	}{
		{
		input: []byte{'a', 'e', 'c'},
		output: 3,
		error_wait: nil,
	},
	{
		input: []byte{'a', 'e', 'c', '2', '1'},
		output: 5,
		error_wait: nil,
	},
	{
		input: []byte{'a'},
		output: 1,
		error_wait: nil,
	},
}
for _, test_for := range test{
	count, err := GetUTFLength(test_for.input)
	if count != test_for.output || err != test_for.error_wait{
		t.Fatal("WA")
	}
}
	t.Run("error", func(t *testing.T){
		count, err := GetUTFLength([]byte{0xff, 0xfe, 0xfd})
		if count != 0 || err.Error() != "invalid utf8"{
			t.Fatal("Не уловил ошибку")
		}
	})
}



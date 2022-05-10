package main

func TestSoma(t *testing.T) {

	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Soma(15, 15) deveria ser 30, mas foi %d", total)
	}
}
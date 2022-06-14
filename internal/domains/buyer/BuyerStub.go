// Gere um Store Stub cuja função “Read” retorne dois produtos com as especificações
// desejadas. Verifique se GetAll() retorna as informações exatamente como o esperado. Para
// isto:
// 1. Dentro da pasta /internal/products, crie um arquivo repository_test.go com o teste
// projetado.

package buyer

type StubSearchBuyer struct {
	searchTestAll bool
}

func (c StubSearchBuyer) SearchTestAll() []Buyer {

	//Esse método FOI EXECUTADO
	//Se tiver validação de método é MOCK
	c.searchTestAll = true

	//Se NAO TIVER validação de método é STUB
	//array de buyers
	var buyers []Buyer = []Buyer{}

	buyers = append(buyers, Buyer{Id: 1, CardNumber: 123456, FirstName: "Jessica", LasttName: "Cruz"})
	buyers = append(buyers, Buyer{Id: 2, CardNumber: 654321, FirstName: "Reimer", LasttName: "Birro"})
	buyers = append(buyers, Buyer{Id: 3, CardNumber: 456321, FirstName: "Manu", LasttName: "Moreira"})

	return buyers
}

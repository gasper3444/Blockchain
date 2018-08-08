package blockchain

type Blockchain struct {
	Chain []Block			//an array of blocks
}

func (chain *Blockchain) Add(blk Block) {
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	chainDuplicate := make([]Block, len(chain.Chain)+1)				//an array of length + 1
	for i, v:= range chain.Chain {
		 chainDuplicate[i]= v
	}
	chainDuplicate[len(chain.Chain)] = blk    //duplicate chain ready
	chain.Chain = chainDuplicate
}

func compareMySlices(arr1 []byte, arr2 []byte) bool {
	if len(arr1) != len(arr2) {
		return false
	} else {
		for i, v := range arr1 {
			if v == arr2[i] {
				continue
			} else {
				return false
			}
		}
		return true
	}
}

func (chain Blockchain) IsValid() bool {
	nullArr := make([]byte, 32)
	var cond1, cond2, cond3, cond4, cond5, cond6 bool
	cond1 = compareMySlices(chain.Chain[0].PrevHash,nullArr) && (chain.Chain[0].Generation == 0)
	cond2= true
	diffLevel := chain.Chain[0].Difficulty
	for _, v := range chain.Chain {
		if v.Difficulty == diffLevel {
			continue
		} else {
			cond2 = false
			break
		}
	}

	cond3= true
	for i:=0 ; i<(len(chain.Chain)-1) ; i++ {
		if chain.Chain[i].Generation == (chain.Chain[i+1].Generation - 1) {
			continue
		} else {
			cond3 = false
			break
		}
	}

	cond4 = true
	for i:=0 ; i<len(chain.Chain)-1 ; i++ {
		if compareMySlices(chain.Chain[i].Hash, chain.Chain[i+1].PrevHash) {
			continue
		} else {
			cond4 = false
			break
		}
	}

	cond5= true
	for i:=0 ; i<len((chain.Chain)) ; i++ {
		if compareMySlices(chain.Chain[i].Hash, chain.Chain[i].CalcHash()) {
			continue
		} else {
			cond5 = false
			break
		}
	}

	cond6= true
	for i:=0 ; i<(len(chain.Chain)) ; i++ {
		if chain.Chain[i].ValidHash() == true {
			continue
		} else {
			cond6 = false
			break
		}
	}
	if cond1 && cond2 && cond3 && cond4 && cond5 && cond6 {
		return true
	} else {
		return false
	}

}

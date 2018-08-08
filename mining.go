package blockchain

import (
	"work_queue"
	"math"
)

//queue.Enqueue(miningWorker) : must be called.
type miningWorker struct {
	// TODO. Should implement work_queue.Worker
	//miningWorker implements Worker inteface (i.e contains Run() method)
	start uint64		//range values
	end uint64
	blk Block
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

func (mw miningWorker) Run() interface{} {

	 var m MiningResult
	 i := mw.start
	 for i <= mw.end {
		 mw.blk.Proof= i
		 mw.blk.Hash= mw.blk.CalcHash()
		 if mw.blk.ValidHash() == true {
 			m.Proof= i
 			m.Found= true
			return m
	 } else {
		 	i+=1
	 	}
 	}
 	m.Found= false
 	return m
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {

	myQueue := work_queue.Create(uint(workers),uint(chunks))
	chunkSize := float64((end - start) / chunks)
	chunkSizeNew := uint64(math.Ceil(chunkSize))

	var res MiningResult

	totChunk := 0

	//1st chunk
	mineWorker := miningWorker{start, start+chunkSizeNew, blk}
	totChunk +=1
	myQueue.Enqueue(mineWorker)
	j := start + chunkSizeNew

	//middle chunks
	ch := uint64(0)
	for ch < chunks - 2 {
		mineWorker2 := miningWorker{j+1, j + chunkSizeNew, blk}
		myQueue.Enqueue(mineWorker2)
		j += chunkSizeNew
		ch +=1
		totChunk +=1
	}

	//remaining chunks
	mineWorker3 := miningWorker{j+1, end, blk}
	myQueue.Enqueue(mineWorker3)
	totChunk +=1
	for v:= range myQueue.Results {
		u := v.(MiningResult)
		if u.Found == true {
			myQueue.Shutdown()
			res= u
			return res
		} else {
			continue
		}
	}
	return res
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << (8 * blk.Difficulty)) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}

package aoc2020

type RfidContextInterface interface {
	CalculateEncryptionKey() uint32
}

type RfidContext struct {
	cardPublicKey uint32
	doorPublicKey uint32
}

const _subjectNumber = 7
const _divisor = 20201227

func _calculateLoopSize(publicKey uint32) uint32 {
	var tmp uint32 = 1
	var loopSize uint32 = 0

	for tmp != publicKey {
		tmp = tmp * _subjectNumber
		tmp = tmp % _divisor
		loopSize = loopSize + 1
	}

	return loopSize
}

func _calculateEncryptionKey(loopSize uint32, subjectNumber uint32) uint32 {
	var value uint64 = 1

	for i := 0; i < int(loopSize); i++ {
		value = value * uint64(subjectNumber)
		value = value % uint64(_divisor)
	}

	return uint32(value)
}

func (ctx *RfidContext) CalculateEncryptionKey() uint32 {
	cardLoopSize := _calculateLoopSize(ctx.cardPublicKey)
	doorLoopSize := _calculateLoopSize(ctx.doorPublicKey)
	cardEncryptionKey := _calculateEncryptionKey(doorLoopSize, ctx.cardPublicKey)
	doorEncryptionKey := _calculateEncryptionKey(cardLoopSize, ctx.doorPublicKey)

	if cardEncryptionKey == doorEncryptionKey {
		return cardEncryptionKey
	}

	return 0
}

func NewRfidContext(cardPublicKey uint32, doorPublicKey uint32) RfidContextInterface {
	return &RfidContext{cardPublicKey: cardPublicKey, doorPublicKey: doorPublicKey}
}

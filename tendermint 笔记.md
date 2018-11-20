## Tendermint笔记
[TOC]

### 概览

- Byzantine Fault Tolerant State Machine Replication 拜占庭容错状态机拷贝

- hash-linked batches of transactions 一批事务组成一个链表，这个链表的每一个元素也就是每一个事务中有一个字段用来存放前一个元素的hash值，这样这一批事务就由这样一个链表表示

- blocks 块 上条所说的事务批也就被称为块

- blockchain 区块链就由这些块组成

- Height 块的唯一索引，这个值是唯一并且严格单调的

- 一些被赋予权重的validator或者说验证者成为一个集合，块就由这个集合里的成员提交

- 这个验证者集合的成员和权重是随时间而变的

- 只要这些验证者中不超过1/3的成员是恶意的或者有缺陷的，这个区块链就是安全并且活跃的。

- 一个commit指一个带签名的消息的集合，这些消息来自当前验证者集的成员，这些成员的权重之和超过总权重的2/3

- 验证者各自对块进行提议和表决，这里的表决，就是投赞成票，收到足够的赞成票或者说表决数之后，这一块就被认为提交了

- 这些表决信息会包含在下一块中，毕竟当前正在处理的块已经创建，而区块链当中块已经创建便不能更改，没有包含在表决信息当中的验证者就被忽略掉了，不论是没有投票赞成还是没有参加表决

- 块一旦提交，就可以由一个应用来进行一些处理，比方说返回块中事务的一些结果

- 这些应用也可以返回一些块之外的信息，比如验证者集合的变化，以及最近状态的加密摘要

- Tendermint 用来对区块链的最近状态进行验证和认证

- 因此，在块的header部分包含了一些加密的信息作为承诺

- 这些信息包括块的目录(所包含的事务)，提交这一块的验证者，以及由应用返回的其他一些结果。需要注意的是应用返回的结果只能包含在下一个块中，因为应用只有在块提交之后才会对块进行处理，而块一经创建便不可更改

- 而事务结果和验证者集合并不直接包含在块中，只是一个加密的摘要(merkel树的根)

- 因此，验证一个块需要一个单独的数据结构来存储这些信息，这些信息称为state

- 块验证需要访问前一个块

### Blockchain--数据结构

#### Block

```go
type Block struct {
    Header      Header
    Txs         Data
    Evidence    EvidenceData
    LastCommit  Commit
}
```

- Header 就是Header
- Txs 代表事务，是个Data类型
- Evidence 是一个list，指的是一些不合法的行为，比方说签名不符合的表决
- LastCommit 顾名思义，上一块的Commit，也就是上一块的表决数据，如前所述

##### Header

块头部是一些元数据，关于块本身、共识、承诺，还有应用返回的结果

```go
type Header struct {
	// basic block info
	Version  Version
	ChainID  string
	Height   int64
	Time     Time
	NumTxs   int64
	TotalTxs int64

	// prev block info
	LastBlockID BlockID

	// hashes of block data
	LastCommitHash []byte // commit from validators from the last block
	DataHash       []byte // Merkle root of transactions

	// hashes from the app output from the prev block
	ValidatorsHash     []byte // validators for the current block
	NextValidatorsHash []byte // validators for the next block
	ConsensusHash      []byte // consensus params for current block
	AppHash            []byte // state after txs from the previous block
	LastResultsHash    []byte // root hash of all results from the txs from the previous block

	// consensus info
	EvidenceHash    []byte // evidence included in the block
	ProposerAddress []byte // original proposer of the block

```

###### Version

包括区块链和应用的协议版本

```go
type Version struct {
    Block   uint64
    App     uint64
}
```

###### BlockID

 `BlockID` 包含块的两个不同的Merkle树根。第一个，作为块的主hash，是头部所有域的Merkle树根。第二个，在共识期间用来保护块的安全传播，是完全序列化的块切分之后的一个Merkle树根。

```go
type BlockID struct {
    Hash []byte
    Parts PartsHeader
}

type PartsHeader struct {
    Hash []byte
    Total int32   //序列化之后的block切分成了Total这么多part
}
```

###### Time

Time是一个Google.Protobuf.WellKnownTypes.Timestamp的类型，其中包括两个整型数，一个用来表示秒，一个用来表示纳秒，当然都是从Epoch开始

##### Data

是一个事务list的封装

```go
type Data struct {
    Txs [][]byte
}
```

##### Commit

```go
type Commit struct {
    BlockID     BlockID
    Precommits  []Vote
}
```

如前所述，这是前一块的commit信息

###### Vote

一个表决是一个验证者的签名信息，验证者都是针对某一特定块而言，也就是说不同的块参与表决的验证者是不必相同的。

```go
type Vote struct {
	Type             SignedMsgType  // byte
	Height           int64
	Round            int
	Timestamp        time.Time
	BlockID          BlockID
	ValidatorAddress Address
	ValidatorIndex   int
	Signature        []byte
}
```

如前所述，一个Commit中包含了块的id，以及对这个块进行了表决的验证者列表，这里说的列表就是通常理解的列表，用slice而不是list来保存。

tendermint的投票包括两轮，所以`vote.Type == 1`就表示是*prevote*轮，而`vote.Type == 2`则表示*precommit*轮

###### Signature

ED25519签名，64个字节的裸数据

##### EvidenceData

```go
type EvidenceData struct {
    Evidence []Evidence
}
```

###### Evidence

注意，这个`Evidence`是`[]`后面那个`Evidence`

`Evidence`是个接口，所以：

```go
// amino name: "tendermint/DuplicateVoteEvidence"
type DuplicateVoteEvidence struct {
	PubKey PubKey
	VoteA  Vote
	VoteB  Vote
}
```

Evidence保存的是有冲突的表决，所谓冲突并不是赞成和否决的冲突，因为在Vote列表中并不存在否决提议的验证者，而是Vote中的一些数据的不一致，比方说签名，要记得一个公钥实际上唯一标识了一个验证者。

### Blockchain -- Validation

#### Header

##### Version

```go
block.Version.Block == state.Version.Block
block.Version.App == state.Version.App
```

##### ChainID

```go
len(block.ChainID) < 50
```

##### Height

```go
block.Header.Height > 0
block.Header.Height == prevBlock.Header.Height + 1
```

##### Time

```go
block.Header.Timestamp >= prevBlock.Header.Timestamp + 1 ms
block.Header.Timestamp == MedianTime(block.LastCommit, state.LastValidators)
```

时间戳必须是单调的，并且是加权的中值，注意不是平均值。另外，一个表决的时间戳必须至少比要表决的那一块的时间戳大一毫秒。

另外：

```go
if block.Header.Height == 1 {
    block.Header.Timestamp == genesisTime
}
```

事实上有这么个关系：

```go
if block.Height == 1 { //姑且先忽略这里的block跟上面的block的关系,只要注意genesisTime的来源就行  
        genesisTime := state.LastBlockTime
    	...
}
```

##### NumTxs

```go
block.Header.NumTxs == len(block.Txs.Txs)
```

回想一下之前的定义：

```go
type Block struct {
    Header      Header
    Txs         Data
    Evidence    EvidenceData
    LastCommit  Commit
}
type Data struct {
    Txs [][]byte
}
type Header struct {
	...	
	NumTxs   int64
	TotalTxs int64
    ...
}
```

##### TotalTxs

```go
block.Header.TotalTxs == prevBlock.Header.TotalTxs + block.Header.NumTxs
```

所以`TotalTxs`是个累加值，表示了区块链中所有的事务的数量。对于第一个块，显然有`block.Header.TotalTxs = block.Header.NumberTxs`

##### LastBlockID

```go
type Header struct {
	...
	// prev block info
	LastBlockID BlockID
    ...
}
type BlockID struct {
    Hash []byte
    Parts PartsHeader
}

type PartsHeader struct {
    Hash []byte
    Total int32
}


prevBlockParts := MakeParts(prevBlock, state.LastConsensusParams.BlockGossip.BlockPartSize)
block.Header.LastBlockID == BlockID {
    Hash: SimpleMerkleRoot(prevBlock.Header),
    PartsHeader{
        Hash: SimpleMerkleRoot(prevBlockParts),
        Total: len(prevBlockParts),
    },
}
```

自然，第一块的 `block.Header.LastBlockID == BlockID{}`.

另外，`state.LastConsensusParams`可能会被应用改变。

**还有别的内容，并不费解，所以直接看原文档即可**

### State

```go
type State struct {
    Version     Version
    LastResults []Result
    AppHash []byte

    LastValidators []Validator
    Validators []Validator
    NextValidators []Validator

    ConsensusParams ConsensusParams
}
```

#### Result

```go
type Result struct {
    Code uint32
    Data []byte
}
```

#### Validator

```go
type Validator struct {
    Address     []byte
    PubKey      PubKey
    VotingPower int64
}
```

#### ConsensusParams

```go
type ConsensusParams struct {
	BlockSize
	Evidence
	Validator
}

type BlockSize struct {
	MaxBytes        int64
	MaxGas          int64
}

type Evidence struct {
	MaxAge int64
}

type Validator struct {
	PubKeyTypes []string
}
```


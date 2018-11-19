## Tendermint笔记
### 概览

- Byzantine Fault Tolerant State Machine Replication 拜占庭容错状态机拷贝

- hash-linked batches of transactions 一批事务组成一个链表，这个链表的每一个元素也就是每一个事务中有一个字段用来存放前一个元素的hash值，这样这一批事务就由这样一个链表表示

- blocks 块 上条所说的事务批也就被称为块

- blockchain 区块链就由这些块组成

- Height 块的唯一索引，这个值是唯一并且严格单调的

- 一些被赋予权重的validator或者说验证者成为一个集合，块就由这个集合里的成员提交

- Membership and weighting within this validator set may change over time.

  这个验证者集合的成员和权重是随时间而变的

- 只要这些验证者中不超过1/3的成员是恶意的或者有缺陷的，这个区块链就是安全并且活跃的。

- 一次commit指一个带签名的消息的集合，这些消息来自当前验证者集的成员，这些成员的权重之和超过总权重的2/3

- 验证者各自对块进行提议和表决，收到足够的票数之后，这一块就被认为提交了

- 这些表决信息会包含在下一块中，毕竟当前正在处理的块已经创建，而区块链当中块已经创建便不能更改

- 块一旦提交，就可以由一个应用来进行一些处理，比方说返回块中事务的一些结果

- 这些应用也可以返回一些块之外的信息，比如验证者集合的变化，以及最近状态的加密摘要

- Tendermint 用来对区块链的最近状态进行验证和认证

- 因此，在块的header部分包含了一些加密的信息作为承诺

- 这些信息包括块的目录(所包含的事务)，提交这一块的验证者，以及由应用返回的其他一些结果。需要注意的是应用返回的结果只能包含在下一个块中，因为应用只有在块提交之后才会对块进行处理，而块一经创建便不可更改

- 而事务结果和验证者集合并不直接包含在块中，只是一个加密的摘要(merkel树的根)

- Hence, verification of a block requires a separate data structure to store this information因此，验证一个块需要一个单独的数据结构来存储这些信息，这些信息称为state

-  Block verification also requires access to the previous block.块验证需要访问前一个块

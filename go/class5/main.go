package class5

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BlockInfo 定义区块信息结构体
type BlockInfo struct {
	Number    *big.Int
	Hash      string
	Timestamp time.Time
	TxCount   int
}

type Config struct {
	PrivateKeyHex string `json:"privateKeyHex"`
}

func Main() {
	// Sepolia测试网络的RPC端点
	sepoliaRPC := "https://sepolia.infura.io/v3/8acc77573b834e378cb6bfd0da3b3ae1"

	// 打开JSON文件
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error Decode file:", err)
		return
	}

	// 连接到Sepolia网络
	client, err := ethclient.Dial(sepoliaRPC)
	if err != nil {
		log.Fatal("Failed to connect to Sepolia network:", err)
	}
	defer client.Close()

	// 检查连接状态
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 获取最新区块号
	lastBlockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatal("Failed to get latest block number:", err)
	}

	blockNumber := big.NewInt(int64(lastBlockNumber - 1))
	// 查询区块信息
	blockInfo, err := getBlockInfo(client, blockNumber)
	if err != nil {
		log.Fatal("获取区块信息失败:", err)
	}

	fmt.Printf("Connected to Sepolia Test Network\n")
	fmt.Printf("Latest Block Number: %d\n", blockNumber)

	fmt.Printf("区块 #%d 信息:\n", blockInfo.Number)
	fmt.Printf("  哈希: %s\n", blockInfo.Hash)
	fmt.Printf("  时间戳: %s\n", blockInfo.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Printf("  交易数量: %d\n", blockInfo.TxCount)

	// 钱包地址 (请替换为实际地址)
	walletAddress := common.HexToAddress("0xDE1123a5Fd6E34bCFD6F1C53677B03953760d53a")

	// 获取余额 (单位: wei)
	balance, err := client.BalanceAt(context.Background(), walletAddress, nil)
	if err != nil {
		log.Fatal("获取余额失败:", err)
	}
	etherBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	fmt.Printf("余额: %.6f ETH\n", etherBalance)

	class5_1(client, config.PrivateKeyHex)
}

// getBlockInfo 获取指定区块号的区块信息
func getBlockInfo(client *ethclient.Client, blockNumber *big.Int) (*BlockInfo, error) {
	// 创建上下文并设置超时
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 获取区块头信息
	block, err := client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("获取区块失败: %v", err)
	}

	// 提取区块信息
	blockInfo := &BlockInfo{
		Number:    block.Number(),
		Hash:      block.Hash().Hex(),
		Timestamp: time.Unix(int64(block.Time()), 0),
		TxCount:   len(block.Transactions()),
	}

	return blockInfo, nil
}

func class5_1(client *ethclient.Client, privateKeyHex string) {
	// 发送方私钥 (请替换为你的私钥)
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("私钥解析失败:", err)
	}

	// 获取发送方地址
	publicKey := privateKey.Public()
	// fmt.Printf("Public Key: %v\n", publicKey)
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法获取公钥")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	toAddress := common.HexToAddress("0x35351B44E100691566E9A9B22f10943B2c3288C7")
	// 获取发送方nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("获取nonce失败:", err)
	}

	// 设置转账金额 (0.001 ETH)
	value := big.NewInt(100000) // 0.001 ETH (单位: wei)

	// 设置Gas限制
	gasLimit := uint64(21000)

	// 获取Gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	fmt.Printf("Gas Price: %s\n", gasPrice.String())
	if err != nil {
		log.Fatal("获取Gas价格失败:", err)
	}

	// 构造交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// 获取链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("获取链ID失败:", err)
	}

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("交易签名失败:", err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("发送交易失败:", err)
	}

	fmt.Printf("交易已发送! 交易哈希: %s\n", signedTx.Hash().Hex())
	fmt.Printf("发送方: %s\n", fromAddress.Hex())
	fmt.Printf("接收方: %s\n", toAddress.Hex())
	fmt.Printf("转账金额: %s ETH\n", weiToEther(value))

	// class5_2
	// 获取发送方nonce_2
	nonce_2, err2 := client.PendingNonceAt(context.Background(), fromAddress)
	if err2 != nil {
		log.Fatal("获取nonce2失败:", err2)
	}

	chainID2, err2 := client.NetworkID(context.Background())
	if err2 != nil {
		log.Fatal("获取链ID失败:", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID2)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce_2))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// input := "1.0"
	address, tx, instance, err := DeployClass5(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	_ = instance
}

// weiToEther 将wei转换为ether
func weiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(100)
	f.SetInt(wei)
	ether := new(big.Float).Quo(f, big.NewFloat(1e18))
	return ether
}

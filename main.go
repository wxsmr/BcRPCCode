package main

import (
	"fmt"
	"encoding/json"
	"BcRPCCode/entity"
	"time"
	"net/http"
	"BcRPCCode/utils"
	"bytes"
	"io/ioutil"
)

//常量
const RPCURL = "http://127.0.0.1:8332"
const RPCUSER = "user"
const RPCPASSWORD = "pwd"

func main() {
	fmt.Print("Hello world")
	/**
	 * 1、准备要进行rpc通信时的json数据
	 * 2、使用http链接的post请求，发送json数据
	 * 3、接收http响应
	 * 4、根据响应的结果，进行判断和处理
	      状态码200正常
	      非200不正常
	 */

	//1、准备要发送的json数据
	/**
	 * {
	 *   "id":编号,
	 *   "method":"功能方法或者命令",
	 *   "jsonrpc":"rpc版本2.0",
	 *   "params":[携带的参数,数组形式]
	 * }
	 */
	//json操作：序列化、反序列化
	rpcReq := entity.RPCRequest{}
	rpcReq.Id = time.Now().Unix()
	rpcReq.Jsonrpc = "2.0"
	rpcReq.Method = "getblockcount" //获取当前节点的区块的数量
	//rpcReq.Method = "getbestblockhash" //获取当前节点的最新区块的hash值
	//对结构体类型进行序列化
	rpcBytes, err := json.Marshal(&rpcReq)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("准备好的json格式的数据：", string(rpcBytes))

	//2、发送一个post类型的请求
	//client：客户端，客户端用于发起请求
	client := http.Client{} //实例化一个客户端

	//实例化一个请求
	request, err := http.NewRequest("POST", RPCURL, bytes.NewBuffer(rpcBytes))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//给post请求添加请求头设置信息
	// key --> value
	request.Header.Add("Encoding", "UTF-8")
	request.Header.Add("Content-Type", "application/json")
	//权限认证设置
	request.Header.Add("Authorization", "Basic "+utils.Base64Str(RPCUSER+":"+RPCPASSWORD))

	//使用客户端发送请求
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//通过response，获取响应的数据
	code := response.StatusCode
	if code == 200 {
		fmt.Println("恭喜，请求成功")
		resultBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("rpc调用的响应结果:" + string(resultBytes))
		//json的反序列化
		var result entity.RPCResult
		err = json.Unmarshal(resultBytes, &result)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//反序列化正常，没有出现错误
		fmt.Println("功能调用结果：", result.Result)
	} else {
		fmt.Println("抱歉，请求失败")
	}
}

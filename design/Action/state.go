package Action

// 状态模式: 本质上就是为了避免写一大堆if-else语句, 状态需要定义一个interface, 然后每个子状态需要重写所有接口
// 不太符合开闭原则, 会导致大量的代码混乱

import (
	"fmt"
	"strconv"
)

type CandyMachinePattern struct {
	status IStatus
	count  int
	coin   int
}

func (c *CandyMachinePattern) CoinOperated() {
	c.status.CoinOperated()
}
func (c *CandyMachinePattern) CoinReturn() {
	c.status.CoinReturn()
}
func (c *CandyMachinePattern) SugarExtraction() {
	c.status.SugarExtraction()
}
func (c *CandyMachinePattern) PrintStatus(msg string) {
	fmt.Println(msg + " 当前有" + strconv.Itoa(c.coin) + "个币，有" + strconv.Itoa(c.count) + "个糖果")
}

// IStatus +-------------------------------------------------------
type IStatus interface {
	CoinOperated()
	CoinReturn()
	SugarExtraction()
}

// SoldOutStatus +-------------------------------------------------------
type SoldOutStatus struct {
	c *CandyMachinePattern
}

func (s *SoldOutStatus) CoinOperated() {
	s.c.PrintStatus("err:糖果已售罄，投币失败！")
}
func (s *SoldOutStatus) CoinReturn() {
	s.c.PrintStatus("err:无币，退币失败！")
}
func (s *SoldOutStatus) SugarExtraction() {
	s.c.PrintStatus("err:无币，退币失败！")
}

// NoCoinStatus +-------------------------------------------------------
type NoCoinStatus struct {
	c *CandyMachinePattern
}

func (s *NoCoinStatus) CoinOperated() {
	s.c.coin++
	s.c.status = &HasCoinStatus{
		c: s.c,
	}
	s.c.PrintStatus("success:投币成功！")
}
func (s *NoCoinStatus) CoinReturn() {
	s.c.PrintStatus("err:无币可退！")
}
func (s *NoCoinStatus) SugarExtraction() {
	s.c.PrintStatus("err:未投币，请先投币！")
}

// HasCoinStatus +-------------------------------------------------------
type HasCoinStatus struct {
	c *CandyMachinePattern
}

func (s *HasCoinStatus) CoinOperated() {
	s.c.PrintStatus("err:已投币，请先取糖果！")
}
func (s *HasCoinStatus) CoinReturn() {
	s.c.coin--
	s.c.status = &NoCoinStatus{
		c: s.c,
	}
	s.c.PrintStatus("success:退币成功！")
}
func (s *HasCoinStatus) SugarExtraction() {
	s.c.count--
	if s.c.count <= 0 {
		s.c.status = &SoldOutStatus{
			c: s.c,
		}
	} else {
		s.c.status = &NoCoinStatus{
			c: s.c,
		}
	}
	s.c.PrintStatus("success:取糖果成功！")
}

// StartPattern +-------------------------------------------------------
func RunState() {
	fmt.Println("state-pattern | CandyMachinePattern")
	c := CandyMachinePattern{
		count: 1,
		coin:  0,
	}
	c.status = &NoCoinStatus{
		c: &c,
	}
	c.PrintStatus("设备初始化成功！")
	c.CoinOperated()
	c.CoinReturn()
	c.CoinOperated()
	c.SugarExtraction()

	c.CoinOperated()
	c.SugarExtraction()
	c.CoinReturn()
	c.CoinReturn()
}

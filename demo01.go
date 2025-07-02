package main
import "fmt"

//APN结构体
type APN struct{
	Name  string
	DataLimit  int
	DataUsage  int
	ExpiryDate string
}

//SIM卡结构体
type SIM struct{
	ICCID  string
	IMSI  string
	MSISDN  string
	TrafficCap  int
	TrafficUsage  int
	ExpiryDate  string
	Status  string
	APNs  []APN
}

//创建一个SIM卡实体对象。
func NewSIM (iccid, imsi, msisdn string) * SIM{
	return &SIM{
		ICCID:  iccid,
		IMSI:  imsi,
		MSISDN:  msisdn,
		TrafficCap:  0,
        TrafficUsage: 0,
		ExpiryDate:  "2099-06-30",
		Status:  "未启用",
		APNs:  []APN{},
	}
}

//添加APN到SIM卡
func (s *SIM) AddAPN(name string, limit int) {
	s.APNs = append(s.APNs, APN{
		Name: name,
		DataLimit: limit,
	})
}

//修改sim卡状态。激活SIM卡
func (s *SIM) Activate(){
	if s.Status == "未启用"{
	s.Status = "已激活"	
	}
	fmt.Println("SIM卡状态:", s.Status)
}

//变更卡的流量上限。更新APN流量上限，同步SIM卡总上限
func (s *SIM) UpdateAPNLimits(LimitMap map [string]int){
	maxLimit:= 0
	for i, _:= range s.APNs{
		if newLimit, ok:= LimitMap [s.APNs[i].Name]; ok{
		s.APNs[i].DataLimit = newLimit

		if newLimit > maxLimit{
			maxLimit = newLimit
		}
		}
	}
	s.TrafficCap = maxLimit
	fmt.Println("SIM卡总流量上限:", s.TrafficCap, "KB")
	for _, apn:= range s.APNs{
		fmt.Printf("%s 流量上限: %dKB\n", apn.Name, apn.DataLimit)
	}
}

//停用SIM卡情况1
func (s *SIM) UpdateUsage(totalUsage int, usageMap map[string]int){
	s.TrafficUsage = totalUsage
	isOverLimit := false

	for i := range s.APNs {
		if used, ok := usageMap[s.APNs[i].Name]; ok {
			s.APNs[i].DataUsage = used
			if used >= s.APNs[i].DataLimit {
				 isOverLimit = true 
			}
		}
	}
	if isOverLimit {
        s.Status = "停用"
        fmt.Printf("当前流量检查后SIM卡状态: %s \n", s.Status)
    } else {
        fmt.Printf("当前流量检查后SIM卡状态: 正常 \n")
    }
}

//修改到期时间。APN设置新时间，SIM卡取最大值
func (s*SIM) UpdateExpiryDates(dateMap map[string]string){
	latestDate:="0000-00-00"
	for i := range s.APNs {
		if newDate, ok := dateMap[s.APNs[i].Name]; ok {
			s.APNs[i].ExpiryDate = newDate
			if newDate > latestDate {
				latestDate = newDate
			}
		}
	}
	s.ExpiryDate = latestDate
	fmt.Printf("SIM卡到期时间: %s\n", s.ExpiryDate)
	for _, apn := range s.APNs {
		fmt.Printf("APN %s 到期时间: %s\n", apn.Name, apn.ExpiryDate)
	}
}

//SIM卡停用情况2
func (s *SIM) CheckExpiry(currentDate string) {
    if currentDate >= s.ExpiryDate {
        s.Status = "停用"
        fmt.Println("当前日期检查后SIM卡状态: 停用")
    } else {
        fmt.Println("当前日期检查后SIM卡状态: 正常")
    }
}


func main(){
	sim := NewSIM("iccid111", "imsi222", "msisdn333")
	sim.AddAPN("apn1", 1000)
	sim.AddAPN("apn2", 2000)
	fmt.Println("\n--- 激活SIM卡 ---")
	sim.Activate()


	fmt.Println("\n--- 更新APN流量上限 ---")
	limits := map[string]int{
		"apn1": 1500,
		"apn2": 3000,
	}
	sim.UpdateAPNLimits(limits)


	fmt.Println("\n--- 更新流量使用量 ---")
	usage := map[string]int{
		"apn1": 1600,          // APN1超限
		"apn2": 2000,
	}
	sim.UpdateUsage(3600, usage)


	fmt.Println("\n--- 更新到期时间 ---")
	dates := map[string]string{
		"apn1": "2025-07-02",
		"apn2": "2026-09-01",
	}
	sim.UpdateExpiryDates(dates)


	fmt.Println("\n--- 检查过期 ---")
	sim.CheckExpiry("2025-07-02")
}
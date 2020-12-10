package subsystems

// 传递资源限制配置的结构体
type ResourceConfig struct {
	MemoryLimit string // 内存限制
	CpuShare    string // CPU时间片权重
	CpuSet      string // CPU 核心数
}

type Subsystem interface {
	Name() string                               // 返回subsystem的名字，比如 cup memory
	Set(path string, res *ResourceConfig) error // 设置某个cgroup在这个subsystem中的资源限制
	Apply(path string, pid int) error           // 将进程添加到某个cgroup中
	Remove(path string) error                   // 移除某个cgroup
}

// 通过不同的subsystem初始化实例创建资源限制处理链数组

var (
	SubsystemsIns = []Subsystem{
		&CpusetSubSystem{},
		&MemorySubSystem{},
		&CpuSubSystem{},
	}
)

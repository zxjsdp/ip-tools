package config

const (
	GUITitle   = "IP Tools"
	GUIVersion = "v0.0.1"
	Title  = GUITitle + " " + GUIVersion
	Help = `IP 转换工具。
1. 从任意文本中提取 IP。
2. 根据 IP 获取 IP 地址范围，可用于 gscan_quic 等工具进行扫描。

例如：
输入：
	"41.188.12.16 &*(41.188.12.29_#$ 58.28.63.15 &!^58.28.64.19)*#66.199.151.143"
输出：
	41.188.12.0-255
	58.28.63.0-255
	58.28.64.0-255
	66.199.151.0-255`
)

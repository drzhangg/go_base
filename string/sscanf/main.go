package main

import (
	"fmt"
	"strings"
)

func main() {
	output := `
Filesystem    1K-blocks    Used Available Use% Mounted on
/dev/root          4096    4096         0 100% /
`

	// 将输出按行分割
	lines := strings.Split(strings.TrimSpace(output), "\n")

	// 假设第二行是数据行
	if len(lines) > 1 {
		var filesystem, mountedOn string
		var total, used, available int
		var usePercent string

		// 解析第二行
		_, err := fmt.Sscanf(lines[1], "%s %d %d %d %s %s", &filesystem, &total, &used, &available, &usePercent, &mountedOn)
		if err != nil {
			fmt.Printf("解析错误: %v\n", err)
			return
		}

		fmt.Printf("文件系统: %s, 总计: %d, 已用: %d, 可用: %d, 使用率: %s, 挂载点: %s\n",
			filesystem, total, used, available, usePercent, mountedOn)
	}
}

.PHONY: watch-rs

# 定义变量，方便后续修改和引用
LEETCODE_RS_DIR := leetcode-rs
EXAMPLE_NAME := noah

# 目标：watch-rs
# 使用 cargo watch 在 leetcode-rs 目录下分别监控 src 和 examples 目录中的变化
watch:
	@cd $(LEETCODE_RS_DIR) && cargo watch -w src/ -w examples/ -x "run --example $(EXAMPLE_NAME)"

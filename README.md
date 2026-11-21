# CPA-OpenCode-ex

本库基于 [opencode-go-cliproxyapi](https://github.com/massiveits/opencode-go-cliproxyapi)。基础安装与配置见[上游使用说明（英文）](README-source.md)。

## 使用本库的改进

在兼容的 [CPA-Core-ex](https://github.com/williamxhero/CPA-Core-ex) 中加载 `cpa-credential-opencode`，将不同协议的模型配置到同一个 `opencode-go` Provider。独立构建时，插件与宿主须使用兼容 SDK；统一发布使用 [CPA-Manager-Plus-ex 的 OpenCode Go 插件目录](https://github.com/williamxhero/CPA-Manager-Plus-ex/tree/main/providers/opencode-go)。

## 相对上游的主要改进

- **单一 Provider 与多协议路由**：上游需要为 Chat Completions、Messages 和 Responses 分别维护配置，密钥轮换和冷却也无法共享。本 fork 以 opencode-go 统一模型命名空间和密钥池，在插件内完成协议翻译、路由和调度。
- **凭证管理与安全命名**：上游凭证缺少稳定的管理入口，展示名可能泄露密钥片段。本 fork 提供管理 API；有显式别名就显示别名，否则显示掩码，认证文件名不再携带密钥片段。
- **原生额度查询**：上游没有把额度查询绑定到当前选中的真实凭证。本 fork 注册原生 quota provider，支持滚动、周、月窗口，并将查询结果明确归属到选定凭证。
- **动态模型目录**：上游模型列表容易和 OpenCode Go 实际目录脱节，凭证模型接口为空时也缺少可诊断的恢复路径。本 fork 按成功刷新同步完整模型目录，覆盖 Space Bunny 等模型，吸收有证据的 fork 修复，并用自动回退部署保护凭证和现有排除规则。对应 [issue #2](https://github.com/williamxhero/CPA-OpenCode-ex/issues/2)。
- **宿主身份围栏**：上游回调范围过宽，宿主难以确认额度和刷新请求来自哪个身份。本 fork 保留选定凭证的回调范围和调用上下文，让宿主可以对身份、来源和能力做持久化门禁。

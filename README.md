<img align="right" src="./ex-badge.svg" alt="EX fork" width="76" />

# OpenCode Go Plugin-EX 主要能力

本页只列相对上游 massiveits/opencode-go-cliproxyapi 的主要功能。上游插件能完成基本转发，但多协议配置、凭证识别和宿主额度查询仍然割裂；本 fork 将它们收敛为一个可管理的 Provider。

- **单一 Provider 与多协议路由**：上游需要为 Chat Completions、Messages 和 Responses 分别维护配置，密钥轮换和冷却也无法共享。本 fork 以 opencode-go 统一模型命名空间和密钥池，在插件内完成协议翻译、路由和调度。
- **凭证管理与安全命名**：上游凭证缺少稳定的管理入口，展示名可能泄露密钥片段。本 fork 提供管理 API；有显式别名就显示别名，否则显示掩码，认证文件名不再携带密钥片段。
- **原生额度查询**：上游没有把额度查询绑定到当前选中的真实凭证。本 fork 注册原生 quota provider，支持滚动、周、月窗口，并将查询结果明确归属到选定凭证。
- **动态模型目录**：上游模型列表容易和 OpenCode Go 实际目录脱节。本 fork 按成功刷新同步模型目录，覆盖 Space Bunny 等模型，并在上游暂时不可用时按策略保留或拒绝旧目录。
- **宿主身份围栏**：上游回调范围过宽，宿主难以确认额度和刷新请求来自哪个身份。本 fork 保留选定凭证的回调范围和调用上下文，让宿主可以对身份、来源和能力做持久化门禁。

本页只说明 EX fork 的增量；插件构建、配置和协议细节仍以仓库正文为准。

<img align="right" src="./ex-badge.svg" alt="EX fork" width="76" />

# OpenCode Go Plugin-EX 主要能力

## 开始使用

OpenCode Go 凭证插件运行在兼容的 CPA Core 宿主内，统一处理协议路由、凭证和原生额度查询。运行标识为 `cpa-credential-opencode`，Provider 为 `opencode-go`。

- 上游配置及协议说明：[README-source.md](README-source.md)。
- CPA-Manager-Plus-ex 的 `providers/opencode-go` 纳入统一发布；本仓库保留独立插件源码。

需要 Go 1.26.7 或更高。构建共享库还需要启用 CGO，并提供对应平台的 C 编译器。

```powershell
go test ./...
$env:CGO_ENABLED = '1'
go build -buildmode=c-shared -o cpa-credential-opencode.dll .
```

插件与宿主必须使用兼容 SDK 构建。发布到本机 CPA 前，先完成独立加载验证，再通过带备份、健康检查和失败回退的部署脚本安装。

## Fork 功能记录

本页只列相对上游 massiveits/opencode-go-cliproxyapi 的主要功能。上游插件能完成基本转发，但多协议配置、凭证识别和宿主额度查询仍然割裂；本 fork 将它们收敛为一个可管理的 Provider。

- **单一 Provider 与多协议路由**：上游需要为 Chat Completions、Messages 和 Responses 分别维护配置，密钥轮换和冷却也无法共享。本 fork 以 opencode-go 统一模型命名空间和密钥池，在插件内完成协议翻译、路由和调度。
- **凭证管理与安全命名**：上游凭证缺少稳定的管理入口，展示名可能泄露密钥片段。本 fork 提供管理 API；有显式别名就显示别名，否则显示掩码，认证文件名不再携带密钥片段。
- **原生额度查询**：上游没有把额度查询绑定到当前选中的真实凭证。本 fork 注册原生 quota provider，支持滚动、周、月窗口，并将查询结果明确归属到选定凭证。
- **动态模型目录**：上游模型列表容易和 OpenCode Go 实际目录脱节，凭证模型接口为空时也缺少可诊断的恢复路径。本 fork 按成功刷新同步完整模型目录，覆盖 Space Bunny 等模型，吸收有证据的 fork 修复，并用自动回退部署保护凭证和现有排除规则。对应 [issue #2](https://github.com/williamxhero/CPA-OpenCode-ex/issues/2)。
- **宿主身份围栏**：上游回调范围过宽，宿主难以确认额度和刷新请求来自哪个身份。本 fork 保留选定凭证的回调范围和调用上下文，让宿主可以对身份、来源和能力做持久化门禁。

本页只说明 EX fork 的增量；插件构建、配置和协议细节仍以仓库正文为准。

> 完成 issue 归档更新时间：2026-11-09 08:34:40 SGT（UTC+08:00）。已列入上述条目的 issue 后续不再重复归档。

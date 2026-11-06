## 本 fork 的功能

相对上游 `massiveits/opencode-go-cliproxyapi`，这条主干增加：

- 插件标识 `cpa-credential-opencode`，提供商仍是 `opencode-go`。
- 凭证管理路由。有别名用别名，否则显示掩码；认证文件名不带密钥片段。
- 原生额度：滚动、周、月三个窗口，查询落在当前选中的凭证上。
- 模型目录路由，包括 Space Bunny 模型。
- 选定凭证的额度回调范围，供宿主做身份围栏。

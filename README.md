<p align="center">
  <img src="https://raw.githubusercontent.com/kwhitestone/prism-fusion/master/src/admin/public/favicon.svg" width="80" alt="Prism Example Site" />
</p>

<h1 align="center">Prism Example Site</h1>

<p align="center">
  <strong>基于 <a href="https://github.com/kwhitestone/prism-fusion">Prism Fusion</a> 框架的业务应用示例</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go" alt="Go" />
  <img src="https://img.shields.io/badge/Vue-3.5-4FC08D?style=flat-square&logo=vue.js" alt="Vue" />
  <img src="https://img.shields.io/badge/Vite-7-646CFF?style=flat-square&logo=vite" alt="Vite" />
  <img src="https://img.shields.io/badge/Auth-Builtin_JWT-blue?style=flat-square" alt="Auth" />
  <img src="https://img.shields.io/badge/RBAC-Builtin-orange?style=flat-square" alt="RBAC" />
</p>

---

演示如何基于 Prism Fusion 框架快速构建全栈业务系统 —— 通过 git submodule 引用框架、注册业务插件，即可获得完整的认证、权限、仪表盘等能力。

## 项目结构

```
prism-example-site/
├── prism-fusion/                  # 框架（git submodule，勿直接修改）
├── app/
│   ├── src/server/                # Go 后端
│   │   ├── addons/                # 业务插件（后端）
│   │   │   ├── dashboard/         #   数据总览
│   │   │   ├── messages/          #   消息记录
│   │   │   ├── site-info/         #   站点信息
│   │   │   └── example/           #   示例插件
│   │   ├── go.mod                 # 业务模块
│   │   ├── go.work                # workspace（引用框架）
│   │   └── main.go                # 入口（导入框架 + 业务插件）
│   ├── src/admin/                 # Vue 前端
│   │   ├── src/addons/            # 业务插件（前端）
│   │   │   ├── dashboard/         #   数据总览页面
│   │   │   ├── messages/          #   消息页面
│   │   │   ├── site-info/         #   站点信息页面
│   │   │   └── example/           #   示例插件页面
│   │   ├── pnpm-workspace.yaml    # workspace（引用框架前端）
│   │   ├── vite.config.ts         # Vite 配置
│   │   └── package.json
│   └── Dockerfile                 # 多阶段生产构建
├── docker-compose.yaml            # 应用服务编排
├── .env.example                   # 环境变量模板
└── README.md
```

## 框架集成方式

本项目通过以下机制引用 Prism Fusion 框架：

| 层 | 机制 | 配置文件 |
|----|------|---------|
| Go 后端 | `go.work` 多模块工作空间 | `app/src/server/go.work` |
| Vue 前端 | pnpm workspace 包引用 | `app/src/admin/pnpm-workspace.yaml` |
| Vite | `@` alias 指向框架 src | `app/src/admin/vite.config.ts` |
| 部署 | git submodule | `.gitmodules` |

业务代码中使用 `@biz/` 别名引用业务模块，`@/` 引用框架模块。

## 前置要求

| 工具 | 版本 |
|------|------|
| Go | >= 1.25 |
| Node.js | >= 22 |
| pnpm | >= 9 |
| Docker + Compose | 部署时需要 |

## 本地开发

### 1. 克隆（含子模块）

```bash
git clone --recurse-submodules https://github.com/kwhitestone/prism-example-site.git
cd prism-example-site
cp .env.example .env   # 按需修改
```

> 如果已克隆但未拉取子模块：`git submodule update --init --recursive`

### 2. 启动后端

```bash
cd app/src/server
go mod tidy
go run main.go
# 监听 :3380
```

- API 文档 (ReDoc)：http://localhost:3380/redoc
- API 文档 (Scalar)：http://localhost:3380/scalar
- OpenAPI JSON：http://localhost:3380/openapi.json

> 默认使用 SQLite（零配置），数据库文件 `prism_example_site.db` 自动创建。

### 3. 启动前端

```bash
cd app/src/admin
pnpm install
pnpm dev
# 访问 http://localhost:3388
```

> Vite 开发服务器已配置代理，`/api` 请求转发到后端 `:3380`。

## 业务插件

本项目使用框架内置的 `auth`（JWT 认证）和 `rbac`（角色权限）插件，并扩展了以下业务插件：

| 插件 | 说明 |
|------|------|
| `dashboard` | 数据总览仪表盘 |
| `messages` | 消息记录管理 |
| `site-info` | 站点信息管理 |
| `example` | 插件开发示例 |

**添加新插件**只需：

1. 后端：在 `app/src/server/addons/` 下创建插件目录，实现 `Plugin` 接口
2. 在 `app/src/server/addons/addons.go` 中添加 `import _ "whitestone.top/prism-example-site/addons/my-plugin"`
3. 前端：在 `app/src/admin/src/addons/` 下创建插件目录，导出 `PluginModule`（自动发现，无需配置）

## 认证与权限配置

本项目使用框架内置的 JWT 认证（`auth.provider: builtin`）和角色权限管理（`rbac.provider: builtin`）。

如需自定义 auth/rbac provider，需同时修改以下配置：

| 文件 | 配置项 | 说明 |
|------|--------|------|
| `app/src/server/config.yaml` | `auth.provider` / `rbac.provider` | 后端 provider 选择 |
| `app/src/admin/public/platform-config.json` | `AuthProvider` / `RBACProvider` | 前端 provider 选择（决定登录组件和路由来源） |
| `app/src/admin/src/main.ts` | 插件 import + `registerExternalPlugins` | 前端插件注册 |

> **重要**：`platform-config.json` 中的 `AuthProvider` 和 `RBACProvider` 必须与后端 `config.yaml` 保持一致，否则前端会加载错误的登录组件或路由逻辑。

### 路由机制：静态路由 vs 动态路由

登录后的页面路由由两部分组成：

#### 静态路由（前端定义）

由各业务插件在 `router/index.ts` 中硬编码，注册插件时自动加入路由表。例如 `dashboard` 插件定义了 `/dashboard` 路由。

**适用场景**：所有用户都能看到的固定页面，不需要按角色区分。

**定义方式**：在 `app/src/admin/src/addons/<plugin>/router/index.ts` 中导出路由数组，插件的 `index.ts` 通过 `routes` 字段注册。

#### 动态路由（后端下发）

框架在登录后会调用 `getAsyncRoutes()` 获取动态路由并注入路由表——这是**框架行为**。但框架默认返回空数组（不加载任何动态路由）。

实际的后端调用由 **rbac 插件** 提供：builtin rbac 插件在初始化时通过 `setAsyncRoutesProvider()` 将默认策略替换为调用 `/api/v1/addons/rbac/async-routes` 接口，后端根据当前用户角色返回该角色可访问的菜单列表。

**适用场景**：需要按角色控制可见页面（如管理员可见"用户管理"，普通用户不可见）。

**定义方式**：通过后端 RBAC 菜单表管理，管理员在后台配置菜单 → 分配给角色 → 前端按角色加载。

#### 当前状态

本项目使用 builtin RBAC，后端仅种子化了角色（管理员 / 普通用户），未预置菜单数据，因此**动态路由为空**，所有页面完全来自前端插件的静态路由。如需按角色控制页面可见性，可通过后端 API 向 `menus` 表插入菜单记录。

登录后默认跳转地址由 `dashboard` 插件的路由决定（`/` → redirect → `/dashboard`）。如需更改默认首页，修改 `app/src/admin/src/addons/dashboard/router/index.ts` 中的 `redirect` 和 `path`。

## 打包部署

### Docker 构建

```bash
# 从项目根目录执行
docker build -f app/Dockerfile -t prism-example-site .
```

### Docker Compose 部署

```bash
cp .env.example .env
# 编辑 .env，填写数据库配置
docker compose up -d
```

Compose 包含服务：
- **prism-example-site** — 主应用（端口 `3380`）

## 环境变量

| 变量 | 默认值 | 说明 |
|------|--------|------|
| `GATEWAY_PORT` | `3380` | 应用监听端口 |
| `DB_HOST` | （空，使用 SQLite） | MySQL 主机地址 |
| `DB_PORT` | `3306` | MySQL 端口 |
| `DB_NAME` | `prism_example_site` | 数据库名 |
| `DB_USER` | `root` | 数据库用户 |
| `DB_PASSWORD` | — | 数据库密码 |

完整变量见 [.env.example](.env.example)。

## 许可证

[MIT License](prism-fusion/LICENSE) © 2025-present [kwhitestone](https://github.com/kwhitestone)

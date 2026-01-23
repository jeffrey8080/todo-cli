# 📋 Todo CLI 命令完全使用指南

## 📊 命令总览

| 命令 | 用途 | 别名 | 核心功能 |
| :--- | :--- | :--- | :--- |
| **add** | 添加新任务 | 无 | 创建待办事项，可设置优先级和紧急标记 |
| **list** | 列出任务 | `ls`, `l` | 查看任务列表，支持筛选和排序 |
| **done** | 标记完成 | 无 | 将任务标记为已完成 |
| **edit** | 编辑任务 | `modify`, `update`, `e` | 修改任务属性，支持交互式和直接修改 |
| **delete** | 删除任务 | `del`, `rm` | 删除单个或多个任务，支持批量操作 |
| **completion** | Shell补全 | 无 | 生成自动补全脚本 |
| **help** | 查看帮助 | 无 | 查看命令帮助信息 |

---

## 🌟 **全局标志（所有命令可用）**

这些标志可以用于**任何命令**的任何位置：

| 标志 | 简写 | 说明 | 示例 |
| :--- | :--- | :--- | :--- |
| `--file` | `-f` | 指定不同的存储文件 | `todo -f work.json list` |
| `--help` | `-h` | 显示帮助信息 | `todo add --help` |
| `--version` | `-V` | 显示版本信息 | `todo --version` |

### 💡 全局标志特性
1. **位置灵活**：可以放在命令的任何位置
2. **作用全局**：对所有子命令都有效
3. **优先级高**：在命令标志之前解析

---

## 1️⃣ **add 命令 - 添加任务**

### 🎯 基本语法
```bash
todo add "任务描述" [选项]
```

### ⚙️ 可用选项
| 选项 | 简写 | 说明 | 可选值 | 默认值 |
| :--- | :--- | :--- | :--- | :--- |
| `--priority` | `-p` | 设置任务优先级 | `low`, `normal`, `high` | `normal` |
| `--urgent` | `-u` | 标记为紧急任务 | 布尔值 | `false` |

### 📝 详细说明

**任务描述规则**：
- 可以包含空格，需要用**双引号**包裹
- 如果不加引号，多个单词会被自动拼接
- 支持中英文和各种符号
- 紧急任务会自动添加 `[紧急] ` 前缀

**优先级系统**：
- `low`（低）：日常任务，不紧急
- `normal`（普通）：一般任务，默认值
- `high`（高）：重要任务，需要优先处理

### 🚀 使用示例
```bash
# 基本添加
todo add "购买明天的食材"
todo add "学习Golang数组和切片"

# 设置优先级
todo add "完成项目报告" --priority high
todo add "阅读文档" -p low

# 标记紧急
todo add "服务器需要重启" --urgent
todo add "紧急客户支持" -u

# 组合使用
todo add "发布新版本" -p high -u

# 长描述任务
todo add "本周五下午2点与市场部开会讨论Q2推广计划，需要准备PPT"

# 使用全局标志指定文件
todo -f work.json add "工作相关任务"
todo --file personal.json add "个人事务"
```

### ✅ 输出示例
```
✅ 已添加任务 #1: 购买明天的食材
✅ 已添加任务 #2: [紧急] 服务器需要重启
✅ 已添加任务 #3: 完成项目报告
```

### 🔍 查看帮助
```bash
todo add --help
todo add -h
```

---

## 2️⃣ **list 命令 - 列出任务**

### 🎯 基本语法
```bash
todo list [选项]
todo ls [选项]      # 使用别名
todo l [选项]       # 更短的别名
```

### ⚙️ 可用选项
| 选项 | 简写 | 说明 | 可选值 | 默认值 |
| :--- | :--- | :--- | :--- | :--- |
| `--all` | `-a` | 显示所有任务（包括已完成） | 布尔值 | `false` |
| `--done` | `-d` | 只显示已完成的任务 | 布尔值 | `false` |
| `--sort` | `-s` | 排序方式 | `id`, `time`, `priority` | `id` |

### 📝 详细说明

**显示模式**：
- **默认模式**：只显示**未完成**的任务
- **全部模式**（`--all`）：显示所有任务，包括已完成和未完成
- **完成模式**（`--done`）：只显示已完成的任务

**排序方式**：
- `id`：按任务ID排序（创建顺序）
- `time`：按创建时间排序（从新到旧）
- `priority`：按优先级排序（high → normal → low）

**显示格式**：
```
ID  状态  优先级  任务              创建时间
--  ----  ------  ----              --------
1   待办  normal  学习Cobra库       2024-01-15 10:30
2   ✅完成 high    [紧急]修复bug     2024-01-15 11:15
```

**状态图标说明**：
- ⏳ **待办**：任务未完成
- ✅ **完成**：任务已完成
- 🚨 **紧急**：任务有紧急标记（显示为 `[紧急] 任务描述`）

### 🚀 使用示例
```bash
# 基本查看（只显示未完成）
todo list
todo ls
todo l

# 显示所有任务
todo list --all
todo ls -a

# 只查看已完成
todo list --done
todo ls -d

# 按时间排序
todo list --sort time
todo ls -s time

# 按优先级排序
todo list --sort priority
todo ls -s priority

# 组合使用
todo list --all --sort priority
todo ls -a -s priority

# 使用全局标志查看其他文件
todo -f archive.json list --all
todo --file work.json ls
```

### ⚠️ 注意事项
1. `--all` 和 `--done` 是**互斥的**，不能同时使用
2. 如果没有任务，会显示 `📭 没有待办事项`
3. 紧急任务在描述前有 `[紧急] ` 标记

### 🔍 查看帮助
```bash
todo list --help
todo ls -h
```

---

## 3️⃣ **done 命令 - 标记任务完成**

### 🎯 基本语法
```bash
todo done <任务ID>
```

### 📋 参数说明
| 参数 | 必需 | 说明 | 示例 |
| :--- | :--- | :--- | :--- |
| `任务ID` | **是** | 要标记完成的任务ID | `1`, `3`, `5` |

### 📝 详细说明

**工作原理**：
1. 读取指定ID的任务
2. 将其 `Done` 状态设置为 `true`
3. 保存到文件
4. 显示成功消息

**错误处理**：
- 如果ID不存在：`未找到任务 #<ID>`
- 如果ID不是数字：`任务ID必须是数字`
- 如果任务已是完成状态：`⚠️ 任务 #<ID> 已经是完成状态`

### 🚀 使用示例
```bash
# 标记单个任务完成
todo done 1
todo done 3

# 连续标记多个任务（需要多次执行）
todo done 2
todo done 4
todo done 5

# 使用全局标志操作其他文件
todo -f work.json done 1
todo --file personal.json done 3
```

### ✅ 输出示例
```
🎉 已完成任务 #1: 学习Cobra库
⚠️  任务 #3 已经是完成状态
❌ 未找到任务 #10
```

### 💡 实用技巧
1. 先用 `todo list` 查看任务ID
2. 使用 `todo list --done` 查看所有已完成的任务
3. 标记错误时，可用 `todo edit <ID> -t` 重新标记为待办

### 🔍 查看帮助
```bash
todo done --help
```

---

# Todo Edit 命令使用说明

## 命令概述

`todo edit` 命令用于编辑已存在的待办事项。该命令允许您修改任务的描述、优先级、完成状态和紧急标记。

## 命令格式

```bash
todo edit <任务ID> [选项]
```

## 命令别名

- `todo modify`
- `todo update`
- `todo e`

## 选项参数

| 选项 | 短选项 | 参数 | 说明 |
|------|--------|------|------|
| `--desc` | `-d` | 字符串 | 设置新的任务描述 |
| `--priority` | `-p` | low/normal/high | 设置任务优先级 |
| `--toggle` | `-t` | 无 | 切换任务的完成状态 |
| `--urgent` | `-u` | 无 | 标记任务为紧急 |
| `--normal` | `-n` | 无 | 取消任务的紧急标记 |

## 优先级说明

| 优先级 | 说明 | 显示图标 |
|--------|------|----------|
| `low` | 低优先级 | 📉 |
| `normal` | 普通优先级（默认） | 📊 |
| `high` | 高优先级 | 📈 |

## 使用示例

### 基础使用

```bash
# 查看任务ID为1的详细信息（未指定修改选项时）
todo edit 1

# 修改任务描述
todo edit 1 -d "完成项目报告"

# 修改任务优先级为高
todo edit 2 -p high

# 切换任务完成状态
todo edit 3 -t
```

### 组合使用

```bash
# 同时修改描述和优先级
todo edit 1 -d "更新项目文档" -p high

# 修改描述并标记为紧急
todo edit 2 -d "紧急修复bug" -u

# 修改描述、切换状态并取消紧急标记
todo edit 3 -d "重新安排会议" -t -n
```

### 紧急标记操作

```bash
# 标记任务为紧急（添加 [紧急] 前缀）
todo edit 4 -u

# 取消任务的紧急标记
todo edit 4 -n

# 查看紧急标记的任务在列表中会有 🚨 图标
todo list
```

## 特殊注意事项

1. **任务ID**：必须是已存在的任务数字ID
2. **参数验证**：
    - 优先级必须是 `low`、`normal` 或 `high`
    - `--urgent` 和 `--normal` 不能同时使用
3. **修改确认**：所有修改会立即生效，请谨慎操作
4. **显示格式**：
    - 已完成任务显示 ✅ 图标
    - 待办任务显示 ⏳ 图标
    - 紧急任务显示 🚨 图标

## 错误处理

| 错误情况 | 错误信息 | 解决方法 |
|----------|----------|----------|
| 任务ID不存在 | `未找到任务 #<ID>` | 使用 `todo list` 查看有效ID |
| 任务ID不是数字 | `任务ID必须是数字` | 提供有效的数字ID |
| 无效优先级 | `优先级必须是 low, normal 或 high` | 使用正确的优先级值 |
| 同时使用 -u 和 -n | `不能同时使用 --urgent 和 --normal` | 选择其中一个选项 |

## 示例工作流

```bash
# 1. 查看所有任务
todo list

# 2. 查看任务详情（假设ID为2）
todo edit 2

# 3. 修改任务描述和优先级
todo edit 2 -d "准备会议材料" -p high

# 4. 完成任务
todo edit 2 -t

# 5. 标记任务为紧急（如果需要重新激活）
todo edit 2 -u
```


## 5️⃣ **delete 命令 - 删除任务**

### 🎯 基本语法
```bash
todo delete <任务ID...> [选项]
todo del <任务ID...> [选项]      # 使用别名
todo rm <任务ID...> [选项]       # 使用别名
```

### ⚙️ 可用选项
| 选项 | 简写 | 说明 | 默认值 |
| :--- | :--- | :--- | :--- |
| `--all` | `-a` | 删除所有任务 | `false` |
| `--force` | `-y` | 强制删除（不确认） | `false` |

### 📝 ID格式说明
支持多种ID格式，灵活组合：

| 格式 | 示例 | 说明 | 效果 |
| :--- | :--- | :--- | :--- |
| 单个ID | `1` | 删除任务1 | 删除#1 |
| 多个ID | `1 3 5` | 删除任务1,3,5 | 删除#1,#3,#5 |
| 逗号分隔 | `1,3,5` | 删除任务1,3,5 | 删除#1,#3,#5 |
| ID范围 | `1-5` | 删除任务1到5 | 删除#1,#2,#3,#4,#5 |
| 混合格式 | `1,3,5-7` | 删除任务1,3,5,6,7 | 删除#1,#3,#5,#6,#7 |

### 📋 详细说明

**安全机制**：
1. **默认需要确认**：删除前会列出要删除的任务，需要输入 `y` 确认
2. **强制删除**：使用 `-y` 选项跳过确认
3. **重新编号**：删除后自动重新编号，保持ID连续

**删除所有任务**：
- 需要额外确认（除非使用 `-y`）
- 清空整个任务列表
- 创建新的空文件

### 🚀 使用示例
```bash
# 删除单个任务（需要确认）
todo delete 1
# 显示：将要删除任务 #1: 学习Cobra库，确认吗？ (y/N):

# 删除多个任务
todo delete 1 3 5
todo del 2,4,6
todo rm 1-3

# 混合格式删除
todo delete 1,3,5-7
todo del "1 3 5-7"

# 强制删除（不确认）
todo delete 2 --force
todo del 1-3 -y

# 删除所有任务
todo delete --all
# 显示：⚠️  即将删除所有 5 个任务，确定吗？ (y/N):

# 强制删除所有
todo delete --all --force
todo del -a -y

# 使用全局标志
todo -f work.json delete 1
todo --file test.json del --all -y
```

### ✅ 输出示例
```
🗑️  已删除 3 个任务，剩余 2 个任务
⚠️  操作已取消
❌ 未找到任务 #10
📭 没有可删除的任务
```

### ⚠️ 注意事项
1. `--all` 不能和具体ID同时使用
2. ID范围必须有效（如 `1-3`，不能 `3-1`）
3. 删除后ID会重新编号，但创建时间保持不变
4. 删除操作**不可逆**，请谨慎使用

### 🔍 查看帮助
```bash
todo delete --help
todo del -h
```

---

## 6️⃣ **completion 命令 - Shell自动补全**

### 🎯 基本语法
```bash
todo completion <shell类型>
```

### 📋 支持的Shell类型
| Shell类型 | 说明 | 配置文件 |
| :--- | :--- | :--- |
| `bash` | Bash Shell | `~/.bashrc` 或 `~/.bash_profile` |
| `zsh` | Z Shell | `~/.zshrc` |
| `fish` | Fish Shell | `~/.config/fish/completions/` |
| `powershell` | PowerShell | Profile文件 |

### 📝 详细说明

**自动补全功能**：
- **命令补全**：输入 `todo a` + Tab → `todo add`
- **选项补全**：输入 `todo add -` + Tab → 显示所有可用选项
- **参数建议**：根据上下文提供建议

**生成脚本**：
- 生成对应Shell的补全脚本
- 可以保存到文件或直接执行
- 需要手动加载或添加到配置文件中

### 🚀 使用示例
```bash
# 查看支持的shell
todo completion --help

# 生成Bash补全脚本（显示到终端）
todo completion bash

# 生成并保存到文件
todo completion bash > ~/.todo-completion.bash

# 加载到当前会话
source ~/.todo-completion.bash

# 生成Zsh补全
todo completion zsh > ~/.zsh/completions/_todo

# 生成PowerShell补全（Windows）
todo completion powershell > todo-completion.ps1

# 永久生效（以Bash为例）
echo "source ~/.todo-completion.bash" >> ~/.bashrc
```

### 📋 各Shell配置方法

#### **Bash**
```bash
# 生成脚本
todo completion bash > ~/.todo-completion.bash

# 临时生效
source ~/.todo-completion.bash

# 永久生效
echo 'source ~/.todo-completion.bash' >> ~/.bashrc

# 重新加载配置
source ~/.bashrc
```

#### **Zsh**
```bash
# 生成脚本到正确位置
todo completion zsh > ~/.zsh/completions/_todo

# 确保补全目录在fpath中
echo 'fpath=(~/.zsh/completions $fpath)' >> ~/.zshrc

# 重新初始化补全系统
autoload -U compinit && compinit

# 重新加载配置
source ~/.zshrc
```

#### **Fish**
```bash
# 生成脚本到Fish补全目录
todo completion fish > ~/.config/fish/completions/todo.fish

# 重新启动Fish或重新加载配置
source ~/.config/fish/config.fish
```

#### **PowerShell**
```powershell
# 生成脚本
todo completion powershell > todo-completion.ps1

# 加载脚本
. .\todo-completion.ps1

# 永久生效（添加到Profile）
echo ". ~/todo-completion.ps1" >> $PROFILE
```

### ✅ 补全效果示例
```bash
# 命令补全
todo [Tab]      # 显示: add delete done edit help list
todo a[Tab]     # 自动补全为: todo add
todo l[Tab]     # 自动补全为: todo list

# 选项补全
todo add -[Tab] # 显示: -h --help -p --priority -u --urgent
todo list -[Tab] # 显示: -a --all -d --done -h --help -s --sort

# 参数建议（某些Shell支持）
todo done [Tab] # 显示可用的任务ID
todo delete [Tab] # 显示可用的任务ID
```

---

## 7️⃣ **help 命令 - 查看帮助**

### 🎯 基本语法
```bash
todo help [命令名]
todo --help [命令名]
todo -h [命令名]
```

### 📝 详细说明

**多层帮助系统**：
1. **全局帮助**：`todo help` 或 `todo --help` - 显示所有可用命令
2. **命令帮助**：`todo help <命令>` 或 `todo <命令> --help` - 显示特定命令的详细帮助
3. **选项帮助**：在命令帮助中显示所有可用选项

**帮助内容包含**：
- 命令用途和描述
- 使用语法
- 可用选项和参数
- 使用示例
- 全局选项说明

### 🚀 使用示例
```bash
# 查看全局帮助（所有命令）
todo help
todo --help
todo -h

# 查看特定命令帮助
todo help add
todo list --help
todo edit -h

# 查看所有详细信息
todo help --all

# 查看不存在的命令帮助
todo help nonexistent
# 显示：未知命令 "nonexistent" for "todo"
```

### 📋 帮助信息结构示例
```
add - 添加一个新的待办事项

添加一个待办事项到列表中。
可以指定优先级和紧急标志。

使用方法:
  todo add [任务描述] [选项]

示例:
  todo add "买牛奶"             # 添加简单任务
  todo add "写报告" -p high      # 高优先级任务
  todo add "紧急会议" -u         # 紧急任务

选项:
  -h, --help                显示帮助信息
  -p, --priority string     任务优先级 (low, normal, high) (默认 "normal")
  -u, --urgent              标记为紧急任务

全局选项:
  -f, --file string   指定待办事项存储文件 (默认 "todos.json")
```


这个Todo CLI工具提供了完整、灵活的任务管理功能。通过合理使用命令、选项和技巧，你可以高效地管理工作和个人任务。
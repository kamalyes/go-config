
## 🌍 环境管理

### 内置环境类型

#### 标准环境（9 个）

| 环境类型 | 常量 | 支持的别名 |
|---------|------|-----------|
| 开发环境 | `EnvDevelopment` | dev, develop, development |
| 本地环境 | `EnvLocal` | local, localhost |
| 测试环境 | `EnvTest` | test, testing, qa, sit |
| 预发布环境 | `EnvStaging` | staging, stage, stg, pre, preprod, fat, gray, canary |
| 生产环境 | `EnvProduction` | prod, production, prd, release, live, online, master |
| 调试环境 | `EnvDebug` | debug, debugging, dbg |
| 演示环境 | `EnvDemo` | demo, demonstration, showcase, preview, sandbox |
| UAT环境 | `EnvUAT` | uat, acceptance, user-acceptance, beta |
| 集成环境 | `EnvIntegration` | integration, int, ci, integration-test |

#### 全球化环境（47 个国家/地区）

**亚洲（20 个）**
- 中国 (china, cn, chn)
- 日本 (japan, jp, jpn)
- 韩国 (korea, kr, kor)
- 印度 (india, in, ind)
- 新加坡 (singapore, sg, sgp)
- 泰国 (thailand, th, tha)
- 越南 (vietnam, vn, vnm)
- 马来西亚 (malaysia, my, mys)
- 印度尼西亚 (indonesia, id, idn)
- 菲律宾 (philippines, ph, phl)
- 缅甸 (myanmar, mm, mmr)
- 老挝 (laos, la, lao)
- 柬埔寨 (cambodia, kh, khm)
- 巴基斯坦 (pakistan, pk, pak)
- 孟加拉国 (bangladesh, bd, bgd)
- 斯里兰卡 (srilanka, lk, lka)
- 尼泊尔 (nepal, np, npl)
- 香港 (hongkong, hk, hkg)
- 台湾 (taiwan, tw, twn)
- 澳门 (macao, mo, mac)

**欧洲（16 个）**
- 英国 (uk, gb, gbr)
- 德国 (germany, de, deu)
- 法国 (france, fr, fra)
- 意大利 (italy, it, ita)
- 西班牙 (spain, es, esp)
- 荷兰 (netherlands, nl, nld)
- 比利时 (belgium, be, bel)
- 瑞士 (switzerland, ch, che)
- 奥地利 (austria, at, aut)
- 瑞典 (sweden, se, swe)
- 挪威 (norway, no, nor)
- 丹麦 (denmark, dk, dnk)
- 芬兰 (finland, fi, fin)
- 波兰 (poland, pl, pol)
- 俄罗斯 (russia, ru, rus)
- 土耳其 (turkey, tr, tur)

**美洲（8 个）**
- 美国 (usa, us)
- 加拿大 (canada, ca, can)
- 墨西哥 (mexico, mx, mex)
- 巴西 (brazil, br, bra)
- 阿根廷 (argentina, ar, arg)
- 智利 (chile, cl, chl)
- 哥伦比亚 (colombia, co, col)
- 秘鲁 (peru, pe, per)

**其他地区（11 个）**
- 澳大利亚 (australia, au, aus)
- 新西兰 (newzealand, nz, nzl)
- 南非 (southafrica, za, zaf)
- 埃及 (egypt, eg, egy)
- 尼日利亚 (nigeria, ng, nga)
- 肯尼亚 (kenya, ke, ken)
- 阿联酋 (uae, ae, are)
- 沙特阿拉伯 (saudiarabia, sa, sau)
- 以色列 (israel, il, isr)
- 卡塔尔 (qatar, qa, qat)

### 配置文件命名规则

配置文件命名格式：`{prefix}-{env-suffix}.{ext}`

```bash
# 标准环境
gateway-xl-dev.yaml          # 开发环境
gateway-xl-prod.yaml         # 生产环境
gateway-xl-staging.yaml      # 预发布环境

# 国家/地区环境
gateway-xl-china.yaml        # 中国环境
gateway-xl-cn.yaml           # 中国环境（别名）
gateway-xl-japan.yaml        # 日本环境
gateway-xl-usa.yaml          # 美国环境
```

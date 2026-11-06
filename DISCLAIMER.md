<!--
Abate —— 免责声明
Copyright (C) 2026 小黎的部落（www.qingdujian.com）
AGPLv3; 商业用途须事先获得小黎的部落书面授权
-->
# Disclaimer / 免责声明

**Abate — Chaos Experiment Specification (CES) Reference Implementation**

Copyright (c) 小黎的部落 www.qingdujian.com. Licensed under AGPL-3.0-only + commercial license.

---

## English

### 1. No Warranty

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

### 2. Chaos Engineering Warning

Abate is a chaos engineering tool. It injects faults (latency, process kill, network partition, resource exhaustion, data corruption, clock skew) into target systems **by design**. Running fault injection experiments against any system you do not own or do not have explicit authorization to test may:

- cause service outages, data loss, or data corruption;
- violate laws, regulations, or contracts;
- breach terms of service of third-party platforms.

**You are solely responsible for obtaining authorization before running any experiment.** The authors accept no liability for any damage caused directly or indirectly by the use or misuse of this software.

### 3. Support Policy — "We Carry, We Don't Guarantee a Win"

Abate is open source, not open-ended support:

- **We carry, we don't guarantee a win.** Abate is the carry (main damage dealer) of its own architecture — but no software guarantees a win in your environment.
- **Fix it yourself.** If something breaks, fix it yourself. If you find a bug, fix it yourself.
- **Build it yourself.** If you need a plugin, an adapter, or a feature — build it yourself.
- **No support channel.** The maintainers do not answer questions about this open source project, in any form, on any channel. There is no schedule reserved for this.

Pull requests may or may not be reviewed. Issues may or may not be answered. There is no SLA, no roadmap commitment, and no response guarantee of any kind.

---

## 中文

### 1. 无担保声明

本软件按「现状」提供，不作任何明示或默示的担保，包括但不限于适销性、特定用途适用性及不侵权。在任何情况下，作者或版权持有人均不对因使用或无法使用本软件而产生的任何索赔、损害或其他责任（包括合同之诉、侵权之诉或其他）承担责任。

### 2. 混沌工程风险警告

Abate 是混沌工程工具，其设计目的就是向目标系统注入故障（延迟、进程终止、网络分区、资源耗尽、数据损坏、时钟偏移）。对任何你不拥有或未获得明确授权的系统运行故障注入实验，可能：

- 导致服务中断、数据丢失或数据损坏；
- 违反法律法规或合同约定；
- 违反第三方平台服务条款。

**运行任何实验前，获得授权的责任完全由你自行承担。** 作者不对因使用或滥用本软件直接或间接造成的任何损害承担责任。

### 3. 支持政策——「包C不包赢」

Abate 开源，但不开无限支持：

- **包C不包赢。** Abate 在自己的架构里是主C（核心输出位），但没有任何软件能保证在你的环境里赢。
- **有问题自己修。** 有问题自己修，有 bug 自己修。
- **需要什么自己动手。** 需要什么插件自己做，一切需要什么自己动手。
- **不回复开源问题。** 开源方不以任何形式回复开源问题，没有时间档期。

Pull Request 可能看也可能不看，Issue 可能回也可能不回。无 SLA、无路线图承诺、无任何形式的回复保证。

---

本声明与 [LICENSE](LICENSE)（AGPL-3.0-only + 商业授权）及 [THIRD-PARTY-NOTICES.md](THIRD-PARTY-NOTICES.md) 共同构成 Abate 项目的完整许可与免责框架。

/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

// htmlInlineCSS is the complete inline stylesheet. It uses local system fonts
// only and contains no remote reference, no external font and no remote sheet.
const htmlInlineCSS = `:root{color-scheme:light;--ink:#1b1f23;--muted:#57606a;--line:#d0d7de;--bg:#ffffff;--panel:#f6f8fa}
*{box-sizing:border-box}
body{margin:0;padding:2rem;background:var(--bg);color:var(--ink);font:14px/1.5 system-ui,-apple-system,"Segoe UI",sans-serif}
main{max-width:1100px;margin:0 auto}
h1{font-size:1.5rem;margin:0 0 .25rem}
h2{font-size:1.1rem;margin:1.5rem 0 .5rem}
p{margin:.25rem 0}
.digest{color:var(--muted);font-family:ui-monospace,monospace}
table{border-collapse:collapse;width:100%;font-size:.9rem}
th,td{border:1px solid var(--line);padding:.35rem .5rem;text-align:left;vertical-align:top}
th{background:var(--panel)}
code{font-family:ui-monospace,monospace;word-break:break-all}
.gauge{display:flex;flex-wrap:wrap;gap:1rem}
.gauge div{border:1px solid var(--line);padding:.5rem .75rem;min-width:8rem}
.gauge b{display:block;font-size:1.2rem}
[data-severity="error"]{border-left:3px solid #cf222e}
[data-severity="warning"]{border-left:3px solid #bf8700}
[data-severity="info"]{border-left:3px solid #0969da}
footer{margin-top:2rem;padding-top:.5rem;border-top:1px solid var(--line);color:var(--muted);font-family:ui-monospace,monospace;word-break:break-all}`

// htmlInlineScript is the complete inline behavior script. It annotates local
// table rows only and performs no network request and no dynamic evaluation.
const htmlInlineScript = `(function () {
  var rows = document.querySelectorAll('[data-severity]');
  for (var i = 0; i < rows.length; i += 1) {
    rows[i].setAttribute('data-row', String(i + 1));
  }
  var phases = document.querySelectorAll('[data-phase]');
  for (var j = 0; j < phases.length; j += 1) {
    phases[j].setAttribute('data-step', String(j));
  }
}());`

/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

// pyRecipe covers decorators, async contexts, match/case, comprehensions,
// generators, type annotations, classes, lambdas and exception finalizers.
func pyRecipe() Recipe {
	return Recipe{
		Language: "py",
		Ext:      "py",
		Tags:     []string{"func", "class", "exception", "async", "comprehension", "match", "typing", "generator"},
		Header:   "\n",
		AnchorClean: `def anchor_probe():
    return 1 + 2 == 3

`,
		AnchorFull: `def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

`,
		CleanBlock: "def step@@I@@_@@S@@():\n    return @@N@@ + @@S@@\n",
		Blocks: []BlockTemplate{
			{Text: `def traced@@I@@_@@S@@(value: int) -> int:
    return value + @@N@@


def decorate@@I@@_@@S@@(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate@@I@@_@@S@@
def doubled@@I@@_@@S@@(value: int) -> int:
    return value + value
`},
			{Text: `async def fetch@@I@@_@@S@@(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
`},
			{Text: `def summarize@@I@@_@@S@@(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > @@N@@:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
`},
			{Text: `def expand@@I@@_@@S@@(values):
    for item in values:
        if item > @@N@@:
            yield item * @@M@@


def collect@@I@@_@@S@@(values):
    return list(expand@@I@@_@@S@@(values))
`},
			{Text: `class Slot@@I@@_@@S@@:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{@@N@@}"

    @property
    def doubled(self) -> int:
        return self.size * 2
`},
			{Text: `def project@@I@@_@@S@@(values):
    select = lambda item: item + @@N@@
    return [select(item) for item in values if item != 0]
`},
			{Text: `def guarded@@I@@_@@S@@(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
`},
			{Text: `def pair@@I@@_@@S@@(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= @@N@@}
`},
		},
	}
}


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand426_0(values):
    for item in values:
        if item > 14:
            yield item * 63


def collect426_0(values):
    return list(expand426_0(values))
async def fetch426_1(gate):
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
def project426_2(values):
    select = lambda item: item + 28
    return [select(item) for item in values if item != 0]
def expand426_3(values):
    for item in values:
        if item > 35:
            yield item * 35


def collect426_3(values):
    return list(expand426_3(values))
async def fetch426_4(gate):
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
def pair426_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 49}
def pair426_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 56}
def project426_7(values):
    select = lambda item: item + 63
    return [select(item) for item in values if item != 0]
class Slot426_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{70}"

    @property
    def doubled(self) -> int:
        return self.size * 2

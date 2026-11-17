
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair120_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 13}
def expand120_1(values):
    for item in values:
        if item > 20:
            yield item * 67


def collect120_1(values):
    return list(expand120_1(values))
def pair120_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 27}
def project120_3(values):
    select = lambda item: item + 34
    return [select(item) for item in values if item != 0]
class Slot120_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{41}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch120_5(gate):
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

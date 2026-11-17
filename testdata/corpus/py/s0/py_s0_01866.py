
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot848_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{68}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot848_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{75}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch848_2(gate):
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
def expand848_3(values):
    for item in values:
        if item > 89:
            yield item * 61


def collect848_3(values):
    return list(expand848_3(values))
def expand848_4(values):
    for item in values:
        if item > 96:
            yield item * 11


def collect848_4(values):
    return list(expand848_4(values))


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch506_0(gate):
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
class Slot506_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{91}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch506_2(gate):
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
def expand506_3(values):
    for item in values:
        if item > 8:
            yield item * 46


def collect506_3(values):
    return list(expand506_3(values))

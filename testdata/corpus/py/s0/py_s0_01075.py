
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot497_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{64}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch497_1(gate):
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
def expand497_2(values):
    for item in values:
        if item > 78:
            yield item * 33


def collect497_2(values):
    return list(expand497_2(values))

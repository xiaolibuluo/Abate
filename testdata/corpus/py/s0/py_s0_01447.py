
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand652_0(values):
    for item in values:
        if item > 42:
            yield item * 59


def collect652_0(values):
    return list(expand652_0(values))
def expand652_1(values):
    for item in values:
        if item > 49:
            yield item * 9


def collect652_1(values):
    return list(expand652_1(values))
class Slot652_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{56}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot652_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{63}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch652_4(gate):
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

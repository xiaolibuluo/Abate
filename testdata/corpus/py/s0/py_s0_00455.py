
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand210_0(values):
    for item in values:
        if item > 19:
            yield item * 15


def collect210_0(values):
    return list(expand210_0(values))
def project210_1(values):
    select = lambda item: item + 26
    return [select(item) for item in values if item != 0]
async def fetch210_2(gate):
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
class Slot210_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{40}"

    @property
    def doubled(self) -> int:
        return self.size * 2

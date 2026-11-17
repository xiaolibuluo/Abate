
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch305_0(gate):
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
class Slot305_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{97}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project305_2(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]
def expand305_3(values):
    for item in values:
        if item > 14:
            yield item * 42


def collect305_3(values):
    return list(expand305_3(values))

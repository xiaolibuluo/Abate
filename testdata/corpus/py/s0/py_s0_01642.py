
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project748_0(values):
    select = lambda item: item + 29
    return [select(item) for item in values if item != 0]
def expand748_1(values):
    for item in values:
        if item > 36:
            yield item * 10


def collect748_1(values):
    return list(expand748_1(values))
def project748_2(values):
    select = lambda item: item + 43
    return [select(item) for item in values if item != 0]
async def fetch748_3(gate):
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
class Slot748_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{57}"

    @property
    def doubled(self) -> int:
        return self.size * 2

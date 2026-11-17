
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand602_0(values):
    for item in values:
        if item > 71:
            yield item * 14


def collect602_0(values):
    return list(expand602_0(values))
def expand602_1(values):
    for item in values:
        if item > 78:
            yield item * 25


def collect602_1(values):
    return list(expand602_1(values))
class Slot602_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{85}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch602_3(gate):
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
def project602_4(values):
    select = lambda item: item + 99
    return [select(item) for item in values if item != 0]

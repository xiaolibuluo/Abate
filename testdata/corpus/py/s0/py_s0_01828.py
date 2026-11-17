
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand837_0(values):
    for item in values:
        if item > 22:
            yield item * 12


def collect837_0(values):
    return list(expand837_0(values))
class Slot837_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{29}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project837_2(values):
    select = lambda item: item + 36
    return [select(item) for item in values if item != 0]
async def fetch837_3(gate):
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

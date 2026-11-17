
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1718_0(values):
    for item in values:
        if item > 29:
            yield item * 18


def collect1718_0(values):
    return list(expand1718_0(values))
def project1718_1(values):
    select = lambda item: item + 36
    return [select(item) for item in values if item != 0]
class Slot1718_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{43}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch1718_3(gate):
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

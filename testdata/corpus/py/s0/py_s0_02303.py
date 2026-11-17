
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1050_0(values):
    for item in values:
        if item > 75:
            yield item * 39


def collect1050_0(values):
    return list(expand1050_0(values))
def expand1050_1(values):
    for item in values:
        if item > 82:
            yield item * 50


def collect1050_1(values):
    return list(expand1050_1(values))
async def fetch1050_2(gate):
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
def expand1050_3(values):
    for item in values:
        if item > 96:
            yield item * 11


def collect1050_3(values):
    return list(expand1050_3(values))
def project1050_4(values):
    select = lambda item: item + 6
    return [select(item) for item in values if item != 0]
def expand1050_5(values):
    for item in values:
        if item > 13:
            yield item * 33


def collect1050_5(values):
    return list(expand1050_5(values))
class Slot1050_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{20}"

    @property
    def doubled(self) -> int:
        return self.size * 2

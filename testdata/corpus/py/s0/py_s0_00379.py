
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot175_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{49}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project175_1(values):
    select = lambda item: item + 56
    return [select(item) for item in values if item != 0]
def expand175_2(values):
    for item in values:
        if item > 63:
            yield item * 36


def collect175_2(values):
    return list(expand175_2(values))
def expand175_3(values):
    for item in values:
        if item > 70:
            yield item * 47


def collect175_3(values):
    return list(expand175_3(values))
class Slot175_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{77}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch175_5(gate):
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
def guarded175_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()

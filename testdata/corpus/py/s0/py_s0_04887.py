
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2262_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{20}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch2262_1(gate):
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
def project2262_2(values):
    select = lambda item: item + 34
    return [select(item) for item in values if item != 0]
def guarded2262_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2262_4(values):
    for item in values:
        if item > 48:
            yield item * 27


def collect2262_4(values):
    return list(expand2262_4(values))
async def fetch2262_5(gate):
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

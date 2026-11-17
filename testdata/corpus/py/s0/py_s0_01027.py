
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded476_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand476_1(values):
    for item in values:
        if item > 89:
            yield item * 58


def collect476_1(values):
    return list(expand476_1(values))
def project476_2(values):
    select = lambda item: item + 96
    return [select(item) for item in values if item != 0]
def project476_3(values):
    select = lambda item: item + 6
    return [select(item) for item in values if item != 0]
def project476_4(values):
    select = lambda item: item + 13
    return [select(item) for item in values if item != 0]
async def fetch476_5(gate):
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
class Slot476_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{27}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot476_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{34}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair476_8(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 41}


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2015_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2015_1(values):
    for item in values:
        if item > 17:
            yield item * 34


def collect2015_1(values):
    return list(expand2015_1(values))
async def fetch2015_2(gate):
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
class Slot2015_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{31}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot2015_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{38}"

    @property
    def doubled(self) -> int:
        return self.size * 2

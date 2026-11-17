
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot709_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{7}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded709_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced709_2(value: int) -> int:
    return value + 21


def decorate709_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate709_2
def doubled709_2(value: int) -> int:
    return value + value
async def fetch709_3(gate):
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
class Slot709_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{35}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand709_5(values):
    for item in values:
        if item > 42:
            yield item * 25


def collect709_5(values):
    return list(expand709_5(values))

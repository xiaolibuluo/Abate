
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1730_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1730_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{95}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced1730_2(value: int) -> int:
    return value + 5


def decorate1730_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1730_2
def doubled1730_2(value: int) -> int:
    return value + value
async def fetch1730_3(gate):
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
class Slot1730_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{19}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1730_5(values):
    for item in values:
        if item > 26:
            yield item * 35


def collect1730_5(values):
    return list(expand1730_5(values))


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1439_0(value: int) -> int:
    return value + 88


def decorate1439_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1439_0
def doubled1439_0(value: int) -> int:
    return value + value
class Slot1439_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{95}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1439_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1439_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{12}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch1439_4(gate):
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
def expand1439_5(values):
    for item in values:
        if item > 26:
            yield item * 11


def collect1439_5(values):
    return list(expand1439_5(values))


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1927_0(gate):
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
def traced1927_1(value: int) -> int:
    return value + 37


def decorate1927_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1927_1
def doubled1927_1(value: int) -> int:
    return value + value
class Slot1927_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{44}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch1927_3(gate):
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
def guarded1927_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1927_5(values):
    for item in values:
        if item > 65:
            yield item * 11


def collect1927_5(values):
    return list(expand1927_5(values))

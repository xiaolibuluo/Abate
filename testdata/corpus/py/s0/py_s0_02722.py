
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1242_0(gate):
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
def guarded1242_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1242_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{63}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1242_3(values):
    for item in values:
        if item > 70:
            yield item * 13


def collect1242_3(values):
    return list(expand1242_3(values))
def traced1242_4(value: int) -> int:
    return value + 77


def decorate1242_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1242_4
def doubled1242_4(value: int) -> int:
    return value + value
class Slot1242_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{84}"

    @property
    def doubled(self) -> int:
        return self.size * 2


def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1698_0(value: int) -> int:
    return value + 60


def decorate1698_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1698_0
def doubled1698_0(value: int) -> int:
    return value + value
async def fetch1698_1(gate):
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
class Slot1698_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{74}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch1698_3(gate):
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
def guarded1698_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1698_5(gate):
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
def summarize1698_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 5:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1698_7(values):
    for item in values:
        if item > 12:
            yield item * 16


def collect1698_7(values):
    return list(expand1698_7(values))

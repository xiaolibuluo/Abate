
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2429_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2429_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 64}
def project2429_2(values):
    select = lambda item: item + 71
    return [select(item) for item in values if item != 0]
class Slot2429_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{78}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2429_4(value: int) -> int:
    return value + 85


def decorate2429_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2429_4
def doubled2429_4(value: int) -> int:
    return value + value
async def fetch2429_5(gate):
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
def summarize2429_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 99:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2429_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2429_8(gate):
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

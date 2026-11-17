
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair2426_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 18}
def pair2426_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 25}
async def fetch2426_2(gate):
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
def guarded2426_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2426_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2426_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{53}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2426_6(value: int) -> int:
    return value + 60


def decorate2426_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2426_6
def doubled2426_6(value: int) -> int:
    return value + value
def traced2426_7(value: int) -> int:
    return value + 67


def decorate2426_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2426_7
def doubled2426_7(value: int) -> int:
    return value + value
def summarize2426_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 74:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

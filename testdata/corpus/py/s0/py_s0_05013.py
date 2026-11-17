
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2328_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2328_1(value: int) -> int:
    return value + 12


def decorate2328_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2328_1
def doubled2328_1(value: int) -> int:
    return value + value
def summarize2328_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 19:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair2328_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 26}
async def fetch2328_4(gate):
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
class Slot2328_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{40}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def traced2328_6(value: int) -> int:
    return value + 47


def decorate2328_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2328_6
def doubled2328_6(value: int) -> int:
    return value + value
def pair2328_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 54}

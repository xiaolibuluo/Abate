
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2520_0(gate):
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
def summarize2520_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 83:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2520_2(values):
    for item in values:
        if item > 90:
            yield item * 42


def collect2520_2(values):
    return list(expand2520_2(values))
class Slot2520_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{97}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded2520_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2520_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2520_6(value: int) -> int:
    return value + 21


def decorate2520_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2520_6
def doubled2520_6(value: int) -> int:
    return value + value
def traced2520_7(value: int) -> int:
    return value + 28


def decorate2520_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2520_7
def doubled2520_7(value: int) -> int:
    return value + value

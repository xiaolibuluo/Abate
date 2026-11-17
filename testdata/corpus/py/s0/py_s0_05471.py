
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2528_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2528_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 90}
def pair2528_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 97}
async def fetch2528_3(gate):
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
def project2528_4(values):
    select = lambda item: item + 14
    return [select(item) for item in values if item != 0]
def summarize2528_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 21:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2528_6(value: int) -> int:
    return value + 28


def decorate2528_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2528_6
def doubled2528_6(value: int) -> int:
    return value + value
def traced2528_7(value: int) -> int:
    return value + 35


def decorate2528_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2528_7
def doubled2528_7(value: int) -> int:
    return value + value
class Slot2528_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{42}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def pair2528_9(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 49}

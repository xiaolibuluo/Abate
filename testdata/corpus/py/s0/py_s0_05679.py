
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2619_0(value: int) -> int:
    return value + 5


def decorate2619_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2619_0
def doubled2619_0(value: int) -> int:
    return value + value
async def fetch2619_1(gate):
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
def traced2619_2(value: int) -> int:
    return value + 19


def decorate2619_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2619_2
def doubled2619_2(value: int) -> int:
    return value + value
def expand2619_3(values):
    for item in values:
        if item > 26:
            yield item * 14


def collect2619_3(values):
    return list(expand2619_3(values))
def summarize2619_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 33:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded2619_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot2619_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{47}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project2619_7(values):
    select = lambda item: item + 54
    return [select(item) for item in values if item != 0]

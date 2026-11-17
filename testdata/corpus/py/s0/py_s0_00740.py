
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize337_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 21:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced337_1(value: int) -> int:
    return value + 28


def decorate337_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate337_1
def doubled337_1(value: int) -> int:
    return value + value
def traced337_2(value: int) -> int:
    return value + 35


def decorate337_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate337_2
def doubled337_2(value: int) -> int:
    return value + value
def summarize337_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 42:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot337_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{49}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch337_5(gate):
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
def guarded337_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced337_7(value: int) -> int:
    return value + 70


def decorate337_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate337_7
def doubled337_7(value: int) -> int:
    return value + value

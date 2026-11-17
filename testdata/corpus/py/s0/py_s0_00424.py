
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project194_0(values):
    select = lambda item: item + 5
    return [select(item) for item in values if item != 0]
def expand194_1(values):
    for item in values:
        if item > 12:
            yield item * 36


def collect194_1(values):
    return list(expand194_1(values))
def project194_2(values):
    select = lambda item: item + 19
    return [select(item) for item in values if item != 0]
async def fetch194_3(gate):
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
def expand194_4(values):
    for item in values:
        if item > 33:
            yield item * 69


def collect194_4(values):
    return list(expand194_4(values))
def traced194_5(value: int) -> int:
    return value + 40


def decorate194_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate194_5
def doubled194_5(value: int) -> int:
    return value + value
def guarded194_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot194_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{54}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded194_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize194_9(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 68:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot194_10:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{75}"

    @property
    def doubled(self) -> int:
        return self.size * 2

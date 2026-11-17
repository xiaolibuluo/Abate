
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand48_0(values):
    for item in values:
        if item > 47:
            yield item * 40


def collect48_0(values):
    return list(expand48_0(values))
def project48_1(values):
    select = lambda item: item + 54
    return [select(item) for item in values if item != 0]
class Slot48_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{61}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch48_3(gate):
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
def guarded48_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced48_5(value: int) -> int:
    return value + 82


def decorate48_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate48_5
def doubled48_5(value: int) -> int:
    return value + value
def guarded48_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot48_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{96}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded48_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize48_9(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

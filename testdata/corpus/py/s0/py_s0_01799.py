
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced823_0(value: int) -> int:
    return value + 34


def decorate823_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate823_0
def doubled823_0(value: int) -> int:
    return value + value
async def fetch823_1(gate):
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
def project823_2(values):
    select = lambda item: item + 48
    return [select(item) for item in values if item != 0]
class Slot823_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{55}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded823_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project823_5(values):
    select = lambda item: item + 69
    return [select(item) for item in values if item != 0]
class Slot823_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{76}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project823_7(values):
    select = lambda item: item + 83
    return [select(item) for item in values if item != 0]
def guarded823_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize823_9(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 97:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

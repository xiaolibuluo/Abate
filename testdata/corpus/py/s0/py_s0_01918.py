
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand873_0(values):
    for item in values:
        if item > 5:
            yield item * 20


def collect873_0(values):
    return list(expand873_0(values))
def expand873_1(values):
    for item in values:
        if item > 12:
            yield item * 31


def collect873_1(values):
    return list(expand873_1(values))
async def fetch873_2(gate):
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
def guarded873_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot873_4:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{33}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot873_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{40}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded873_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize873_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 54:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded873_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot873_9:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{68}"

    @property
    def doubled(self) -> int:
        return self.size * 2

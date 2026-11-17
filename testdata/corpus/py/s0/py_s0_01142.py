
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded528_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded528_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project528_2(values):
    select = lambda item: item + 93
    return [select(item) for item in values if item != 0]
def expand528_3(values):
    for item in values:
        if item > 100:
            yield item * 17


def collect528_3(values):
    return list(expand528_3(values))
def project528_4(values):
    select = lambda item: item + 10
    return [select(item) for item in values if item != 0]
def summarize528_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 17:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot528_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{24}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch528_7(gate):
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
def expand528_8(values):
    for item in values:
        if item > 38:
            yield item * 11


def collect528_8(values):
    return list(expand528_8(values))
async def fetch528_9(gate):
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

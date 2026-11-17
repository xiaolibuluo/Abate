
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded47_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand47_1(values):
    for item in values:
        if item > 41:
            yield item * 44


def collect47_1(values):
    return list(expand47_1(values))
def summarize47_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 48:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project47_3(values):
    select = lambda item: item + 55
    return [select(item) for item in values if item != 0]
def project47_4(values):
    select = lambda item: item + 62
    return [select(item) for item in values if item != 0]
async def fetch47_5(gate):
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
def project47_6(values):
    select = lambda item: item + 76
    return [select(item) for item in values if item != 0]
def guarded47_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot47_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{90}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch47_9(gate):
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
async def fetch47_10(gate):
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

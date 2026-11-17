
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1751_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1751_1(values):
    for item in values:
        if item > 77:
            yield item * 16


def collect1751_1(values):
    return list(expand1751_1(values))
class Slot1751_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{84}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand1751_3(values):
    for item in values:
        if item > 91:
            yield item * 38


def collect1751_3(values):
    return list(expand1751_3(values))
def summarize1751_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 98:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1751_5(values):
    for item in values:
        if item > 8:
            yield item * 60


def collect1751_5(values):
    return list(expand1751_5(values))
async def fetch1751_6(gate):
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
def pair1751_7(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 22}
class Slot1751_8:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{29}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded1751_9(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()

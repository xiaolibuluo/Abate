
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1213_0(values):
    for item in values:
        if item > 60:
            yield item * 21


def collect1213_0(values):
    return list(expand1213_0(values))
def summarize1213_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 67:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1213_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1213_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{81}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def project1213_4(values):
    select = lambda item: item + 88
    return [select(item) for item in values if item != 0]
async def fetch1213_5(gate):
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
class Slot1213_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{5}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot1213_7:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{12}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch1213_8(gate):
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

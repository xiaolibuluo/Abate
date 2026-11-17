
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot961_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{82}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot961_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{89}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch961_2(gate):
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
def project961_3(values):
    select = lambda item: item + 6
    return [select(item) for item in values if item != 0]
def project961_4(values):
    select = lambda item: item + 13
    return [select(item) for item in values if item != 0]
def project961_5(values):
    select = lambda item: item + 20
    return [select(item) for item in values if item != 0]
def summarize961_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 27:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch961_7(gate):
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

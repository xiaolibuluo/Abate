
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded473_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project473_1(values):
    select = lambda item: item + 50
    return [select(item) for item in values if item != 0]
class Slot473_2:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{57}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot473_3:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{64}"

    @property
    def doubled(self) -> int:
        return self.size * 2
async def fetch473_4(gate):
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
def summarize473_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 78:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
class Slot473_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{85}"

    @property
    def doubled(self) -> int:
        return self.size * 2

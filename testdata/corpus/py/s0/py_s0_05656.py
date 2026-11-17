
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

class Slot2606_0:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{30}"

    @property
    def doubled(self) -> int:
        return self.size * 2
class Slot2606_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{37}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def expand2606_2(values):
    for item in values:
        if item > 44:
            yield item * 34


def collect2606_2(values):
    return list(expand2606_2(values))
def guarded2606_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2606_4(gate):
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
def project2606_5(values):
    select = lambda item: item + 65
    return [select(item) for item in values if item != 0]

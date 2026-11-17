
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded810_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot810_1:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{66}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded810_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch810_3(gate):
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
def expand810_4(values):
    for item in values:
        if item > 87:
            yield item * 50


def collect810_4(values):
    return list(expand810_4(values))
def expand810_5(values):
    for item in values:
        if item > 94:
            yield item * 61


def collect810_5(values):
    return list(expand810_5(values))
def guarded810_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()

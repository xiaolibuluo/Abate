
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded174_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project174_1(values):
    select = lambda item: item + 43
    return [select(item) for item in values if item != 0]
def summarize174_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 50:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded174_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand174_4(values):
    for item in values:
        if item > 64:
            yield item * 51


def collect174_4(values):
    return list(expand174_4(values))
class Slot174_5:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{71}"

    @property
    def doubled(self) -> int:
        return self.size * 2
def guarded174_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()

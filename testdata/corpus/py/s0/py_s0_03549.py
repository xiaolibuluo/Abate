
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1636_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1636_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1636_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1636_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 51:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1636_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 58:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1636_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
class Slot1636_6:
    def __init__(self, size: int) -> None:
        self.size = size

    def describe(self) -> str:
        return f"slot-{self.size}-{72}"

    @property
    def doubled(self) -> int:
        return self.size * 2

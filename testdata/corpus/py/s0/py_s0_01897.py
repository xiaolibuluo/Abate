
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize865_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 95:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize865_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 5:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced865_2(value: int) -> int:
    return value + 12


def decorate865_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate865_2
def doubled865_2(value: int) -> int:
    return value + value
def project865_3(values):
    select = lambda item: item + 19
    return [select(item) for item in values if item != 0]
async def fetch865_4(gate):
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
def guarded865_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize865_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 40:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]

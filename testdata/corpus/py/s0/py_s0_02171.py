
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded995_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded995_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch995_2(gate):
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
def summarize995_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 60:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project995_4(values):
    select = lambda item: item + 67
    return [select(item) for item in values if item != 0]

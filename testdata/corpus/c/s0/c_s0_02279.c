
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1539_0)(int value);

typedef struct Slot1539_0 {
    int slots[16];
    predicate1539_0 accept;
} Slot1539_0;

int slot_accept1539_0(Slot1539_0 *store, int value) {
    return store->accept(value) + 30;
}
#if defined(__STDC_VERSION__)
int variant1539_1(int value, ...) {
    return value + 37;
}
#else
int variant1539_1(int value) {
    return value + 37;
}
#endif
#define MASK1539_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1539_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1539_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1539_3 { KIND_EMPTY1539_3 = 0, KIND_FULL1539_3 = 51 };

union Cell1539_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1539_3(union Cell1539_3 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL1539_3 : KIND_EMPTY1539_3;
}
#define MASK1539_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1539_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1539_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1539_5(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 65;
}

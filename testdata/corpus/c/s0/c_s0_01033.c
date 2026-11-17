
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind708_0 { KIND_EMPTY708_0 = 0, KIND_FULL708_0 = 91 };

union Cell708_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind708_0(union Cell708_0 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL708_0 : KIND_EMPTY708_0;
}
#if defined(__STDC_VERSION__)
int variant708_1(int value, ...) {
    return value + 98;
}
#else
int variant708_1(int value) {
    return value + 98;
}
#endif
typedef int (*predicate708_2)(int value);

typedef struct Slot708_2 {
    int slots[16];
    predicate708_2 accept;
} Slot708_2;

int slot_accept708_2(Slot708_2 *store, int value) {
    return store->accept(value) + 8;
}
int scan708_3(const int *begin, const int *end) {
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
    return total + 15;
}
#define MASK708_4(v) (((v) << 4) | ((v) & 0x0F))

int classify708_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK708_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

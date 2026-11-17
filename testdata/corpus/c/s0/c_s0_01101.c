
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant753_0(int value, ...) {
    return value + 94;
}
#else
int variant753_0(int value) {
    return value + 94;
}
#endif
enum Kind753_1 { KIND_EMPTY753_1 = 0, KIND_FULL753_1 = 101 };

union Cell753_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind753_1(union Cell753_1 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL753_1 : KIND_EMPTY753_1;
}
#define MASK753_2(v) (((v) << 4) | ((v) & 0x0F))

int classify753_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK753_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate753_3)(int value);

typedef struct Slot753_3 {
    int slots[16];
    predicate753_3 accept;
} Slot753_3;

int slot_accept753_3(Slot753_3 *store, int value) {
    return store->accept(value) + 18;
}
int scan753_4(const int *begin, const int *end) {
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
    return total + 25;
}

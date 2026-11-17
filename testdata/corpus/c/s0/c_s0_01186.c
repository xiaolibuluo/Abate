
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant813_0(int value, ...) {
    return value + 98;
}
#else
int variant813_0(int value) {
    return value + 98;
}
#endif
typedef int (*predicate813_1)(int value);

typedef struct Slot813_1 {
    int slots[16];
    predicate813_1 accept;
} Slot813_1;

int slot_accept813_1(Slot813_1 *store, int value) {
    return store->accept(value) + 8;
}
#define MASK813_2(v) (((v) << 4) | ((v) & 0x0F))

int classify813_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK813_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind813_3 { KIND_EMPTY813_3 = 0, KIND_FULL813_3 = 22 };

union Cell813_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind813_3(union Cell813_3 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL813_3 : KIND_EMPTY813_3;
}
int scan813_4(const int *begin, const int *end) {
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
    return total + 29;
}
#if defined(__STDC_VERSION__)
int variant813_5(int value, ...) {
    return value + 36;
}
#else
int variant813_5(int value) {
    return value + 36;
}
#endif
enum Kind813_6 { KIND_EMPTY813_6 = 0, KIND_FULL813_6 = 43 };

union Cell813_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind813_6(union Cell813_6 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL813_6 : KIND_EMPTY813_6;
}

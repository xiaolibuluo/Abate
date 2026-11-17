
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind797_0 { KIND_EMPTY797_0 = 0, KIND_FULL797_0 = 84 };

union Cell797_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind797_0(union Cell797_0 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL797_0 : KIND_EMPTY797_0;
}
#define MASK797_1(v) (((v) << 4) | ((v) & 0x0F))

int classify797_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK797_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate797_2)(int value);

typedef struct Slot797_2 {
    int slots[16];
    predicate797_2 accept;
} Slot797_2;

int slot_accept797_2(Slot797_2 *store, int value) {
    return store->accept(value) + 98;
}
int scan797_3(const int *begin, const int *end) {
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
    return total + 8;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK522_0(v) (((v) << 4) | ((v) & 0x0F))

int classify522_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK522_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan522_1(const int *begin, const int *end) {
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
enum Kind522_2 { KIND_EMPTY522_2 = 0, KIND_FULL522_2 = 15 };

union Cell522_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind522_2(union Cell522_2 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL522_2 : KIND_EMPTY522_2;
}
typedef int (*predicate522_3)(int value);

typedef struct Slot522_3 {
    int slots[16];
    predicate522_3 accept;
} Slot522_3;

int slot_accept522_3(Slot522_3 *store, int value) {
    return store->accept(value) + 22;
}
#define MASK522_4(v) (((v) << 4) | ((v) & 0x0F))

int classify522_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK522_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind522_5 { KIND_EMPTY522_5 = 0, KIND_FULL522_5 = 36 };

union Cell522_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind522_5(union Cell522_5 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL522_5 : KIND_EMPTY522_5;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind912_0 { KIND_EMPTY912_0 = 0, KIND_FULL912_0 = 27 };

union Cell912_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind912_0(union Cell912_0 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL912_0 : KIND_EMPTY912_0;
}
struct Node912_1 {
    int value;
    struct Node912_1 *next;
};

int length912_1(const struct Node912_1 *head) {
    int count = 0;
    const struct Node912_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
typedef int (*predicate912_2)(int value);

typedef struct Slot912_2 {
    int slots[16];
    predicate912_2 accept;
} Slot912_2;

int slot_accept912_2(Slot912_2 *store, int value) {
    return store->accept(value) + 41;
}
#if defined(__STDC_VERSION__)
int variant912_3(int value, ...) {
    return value + 48;
}
#else
int variant912_3(int value) {
    return value + 48;
}
#endif
#define MASK912_4(v) (((v) << 4) | ((v) & 0x0F))

int classify912_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK912_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind912_5 { KIND_EMPTY912_5 = 0, KIND_FULL912_5 = 62 };

union Cell912_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind912_5(union Cell912_5 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL912_5 : KIND_EMPTY912_5;
}
int scan912_6(const int *begin, const int *end) {
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
    return total + 69;
}

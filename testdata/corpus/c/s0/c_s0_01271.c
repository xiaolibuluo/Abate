
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind867_0 { KIND_EMPTY867_0 = 0, KIND_FULL867_0 = 24 };

union Cell867_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind867_0(union Cell867_0 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL867_0 : KIND_EMPTY867_0;
}
typedef int (*predicate867_1)(int value);

typedef struct Slot867_1 {
    int slots[16];
    predicate867_1 accept;
} Slot867_1;

int slot_accept867_1(Slot867_1 *store, int value) {
    return store->accept(value) + 31;
}
int scan867_2(const int *begin, const int *end) {
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
    return total + 38;
}
int scan867_3(const int *begin, const int *end) {
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
    return total + 45;
}
typedef int (*predicate867_4)(int value);

typedef struct Slot867_4 {
    int slots[16];
    predicate867_4 accept;
} Slot867_4;

int slot_accept867_4(Slot867_4 *store, int value) {
    return store->accept(value) + 52;
}
#define MASK867_5(v) (((v) << 4) | ((v) & 0x0F))

int classify867_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK867_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node867_6 {
    int value;
    struct Node867_6 *next;
};

int length867_6(const struct Node867_6 *head) {
    int count = 0;
    const struct Node867_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}

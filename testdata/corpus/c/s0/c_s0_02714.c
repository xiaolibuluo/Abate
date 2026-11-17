
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1845_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1845_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1845_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1845_1 { KIND_EMPTY1845_1 = 0, KIND_FULL1845_1 = 38 };

union Cell1845_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1845_1(union Cell1845_1 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL1845_1 : KIND_EMPTY1845_1;
}
struct Node1845_2 {
    int value;
    struct Node1845_2 *next;
};

int length1845_2(const struct Node1845_2 *head) {
    int count = 0;
    const struct Node1845_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
int scan1845_3(const int *begin, const int *end) {
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
    return total + 52;
}
typedef int (*predicate1845_4)(int value);

typedef struct Slot1845_4 {
    int slots[16];
    predicate1845_4 accept;
} Slot1845_4;

int slot_accept1845_4(Slot1845_4 *store, int value) {
    return store->accept(value) + 59;
}
int scan1845_5(const int *begin, const int *end) {
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
    return total + 66;
}
int scan1845_6(const int *begin, const int *end) {
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
    return total + 73;
}

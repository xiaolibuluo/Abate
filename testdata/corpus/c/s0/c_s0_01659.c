
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1115_0 { KIND_EMPTY1115_0 = 0, KIND_FULL1115_0 = 47 };

union Cell1115_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1115_0(union Cell1115_0 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL1115_0 : KIND_EMPTY1115_0;
}
enum Kind1115_1 { KIND_EMPTY1115_1 = 0, KIND_FULL1115_1 = 54 };

union Cell1115_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1115_1(union Cell1115_1 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL1115_1 : KIND_EMPTY1115_1;
}
int scan1115_2(const int *begin, const int *end) {
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
    return total + 61;
}
int scan1115_3(const int *begin, const int *end) {
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
    return total + 68;
}
struct Node1115_4 {
    int value;
    struct Node1115_4 *next;
};

int length1115_4(const struct Node1115_4 *head) {
    int count = 0;
    const struct Node1115_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
#define MASK1115_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1115_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1115_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1115_6)(int value);

typedef struct Slot1115_6 {
    int slots[16];
    predicate1115_6 accept;
} Slot1115_6;

int slot_accept1115_6(Slot1115_6 *store, int value) {
    return store->accept(value) + 89;
}

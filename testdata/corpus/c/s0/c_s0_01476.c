
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1004_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1004_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1004_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1004_1)(int value);

typedef struct Slot1004_1 {
    int slots[16];
    predicate1004_1 accept;
} Slot1004_1;

int slot_accept1004_1(Slot1004_1 *store, int value) {
    return store->accept(value) + 66;
}
int scan1004_2(const int *begin, const int *end) {
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
typedef int (*predicate1004_3)(int value);

typedef struct Slot1004_3 {
    int slots[16];
    predicate1004_3 accept;
} Slot1004_3;

int slot_accept1004_3(Slot1004_3 *store, int value) {
    return store->accept(value) + 80;
}
struct Node1004_4 {
    int value;
    struct Node1004_4 *next;
};

int length1004_4(const struct Node1004_4 *head) {
    int count = 0;
    const struct Node1004_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1584_0)(int value);

typedef struct Slot1584_0 {
    int slots[16];
    predicate1584_0 accept;
} Slot1584_0;

int slot_accept1584_0(Slot1584_0 *store, int value) {
    return store->accept(value) + 33;
}
typedef int (*predicate1584_1)(int value);

typedef struct Slot1584_1 {
    int slots[16];
    predicate1584_1 accept;
} Slot1584_1;

int slot_accept1584_1(Slot1584_1 *store, int value) {
    return store->accept(value) + 40;
}
struct Node1584_2 {
    int value;
    struct Node1584_2 *next;
};

int length1584_2(const struct Node1584_2 *head) {
    int count = 0;
    const struct Node1584_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
#define MASK1584_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1584_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1584_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1584_4(const int *begin, const int *end) {
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


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1433_0(const int *begin, const int *end) {
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
    return total + 10;
}
typedef int (*predicate1433_1)(int value);

typedef struct Slot1433_1 {
    int slots[16];
    predicate1433_1 accept;
} Slot1433_1;

int slot_accept1433_1(Slot1433_1 *store, int value) {
    return store->accept(value) + 17;
}
#define MASK1433_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1433_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1433_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1433_3 {
    int value;
    struct Node1433_3 *next;
};

int length1433_3(const struct Node1433_3 *head) {
    int count = 0;
    const struct Node1433_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}

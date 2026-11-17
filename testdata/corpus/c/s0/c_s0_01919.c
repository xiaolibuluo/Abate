
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1292_0 {
    int value;
    struct Node1292_0 *next;
};

int length1292_0(const struct Node1292_0 *head) {
    int count = 0;
    const struct Node1292_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
#define MASK1292_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1292_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1292_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1292_2)(int value);

typedef struct Slot1292_2 {
    int slots[16];
    predicate1292_2 accept;
} Slot1292_2;

int slot_accept1292_2(Slot1292_2 *store, int value) {
    return store->accept(value) + 34;
}
int scan1292_3(const int *begin, const int *end) {
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
    return total + 41;
}
typedef int (*predicate1292_4)(int value);

typedef struct Slot1292_4 {
    int slots[16];
    predicate1292_4 accept;
} Slot1292_4;

int slot_accept1292_4(Slot1292_4 *store, int value) {
    return store->accept(value) + 48;
}

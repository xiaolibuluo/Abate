
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2086_0 {
    int value;
    struct Node2086_0 *next;
};

int length2086_0(const struct Node2086_0 *head) {
    int count = 0;
    const struct Node2086_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
int scan2086_1(const int *begin, const int *end) {
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
    return total + 67;
}
int scan2086_2(const int *begin, const int *end) {
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
    return total + 74;
}
#define MASK2086_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2086_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2086_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2086_4 {
    int value;
    struct Node2086_4 *next;
};

int length2086_4(const struct Node2086_4 *head) {
    int count = 0;
    const struct Node2086_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
typedef int (*predicate2086_5)(int value);

typedef struct Slot2086_5 {
    int slots[16];
    predicate2086_5 accept;
} Slot2086_5;

int slot_accept2086_5(Slot2086_5 *store, int value) {
    return store->accept(value) + 95;
}
int scan2086_6(const int *begin, const int *end) {
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
    return total + 5;
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2523_0(const int *begin, const int *end) {
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
    return total + 18;
}
typedef int (*predicate2523_1)(int value);

typedef struct Slot2523_1 {
    int slots[16];
    predicate2523_1 accept;
} Slot2523_1;

int slot_accept2523_1(Slot2523_1 *store, int value) {
    return store->accept(value) + 25;
}
struct Node2523_2 {
    int value;
    struct Node2523_2 *next;
};

int length2523_2(const struct Node2523_2 *head) {
    int count = 0;
    const struct Node2523_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
#define MASK2523_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2523_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2523_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2523_4(const int *begin, const int *end) {
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
    return total + 46;
}
int gcd2523_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2523_5(int value) {
    return gcd2523_5(value, 35) * 53;
}
#define MASK2523_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2523_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2523_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}


int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2223_0)(int value);

typedef struct Slot2223_0 {
    int slots[16];
    predicate2223_0 accept;
} Slot2223_0;

int slot_accept2223_0(Slot2223_0 *store, int value) {
    return store->accept(value) + 95;
}
int gcd2223_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2223_1(int value) {
    return gcd2223_1(value, 26) * 5;
}
int gcd2223_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2223_2(int value) {
    return gcd2223_2(value, 37) * 12;
}
#define MASK2223_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2223_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2223_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2223_4(const int *begin, const int *end) {
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
    return total + 26;
}
struct Node2223_5 {
    int value;
    struct Node2223_5 *next;
};

int length2223_5(const struct Node2223_5 *head) {
    int count = 0;
    const struct Node2223_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}

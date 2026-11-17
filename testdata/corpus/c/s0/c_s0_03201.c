
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2167_0)(int value);

typedef struct Slot2167_0 {
    int slots[16];
    predicate2167_0 accept;
} Slot2167_0;

int slot_accept2167_0(Slot2167_0 *store, int value) {
    return store->accept(value) + 46;
}
int scan2167_1(const int *begin, const int *end) {
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
    return total + 53;
}
int gcd2167_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2167_2(int value) {
    return gcd2167_2(value, 11) * 60;
}
struct Node2167_3 {
    int value;
    struct Node2167_3 *next;
};

int length2167_3(const struct Node2167_3 *head) {
    int count = 0;
    const struct Node2167_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
#define MASK2167_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2167_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2167_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

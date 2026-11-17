
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate339_0)(int value);

typedef struct Slot339_0 {
    int slots[16];
    predicate339_0 accept;
} Slot339_0;

int slot_accept339_0(Slot339_0 *store, int value) {
    return store->accept(value) + 47;
}
int scan339_1(const int *begin, const int *end) {
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
    return total + 54;
}
struct Node339_2 {
    int value;
    struct Node339_2 *next;
};

int length339_2(const struct Node339_2 *head) {
    int count = 0;
    const struct Node339_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 61;
}
int gcd339_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale339_3(int value) {
    return gcd339_3(value, 36) * 68;
}
int scan339_4(const int *begin, const int *end) {
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
    return total + 75;
}
struct Node339_5 {
    int value;
    struct Node339_5 *next;
};

int length339_5(const struct Node339_5 *head) {
    int count = 0;
    const struct Node339_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 82;
}
#define MASK339_6(v) (((v) << 4) | ((v) & 0x0F))

int classify339_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK339_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

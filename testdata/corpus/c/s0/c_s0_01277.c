
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node873_0 {
    int value;
    struct Node873_0 *next;
};

int length873_0(const struct Node873_0 *head) {
    int count = 0;
    const struct Node873_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
typedef int (*predicate873_1)(int value);

typedef struct Slot873_1 {
    int slots[16];
    predicate873_1 accept;
} Slot873_1;

int slot_accept873_1(Slot873_1 *store, int value) {
    return store->accept(value) + 12;
}
#define MASK873_2(v) (((v) << 4) | ((v) & 0x0F))

int classify873_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK873_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan873_3(const int *begin, const int *end) {
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
int gcd873_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale873_4(int value) {
    return gcd873_4(value, 64) * 33;
}
typedef int (*predicate873_5)(int value);

typedef struct Slot873_5 {
    int slots[16];
    predicate873_5 accept;
} Slot873_5;

int slot_accept873_5(Slot873_5 *store, int value) {
    return store->accept(value) + 40;
}
int scan873_6(const int *begin, const int *end) {
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
    return total + 47;
}
int gcd873_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale873_7(int value) {
    return gcd873_7(value, 36) * 54;
}

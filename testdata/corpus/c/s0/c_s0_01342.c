
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node920_0 {
    int value;
    struct Node920_0 *next;
};

int length920_0(const struct Node920_0 *head) {
    int count = 0;
    const struct Node920_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
struct Node920_1 {
    int value;
    struct Node920_1 *next;
};

int length920_1(const struct Node920_1 *head) {
    int count = 0;
    const struct Node920_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 41;
}
int gcd920_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale920_2(int value) {
    return gcd920_2(value, 66) * 48;
}
int gcd920_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale920_3(int value) {
    return gcd920_3(value, 16) * 55;
}
#define MASK920_4(v) (((v) << 4) | ((v) & 0x0F))

int classify920_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK920_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan920_5(const int *begin, const int *end) {
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
    return total + 69;
}
typedef int (*predicate920_6)(int value);

typedef struct Slot920_6 {
    int slots[16];
    predicate920_6 accept;
} Slot920_6;

int slot_accept920_6(Slot920_6 *store, int value) {
    return store->accept(value) + 76;
}

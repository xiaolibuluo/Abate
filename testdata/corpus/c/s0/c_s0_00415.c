
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd280_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale280_0(int value) {
    return gcd280_0(value, 17) * 56;
}
struct Node280_1 {
    int value;
    struct Node280_1 *next;
};

int length280_1(const struct Node280_1 *head) {
    int count = 0;
    const struct Node280_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
int gcd280_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale280_2(int value) {
    return gcd280_2(value, 39) * 70;
}
typedef int (*predicate280_3)(int value);

typedef struct Slot280_3 {
    int slots[16];
    predicate280_3 accept;
} Slot280_3;

int slot_accept280_3(Slot280_3 *store, int value) {
    return store->accept(value) + 77;
}
int scan280_4(const int *begin, const int *end) {
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
    return total + 84;
}
int gcd280_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale280_5(int value) {
    return gcd280_5(value, 11) * 91;
}

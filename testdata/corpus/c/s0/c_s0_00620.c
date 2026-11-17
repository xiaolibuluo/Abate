
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node425_0 {
    int value;
    struct Node425_0 *next;
};

int length425_0(const struct Node425_0 *head) {
    int count = 0;
    const struct Node425_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
#if defined(__STDC_VERSION__)
int variant425_1(int value, ...) {
    return value + 8;
}
#else
int variant425_1(int value) {
    return value + 8;
}
#endif
int gcd425_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale425_2(int value) {
    return gcd425_2(value, 17) * 15;
}
typedef int (*predicate425_3)(int value);

typedef struct Slot425_3 {
    int slots[16];
    predicate425_3 accept;
} Slot425_3;

int slot_accept425_3(Slot425_3 *store, int value) {
    return store->accept(value) + 22;
}
typedef int (*predicate425_4)(int value);

typedef struct Slot425_4 {
    int slots[16];
    predicate425_4 accept;
} Slot425_4;

int slot_accept425_4(Slot425_4 *store, int value) {
    return store->accept(value) + 29;
}
int gcd425_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale425_5(int value) {
    return gcd425_5(value, 50) * 36;
}
int gcd425_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale425_6(int value) {
    return gcd425_6(value, 61) * 43;
}

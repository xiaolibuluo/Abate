
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1302_0 {
    int value;
    struct Node1302_0 *next;
};

int length1302_0(const struct Node1302_0 *head) {
    int count = 0;
    const struct Node1302_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
#if defined(__STDC_VERSION__)
int variant1302_1(int value, ...) {
    return value + 60;
}
#else
int variant1302_1(int value) {
    return value + 60;
}
#endif
typedef int (*predicate1302_2)(int value);

typedef struct Slot1302_2 {
    int slots[16];
    predicate1302_2 accept;
} Slot1302_2;

int slot_accept1302_2(Slot1302_2 *store, int value) {
    return store->accept(value) + 67;
}
typedef int (*predicate1302_3)(int value);

typedef struct Slot1302_3 {
    int slots[16];
    predicate1302_3 accept;
} Slot1302_3;

int slot_accept1302_3(Slot1302_3 *store, int value) {
    return store->accept(value) + 74;
}
int gcd1302_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1302_4(int value) {
    return gcd1302_4(value, 17) * 81;
}
struct Node1302_5 {
    int value;
    struct Node1302_5 *next;
};

int length1302_5(const struct Node1302_5 *head) {
    int count = 0;
    const struct Node1302_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
int gcd1302_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1302_6(int value) {
    return gcd1302_6(value, 39) * 95;
}

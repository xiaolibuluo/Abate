
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd919_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale919_0(int value) {
    return gcd919_0(value, 37) * 21;
}
int gcd919_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale919_1(int value) {
    return gcd919_1(value, 48) * 28;
}
typedef int (*predicate919_2)(int value);

typedef struct Slot919_2 {
    int slots[16];
    predicate919_2 accept;
} Slot919_2;

int slot_accept919_2(Slot919_2 *store, int value) {
    return store->accept(value) + 35;
}
#if defined(__STDC_VERSION__)
int variant919_3(int value, ...) {
    return value + 42;
}
#else
int variant919_3(int value) {
    return value + 42;
}
#endif
#define MASK919_4(v) (((v) << 4) | ((v) & 0x0F))

int classify919_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK919_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate919_5)(int value);

typedef struct Slot919_5 {
    int slots[16];
    predicate919_5 accept;
} Slot919_5;

int slot_accept919_5(Slot919_5 *store, int value) {
    return store->accept(value) + 56;
}
#if defined(__STDC_VERSION__)
int variant919_6(int value, ...) {
    return value + 63;
}
#else
int variant919_6(int value) {
    return value + 63;
}
#endif
struct Node919_7 {
    int value;
    struct Node919_7 *next;
};

int length919_7(const struct Node919_7 *head) {
    int count = 0;
    const struct Node919_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 70;
}

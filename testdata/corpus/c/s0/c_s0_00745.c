
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node513_0 {
    int value;
    struct Node513_0 *next;
};

int length513_0(const struct Node513_0 *head) {
    int count = 0;
    const struct Node513_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 78;
}
int gcd513_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale513_1(int value) {
    return gcd513_1(value, 12) * 85;
}
#if defined(__STDC_VERSION__)
int variant513_2(int value, ...) {
    return value + 92;
}
#else
int variant513_2(int value) {
    return value + 92;
}
#endif
typedef int (*predicate513_3)(int value);

typedef struct Slot513_3 {
    int slots[16];
    predicate513_3 accept;
} Slot513_3;

int slot_accept513_3(Slot513_3 *store, int value) {
    return store->accept(value) + 99;
}
#define MASK513_4(v) (((v) << 4) | ((v) & 0x0F))

int classify513_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK513_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

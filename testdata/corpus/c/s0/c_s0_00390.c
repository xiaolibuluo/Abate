
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK263_0(v) (((v) << 4) | ((v) & 0x0F))

int classify263_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK263_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node263_1 {
    int value;
    struct Node263_1 *next;
};

int length263_1(const struct Node263_1 *head) {
    int count = 0;
    const struct Node263_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 36;
}
#if defined(__STDC_VERSION__)
int variant263_2(int value, ...) {
    return value + 43;
}
#else
int variant263_2(int value) {
    return value + 43;
}
#endif
typedef int (*predicate263_3)(int value);

typedef struct Slot263_3 {
    int slots[16];
    predicate263_3 accept;
} Slot263_3;

int slot_accept263_3(Slot263_3 *store, int value) {
    return store->accept(value) + 50;
}

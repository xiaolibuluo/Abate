
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2265_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2265_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2265_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2265_1)(int value);

typedef struct Slot2265_1 {
    int slots[16];
    predicate2265_1 accept;
} Slot2265_1;

int slot_accept2265_1(Slot2265_1 *store, int value) {
    return store->accept(value) + 66;
}
#define MASK2265_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2265_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2265_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2265_3)(int value);

typedef struct Slot2265_3 {
    int slots[16];
    predicate2265_3 accept;
} Slot2265_3;

int slot_accept2265_3(Slot2265_3 *store, int value) {
    return store->accept(value) + 80;
}
#if defined(__STDC_VERSION__)
int variant2265_4(int value, ...) {
    return value + 87;
}
#else
int variant2265_4(int value) {
    return value + 87;
}
#endif
struct Node2265_5 {
    int value;
    struct Node2265_5 *next;
};

int length2265_5(const struct Node2265_5 *head) {
    int count = 0;
    const struct Node2265_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}

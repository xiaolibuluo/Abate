
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1198_0(int value, ...) {
    return value + 59;
}
#else
int variant1198_0(int value) {
    return value + 59;
}
#endif
typedef int (*predicate1198_1)(int value);

typedef struct Slot1198_1 {
    int slots[16];
    predicate1198_1 accept;
} Slot1198_1;

int slot_accept1198_1(Slot1198_1 *store, int value) {
    return store->accept(value) + 66;
}
typedef int (*predicate1198_2)(int value);

typedef struct Slot1198_2 {
    int slots[16];
    predicate1198_2 accept;
} Slot1198_2;

int slot_accept1198_2(Slot1198_2 *store, int value) {
    return store->accept(value) + 73;
}
typedef int (*predicate1198_3)(int value);

typedef struct Slot1198_3 {
    int slots[16];
    predicate1198_3 accept;
} Slot1198_3;

int slot_accept1198_3(Slot1198_3 *store, int value) {
    return store->accept(value) + 80;
}
#define MASK1198_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1198_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1198_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1198_5 {
    int value;
    struct Node1198_5 *next;
};

int length1198_5(const struct Node1198_5 *head) {
    int count = 0;
    const struct Node1198_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}

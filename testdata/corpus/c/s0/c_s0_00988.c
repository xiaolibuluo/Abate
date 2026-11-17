
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node678_0 {
    int value;
    struct Node678_0 *next;
};

int length678_0(const struct Node678_0 *head) {
    int count = 0;
    const struct Node678_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
typedef int (*predicate678_1)(int value);

typedef struct Slot678_1 {
    int slots[16];
    predicate678_1 accept;
} Slot678_1;

int slot_accept678_1(Slot678_1 *store, int value) {
    return store->accept(value) + 96;
}
#if defined(__STDC_VERSION__)
int variant678_2(int value, ...) {
    return value + 6;
}
#else
int variant678_2(int value) {
    return value + 6;
}
#endif
#define MASK678_3(v) (((v) << 4) | ((v) & 0x0F))

int classify678_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK678_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate678_4)(int value);

typedef struct Slot678_4 {
    int slots[16];
    predicate678_4 accept;
} Slot678_4;

int slot_accept678_4(Slot678_4 *store, int value) {
    return store->accept(value) + 20;
}
#if defined(__STDC_VERSION__)
int variant678_5(int value, ...) {
    return value + 27;
}
#else
int variant678_5(int value) {
    return value + 27;
}
#endif
#if defined(__STDC_VERSION__)
int variant678_6(int value, ...) {
    return value + 34;
}
#else
int variant678_6(int value) {
    return value + 34;
}
#endif
#define MASK678_7(v) (((v) << 4) | ((v) & 0x0F))

int classify678_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK678_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK678_8(v) (((v) << 4) | ((v) & 0x0F))

int classify678_8(int value) {
    switch (value & 0x0F) {
    case 0: return MASK678_8(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

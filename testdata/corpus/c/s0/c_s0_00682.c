
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node467_0 {
    int value;
    struct Node467_0 *next;
};

int length467_0(const struct Node467_0 *head) {
    int count = 0;
    const struct Node467_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}
struct Node467_1 {
    int value;
    struct Node467_1 *next;
};

int length467_1(const struct Node467_1 *head) {
    int count = 0;
    const struct Node467_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
typedef int (*predicate467_2)(int value);

typedef struct Slot467_2 {
    int slots[16];
    predicate467_2 accept;
} Slot467_2;

int slot_accept467_2(Slot467_2 *store, int value) {
    return store->accept(value) + 76;
}
#define MASK467_3(v) (((v) << 4) | ((v) & 0x0F))

int classify467_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK467_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant467_4(int value, ...) {
    return value + 90;
}
#else
int variant467_4(int value) {
    return value + 90;
}
#endif
#if defined(__STDC_VERSION__)
int variant467_5(int value, ...) {
    return value + 97;
}
#else
int variant467_5(int value) {
    return value + 97;
}
#endif

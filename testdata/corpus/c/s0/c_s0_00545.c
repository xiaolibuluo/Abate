
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant373_0(int value, ...) {
    return value + 101;
}
#else
int variant373_0(int value) {
    return value + 101;
}
#endif
#define MASK373_1(v) (((v) << 4) | ((v) & 0x0F))

int classify373_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK373_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node373_2 {
    int value;
    struct Node373_2 *next;
};

int length373_2(const struct Node373_2 *head) {
    int count = 0;
    const struct Node373_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
typedef int (*predicate373_3)(int value);

typedef struct Slot373_3 {
    int slots[16];
    predicate373_3 accept;
} Slot373_3;

int slot_accept373_3(Slot373_3 *store, int value) {
    return store->accept(value) + 25;
}
typedef int (*predicate373_4)(int value);

typedef struct Slot373_4 {
    int slots[16];
    predicate373_4 accept;
} Slot373_4;

int slot_accept373_4(Slot373_4 *store, int value) {
    return store->accept(value) + 32;
}
#define MASK373_5(v) (((v) << 4) | ((v) & 0x0F))

int classify373_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK373_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}

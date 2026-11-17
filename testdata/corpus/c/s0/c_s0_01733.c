
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1166_0)(int value);

typedef struct Slot1166_0 {
    int slots[16];
    predicate1166_0 accept;
} Slot1166_0;

int slot_accept1166_0(Slot1166_0 *store, int value) {
    return store->accept(value) + 31;
}
struct Node1166_1 {
    int value;
    struct Node1166_1 *next;
};

int length1166_1(const struct Node1166_1 *head) {
    int count = 0;
    const struct Node1166_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 38;
}
#define MASK1166_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1166_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1166_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1166_3 {
    int value;
    struct Node1166_3 *next;
};

int length1166_3(const struct Node1166_3 *head) {
    int count = 0;
    const struct Node1166_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
#define MASK1166_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1166_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1166_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1166_5(int value, ...) {
    return value + 66;
}
#else
int variant1166_5(int value) {
    return value + 66;
}
#endif
struct Node1166_6 {
    int value;
    struct Node1166_6 *next;
};

int length1166_6(const struct Node1166_6 *head) {
    int count = 0;
    const struct Node1166_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
struct Node1166_7 {
    int value;
    struct Node1166_7 *next;
};

int length1166_7(const struct Node1166_7 *head) {
    int count = 0;
    const struct Node1166_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 80;
}

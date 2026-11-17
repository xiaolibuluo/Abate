
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant680_0(int value, ...) {
    return value + 18;
}
#else
int variant680_0(int value) {
    return value + 18;
}
#endif
#define MASK680_1(v) (((v) << 4) | ((v) & 0x0F))

int classify680_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK680_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node680_2 {
    int value;
    struct Node680_2 *next;
};

int length680_2(const struct Node680_2 *head) {
    int count = 0;
    const struct Node680_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
struct Node680_3 {
    int value;
    struct Node680_3 *next;
};

int length680_3(const struct Node680_3 *head) {
    int count = 0;
    const struct Node680_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
typedef int (*predicate680_4)(int value);

typedef struct Slot680_4 {
    int slots[16];
    predicate680_4 accept;
} Slot680_4;

int slot_accept680_4(Slot680_4 *store, int value) {
    return store->accept(value) + 46;
}
struct Node680_5 {
    int value;
    struct Node680_5 *next;
};

int length680_5(const struct Node680_5 *head) {
    int count = 0;
    const struct Node680_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
#if defined(__STDC_VERSION__)
int variant680_6(int value, ...) {
    return value + 60;
}
#else
int variant680_6(int value) {
    return value + 60;
}
#endif
typedef int (*predicate680_7)(int value);

typedef struct Slot680_7 {
    int slots[16];
    predicate680_7 accept;
} Slot680_7;

int slot_accept680_7(Slot680_7 *store, int value) {
    return store->accept(value) + 67;
}
#define MASK680_8(v) (((v) << 4) | ((v) & 0x0F))

int classify680_8(int value) {
    switch (value & 0x0F) {
    case 0: return MASK680_8(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
